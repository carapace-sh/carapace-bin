package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpt",
	Short: "GUID partition table maintenance utility",
	Long:  "https://keith.github.io/xcode-manpages/gpt.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("f", false, "Open the device with O_SHLOCK|O_RDWR mode")
	rootCmd.Flags().String("p", "", "Change the number of partitions the GPT can accommodate")
	rootCmd.Flags().Bool("r", false, "Open the device for reading only")
	rootCmd.Flags().Bool("v", false, "Increase verbosity level")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "Add a new partition to an existing table",
			"create", "Create a new GPT",
			"destroy", "Destroy the GPT",
			"init", "Initialize a GPT",
			"label", "Label a partition",
			"recover", "Recover the GPT",
			"remove", "Remove a partition from the table",
			"show", "Show the partition table",
		),
	)
}