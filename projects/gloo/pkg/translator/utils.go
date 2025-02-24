package translator

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	errors "github.com/rotisserie/eris"

	envoyal "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// returns the name of the cluster created for a given upstream
func UpstreamToClusterName(upstream *core.ResourceRef) string {

	// For non-namespaced resources, return only name
	if upstream.GetNamespace() == "" {
		return upstream.GetName()
	}

	// Don't use dots in the name as it messes up prometheus stats
	return fmt.Sprintf("%s_%s", upstream.GetName(), upstream.GetNamespace())
}

// returns the ref of the upstream for a given cluster
func ClusterToUpstreamRef(cluster string) (*core.ResourceRef, error) {

	split := strings.Split(cluster, "_")
	if len(split) > 2 || len(split) < 1 {
		return nil, errors.Errorf("unable to convert cluster %s back to upstream ref", cluster)
	}

	ref := &core.ResourceRef{
		Name: split[0],
	}

	if len(split) == 2 {
		ref.Namespace = split[1]
	}
	return ref, nil
}

func NewFilterWithTypedConfig(name string, config proto.Message) (*envoy_config_listener_v3.Filter, error) {

	s := &envoy_config_listener_v3.Filter{
		Name: name,
	}

	if config != nil {
		marshalledConf, err := utils.MessageToAny(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return &envoy_config_listener_v3.Filter{}, err
		}

		s.ConfigType = &envoy_config_listener_v3.Filter_TypedConfig{
			TypedConfig: marshalledConf,
		}
	}

	return s, nil
}

func NewAccessLogWithConfig(name string, config proto.Message) (envoyal.AccessLog, error) {
	s := envoyal.AccessLog{
		Name: name,
	}

	if config != nil {
		marshalledConf, err := utils.MessageToAny(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return envoyal.AccessLog{}, err
		}

		s.ConfigType = &envoyal.AccessLog_TypedConfig{
			TypedConfig: marshalledConf,
		}
	}

	return s, nil
}

func ParseTypedConfig(c typedConfigObject, config proto.Message) error {
	any := c.GetTypedConfig()
	if any != nil {
		return ptypes.UnmarshalAny(any, config)
	}
	return nil
}

type typedConfigObject interface {
	GetTypedConfig() *any.Any
}

// IsIpv4Address returns whether
// the provided address is valid IPv4, is pure(unmapped) IPv4, and if there was an error in the bindaddr
// This is used to distinguish between IPv4 and IPv6 addresses
func IsIpv4Address(bindAddress string) (validIpv4, strictIPv4 bool, err error) {
	bindIP := net.ParseIP(bindAddress)
	if bindIP == nil {
		// If bindAddress is not a valid textual representation of an IP address
		return false, false, errors.Errorf("bindAddress %s is not a valid IP address", bindAddress)

	} else if bindIP.To4() == nil {
		// If bindIP is not an IPv4 address, To4 returns nil.
		// so this is not an acceptable ipv4
		return false, false, nil
	}
	return true, isPureIPv4Address(bindAddress), nil

}

// isPureIPv4Address checks the string to see if it is
// ipv4 and not ipv4 mapped into ipv6 space and not ipv6.
// Used as the standard net.Parse smashes everything to ipv6.
// Basically false if ::ffff:0.0.0.0 and true if 0.0.0.0
func isPureIPv4Address(ipString string) bool {
	for i := 0; i < len(ipString); i++ {
		switch ipString[i] {
		case '.':
			return true
		case ':':
			return false
		}
	}
	return false
}

var (
	Logger *LeptonLogger
)

func init() {
	Logger = CreateLogger(DefaultLoggerConfig())
}

func DefaultLoggerConfig() *zap.Config {
	c := zap.NewProductionConfig()
	c.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	return &c
}

func CreateLogger(config *zap.Config) *LeptonLogger {
	if config == nil {
		config = DefaultLoggerConfig()
	}

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	return &LeptonLogger{
		SugaredLogger: l.Sugar(),
		zapLogger:     l,
	}
}

type LeptonLogger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
}

func (l *LeptonLogger) ZapLogger() *zap.Logger {
	return l.zapLogger
}

// Override the default logger's Errorw func to down level context canceled error
func (l *LeptonLogger) Errorw(msg string, keysAndValues ...interface{}) {
	for i := 0; i < len(keysAndValues); i += 2 {
		if keysAndValues[i] != "error" {
			continue
		}
		if err, ok := keysAndValues[i+1].(error); ok {
			if strings.Contains(err.Error(), context.Canceled.Error()) {
				l.SugaredLogger.Warnw(msg, keysAndValues...)
				return
			}
		}
	}

	l.SugaredLogger.Errorw(msg, keysAndValues...)
}
