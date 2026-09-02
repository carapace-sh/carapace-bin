package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage packages for publishing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stageCmd).Standalone()
	stageCmd.Flags().String("access", "", "package access level")
	stageCmd.Flags().Bool("dry-run", false, "report what would be done without making changes")
	stageCmd.Flags().Bool("json", false, "output as json")
	stageCmd.Flags().String("otp", "", "one-time password")
	stageCmd.Flags().Bool("provenance", false, "link package to where it was built")
	stageCmd.Flags().String("tag", "latest", "the tag to add to the package")
	addWorkspaceFlags(stageCmd)

	rootCmd.AddCommand(stageCmd)

	carapace.Gen(stageCmd).PositionalCompletion(
		carapace.ActionValues("approve", "download", "list", "publish", "reject", "view"),
	)
}
