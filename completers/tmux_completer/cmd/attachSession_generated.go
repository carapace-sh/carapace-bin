package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var attachSessionCmd = &cobra.Command{
	Use:   "attach-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachSessionCmd).Standalone()

	attachSessionCmd.Flags().BoolS("E", "E", false, "TODO description")
	attachSessionCmd.Flags().StringS("c", "c", "", "working-directory")
	attachSessionCmd.Flags().BoolS("d", "d", false, "TODO description")
	attachSessionCmd.Flags().StringS("f", "f", "", "flags")
	attachSessionCmd.Flags().BoolS("r", "r", false, "TODO description")
	attachSessionCmd.Flags().StringS("t", "t", "", "target-session")
	attachSessionCmd.Flags().BoolS("x", "x", false, "TODO description")
	rootCmd.AddCommand(attachSessionCmd)
}
