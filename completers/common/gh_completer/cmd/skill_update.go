package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skill_updateCmd = &cobra.Command{
	Use:   "update [<skill>...] [flags]",
	Short: "Update installed skills to their latest versions (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_updateCmd).Standalone()

	skill_updateCmd.Flags().Bool("all", false, "Update all skills without prompting")
	skill_updateCmd.Flags().String("dir", "", "Scan a custom directory for installed skills")
	skill_updateCmd.Flags().Bool("dry-run", false, "Report available updates without modifying files")
	skill_updateCmd.Flags().Bool("force", false, "Re-download even if already up to date")
	skill_updateCmd.Flags().Bool("unpin", false, "Clear pinned version and include pinned skills in update")
	skillCmd.AddCommand(skill_updateCmd)

	carapace.Gen(skill_updateCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})

	// TODO (installed?) skill completion
}
