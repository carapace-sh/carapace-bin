package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "terraform",
	Short: "infrastructure as code software tool",
	Long:  "https://www.terraform.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "main", Title: "Main commands"},
	)

	rootCmd.Flags().StringS("chdir", "chdir", "", "Switch to a different working directory before executing the given subcommand.")
	rootCmd.Flags().BoolS("help", "help", false, "Show this help output, or the help for a specified subcommand.")
	rootCmd.Flags().BoolS("version", "version", false, "An alias for the \"version\" subcommand.")

	rootCmd.Flag("chdir").NoOptDefVal = "."

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chdir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		if f := rootCmd.Flag("chdir"); f.Changed && f != flag {
			return action.Chdir(f.Value.String())
		}
		return action
	})
}
