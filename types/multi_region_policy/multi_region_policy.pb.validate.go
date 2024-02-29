// Code generated by protoc-gen-goten-validate
// File: goten/types/multi_region_policy.proto
// DO NOT EDIT!!!

package multi_region_policy

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

func (obj *MultiRegionPolicy) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if len(obj.EnabledRegions) < 1 {
		return gotenvalidate.NewValidationError("MultiRegionPolicy", "enabledRegions", obj.EnabledRegions, "field must have at least 1 items", nil)
	}
	if obj.DefaultControlRegion == "" {
		return gotenvalidate.NewValidationError("MultiRegionPolicy", "defaultControlRegion", obj.DefaultControlRegion, "field is required", nil)
	}
	for idx, elem := range obj.CriteriaForDisabledSync {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("MultiRegionPolicy", "criteriaForDisabledSync", obj.CriteriaForDisabledSync[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *MultiRegionPolicy_CriteriaForDisabledSync) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
