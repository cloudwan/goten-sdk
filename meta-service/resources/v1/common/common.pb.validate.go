// Code generated by protoc-gen-goten-validate
// File: goten/meta-service/proto/v1/common.proto
// DO NOT EDIT!!!

package common

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
import ()

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
var ()

func (obj *LabelledDomain) GotenValidate() error {
	if obj == nil {
		return nil
	}
	{
		rlen := utf8.RuneCountInString(obj.Label)
		if rlen > 64 {
			return gotenvalidate.NewValidationError("LabelledDomain", "label", obj.Label, "field must contain at most 64 characters", nil)
		}
	}
	if obj.Label == "" {
		return gotenvalidate.NewValidationError("LabelledDomain", "label", obj.Label, "field is required", nil)
	}
	{
		rlen := utf8.RuneCountInString(obj.Domain)
		if rlen > 255 {
			return gotenvalidate.NewValidationError("LabelledDomain", "domain", obj.Domain, "field must contain at most 255 characters", nil)
		}
	}
	if err := gotenvalidate.ValidateAddress(string(obj.Domain)); err != nil {
		return gotenvalidate.NewValidationError("LabelledDomain", "domain", obj.Domain, "field must contain a valid address", nil)
	}
	if obj.Domain == "" {
		return gotenvalidate.NewValidationError("LabelledDomain", "domain", obj.Domain, "field is required", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
