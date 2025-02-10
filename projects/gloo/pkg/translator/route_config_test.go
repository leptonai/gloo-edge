package translator_test

import (
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"

	envoy_header_to_metadata_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/header_to_metadata/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Route Configs", func() {

	DescribeTable("validate route path", func(path string, expectedValue bool) {
		if expectedValue {
			Expect(translator.ValidateRoutePath(path)).ToNot(HaveOccurred())
		} else {
			Expect(translator.ValidateRoutePath(path)).To(HaveOccurred())
		}
	},
		Entry("Hex", "%af", true),
		Entry("Hex Camel", "%Af", true),
		Entry("Hex num", "%00", true),
		Entry("Hex double", "%11", true),
		Entry("Hex with valid", "%af801&*", true),
		Entry("valid with hex", "801&*%af", true),
		Entry("valid with hex and valid", "801&*%af719$@!", true),
		Entry("Hex single", "%0", false),
		Entry("unicode chars", "ƒ©", false),
		Entry("unicode chars", "¥¨˚∫", false),
		Entry("//", "hello/something//", false),
		Entry("/./", "hello/something/./", false),
		Entry("/../", "hello/something/../", false),
		Entry("hex slash upper", "hello/something%2F", false),
		Entry("hex slash lower", "hello/something%2f", false),
		Entry("hash", "hello/something#", false),
		Entry("/..", "hello/../something", false),
		Entry("/.", "hello/./something", false),
	)

	It("Should validate all seperate characters", func() {
		// must allow all "pchar" characters = unreserved / pct-encoded / sub-delims / ":" / "@"
		// https://www.rfc-editor.org/rfc/rfc3986
		// unreserved
		// alpha Upper and Lower
		for i := 'a'; i <= 'z'; i++ {
			Expect(translator.ValidateRoutePath(string(i))).ToNot(HaveOccurred())
			Expect(translator.ValidateRoutePath(strings.ToUpper(string(i)))).ToNot(HaveOccurred())
		}
		// digit
		for i := 0; i < 10; i++ {
			Expect(translator.ValidateRoutePath(strconv.Itoa(i))).ToNot(HaveOccurred())
		}
		unreservedChars := "-._~"
		for _, c := range unreservedChars {
			Expect(translator.ValidateRoutePath(string(c))).ToNot(HaveOccurred())
		}
		// sub-delims
		subDelims := "!$&'()*+,;="
		Expect(len(subDelims)).To(Equal(11))
		for _, c := range subDelims {
			Expect(translator.ValidateRoutePath(string(c))).ToNot(HaveOccurred())
		}
		// pchar
		pchar := ":@"
		for _, c := range pchar {
			Expect(translator.ValidateRoutePath(string(c))).ToNot(HaveOccurred())
		}
		// invalid characters
		invalid := "<>?\\|[]{}\"^%#"
		for _, c := range invalid {
			Expect(translator.ValidateRoutePath(string(c))).To(HaveOccurred())
		}
	})

	DescribeTable("path rewrites", func(s string, pass bool) {
		err := translator.ValidatePrefixRewrite(s)
		if pass {
			Expect(err).ToNot(HaveOccurred())
		} else {
			Expect(err).To(HaveOccurred())
		}
	},
		Entry("allow query parameters", "some/site?a=data&b=location&c=searchterm", true),
		Entry("allow fragments", "some/site#framgentedinfo", true),
		Entry("invalid", "some/site<hello", false),
		Entry("invalid", "some/site{hello", false),
		Entry("invalid", "some/site}hello", false),
		Entry("invalid", "some/site[hello", false),
	)
})

func Test_addMetadataToHeaderMapping(t *testing.T) {
	tests := []struct {
		literalConfig               map[string]string
		endpointMetadataKeyToHeader map[string]string
		expect                      map[string]string
	}{
		{
			// simple literal config case, no dynamic loading expected
			literalConfig: map[string]string{
				"version": "2",
			},
			endpointMetadataKeyToHeader: map[string]string{},
			expect:                      map[string]string{},
		},
		{
			// simple dynamic config case, no dynamic loading expected
			literalConfig: map[string]string{
				"version": "$FromHeader:x-version",
			},
			endpointMetadataKeyToHeader: map[string]string{},
			expect:                      map[string]string{"version": "x-version"},
		},
		{
			// append behavior
			literalConfig: map[string]string{
				"version": "$FromHeader: x-version",
			},
			endpointMetadataKeyToHeader: map[string]string{"a": "b"},
			expect:                      map[string]string{"a": "b", "version": "x-version"},
		},
		{
			// forbit empty header name
			literalConfig: map[string]string{
				"version": "$FromHeader:",
			},
			endpointMetadataKeyToHeader: map[string]string{"a": "b"},
			expect:                      map[string]string{"a": "b"},
		},
	}
	for _, tt := range tests {
		translator.AddMetadataToHeaderMapping(tt.endpointMetadataKeyToHeader, tt.literalConfig)
		assert.Equal(t, tt.endpointMetadataKeyToHeader, tt.expect)
	}
}

func subset(m map[string]string) *v1.Subset {
	return &v1.Subset{
		Values: m,
	}
}

func sort(t *testing.T, c *envoy_header_to_metadata_v3.Config) {
	if c == nil {
		return
	}
	slices.SortFunc(c.RequestRules, func(a, b *envoy_header_to_metadata_v3.Config_Rule) int {
		if a == nil {
			t.Fatal("should not have nil rule")
		}
		if a.Header < b.Header {
			return -1
		}
		return 1
	})
}

func TestNewMetadataToHeaderConfig(t *testing.T) {
	tests := []struct {
		action *v1.RouteAction
		want   *envoy_header_to_metadata_v3.Config
	}{
		{
			action: nil,
			want:   nil,
		},
		{
			action: &v1.RouteAction{Destination: &v1.RouteAction_Single{
				Single: &v1.Destination{Subset: subset(map[string]string{"version": "2"})},
			}},
			want: nil,
		},
		{
			action: &v1.RouteAction{Destination: &v1.RouteAction_Single{
				Single: &v1.Destination{Subset: subset(map[string]string{"version": "$FromHeader:x-version", "a": "$FromHeader:b"})},
			}},
			want: &envoy_header_to_metadata_v3.Config{RequestRules: []*envoy_header_to_metadata_v3.Config_Rule{
				{Header: "x-version", Remove: false, OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb", Type: 0, Key: "version",
				}},
				{Header: "b", Remove: false, OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb", Type: 0, Key: "a",
				}},
			}},
		},
		{
			action: &v1.RouteAction{Destination: &v1.RouteAction_Multi{
				Multi: &v1.MultiDestination{
					Destinations: []*v1.WeightedDestination{
						{Destination: &v1.Destination{Subset: subset(map[string]string{"version": "$FromHeader: x-version", "a": "$FromHeader:b"})}},
						{Destination: &v1.Destination{Subset: subset(map[string]string{"c": "$FromHeader:d", "a": "$FromHeader:b"})}},
					},
				},
			}},
			want: &envoy_header_to_metadata_v3.Config{RequestRules: []*envoy_header_to_metadata_v3.Config_Rule{
				{Header: "x-version", Remove: false, OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb", Type: 0, Key: "version",
				}},
				{Header: "b", Remove: false, OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb", Type: 0, Key: "a",
				}},
				{Header: "d", Remove: false, OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb", Type: 0, Key: "c",
				}},
			}},
		},
	}
	for _, tt := range tests {
		got := translator.NewMetadataToHeaderConfig(tt.action)
		sort(t, got)
		sort(t, tt.want)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewMetadataToHeaderConfig() = %v, want %v", got, tt.want)
		}
	}
}
