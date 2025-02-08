package ratelimit

import (
	editRouteOptions "github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/edit/route/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

func RateLimitConfig(opts *editRouteOptions.RouteEditInput, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {

	cmd := &cobra.Command{
		// Use command constants to aid with replacement.
		Use:     constants.CONFIG_RATELIMIT_COMMAND.Use,
		Aliases: constants.CONFIG_RATELIMIT_COMMAND.Aliases,
		Short:   "Configure rate-limits (Enterprise)",
		Long:    "Configure rate-limits for requests that match this route. This is a Gloo Enterprise feature.",
	}

	cliutils.ApplyOptions(cmd, optionsFunc)

	cmd.AddCommand(RateLimitCustomConfig(opts))
	return cmd
}
