package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dseditgroup",
	Short: "group record manipulation tool",
	Long:  "https://keith.github.io/xcode-manpages/dseditgroup.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "Maintain ComputerLists in parallel with ComputerGroups")
	rootCmd.Flags().StringS("P", "P", "", "Authentication password")
	rootCmd.Flags().StringS("S", "S", "", "SID to add/replace")
	rootCmd.Flags().StringS("T", "T", "", "Group type")
	rootCmd.Flags().StringS("a", "a", "", "Record name to add")
	rootCmd.Flags().StringS("c", "c", "", "Comment")
	rootCmd.Flags().StringS("d", "d", "", "Record name to delete")
	rootCmd.Flags().StringS("f", "f", "", "Change group format")
	rootCmd.Flags().StringS("g", "g", "", "GUID to add/replace")
	rootCmd.Flags().StringS("i", "i", "", "GID to add/replace")
	rootCmd.Flags().StringS("k", "k", "", "Keyword")
	rootCmd.Flags().StringS("m", "m", "", "Username for checkmember")
	rootCmd.Flags().StringS("n", "n", "", "Directory node location")
	rootCmd.Flags().StringS("o", "o", "", "Operation")
	rootCmd.Flags().BoolS("p", "p", false, "Prompt for authentication password")
	rootCmd.Flags().BoolS("q", "q", false, "Disables interactive verification")
	rootCmd.Flags().StringS("r", "r", "", "Realname")
	rootCmd.Flags().StringS("s", "s", "", "Seconds to live")
	rootCmd.Flags().StringS("t", "t", "", "Record type")
	rootCmd.Flags().StringS("u", "u", "", "Admin username")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"m": os.ActionUsers(),
		"o": carapace.ActionValues("read", "create", "delete", "edit", "checkmember"),
		"u": os.ActionUsers(),
	})
}
