package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_createCmd).Standalone()

	volume_createCmd.Flags().String("availability-zone", "", "Create volume in <availability-zone>")
	volume_createCmd.Flags().String("backup", "", "Restore backup to a volume (name or ID) (supported by --os-volume-api-version 3.47 or later)")
	volume_createCmd.Flags().Bool("bootable", false, "Mark volume as bootable")
	volume_createCmd.Flags().String("cluster", "", "Cinder cluster on which the existing volume resides; takes the form: cluster@backend-name#pool.")
	volume_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_createCmd.Flags().String("consistency-group", "", "Consistency group where the new volume belongs to")
	volume_createCmd.Flags().String("description", "", "Volume description")
	volume_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_createCmd.Flags().String("hint", "", "Arbitrary scheduler hint key-value pairs to help creating a volume.")
	volume_createCmd.Flags().String("host", "", "Cinder host on which the existing volume resides; takes the form: host@backend-name#pool.")
	volume_createCmd.Flags().String("image", "", "Use <image> as source of volume (name or ID)")
	volume_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_createCmd.Flags().Bool("non-bootable", false, "Mark volume as non-bootable (default)")
	volume_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_createCmd.Flags().String("property", "", "Set a property to this volume (repeat option to set multiple properties)")
	volume_createCmd.Flags().Bool("read-only", false, "Set volume to read-only access mode")
	volume_createCmd.Flags().Bool("read-write", false, "Set volume to read-write access mode (default)")
	volume_createCmd.Flags().String("remote-source", "", "The attribute(s) of the existing remote volume (admin required) (repeat option to specify multiple attributes, e.g.: '--remote-source source-name=test_name --remote-source source-id=test_id')")
	volume_createCmd.Flags().String("size", "", "Volume size in GB (required unless --snapshot or --source specified)")
	volume_createCmd.Flags().String("snapshot", "", "Use <snapshot> as source of volume (name or ID)")
	volume_createCmd.Flags().String("source", "", "Volume to clone (name or ID)")
	volume_createCmd.Flags().String("source-replicated", "", "==SUPPRESS==")
	volume_createCmd.Flags().String("type", "", "Set the type of volume")
	volume_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volumeCmd.AddCommand(volume_createCmd)
}
