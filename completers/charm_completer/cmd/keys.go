package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Browse or print linked SSH keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keysCmd).Standalone()

	keysCmd.Flags().BoolP("randomart", "r", false, "print SSH 5.1 randomart for each key (the Drunken Bishop algorithm)")
	keysCmd.Flags().BoolP("simple", "s", false, "simple, non-interactive output (good for scripts)")
	rootCmd.AddCommand(keysCmd)
}
