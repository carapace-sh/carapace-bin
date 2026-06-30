package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "abbr",
	Short: "Manage fish abbreviations",
	Long:  "https://fishshell.com/docs/current/cmds/abbr.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "add a new abbreviation")
	rootCmd.Flags().Bool("add", false, "add a new abbreviation")
	rootCmd.Flags().String("color", "", "color output")
	rootCmd.Flags().String("command", "", "only expand with given command")
	rootCmd.Flags().BoolS("e", "e", false, "erase abbreviations")
	rootCmd.Flags().Bool("erase", false, "erase abbreviations")
	rootCmd.Flags().String("function", "", "use function to generate expansion")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("l", "l", false, "list abbreviation names")
	rootCmd.Flags().Bool("list", false, "list abbreviation names")
	rootCmd.Flags().String("position", "", "where abbreviation expands")
	rootCmd.Flags().BoolS("q", "q", false, "query abbreviations")
	rootCmd.Flags().Bool("query", false, "query abbreviations")
	rootCmd.Flags().String("regex", "", "match using regex")
	rootCmd.Flags().String("rename", "", "rename an abbreviation")
	rootCmd.Flags().BoolS("s", "s", false, "show all abbreviations")
	rootCmd.Flags().String("set-cursor", "", "set cursor position marker")
	rootCmd.Flags().Bool("show", false, "show all abbreviations")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":    carapace.ActionValues("always", "never", "auto"),
		"position": carapace.ActionValues("command", "anywhere"),
	})
}
