package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hclfmtCmd = &cobra.Command{
	Use:     "hclfmt",
	Short:   "Recursively find hcl files and rewrite them into a canonical format",
	GroupID: "terragrunt",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hclfmtCmd).Standalone()

	rootCmd.AddCommand(hclfmtCmd)
}
