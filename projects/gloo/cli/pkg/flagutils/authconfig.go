package flagutils

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/spf13/pflag"
)

func AddAuthConfigFlags(set *pflag.FlagSet, ac *options.InputAuthConfig) {
	addVirtualServiceFlagsOIDC(set, &ac.OIDCAuth)
	addVirtualServiceFlagsApiKey(set, &ac.ApiKeyAuth)
	addVirtualServiceFlagsOpa(set, &ac.OpaAuth)
}

func addVirtualServiceFlagsOIDC(set *pflag.FlagSet, oidc *options.OIDCAuth) {
	// TODO: add support for authorization when it is supported for ratelimit
	//set.StringVar(&Options.RateLimits.AuthorizedHeader, "rate-limit-authorize-header", "", "header name used to authorize requests")
	set.BoolVar(&oidc.Enable, "enable-oidc-auth", false, "enable oidc auth features for this virtual service")
	set.StringVar(&oidc.ClientId, "oidc-auth-client-id", "", "client id as registered with id provider")
	set.StringVar(&oidc.ClientSecretRef.Name, "oidc-auth-client-secret-name", "", "name of the 'client secret' secret")
	set.StringVar(&oidc.ClientSecretRef.Namespace, "oidc-auth-client-secret-namespace", "", "namespace of the 'client secret' secret")
	set.StringVar(&oidc.IssuerUrl, "oidc-auth-issuer-url", "", "the url of the issuer")
	set.StringToStringVar(&oidc.AuthEndpointQueryParams, "auth-endpoint-query-params", nil, "additional static query parameters to include in authorization request to identity provider")
	set.StringVar(&oidc.AppUrl, "oidc-auth-app-url", "", "the public url of your app")
	set.StringVar(&oidc.CallbackPath, "oidc-auth-callback-path", "/oidc-gloo-callback", "the callback path. relative to the app url.")
	set.StringSliceVar(&oidc.Scopes, "oidc-scope", nil, "scopes to request in addition to 'openid'. optional.")

}

func addVirtualServiceFlagsApiKey(set *pflag.FlagSet, apiKey *options.ApiKeyAuth) {
	// TODO: add support for authorization when it is supported for ratelimit
	//set.StringVar(&Options.RateLimits.AuthorizedHeader, "rate-limit-authorize-header", "", "header name used to authorize requests")
	set.BoolVar(&apiKey.Enable, "enable-apikey-auth", false, "enable apikey auth features for this virtual service")
	set.StringSliceVar(&apiKey.Labels, "apikey-label-selector", []string{}, "apikey label selector to identify valid apikeys for this virtual service; a comma-separated list of labels (key=value)")
	set.StringVar(&apiKey.SecretNamespace, "apikey-secret-namespace", "", "namespace to search for an individual apikey secret")
	set.StringVar(&apiKey.SecretName, "apikey-secret-name", "", "name to search for in provided namespace for an individual apikey secret")
}

func addVirtualServiceFlagsOpa(set *pflag.FlagSet, opa *options.OpaAuth) {
	// TODO: add support for authorization when it is supported for ratelimit
	//set.StringVar(&Options.RateLimits.AuthorizedHeader, "rate-limit-authorize-header", "", "header name used to authorize requests")
	set.BoolVar(&opa.Enable, "enable-opa-auth", false, "enable opa auth features for this virtual service")
	set.StringVar(&opa.Query, "opa-query", "", "The OPA query to evaluate on a request")
	set.StringSliceVar(&opa.Modules, "opa-module-ref", []string{}, "namespace.name references to a config map containing OPA modules")
}
