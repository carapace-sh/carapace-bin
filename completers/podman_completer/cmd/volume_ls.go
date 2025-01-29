package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_lsCmd = &cobra.Command{
	Use:     "ls [options]",
	Short:   "List volumes",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_lsCmd).Standalone()

	volume_lsCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter volume output")
	volume_lsCmd.Flags().String("format", "", "Format volume output using Go template")
	volume_lsCmd.Flags().BoolP("noheading", "n", false, "Do not print headers")
	volume_lsCmd.Flags().BoolP("quiet", "q", false, "Print volume output in quiet mode")
	volumeCmd.AddCommand(volume_lsCmd)
}
