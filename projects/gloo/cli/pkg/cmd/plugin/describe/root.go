package describe

import (
	"fmt"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/plugin/common"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.PLUGIN_DESCRIBE_COMMAND.Use,
		Short: constants.PLUGIN_DESCRIBE_COMMAND.Short,
		RunE: func(cmd *cobra.Command, args []string) error {
			var pluginToDescribe string
			if len(args) > 0 {
				pluginToDescribe = args[0]
			}

			if pluginToDescribe == "" {
				return eris.New("Please specify a plugin to describe.")
			}

			// Print whether the plugin is installed and which versions are available for install.
			genius, err := common.NewPluginGenius(opts.Top.Ctx, cmd, constants.GlooctlPluginBucket, "glooctl-")
			if err != nil {
				return err
			}

			plugin, err := genius.Describe(opts.Top.Ctx, pluginToDescribe)
			if err != nil {
				return err
			}

			fmt.Printf("Plugin: %s\n", pluginToDescribe)
			if plugin.InstalledError != nil {
				fmt.Printf("%s\n", plugin.InstalledError.Error())
			} else {
				fmt.Printf("Install path: %s\n", plugin.InstalledPlugin.FullPath)
				if len(plugin.InstalledPlugin.Warnings) > 0 {
					fmt.Printf("Installation warnings: %s\n", plugin.InstalledPlugin.Warnings)
				}
			}
			if plugin.RemoteError != nil {
				fmt.Printf("Failed to describe remote plugin %s\n", plugin.RemoteError.Error())
			} else {
				fmt.Printf("Available versions: %v\n", plugin.RegistryPlugin.AvailableVersions.ListVersions())
			}

			return nil
		},
	}

	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}
