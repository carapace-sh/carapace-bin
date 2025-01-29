package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_createCmd = &cobra.Command{
	Use:   "create [options] [NAME]",
	Short: "Create a new volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_createCmd).Standalone()

	volume_createCmd.Flags().StringP("driver", "d", "", "Specify volume driver name")
	volume_createCmd.Flags().Bool("ignore", false, "Don't fail if volume already exists")
	volume_createCmd.Flags().StringSliceP("label", "l", []string{}, "Set metadata for a volume (default [])")
	volume_createCmd.Flags().StringSliceP("opt", "o", []string{}, "Set driver specific options (default [])")
	volumeCmd.AddCommand(volume_createCmd)
}
