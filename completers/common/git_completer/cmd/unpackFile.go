package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpackFileCmd = &cobra.Command{
	Use:     "unpack-file",
	Short:   "Creates a temporary file with a blob's contents",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(unpackFileCmd).Standalone()

	rootCmd.AddCommand(unpackFileCmd)
}
