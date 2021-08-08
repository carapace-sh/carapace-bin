package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Print package dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()

	depsCmd.Flags().Bool("dev", false, "Include dev dependencies.")
	depsCmd.Flags().Bool("executables", false, "List all available executables.")
	depsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	depsCmd.Flags().Bool("no-dev", false, "Do not include dev dependencies.")
	depsCmd.Flags().StringP("style", "s", "", "How output should be displayed.")
	rootCmd.AddCommand(depsCmd)

	carapace.Gen(depsCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("compact", "tree", "list"),
	})
}
