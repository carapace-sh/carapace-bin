package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailsplitCmd = &cobra.Command{
	Use:     "mailsplit",
	Short:   "Simple UNIX mbox splitter program",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(mailsplitCmd).Standalone()

	mailsplitCmd.Flags().BoolS("b", "b", false, "if any file doesn’t begin with a From line, assume it is a single mail message")
	mailsplitCmd.Flags().StringS("d", "d", "", "precision for the generated filenames")
	mailsplitCmd.Flags().StringS("f", "f", "", "skip the first <nn> numbers")
	mailsplitCmd.Flags().Bool("keep-cr", false, "do not remove \r from lines ending with \r\n")
	mailsplitCmd.Flags().Bool("mboxrd", false, "input is of the \"mboxrd\" format and \"^>+From \" line escaping is reversed")
	mailsplitCmd.Flags().StringS("o", "o", "", "directory in which to place the individual messages")
	rootCmd.AddCommand(mailsplitCmd)

	carapace.Gen(mailsplitCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionDirectories(),
	})

	carapace.Gen(mailsplitCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(mailsplitCmd).DashAnyCompletion(
		carapace.ActionPositional(mailsplitCmd),
	)
}
