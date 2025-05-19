package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var postNewsCmd = &cobra.Command{
	Use:    "post-news",
	Short:  "Post news to the self-hosted Charm server.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(postNewsCmd).Standalone()

	postNewsCmd.Flags().String("data-dir", "", "Directory to store SQLite db, SSH keys and file data")
	postNewsCmd.Flags().StringP("subject", "s", "", "Subject for news post")
	postNewsCmd.Flags().StringP("tags", "t", "", "Tags for news post, comma separated")
	rootCmd.AddCommand(postNewsCmd)

	carapace.Gen(postNewsCmd).FlagCompletion(carapace.ActionMap{
		"data-dir": carapace.ActionDirectories(),
	})
}
