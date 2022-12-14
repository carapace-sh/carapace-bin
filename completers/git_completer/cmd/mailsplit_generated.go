package cmd

import (
	"github.com/spf13/cobra"
)

var mailsplitCmd = &cobra.Command{
	Use:     "mailsplit",
	Short:   "Simple UNIX mbox splitter program",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {

	rootCmd.AddCommand(mailsplitCmd)
}
