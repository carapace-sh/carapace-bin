// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Manage config files",
	Aliases: []string{"cfg", "toml"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("help", "h", false, "Print help")
	configCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	configCmd.Flags().Bool("no-header", false, "Do not print the table header")
	configCmd.Flags().Bool("no-headers", false, "Do not print the table header")
	configCmd.Flags().Bool("no-truncate", false, "Truncate long terminal output to fit the available width")
	configCmd.Flags().Bool("tracked-configs", false, "List every config file mise has loaded, in any project")
	configCmd.Flags().Bool("truncate", true, "Truncate long terminal output to fit the available width")
	configCmd.Flag("no-headers").Hidden = true
	configCmd.Flag("no-truncate").Hidden = true
	rootCmd.AddCommand(configCmd)
}
