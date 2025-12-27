package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagePolicyCmd = &cobra.Command{
	Use:   "image-policy",
	Short: "Analyze image policy string",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagePolicyCmd).Standalone()

	rootCmd.AddCommand(imagePolicyCmd)
}
