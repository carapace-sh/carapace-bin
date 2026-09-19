package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugobsoleteCmd = &cobra.Command{
	Use:    "debugobsolete",
	Short:  "create arbitrary obsolete marker",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugobsoleteCmd).Standalone()

	debugobsoleteCmd.Flags().StringP("date", "d", "", "record the specified date as commit date")
	debugobsoleteCmd.Flags().String("delete", "", "delete markers specified by indices")
	debugobsoleteCmd.Flags().Bool("exclusive", false, "restrict display to markers only relevant to REV")
	debugobsoleteCmd.Flags().String("flags", "", "markers flag")
	debugobsoleteCmd.Flags().Bool("index", false, "display index of the marker")
	debugobsoleteCmd.Flags().Bool("record-parents", false, "record parent information for the precursor")
	debugobsoleteCmd.Flags().StringArrayP("rev", "r", nil, "display markers relevant to REV")
	debugobsoleteCmd.Flags().StringP("template", "T", "", "display with template")
	debugobsoleteCmd.Flags().StringP("user", "u", "", "record the specified user as committer")
	rootCmd.AddCommand(debugobsoleteCmd)

	carapace.Gen(debugobsoleteCmd).FlagCompletion(carapace.ActionMap{
		"delete": action.ActionShelves(),
		"rev":    action.ActionRevisions(),
	})
}
