package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skill_publishCmd = &cobra.Command{
	Use:   "publish [<directory>] [flags]",
	Short: "Validate and publish skills to a GitHub repository (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_publishCmd).Standalone()

	skill_publishCmd.Flags().Bool("dry-run", false, "Validate without publishing")
	skill_publishCmd.Flags().Bool("fix", false, "Auto-fix issues where possible without publishing (e.g. strip install metadata)")
	skill_publishCmd.Flags().String("tag", "", "Version tag for the release (e.g. v1.0.0)")
	skillCmd.AddCommand(skill_publishCmd)

	carapace.Gen(skill_publishCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
