package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:     "setup",
	Short:   "Explain that setup is no longer required",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "other commands",
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	setupCmd.Flags().Bool("init", false, "Report an error with a hint to use but init instead")
	rootCmd.AddCommand(setupCmd)
}
