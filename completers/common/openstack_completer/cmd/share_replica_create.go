package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a replica of the given share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_createCmd).Standalone()

	share_replica_createCmd.Flags().String("availability-zone", "", "Availability zone in which the replica should be created.")
	share_replica_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_replica_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_replica_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_replica_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_replica_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_replica_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_replica_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_replica_createCmd.Flags().String("property", "", "Set a property to this replica (repeat option to set multiple properties).")
	share_replica_createCmd.Flags().String("scheduler-hint", "", "Scheduler hint for the share replica as key=value pairs, Supported key is only_host.")
	share_replica_createCmd.Flags().String("share-network", "", "Optional network info ID or name.")
	share_replica_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_replica_createCmd.Flags().Bool("wait", false, "Wait for replica creation")
	share_replicaCmd.AddCommand(share_replica_createCmd)
}
