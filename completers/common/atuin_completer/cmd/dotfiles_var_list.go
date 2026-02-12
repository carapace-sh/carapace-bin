package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_var_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_var_listCmd).Standalone()

	dotfiles_var_listCmd.Flags().Bool("exports-only", false, "Show only exported variables")
	dotfiles_var_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	dotfiles_var_listCmd.Flags().StringP("name", "n", "", "Filter variables by name (substring match)")
	dotfiles_var_listCmd.Flags().BoolP("reverse", "r", false, "Sort in reverse (descending) order")
	dotfiles_var_listCmd.Flags().Bool("shell-only", false, "Show only non-exported (shell) variables")
	dotfiles_var_listCmd.Flags().String("sort-by", "", "Sort results by field")
	dotfiles_var_listCmd.Flags().StringP("value", "v", "", "Filter variables by value (substring match)")
	dotfiles_varCmd.AddCommand(dotfiles_var_listCmd)

	carapace.Gen(dotfiles_var_listCmd).FlagCompletion(carapace.ActionMap{
		"sort-by": carapace.ActionValues("name", "value"),
	})
}
