package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var menuconfigCmd = &cobra.Command{
	Use:   "menuconfig",
	Short: "Run menuconfig project configuration tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(menuconfigCmd).Standalone()
	menuconfigCmd.Flags().String("color-scheme", "", "Menuconfig color scheme")
	menuconfigCmd.Flags().String("style", "", "Menuconfig color scheme")
	carapace.Gen(menuconfigCmd).FlagCompletion(carapace.ActionMap{
		"color-scheme": carapace.ActionValues("dark", "light"),
		"style":        carapace.ActionValues("dark", "light"),
	})
	rootCmd.AddCommand(menuconfigCmd)
}
