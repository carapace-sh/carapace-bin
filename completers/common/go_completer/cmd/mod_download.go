package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var mod_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download modules to local cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_downloadCmd).Standalone()
	mod_downloadCmd.Flags().SetInterspersed(false)

	mod_downloadCmd.Flags().BoolS("json", "json", false, "print a sequence of JSON objects")
	mod_downloadCmd.Flags().BoolS("x", "x", false, "print the commands download executes")
	modCmd.AddCommand(mod_downloadCmd)

	carapace.Gen(mod_downloadCmd).PositionalCompletion(
		carapace.Batch(
			golang.ActionModules(golang.ModuleOpts{Direct: true, Indirect: true, IncludeVersion: true}).Suppress(".*go.mod file not found.*"),
			golang.ActionModuleSearch(),
		).ToA(),
	)
}
