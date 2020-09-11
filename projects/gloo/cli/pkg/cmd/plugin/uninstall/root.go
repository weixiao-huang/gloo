package uninstall

import (
	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/plugin/common"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

func RootCmd(_ *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.PLUGIN_UNINSTALL_COMMAND.Use,
		Short: constants.PLUGIN_UNINSTALL_COMMAND.Short,
		RunE: func(cmd *cobra.Command, args []string) error {
			var pluginToUninstall string
			if len(args) > 0 {
				pluginToUninstall = args[0]
			}

			if pluginToUninstall == "" {
				return eris.New("Please specify a plugin to uninstall.")
			}

			return common.NewInstalledPluginRegistry(cmd).Uninstall(pluginToUninstall)
		},
	}

	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}
