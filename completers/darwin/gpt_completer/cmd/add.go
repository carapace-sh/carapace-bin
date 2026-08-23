package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new partition to an existing table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("block", "b", "", "Starting sector number")
	addCmd.Flags().StringP("index", "i", "", "Entry index in the GPT table")
	addCmd.Flags().StringP("size", "s", "", "Size of the partition in sectors")
	addCmd.Flags().StringP("type", "t", "", "Partition type UUID")

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValuesDescribed(
			"00000000-0000-0000-0000-000000000000", "Unused",
			"48465300-0000-11AA-AA11-00306543ECAC", "Apple HFS+",
			"7C3457EF-0000-11AA-AA11-00306543ECAC", "Apple APFS",
		),
	})
}
