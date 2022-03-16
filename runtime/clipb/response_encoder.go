package clipb

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/resource"
)

const (
	MaxTableViewColumnWidth = 200
)

// ResponseEncoder encodes one or more messages in a desired format.
// There are two types of encoders:
//   streaming - Output data as quickly as it provided by SetNames and Add.
//               In case of streaming responses, they output data as soon
//               as it is received.
//   non-streaming - Need to know all data upfront in order to be able
//                   to format it. Data will be produced only after Finish() is called.
type ResponseEncoder interface {
	SetColumns(fieldPaths []object.FieldPath, displayNames []string)
	Add(msg proto.Message) error
	SetPageTokens(next, prev resource.Cursor) error
	SetResponseHeaders(headers metadata.MD) error

	// Close ensures that all data gets written to encoder's stream.
	Close() error
}

type ResponseEncoderFormat string

const (
	// JsonEncoderFormat for streaming requests
	JsonEncoderFormat ResponseEncoderFormat = "json"

	// TableEncoderFormat for non-streaming requests
	TableEncoderFormat ResponseEncoderFormat = "table"
)

func NewResponseEncoder(writer io.Writer, format ResponseEncoderFormat) (ResponseEncoder, error) {
	switch format {
	case JsonEncoderFormat:
		return &encoderTypeCheckingWrapper{
			impl: &jsonResponseEncoder{
				writer: writer,
			},
		}, nil
	case TableEncoderFormat:
		return &encoderTypeCheckingWrapper{
			impl: &tableResponseEncoder{
				writer: writer,
				table:  tablewriter.NewWriter(writer),
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown or unsupported encoder format: %s", format)
	}
}

type encoderTypeCheckingWrapper struct {
	impl ResponseEncoder
}

func (etcw *encoderTypeCheckingWrapper) SetColumns(fieldPaths []object.FieldPath, displayNames []string) {
	if len(displayNames) != len(fieldPaths) {
		panic(fmt.Errorf("field paths count %d different from display names count %d", len(fieldPaths), len(displayNames)))
	}
	etcw.impl.SetColumns(fieldPaths, displayNames)
}

func (ectw *encoderTypeCheckingWrapper) Add(msg proto.Message) error {
	return ectw.impl.Add(msg)
}

func (ectw *encoderTypeCheckingWrapper) SetPageTokens(next, prev resource.Cursor) error {
	return ectw.impl.SetPageTokens(next, prev)
}

func (ectw *encoderTypeCheckingWrapper) SetResponseHeaders(headers metadata.MD) error {
	return ectw.impl.SetResponseHeaders(headers)
}

func (ectw *encoderTypeCheckingWrapper) Close() error {
	return ectw.impl.Close()
}

type jsonResponseEncoder struct {
	paths  []object.FieldPath
	writer io.Writer
}

var _ ResponseEncoder = &jsonResponseEncoder{}

func (jre *jsonResponseEncoder) SetColumns(fieldPaths []object.FieldPath, _ []string) {
	jre.paths = fieldPaths
}

func (jre *jsonResponseEncoder) Add(msg proto.Message) error {
	out, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = jre.writer.Write(append(out, '\n'))
	return err
}

func (jre *jsonResponseEncoder) SetPageTokens(next, prev resource.Cursor) error {
	tokens := map[string]string{}
	if next != nil && !reflect.ValueOf(next).IsNil() {
		tokens["nextPageToken"] = next.String()
	}
	if prev != nil && !reflect.ValueOf(prev).IsNil() {
		tokens["prevPageToken"] = prev.String()
	}
	if len(tokens) == 0 {
		return nil
	}
	out, err := json.Marshal(tokens)
	if err != nil {
		return err
	}
	_, err = jre.writer.Write(append(out, '\n'))
	return err
}

func (jre *jsonResponseEncoder) SetResponseHeaders(headers metadata.MD) error {
	if len(headers) > 0 {
		out, err := json.Marshal(headers)
		if err != nil {
			return err
		}
		_, err = jre.writer.Write(append(out, '\n'))
		return err
	} else {
		return nil
	}
}

func (jre *jsonResponseEncoder) Close() error {
	return nil
}

type tableResponseEncoder struct {
	// Note: tablewriter.Table does no checking for I/O errors
	writer                       io.Writer
	paths                        []object.FieldPath
	table                        *tablewriter.Table
	nextPageToken, prevPageToken resource.Cursor
	responseHeaders              metadata.MD
	clippedColumnsCount          int
}

func (tfe *tableResponseEncoder) SetColumns(fieldPaths []object.FieldPath, displayNames []string) {
	tfe.paths = fieldPaths
	tfe.table.SetHeader(displayNames)
}

func (tfe *tableResponseEncoder) Add(msg proto.Message) error {
	stringFields := make([]string, 0, len(tfe.paths))

	for _, fieldPath := range tfe.paths {
		rawValue, ok := fieldPath.GetSingleRaw(msg)
		if !ok {
			stringFields = append(stringFields, "")
		} else {
			if m, ok := rawValue.(proto.Message); ok {
				data, err := protojson.Marshal(m)
				if err != nil {
					return err
				}
				stringFields = append(stringFields, string(data))
			} else if stringer, ok := rawValue.(fmt.Stringer); ok {
				stringFields = append(stringFields, stringer.String())
			} else {
				stringFields = append(stringFields, fmt.Sprint(rawValue))
			}
		}
	}
	for i, strValue := range stringFields {
		if len(strValue) >= MaxTableViewColumnWidth {
			tfe.clippedColumnsCount++
			stringFields[i] = fmt.Sprintf("%s...", strValue[:MaxTableViewColumnWidth])
		}
	}

	tfe.table.Append(stringFields)
	return nil
}

func (tfe *tableResponseEncoder) SetPageTokens(next, prev resource.Cursor) error {
	tfe.prevPageToken = prev
	tfe.nextPageToken = next
	return nil
}

func (tfe *tableResponseEncoder) SetResponseHeaders(headers metadata.MD) error {
	tfe.responseHeaders = headers
	return nil
}

func (tfe *tableResponseEncoder) Close() error {
	if tfe.clippedColumnsCount > 0 {
		_, _ = os.Stdout.WriteString(fmt.Sprintf(
			"WARNING: %d column values were clipped due to oversize." +
				" In order to display values in full and more readable format, use \"-o json\" option\n\n",
				tfe.clippedColumnsCount))
	}
	if len(tfe.responseHeaders) > 0 {
		for k, vals := range tfe.responseHeaders {
			_, err := tfe.writer.Write([]byte(fmt.Sprintf("%s: %s\n", k, strings.Join(vals, ", "))))
			if err != nil {
				return err
			}
		}
		_, _ = tfe.writer.Write([]byte("\n"))
	}
	// Special case: if response has no fields, then don't print anything
	if len(tfe.paths) != 0 {
		tfe.table.Render()
	}
	if tfe.nextPageToken != nil && !reflect.ValueOf(tfe.nextPageToken).IsNil() {
		_, err := tfe.writer.Write([]byte(fmt.Sprintf("NextPageToken: %s\n", tfe.nextPageToken)))
		if err != nil {
			return err
		}
	}
	if tfe.prevPageToken != nil && !reflect.ValueOf(tfe.prevPageToken).IsNil() {
		_, err := tfe.writer.Write([]byte(fmt.Sprintf("PrevPageToken: %s\n", tfe.prevPageToken)))
		if err != nil {
			return err
		}
	}
	return nil
}
