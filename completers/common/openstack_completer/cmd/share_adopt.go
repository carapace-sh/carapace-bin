package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_adoptCmd = &cobra.Command{
	Use:   "adopt",
	Short: "Adopt a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_adoptCmd).Standalone()

	share_adoptCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_adoptCmd.Flags().String("description", "", "Optional share description.")
	share_adoptCmd.Flags().String("driver-options", "", "Optional driver options as key=value pairs (Default=None).")
	share_adoptCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_adoptCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_adoptCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_adoptCmd.Flags().String("mount-point-name", "", "Optional custom export location.")
	share_adoptCmd.Flags().String("name", "", "Optional share name.")
	share_adoptCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_adoptCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_adoptCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_adoptCmd.Flags().Bool("public", false, "Level of visibility for share.")
	share_adoptCmd.Flags().String("share-server-id", "", "Share server associated with share when using a share type with \"driver_handles_share_servers\" extra_spec set to True.")
	share_adoptCmd.Flags().String("share-type", "", "Optional share type assigned to share.")
	share_adoptCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_adoptCmd.Flags().Bool("wait", false, "Wait until share is adopted")
	shareCmd.AddCommand(share_adoptCmd)
}
