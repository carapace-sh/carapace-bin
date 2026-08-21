package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new share type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_createCmd).Standalone()

	share_type_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_type_createCmd.Flags().String("create-share-from-snapshot-support", "", "Boolean extra spec used for filtering of back ends by their capability to create shares from snapshots.")
	share_type_createCmd.Flags().String("description", "", "Share type description.")
	share_type_createCmd.Flags().String("extra-specs", "", "Extra specs key and value of share type that will be used for share type creation.")
	share_type_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_type_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_type_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_type_createCmd.Flags().String("mount-snapshot-support", "", "Boolean extra spec used for filtering of back ends by their capability to mount share snapshots.")
	share_type_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_type_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_type_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_type_createCmd.Flags().String("public", "", "Make type accessible to the public (default true).")
	share_type_createCmd.Flags().String("revert-to-snapshot-support", "", "Boolean extra spec used for filtering of back ends by their capability to revert shares to snapshots.")
	share_type_createCmd.Flags().String("snapshot-support", "", "Boolean extra spec used for filtering of back ends by their capability to create share snapshots.")
	share_type_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_typeCmd.AddCommand(share_type_createCmd)
}
