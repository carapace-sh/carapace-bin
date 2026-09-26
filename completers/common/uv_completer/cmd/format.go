package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format Python code in the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().Bool("check", false, "Check if files are formatted without applying changes")
	formatCmd.Flags().Bool("diff", false, "Show a diff of formatting changes without applying them")
	formatCmd.Flags().String("exclude-newer", "", "Limit candidate Ruff versions to those released prior to the given date")
	formatCmd.Flags().Bool("no-project", false, "Avoid discovering a project or workspace")
	formatCmd.Flags().Bool("show-version", false, "Display the version of Ruff that will be used for formatting")
	formatCmd.Flags().String("version", "", "The version of Ruff to use for formatting")
	formatCmd.Flag("show-version").Hidden = true
	rootCmd.AddCommand(formatCmd)
}
