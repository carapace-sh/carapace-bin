package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:     "kill [OPTION…] INSTANCE",
	Short:   "Stop a running application",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	killCmd.Flags().BoolP("help", "h", false, "Show help options")
	killCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	killCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(killCmd)

	// TODO positional
}
