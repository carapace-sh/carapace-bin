package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_alias_listCmd).Standalone()

	dotfiles_alias_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	dotfiles_alias_listCmd.Flags().StringP("name", "n", "", "Filter aliases by name (substring match)")
	dotfiles_alias_listCmd.Flags().BoolP("reverse", "r", false, "Sort in reverse (descending) order")
	dotfiles_alias_listCmd.Flags().String("sort-by", "", "Sort results by field")
	dotfiles_alias_listCmd.Flags().StringP("value", "v", "", "Filter aliases by value (substring match)")
	dotfiles_aliasCmd.AddCommand(dotfiles_alias_listCmd)
}
