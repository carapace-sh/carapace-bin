// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List config files currently in use, from lowest to highest precedence",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_lsCmd).Standalone()

	config_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	config_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	config_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	config_lsCmd.Flags().Bool("no-headers", false, "Do not print the table header")
	config_lsCmd.Flags().Bool("no-truncate", false, "Truncate long terminal output to fit the available width")
	config_lsCmd.Flags().Bool("tracked-configs", false, "List every config file mise has loaded, in any project")
	config_lsCmd.Flags().Bool("truncate", true, "Truncate long terminal output to fit the available width")
	config_lsCmd.Flag("no-headers").Hidden = true
	config_lsCmd.Flag("no-truncate").Hidden = true
	configCmd.AddCommand(config_lsCmd)
}
