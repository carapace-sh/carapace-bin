package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whatsnewCmd = &cobra.Command{
	Use:   "whatsnew [version]",
	Short: "Show release notes for new versions of glab.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whatsnewCmd).Standalone()

	whatsnewCmd.Flags().Bool("latest", false, "Show release notes for the latest published release only.")
	whatsnewCmd.Flags().String("since", "", "Show release notes for every release newer than this version.")
	rootCmd.AddCommand(whatsnewCmd)
}
