package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "systeminfo",
	Short: "display detailed configuration information about a computer and its operating system",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/systeminfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("s", "s", false, "specify remote system to query")
	rootCmd.Flags().BoolP("fo", "fo", false, "specify output format (table, list, csv)")
	rootCmd.Flags().BoolP("nh", "nh", false, "no column headers in output")
}
