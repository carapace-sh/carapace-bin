package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detailed information for a volume group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_showCmd).Standalone()

	volume_group_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_group_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_group_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_group_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_group_showCmd.Flags().Bool("no-replication-targets", false, "Do not show replication targets for the group.")
	volume_group_showCmd.Flags().Bool("no-volumes", false, "Do not show volumes included in the group.")
	volume_group_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_group_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_group_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_group_showCmd.Flags().Bool("replication-targets", false, "Show replication targets for the group.")
	volume_group_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_group_showCmd.Flags().Bool("volumes", false, "Show volumes included in the group.")
	volume_groupCmd.AddCommand(volume_group_showCmd)
}
