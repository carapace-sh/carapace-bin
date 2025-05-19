package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imapSendCmd = &cobra.Command{
	Use:     "imap-send",
	Short:   "Send a collection of patches from stdin to an IMAP folder",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(imapSendCmd).Standalone()

	imapSendCmd.Flags().Bool("curl", false, "use libcurl to communicate with the IMAP server")
	imapSendCmd.Flags().Bool("no-curl", false, "do not use libcurl to communicate with the IMAP server")
	imapSendCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	imapSendCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	rootCmd.AddCommand(imapSendCmd)
}
