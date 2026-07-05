package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Create or update a local mirror of distfiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mirrorCmd).Standalone()
	mirrorCmd.Flags().Bool("new", false, "Only mirror distfiles not already present")
	rootCmd.AddCommand(mirrorCmd)
}
