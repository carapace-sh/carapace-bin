package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keymapCmd = &cobra.Command{
	Use:   "keymap",
	Short: "Show what each key does",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keymapCmd).Standalone()

	keymapCmd.Flags().Bool("json", false, "output the response as JSON")
	keymapCmd.Flags().String("request", "", "send the client.keymap request in FILE")

	rootCmd.AddCommand(keymapCmd)

	carapace.Gen(keymapCmd).FlagCompletion(carapace.ActionMap{
		"request": carapace.ActionFiles(),
	})
}
