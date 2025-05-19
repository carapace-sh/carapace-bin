package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix [OPTIONS]",
	Short: "Update files with formatting fixes or other changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().Bool("include-unchanged-files", false, "Fix unchanged files in addition to changed ones. If no paths are specified, all files in the repo will be fixed")
	fixCmd.Flags().StringSliceP("source", "s", nil, "Fix files in specified revision(s) and their descendants")
	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).FlagCompletion(carapace.ActionMap{
		"source": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
