package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Discover agent skills relevant to your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillsCmd).Standalone()

	skillsCmd.Flags().String("format", "", "Output format")
	skillsCmd.Flags().Bool("json", false, "Output as JSON")
	skillsCmd.Flags().Bool("yes", false, "Skip confirmation")

	rootCmd.AddCommand(skillsCmd)

	carapace.Gen(skillsCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
