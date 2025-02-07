// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/parameters.proto

package transformation

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Parameters) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Parameters)
	if !ok {
		that2, ok := that.(Parameters)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for k, v := range m.GetHeaders() {

		if strings.Compare(v, target.GetHeaders()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetPath()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPath()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPath(), target.GetPath()) {
			return false
		}
	}

	return true
}
