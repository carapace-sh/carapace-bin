package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_qos_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new QoS specification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qos_createCmd).Standalone()

	volume_qos_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_qos_createCmd.Flags().String("consumer", "", "Consumer of the QoS.")
	volume_qos_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_qos_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_qos_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_qos_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_qos_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_qos_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_qos_createCmd.Flags().String("property", "", "Set a QoS specification property (repeat option to set multiple properties)")
	volume_qos_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_qosCmd.AddCommand(volume_qos_createCmd)
}
