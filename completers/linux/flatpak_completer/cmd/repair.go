package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repairCmd = &cobra.Command{
	Use:     "repair [OPTION…]",
	Short:   "Repair a flatpak installation",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repairCmd).Standalone()

	repairCmd.Flags().Bool("dry-run", false, "Don't make any changes")
	repairCmd.Flags().BoolP("help", "h", false, "Show help options")
	repairCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	repairCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	repairCmd.Flags().Bool("reinstall-all", false, "Reinstall all refs")
	repairCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	repairCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	repairCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(repairCmd)

	// TODO flag
}
