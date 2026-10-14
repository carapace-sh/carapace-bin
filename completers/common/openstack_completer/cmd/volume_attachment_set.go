package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_attachment_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update an attachment for a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_attachment_setCmd).Standalone()

	volume_attachment_setCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_attachment_setCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_attachment_setCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_attachment_setCmd.Flags().String("host", "", "Name of the host attaching to")
	volume_attachment_setCmd.Flags().String("initiator", "", "IQN of the initiator attaching to")
	volume_attachment_setCmd.Flags().String("ip", "", "IP of the system attaching to")
	volume_attachment_setCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_attachment_setCmd.Flags().String("mountpoint", "", "Mountpoint volume will be attached at")
	volume_attachment_setCmd.Flags().Bool("multipath", false, "Use multipath")
	volume_attachment_setCmd.Flags().Bool("no-multipath", false, "Use multipath")
	volume_attachment_setCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_attachment_setCmd.Flags().String("os-type", "", "OS type")
	volume_attachment_setCmd.Flags().String("platform", "", "Platform type")
	volume_attachment_setCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_attachment_setCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_attachment_setCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_attachmentCmd.AddCommand(volume_attachment_setCmd)
}
