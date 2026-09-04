package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nugetCmd = &cobra.Command{
	Use:   "nuget",
	Short: "manage NuGet packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nugetCmd).Standalone()
	rootCmd.AddCommand(nugetCmd)
}
