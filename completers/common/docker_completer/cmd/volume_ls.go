package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List volumes",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_lsCmd).Standalone()

	volume_lsCmd.Flags().Bool("cluster", false, "Display only cluster volumes, and use cluster volume list formatting")
	volume_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. \"dangling=true\")")
	volume_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	volume_lsCmd.Flags().BoolP("quiet", "q", false, "Only display volume names")
	volumeCmd.AddCommand(volume_lsCmd)
}
