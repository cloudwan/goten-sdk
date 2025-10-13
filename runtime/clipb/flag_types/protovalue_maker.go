package flag_types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

func makeProtoValue(raw string, fd preflect.FieldDescriptor, tp reflect.Type) (preflect.Value, error) {
	if tp.Implements(reflect.TypeOf((*CustomTypeCliValue)(nil)).Elem()) {
		custom := reflect.New(tp.Elem()).Interface().(CustomTypeCliValue)
		if err := custom.SetFromCliFlag(raw); err != nil {
			return preflect.Value{}, err
		}
		if asProtoStringer, ok := custom.(preflect.ProtoStringer); ok {
			return preflect.ValueOfProtoString(asProtoStringer), nil
		}
		return preflect.ValueOfMessage(custom.(proto.Message).ProtoReflect()), nil
	}

	switch fd.Kind() {
	case preflect.Int32Kind, preflect.Sint32Kind, preflect.Sfixed32Kind:
		tmp, err := strconv.ParseInt(raw, 10, 32)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfInt32(int32(tmp)), nil
	case preflect.Int64Kind, preflect.Sint64Kind, preflect.Sfixed64Kind:
		tmp, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfInt64(tmp), nil
	case preflect.Uint32Kind, preflect.Fixed32Kind:
		tmp, err := strconv.ParseUint(raw, 10, 32)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfUint32(uint32(tmp)), nil
	case preflect.Uint64Kind, preflect.Fixed64Kind:
		tmp, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfUint64(tmp), nil
	case preflect.FloatKind:
		tmp, err := strconv.ParseFloat(raw, 32)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfFloat32(float32(tmp)), nil
	case preflect.DoubleKind:
		tmp, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfFloat64(tmp), nil
	case preflect.BoolKind:
		tmp, err := strconv.ParseBool(raw)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfBool(tmp), nil
	case preflect.StringKind:
		return preflect.ValueOfString(raw), nil
	case preflect.BytesKind:
		bytes, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfBytes(bytes), nil
	case preflect.EnumKind:
		enValDes := fd.Enum().Values().ByName(preflect.Name(raw))
		if enValDes == nil {
			return preflect.Value{}, status.Errorf(
				codes.InvalidArgument, "Enum value %s not found in %s", raw, fd.Enum().Name())
		}
		return preflect.ValueOfEnum(enValDes.Number()), nil
	case preflect.MessageKind, preflect.GroupKind:
		if tp.Kind() == reflect.Pointer {
			tp = tp.Elem()
		}
		subMsg := reflect.New(tp).Interface().(proto.Message)
		// those timestamps somewhat dont work in protojson...
		if asTimestamp, isTimestamp := subMsg.(*timestamppb.Timestamp); isTimestamp {
			t, err := time.Parse(time.RFC3339Nano, raw)
			if err != nil {
				return preflect.Value{}, status.Errorf(
					codes.InvalidArgument, "Failed to parse timestamp %s: %s", raw, err)
			}
			asTimestamp.Seconds = t.Unix()
			asTimestamp.Nanos = int32(t.Nanosecond())
		} else if asDuration, isDuration := subMsg.(*durationpb.Duration); isDuration {
			d, err := time.ParseDuration(raw)
			if err != nil {
				return preflect.Value{}, status.Errorf(
					codes.InvalidArgument, "Failed to parse duration %s: %s", raw, err)
			}
			asDuration.Seconds = int64(d / time.Second)
			asDuration.Nanos = int32(d % time.Nanosecond)
		} else if asMoney, isMoney := subMsg.(*money.Money); isMoney {
			// Handle JSON strings like "$0.25" when embedded in JSON
			// (protojson will pass us quoted strings for string values in JSON)
			if isJSONString(raw) {
				moneyStr, err := strconv.Unquote(raw)
				if err == nil {
					// Try to parse as money string
					amount, currency, parseErr := parseMoney(moneyStr)
					if parseErr == nil {
						// Convert to proper JSON object for protojson
						raw = moneyToJSON(amount, currency)
					}
					// If parsing fails, fall through to protojson unmarshal (might be valid JSON object)
				}
			}
			// Try direct string parsing first (for non-JSON context)
			amount, currency, err := parseMoney(raw)
			if err == nil {
				asMoney.CurrencyCode = currency
				asMoney.Units = int64(amount)
				asMoney.Nanos = int32((amount - float64(int64(amount))) * 1e9)
			} else {
				// Fall back to protojson unmarshal for standard JSON format
				if err := protojson.Unmarshal([]byte(raw), subMsg); err != nil {
					return preflect.Value{}, err
				}
			}
		} else {
			// Preprocess JSON to handle Money string fields before protojson
			preprocessed, err := preprocessMoneyStringsInJSON(raw, subMsg.ProtoReflect().Descriptor())
			if err != nil {
				return preflect.Value{}, err
			}
			if err := protojson.Unmarshal([]byte(preprocessed), subMsg); err != nil {
				return preflect.Value{}, err
			}
		}
		return preflect.ValueOfMessage(subMsg.ProtoReflect()), nil
	}

	return preflect.Value{}, status.Errorf(
		codes.InvalidArgument, "Unrecognized kind %s for raw value %s", fd.Name(), raw)
}

// parseMoney parses money strings in multiple formats:
// "$2.50", "2.50$", "2.50 USD", "USD 2.50", "2.50USD", "USD2.50", "0.25" (defaults to USD)
// Returns: (amount, currency_code, error)
func parseMoney(raw string) (float64, string, error) {
	raw = strings.TrimSpace(raw)

	// Handle "$" prefix: "$2.50"
	if strings.HasPrefix(raw, "$") {
		amountStr := strings.TrimPrefix(raw, "$")
		amount, err := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
		if err != nil {
			return 0, "", err
		}
		return amount, "USD", nil
	}

	// Handle "$" suffix: "2.50$"
	if strings.HasSuffix(raw, "$") {
		amountStr := strings.TrimSuffix(raw, "$")
		amount, err := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
		if err != nil {
			return 0, "", err
		}
		return amount, "USD", nil
	}

	// Try to split by space
	parts := strings.Fields(raw)

	// Format: "2.50 USD" or "USD 2.50"
	if len(parts) == 2 {
		// Try first part as amount
		amount, err1 := strconv.ParseFloat(parts[0], 64)
		if err1 == nil {
			return amount, strings.ToUpper(parts[1]), nil
		}

		// Try second part as amount
		amount, err2 := strconv.ParseFloat(parts[1], 64)
		if err2 == nil {
			return amount, strings.ToUpper(parts[0]), nil
		}

		return 0, "", status.Errorf(codes.InvalidArgument,
			"could not parse money from '%s': neither part is a valid number", raw)
	}

	// Single token - try to extract number and currency
	if len(parts) == 1 {
		// Check if it starts with a currency code (e.g., "USD2.5")
		if len(raw) > 3 {
			possibleCurrency := strings.ToUpper(raw[:3])
			if isValidCurrency(possibleCurrency) {
				amountStr := raw[3:]
				amount, err := strconv.ParseFloat(amountStr, 64)
				if err == nil {
					return amount, possibleCurrency, nil
				}
			}
		}

		// Check if it ends with a currency code (e.g., "2.5USD")
		if len(raw) > 3 {
			possibleCurrency := strings.ToUpper(raw[len(raw)-3:])
			if isValidCurrency(possibleCurrency) {
				amountStr := raw[:len(raw)-3]
				amount, err := strconv.ParseFloat(amountStr, 64)
				if err == nil {
					return amount, possibleCurrency, nil
				}
			}
		}

		// No currency found, assume USD
		amount, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return 0, "", err
		}
		return amount, "USD", nil
	}

	return 0, "", status.Errorf(codes.InvalidArgument,
		"invalid money format '%s': expected formats like '2.50 USD', 'USD2.50', '2.50USD', '$2.50', '2.50$', or '0.25'", raw)
}

// isValidCurrency checks if a string is a valid ISO 4217 currency code
// This is a simplified check - just verify it's 3 letters
func isValidCurrency(s string) bool {
	if len(s) != 3 {
		return false
	}
	s = strings.ToUpper(s)
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			return false
		}
	}
	return true
}

// isJSONString checks if a string is a JSON-quoted string (starts and ends with ")
func isJSONString(s string) bool {
	s = strings.TrimSpace(s)
	return len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"'
}

// moneyToJSON converts amount and currency to a JSON object string for protojson
func moneyToJSON(amount float64, currencyCode string) string {
	units := int64(amount)
	nanos := int32((amount - float64(units)) * 1e9)

	// Handle negative amounts correctly (signs must align)
	if amount < 0 && nanos > 0 {
		units--
		nanos = int32(1e9) - nanos
		nanos = -nanos
	}

	return fmt.Sprintf(`{"currencyCode":"%s","units":"%d","nanos":%d}`, currencyCode, units, nanos)
}

// preprocessMoneyStringsInJSON walks through JSON and converts Money string fields to objects
// This is necessary because protojson doesn't support custom unmarshaling for google.type.Money
func preprocessMoneyStringsInJSON(jsonStr string, md preflect.MessageDescriptor) (string, error) {
	// Parse JSON into generic tree
	var node interface{}
	if err := json.Unmarshal([]byte(jsonStr), &node); err != nil {
		return jsonStr, err // Not valid JSON, let protojson handle the error
	}

	// Rewrite Money strings to objects using message descriptor
	rewritten, err := rewriteMoneyFields(node, md)
	if err != nil {
		return "", err
	}

	// Marshal back to JSON
	result, err := json.Marshal(rewritten)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// rewriteMoneyFields recursively walks the JSON tree and converts Money strings to objects
func rewriteMoneyFields(node interface{}, md preflect.MessageDescriptor) (interface{}, error) {
	obj, ok := node.(map[string]interface{})
	if !ok {
		return node, nil // Not an object, return as-is
	}

	// Build lookup maps for field names
	byJSONName := make(map[string]preflect.FieldDescriptor)
	byProtoName := make(map[string]preflect.FieldDescriptor)
	for i := 0; i < md.Fields().Len(); i++ {
		fd := md.Fields().Get(i)
		byJSONName[fd.JSONName()] = fd
		byProtoName[string(fd.Name())] = fd
	}

	// Process each field in the JSON object
	for key, value := range obj {
		// Find field descriptor
		fd, ok := byJSONName[key]
		if !ok {
			fd, ok = byProtoName[key]
			if !ok {
				continue // Unknown field, let protojson handle it
			}
		}

		// Handle repeated fields
		if fd.IsList() {
			arr, isArray := value.([]interface{})
			if !isArray {
				continue
			}
			if fd.Kind() == preflect.MessageKind {
				if isMoneyMessage(fd.Message()) {
					// Repeated Money field
					for i, elem := range arr {
						if strVal, isStr := elem.(string); isStr {
							moneyObj, err := moneyStringToJSONObject(strVal)
							if err != nil {
								return nil, err
							}
							arr[i] = moneyObj
						}
						// If already an object, leave it as-is
					}
				} else {
					// Repeated message field (not Money)
					for i, elem := range arr {
						if subObj, isObj := elem.(map[string]interface{}); isObj {
							rewritten, err := rewriteMoneyFields(subObj, fd.Message())
							if err != nil {
								return nil, err
							}
							arr[i] = rewritten
						}
					}
				}
			}
			obj[key] = arr
			continue
		}

		// Handle singular message fields
		if fd.Kind() == preflect.MessageKind {
			if isMoneyMessage(fd.Message()) {
				// Money field
				if strVal, isStr := value.(string); isStr {
					moneyObj, err := moneyStringToJSONObject(strVal)
					if err != nil {
						return nil, err
					}
					obj[key] = moneyObj
				}
				// If already an object, leave it as-is
			} else {
				// Nested message (not Money)
				if subObj, isObj := value.(map[string]interface{}); isObj {
					rewritten, err := rewriteMoneyFields(subObj, fd.Message())
					if err != nil {
						return nil, err
					}
					obj[key] = rewritten
				}
			}
		}
	}

	return obj, nil
}

// isMoneyMessage checks if a message descriptor is google.type.Money
func isMoneyMessage(md preflect.MessageDescriptor) bool {
	return string(md.FullName()) == "google.type.Money"
}

// moneyStringToJSONObject converts a money string to a JSON object (map[string]interface{})
func moneyStringToJSONObject(s string) (map[string]interface{}, error) {
	amount, currency, err := parseMoney(s)
	if err != nil {
		return nil, err
	}

	units := int64(amount)
	nanos := int32((amount - float64(units)) * 1e9)

	// Handle negative amounts correctly (signs must align)
	if amount < 0 && nanos > 0 {
		units--
		nanos = int32(1e9) - nanos
		nanos = -nanos
	}

	return map[string]interface{}{
		"currencyCode": currency,
		"units":        fmt.Sprintf("%d", units),
		"nanos":        int64(nanos),
	}, nil
}
