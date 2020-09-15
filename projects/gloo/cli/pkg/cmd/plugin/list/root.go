package list

import (
	"fmt"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/plugin/common"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/spf13/cobra"
)

func RootCmd(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.PLUGIN_LIST_COMMAND.Use,
		Short: constants.PLUGIN_LIST_COMMAND.Short,
		RunE: func(cmd *cobra.Command, args []string) error {
			list, warnings, err := common.NewInstalledPluginRegistry(cmd).ListAll()
			if err != nil {
				return err
			}

			if len(list) > 0 {
				fmt.Println("The following plugins are installed:")
				for _, plugin := range list {
					fmt.Print(plugin.Name)
					if len(plugin.Warnings) > 0 {
						for _, err := range plugin.Warnings {
							fmt.Printf("  - %s\n", err.Error())
						}
					} else {
						fmt.Println("")
					}
				}
			}

			if len(warnings) > 0 {
				fmt.Println("\nglooctl encountered the following warnings while searching for installed plugins:")
				for _, warning := range warnings {
					fmt.Println(warning.Error())
				}
			}

			return nil
		},
	}
	return cmd
}
