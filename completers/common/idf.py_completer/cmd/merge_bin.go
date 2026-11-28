package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mergeBinCmd = &cobra.Command{
	Use:   "merge-bin",
	Short: "Merge binaries into a single image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeBinCmd).Standalone()
	mergeBinCmd.Flags().String("flash-offset", "", "Flash offset where the output file will be flashed")
	mergeBinCmd.Flags().StringP("format", "f", "", "Format of the output file")
	mergeBinCmd.Flags().Bool("md5-disable", false, "Disable MD5 checksum in UF2 output")
	mergeBinCmd.Flags().StringP("output", "o", "", "Output filename")
	mergeBinCmd.Flags().String("pad-to-size", "", "Pad final binary file with FF bytes up to this flash size")
	carapace.Gen(mergeBinCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("hex", "raw", "uf2"),
		"output":      carapace.ActionFiles(),
		"pad-to-size": carapace.ActionValues("256KB", "512KB", "1MB", "2MB", "4MB", "8MB", "16MB", "32MB", "64MB", "128MB"),
	})
	rootCmd.AddCommand(mergeBinCmd)
}
