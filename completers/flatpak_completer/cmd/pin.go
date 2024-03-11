package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinCmd = &cobra.Command{
	Use:     "pin [OPTION…] [PATTERN…]",
	Short:   "disable automatic removal of runtimes matching patterns",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinCmd).Standalone()

	pinCmd.Flags().BoolP("help", "h", false, "Show help options")
	pinCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	pinCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	pinCmd.Flags().Bool("remove", false, "Remove matching pins")
	pinCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	pinCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	pinCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(pinCmd)

	// TODO flag
}
