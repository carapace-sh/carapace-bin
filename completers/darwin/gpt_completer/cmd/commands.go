package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(labelCmd)
	rootCmd.AddCommand(recoverCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(showCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new partition to an existing table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new GPT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy the GPT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "label a partition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "recover the GPT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a partition from the table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the partition table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	carapace.Gen(createCmd).Standalone()
	carapace.Gen(destroyCmd).Standalone()
	carapace.Gen(labelCmd).Standalone()
	carapace.Gen(recoverCmd).Standalone()
	carapace.Gen(removeCmd).Standalone()
	carapace.Gen(showCmd).Standalone()

	addCmd.Flags().StringP("block", "b", "", "Starting sector number")
	addCmd.Flags().StringP("index", "i", "", "Entry index in the GPT table")
	addCmd.Flags().StringP("size", "s", "", "Size of the partition in sectors")
	addCmd.Flags().StringP("type", "t", "", "Partition type UUID")

	createCmd.Flags().BoolP("force", "f", false, "Force creation")
	createCmd.Flags().BoolP("protective", "p", false, "Create only protective MBR")

	destroyCmd.Flags().BoolP("recursive", "r", false, "Destroy recursively")

	labelCmd.Flags().Bool("a", false, "Apply to all matching")
	labelCmd.Flags().StringP("file", "f", "", "Read label from file")
	labelCmd.Flags().StringP("label", "l", "", "Label string")

	removeCmd.Flags().Bool("a", false, "Remove all partitions")

	showCmd.Flags().BoolP("list", "l", false, "Show partition details")

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"type": carapace.ActionValues(
			"00000000-0000-0000-0000-000000000000",
			"48465300-0000-11AA-AA11-00306543ECAC",
			"7C3457EF-0000-11AA-AA11-00306543ECAC",
		).Described(
			"Unused",
			"Apple HFS+",
			"Apple APFS",
		),
	})
}