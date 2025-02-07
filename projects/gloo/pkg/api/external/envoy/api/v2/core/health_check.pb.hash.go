// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/api/v2/core/health_check.proto

package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *HealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Timeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Timeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Interval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Interval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetInitialJitter()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("InitialJitter")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetInitialJitter(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("InitialJitter")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIntervalJitter()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IntervalJitter")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIntervalJitter(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IntervalJitter")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetIntervalJitterPercent())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUnhealthyThreshold()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UnhealthyThreshold")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUnhealthyThreshold(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UnhealthyThreshold")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHealthyThreshold()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HealthyThreshold")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHealthyThreshold(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HealthyThreshold")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetReuseConnection()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ReuseConnection")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetReuseConnection(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ReuseConnection")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetNoTrafficInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("NoTrafficInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetNoTrafficInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("NoTrafficInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUnhealthyInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UnhealthyInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUnhealthyInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UnhealthyInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUnhealthyEdgeInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UnhealthyEdgeInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUnhealthyEdgeInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UnhealthyEdgeInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHealthyEdgeInterval()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HealthyEdgeInterval")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHealthyEdgeInterval(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HealthyEdgeInterval")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetEventLogPath())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAlwaysLogHealthCheckFailures())
	if err != nil {
		return 0, err
	}

	switch m.HealthChecker.(type) {

	case *HealthCheck_HttpHealthCheck_:

		if h, ok := interface{}(m.GetHttpHealthCheck()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HttpHealthCheck")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHttpHealthCheck(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HttpHealthCheck")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HealthCheck_TcpHealthCheck_:

		if h, ok := interface{}(m.GetTcpHealthCheck()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TcpHealthCheck")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTcpHealthCheck(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TcpHealthCheck")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HealthCheck_GrpcHealthCheck_:

		if h, ok := interface{}(m.GetGrpcHealthCheck()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GrpcHealthCheck")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGrpcHealthCheck(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GrpcHealthCheck")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HealthCheck_CustomHealthCheck_:

		if h, ok := interface{}(m.GetCustomHealthCheck()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("CustomHealthCheck")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCustomHealthCheck(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("CustomHealthCheck")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HealthCheck_Payload) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck_Payload")); err != nil {
		return 0, err
	}

	switch m.Payload.(type) {

	case *HealthCheck_Payload_Text:

		if _, err = hasher.Write([]byte(m.GetText())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HealthCheck_HttpHealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck_HttpHealthCheck")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHost())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetRequestHeadersToAdd() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetRequestHeadersToRemove() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUseHttp2())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetExpectedStatuses() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetResponseAssertions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseAssertions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseAssertions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseAssertions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HealthCheck_TcpHealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck_TcpHealthCheck")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSend()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Send")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSend(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Send")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetReceive() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HealthCheck_RedisHealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck_RedisHealthCheck")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HealthCheck_GrpcHealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck_GrpcHealthCheck")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAuthority())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HealthCheck_CustomHealthCheck) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.api.v2.core.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core.HealthCheck_CustomHealthCheck")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	switch m.ConfigType.(type) {

	case *HealthCheck_CustomHealthCheck_Config:

		if h, ok := interface{}(m.GetConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Config")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Config")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HealthCheck_CustomHealthCheck_TypedConfig:

		if h, ok := interface{}(m.GetTypedConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TypedConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTypedConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TypedConfig")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}
