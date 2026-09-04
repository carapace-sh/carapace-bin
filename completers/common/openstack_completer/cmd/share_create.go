package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_createCmd).Standalone()

	share_createCmd.Flags().String("availability-zone", "", "Availability zone in which share should be created.")
	share_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_createCmd.Flags().String("description", "", "Optional share description.")
	share_createCmd.Flags().String("encryption-key-ref", "", "Set encryption key reference i.e. UUID of the secret stored in the key manager.")
	share_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_createCmd.Flags().String("mount-point-name", "", "Optional custom export location.")
	share_createCmd.Flags().String("name", "", "Optional share name.")
	share_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_createCmd.Flags().String("property", "", "Set a property to this share (repeat option to set multiple properties)")
	share_createCmd.Flags().String("public", "", "Level of visibility for share.")
	share_createCmd.Flags().String("scheduler-hint", "", "Set Scheduler hints for the share as key=value pairs, possible keys are same_host, different_host.")
	share_createCmd.Flags().String("share-group", "", "Optional share group name or ID in which to create the share.")
	share_createCmd.Flags().String("share-network", "", "Optional network info ID or name.")
	share_createCmd.Flags().String("share-type", "", "The share type to create the share with.")
	share_createCmd.Flags().String("snapshot-id", "", "Optional snapshot ID to create the share from.")
	share_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_createCmd.Flags().Bool("wait", false, "Wait for share creation")
	shareCmd.AddCommand(share_createCmd)
}
