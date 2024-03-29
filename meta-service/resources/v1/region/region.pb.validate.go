// Code generated by protoc-gen-goten-validate
// File: goten/meta-service/proto/v1/region.proto
// DO NOT EDIT!!!

package region

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Errorf
	_ = net.ParseIP
	_ = regexp.Match
	_ = strings.Split
	_ = time.Now
	_ = utf8.RuneCountInString
	_ = url.Parse
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &meta.Meta{}
)

func (obj *Region) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Region", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	{
		rlen := utf8.RuneCountInString(obj.Title)
		if rlen > 64 {
			return gotenvalidate.NewValidationError("Region", "title", obj.Title, "field must contain at most 64 characters", nil)
		}
	}
	if obj.Title == "" {
		return gotenvalidate.NewValidationError("Region", "title", obj.Title, "field is required", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
