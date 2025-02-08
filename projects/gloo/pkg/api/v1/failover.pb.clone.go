// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/failover.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"
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
func (m *Failover) Clone() proto.Message {
	var target *Failover
	if m == nil {
		return target
	}
	target = &Failover{}

	if m.GetPrioritizedLocalities() != nil {
		target.PrioritizedLocalities = make([]*Failover_PrioritizedLocality, len(m.GetPrioritizedLocalities()))
		for idx, v := range m.GetPrioritizedLocalities() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.PrioritizedLocalities[idx] = h.Clone().(*Failover_PrioritizedLocality)
			} else {
				target.PrioritizedLocalities[idx] = proto.Clone(v).(*Failover_PrioritizedLocality)
			}

		}
	}

	if h, ok := interface{}(m.GetPolicy()).(clone.Cloner); ok {
		target.Policy = h.Clone().(*Failover_Policy)
	} else {
		target.Policy = proto.Clone(m.GetPolicy()).(*Failover_Policy)
	}

	return target
}

// Clone function
func (m *LocalityLbEndpoints) Clone() proto.Message {
	var target *LocalityLbEndpoints
	if m == nil {
		return target
	}
	target = &LocalityLbEndpoints{}

	if h, ok := interface{}(m.GetLocality()).(clone.Cloner); ok {
		target.Locality = h.Clone().(*Locality)
	} else {
		target.Locality = proto.Clone(m.GetLocality()).(*Locality)
	}

	if m.GetLbEndpoints() != nil {
		target.LbEndpoints = make([]*LbEndpoint, len(m.GetLbEndpoints()))
		for idx, v := range m.GetLbEndpoints() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.LbEndpoints[idx] = h.Clone().(*LbEndpoint)
			} else {
				target.LbEndpoints[idx] = proto.Clone(v).(*LbEndpoint)
			}

		}
	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(clone.Cloner); ok {
		target.LoadBalancingWeight = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.LoadBalancingWeight = proto.Clone(m.GetLoadBalancingWeight()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *LbEndpoint) Clone() proto.Message {
	var target *LbEndpoint
	if m == nil {
		return target
	}
	target = &LbEndpoint{}

	target.Address = m.GetAddress()

	target.Port = m.GetPort()

	if h, ok := interface{}(m.GetHealthCheckConfig()).(clone.Cloner); ok {
		target.HealthCheckConfig = h.Clone().(*LbEndpoint_HealthCheckConfig)
	} else {
		target.HealthCheckConfig = proto.Clone(m.GetHealthCheckConfig()).(*LbEndpoint_HealthCheckConfig)
	}

	if h, ok := interface{}(m.GetUpstreamSslConfig()).(clone.Cloner); ok {
		target.UpstreamSslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.UpstreamSslConfig)
	} else {
		target.UpstreamSslConfig = proto.Clone(m.GetUpstreamSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.UpstreamSslConfig)
	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(clone.Cloner); ok {
		target.LoadBalancingWeight = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.LoadBalancingWeight = proto.Clone(m.GetLoadBalancingWeight()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *Locality) Clone() proto.Message {
	var target *Locality
	if m == nil {
		return target
	}
	target = &Locality{}

	target.Region = m.GetRegion()

	target.Zone = m.GetZone()

	target.SubZone = m.GetSubZone()

	return target
}

// Clone function
func (m *Failover_PrioritizedLocality) Clone() proto.Message {
	var target *Failover_PrioritizedLocality
	if m == nil {
		return target
	}
	target = &Failover_PrioritizedLocality{}

	if m.GetLocalityEndpoints() != nil {
		target.LocalityEndpoints = make([]*LocalityLbEndpoints, len(m.GetLocalityEndpoints()))
		for idx, v := range m.GetLocalityEndpoints() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.LocalityEndpoints[idx] = h.Clone().(*LocalityLbEndpoints)
			} else {
				target.LocalityEndpoints[idx] = proto.Clone(v).(*LocalityLbEndpoints)
			}

		}
	}

	return target
}

// Clone function
func (m *Failover_Policy) Clone() proto.Message {
	var target *Failover_Policy
	if m == nil {
		return target
	}
	target = &Failover_Policy{}

	if h, ok := interface{}(m.GetOverprovisioningFactor()).(clone.Cloner); ok {
		target.OverprovisioningFactor = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.OverprovisioningFactor = proto.Clone(m.GetOverprovisioningFactor()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *LbEndpoint_HealthCheckConfig) Clone() proto.Message {
	var target *LbEndpoint_HealthCheckConfig
	if m == nil {
		return target
	}
	target = &LbEndpoint_HealthCheckConfig{}

	target.PortValue = m.GetPortValue()

	target.Hostname = m.GetHostname()

	target.Path = m.GetPath()

	target.Method = m.GetMethod()

	return target
}
