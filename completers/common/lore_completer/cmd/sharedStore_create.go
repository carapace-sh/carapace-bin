package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedStore_createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedStore_createCmd).Standalone()

	sharedStore_createCmd.Flags().BoolP("help", "h", false, "Print help")
	sharedStore_createCmd.Flags().String("make-default", "", "Set this as the default shared store in the global config file, defaults to true")
	sharedStore_createCmd.Flags().String("path", "", "Where to create the shared store")
	sharedStoreCmd.AddCommand(sharedStore_createCmd)

	carapace.Gen(sharedStore_createCmd).FlagCompletion(carapace.ActionMap{
		"make-default": carapace.ActionValues("true", "false"),
		"path":         carapace.ActionDirectories(),
	})

	carapace.Gen(sharedStore_createCmd).PositionalCompletion(
		carapace.ActionValues(), // remote-url
	)
}
