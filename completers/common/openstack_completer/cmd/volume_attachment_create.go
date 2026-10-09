package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_attachment_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an attachment for a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_attachment_createCmd).Standalone()

	volume_attachment_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_attachment_createCmd.Flags().Bool("connect", false, "Make an active connection using provided connector info")
	volume_attachment_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_attachment_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_attachment_createCmd.Flags().String("host", "", "Name of the host attaching to")
	volume_attachment_createCmd.Flags().String("initiator", "", "IQN of the initiator attaching to")
	volume_attachment_createCmd.Flags().String("ip", "", "IP of the system attaching to")
	volume_attachment_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_attachment_createCmd.Flags().String("mode", "", "Mode of volume attachment, rw, ro and null, where null indicates we want to honor any existing admin-metadata settings (supported by --os-volume-api-version 3.54 or later)")
	volume_attachment_createCmd.Flags().String("mountpoint", "", "Mountpoint volume will be attached at")
	volume_attachment_createCmd.Flags().Bool("multipath", false, "Use multipath")
	volume_attachment_createCmd.Flags().Bool("no-connect", false, "Do not make an active connection using provided connector info")
	volume_attachment_createCmd.Flags().Bool("no-multipath", false, "Use multipath")
	volume_attachment_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_attachment_createCmd.Flags().String("os-type", "", "OS type")
	volume_attachment_createCmd.Flags().String("platform", "", "Platform type")
	volume_attachment_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_attachment_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_attachment_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_attachmentCmd.AddCommand(volume_attachment_createCmd)
}
