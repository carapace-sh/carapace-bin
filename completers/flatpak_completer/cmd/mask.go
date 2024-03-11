package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var maskCmd = &cobra.Command{
	Use:     "mask",
	Short:   "disable updates and automatic installation matching patterns",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maskCmd).Standalone()

	maskCmd.Flags().BoolP("help", "h", false, "Show help options")
	maskCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	maskCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	maskCmd.Flags().Bool("remove", false, "Remove matching masks")
	maskCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	maskCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	maskCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(maskCmd)

	// TODO flag
}
