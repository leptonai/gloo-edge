// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxylatency/proxylatency.proto

package proxylatency

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *ProxyLatency) Clone() proto.Message {
	var target *ProxyLatency
	if m == nil {
		return target
	}
	target = &ProxyLatency{}

	target.Request = m.GetRequest()

	target.MeasureRequestInternally = m.GetMeasureRequestInternally()

	target.Response = m.GetResponse()

	if h, ok := interface{}(m.GetChargeClusterStat()).(clone.Cloner); ok {
		target.ChargeClusterStat = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.ChargeClusterStat = proto.Clone(m.GetChargeClusterStat()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetChargeListenerStat()).(clone.Cloner); ok {
		target.ChargeListenerStat = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.ChargeListenerStat = proto.Clone(m.GetChargeListenerStat()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetEmitDynamicMetadata()).(clone.Cloner); ok {
		target.EmitDynamicMetadata = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.EmitDynamicMetadata = proto.Clone(m.GetEmitDynamicMetadata()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}
