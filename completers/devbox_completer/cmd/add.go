package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new package to your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	addCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO devbox is currently hardcoded against `nixpkgs` alias which is implicitly added as prefix.
	// This channel also has no sqlitedb so package names cannot be completed
	// carapace.Gen(addCmd).PositionalAnyCompletion(
	// 	carapace.ActionCallback(func(c carapace.Context) carapace.Action {
	// 		return nix.ActionChannelPackages().Invoke(c).Filter(c.Parts).ToA()
	// 	}),
	// )
}
