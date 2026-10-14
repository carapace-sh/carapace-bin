package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_qos_type_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new qos type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_qos_type_createCmd).Standalone()

	share_qos_type_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_qos_type_createCmd.Flags().String("description", "", "QoS type description.")
	share_qos_type_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_qos_type_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_qos_type_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_qos_type_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_qos_type_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_qos_type_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_qos_type_createCmd.Flags().String("spec", "", "Spec key and value of QoS type that will be used for QoS type creation.")
	share_qos_type_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_qos_typeCmd.AddCommand(share_qos_type_createCmd)
}
