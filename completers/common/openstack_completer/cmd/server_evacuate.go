package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_evacuateCmd = &cobra.Command{
	Use:   "evacuate",
	Short: "Evacuate a server to a different host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_evacuateCmd).Standalone()

	server_evacuateCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_evacuateCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_evacuateCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_evacuateCmd.Flags().String("host", "", "Set the preferred host on which to rebuild the evacuated server.")
	server_evacuateCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_evacuateCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_evacuateCmd.Flags().String("password", "", "Set the password on the evacuated instance.")
	server_evacuateCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_evacuateCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_evacuateCmd.Flags().Bool("shared-storage", false, "Indicate that the instance is on shared storage.")
	server_evacuateCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_evacuateCmd.Flags().Bool("wait", false, "Wait for evacuation to complete")
	serverCmd.AddCommand(server_evacuateCmd)
}
