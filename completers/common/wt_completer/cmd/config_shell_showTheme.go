package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_shell_showThemeCmd = &cobra.Command{
	Use:   "show-theme",
	Short: "Show output theme samples",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_shell_showThemeCmd).Standalone()

	config_shell_showThemeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_shellCmd.AddCommand(config_shell_showThemeCmd)
}
