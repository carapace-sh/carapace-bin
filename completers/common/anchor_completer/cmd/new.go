package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().String("anchor-version", "v1", "Anchor template version to generate")
	newCmd.Flags().Bool("force", false, "Create new program even if there is already one")
	newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	newCmd.Flags().StringP("template", "t", "multiple", "Rust program template to use")
	rootCmd.AddCommand(newCmd)
}
