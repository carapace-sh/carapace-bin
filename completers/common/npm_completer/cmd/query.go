package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Retrieve a filtered list of packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()
	queryCmd.Flags().String("before", "", "only install packages available on or before the given date")
	queryCmd.Flags().Bool("expect-results", false, "whether or not to expect results from the command")
	queryCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	queryCmd.Flags().Bool("include-workspace-root", false, "include the workspace root when workspaces are enabled")
	queryCmd.Flags().String("min-release-age", "", "only install packages available more than the given number of days ago")
	queryCmd.Flags().StringArray("min-release-age-exclude", []string{}, "package names or globs exempt from the min-release-age filter")
	queryCmd.Flags().Bool("package-lock-only", false, "only use the package-lock.json")
	addWorkspaceFlags(queryCmd)

	rootCmd.AddCommand(queryCmd)
}
