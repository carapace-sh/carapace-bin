package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set volume properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_setCmd).Standalone()

	volume_setCmd.Flags().Bool("attached", false, "Set volume attachment status to \"attached\" (admin only) (This option simply changes the state of the volume in the database with no regard to actual status, exercise caution when using)")
	volume_setCmd.Flags().Bool("bootable", false, "Mark volume as bootable")
	volume_setCmd.Flags().String("description", "", "New volume description")
	volume_setCmd.Flags().Bool("detached", false, "Set volume attachment status to \"detached\" (admin only) (This option simply changes the state of the volume in the database with no regard to actual status, exercise caution when using)")
	volume_setCmd.Flags().String("image-property", "", "Set an image property on this volume (repeat option to set multiple image properties)")
	volume_setCmd.Flags().String("migration-policy", "", "Migration policy while re-typing volume (\"never\" or \"on-demand\", default is \"never\" ) (available only when --type option is specified)")
	volume_setCmd.Flags().String("name", "", "New volume name")
	volume_setCmd.Flags().Bool("no-property", false, "Remove all properties from <volume> (specify both --no-property and --property to remove the current properties before setting new properties.)")
	volume_setCmd.Flags().Bool("non-bootable", false, "Mark volume as non-bootable")
	volume_setCmd.Flags().String("property", "", "Set a property on this volume (repeat option to set multiple properties)")
	volume_setCmd.Flags().Bool("read-only", false, "Set volume to read-only access mode")
	volume_setCmd.Flags().Bool("read-write", false, "Set volume to read-write access mode")
	volume_setCmd.Flags().String("retype-policy", "", "==SUPPRESS==")
	volume_setCmd.Flags().String("size", "", "Extend volume size in GB")
	volume_setCmd.Flags().String("state", "", "New volume state (\"available\", \"error\", \"creating\", \"deleting\", \"in-use\", \"attaching\", \"detaching\", \"error_deleting\" or \"maintenance\") (admin only) (This option simply changes the state of the volume in the database with no regard to actual status, exercise caution when using)")
	volume_setCmd.Flags().String("type", "", "New volume type (name or ID)")
	volumeCmd.AddCommand(volume_setCmd)
}
