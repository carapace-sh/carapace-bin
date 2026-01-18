package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config [OPTION…] [KEY [VALUE]]",
	Short:   "Manage configuration",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("get", false, "Get configuration for KEY")
	configCmd.Flags().BoolP("help", "h", false, "Show help options")
	configCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	configCmd.Flags().Bool("list", false, "List configuration keys and values")
	configCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	configCmd.Flags().Bool("set", false, "Set configuration for KEY to VALUE")
	configCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	configCmd.Flags().Bool("unset", false, "Unset configuration for KEY")
	configCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	configCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(configCmd)

	// TODO flag

	carapace.Gen(configCmd).PositionalCompletion(
	// TODO keys
	// TODO value when set
	)
}
