package validate

import (
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

type Validator interface {
	GotenValidate() error
}

type CustomValidator interface {
	GotenCustomValidate() error
}

type ValidationError interface {
	error

	ProtoMessageName() string
	FieldName() string
	Value() interface{}
	ErrorMessage() string
	Cause() error
}

type validationError struct {
	protoMessageName string
	fieldName        string
	value            interface{}
	errorMessage     string
	cause            error
}

func NewValidationError(
	protoMessageName string, fieldName string,
	value interface{}, errorMessage string,
	cause error,
) ValidationError {
	return &validationError{
		protoMessageName, fieldName,
		value, errorMessage,
		cause,
	}
}

func (vErr *validationError) ProtoMessageName() string {
	return vErr.protoMessageName
}

func (vErr *validationError) FieldName() string {
	return vErr.fieldName
}

func (vErr *validationError) Value() interface{} {
	return vErr.value
}

func (vErr *validationError) ErrorMessage() string {
	return vErr.errorMessage
}

func (vErr *validationError) Cause() error {
	return vErr.cause
}

func trystringer(v interface{}) string {
	if sv, ok := v.(fmt.Stringer); ok {
		return sv.String()
	} else {
		return fmt.Sprintf("%#v", v)
	}
}

func extractRootCause(verr ValidationError) (fp []string, msg string, value interface{}) {
	fp = append(fp, verr.FieldName())
	if cause := verr.Cause(); cause == nil {
		return fp, verr.ErrorMessage(), verr.Value()
	} else {
		if cverr, ok := cause.(ValidationError); ok {
			cvfp, cmsg, cv := extractRootCause(cverr)
			fp = append(fp, cvfp...)
			return fp, cmsg, cv
		} else {
			return fp, cause.Error(), verr.Value()
		}
	}
}

func (vErr *validationError) Error() string {
	fp, msg, v := extractRootCause(vErr)
	return fmt.Sprintf("validation error: %s.%s: %s (got: %s)", vErr.protoMessageName, strings.Join(fp, "."), msg, trystringer(v))
}

func ValidateEmail(s string) error {
	a, err := mail.ParseAddress(s)
	if err != nil {
		return err
	}

	parts := strings.SplitN(a.Address, "@", 2)
	return ValidateHostname(parts[1])
}

func ValidateAddress(s string) error {
	if net.ParseIP(s) == nil {
		return nil
	}
	if err := ValidateHostname(s); err == nil {
		return nil
	}
	return errors.New("address must be either a valid IP, or a valid hostname")
}

func ValidateHostname(s string) error {
	if len(s) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	s = strings.ToLower(strings.TrimSuffix(s, "."))

	for _, part := range strings.Split(s, ".") {
		if len(part) == 0 || len(part) > 63 {
			return errors.New("hostname part cannot be empty and cannot exceed 63 characters")
		}

		if part[0] == '-' || part[len(part)-1] == '-' {
			return errors.New("hostname part cannot start or end with a hyphen")
		}

		for _, c := range part {
			if (c < 'a' || c > 'z') && (c < '0' || c > '9') && c != '-' {
				return fmt.Errorf("hostname part can contain only letters, digits or hyphens, got %q", string(c))
			}
		}
	}

	return nil
}

func ValidateUUID(s string) error {
	_, err := uuid.Parse(s)
	return err
}

func ValidatePEM(s string) error {
	var block *pem.Block
	rest := []byte(s)
	bytesOk := 0

	for {
		block, rest = pem.Decode([]byte(rest))
		if block == nil {
			if len(rest) > 0 {
				return fmt.Errorf("PEM encoding not satisfied starting at byte %d", bytesOk+1)
			} else {
				return nil
			}
		}
		bytesOk += len(block.Bytes)
	}
}
