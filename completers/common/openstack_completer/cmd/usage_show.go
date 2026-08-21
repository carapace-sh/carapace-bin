package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usage_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show resource usage for a single project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usage_showCmd).Standalone()

	usage_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	usage_showCmd.Flags().String("end", "", "Usage range end date, ex 2012-01-20 (default: tomorrow)")
	usage_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	usage_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	usage_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	usage_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	usage_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	usage_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	usage_showCmd.Flags().String("project", "", "Name or ID of project to show usage for")
	usage_showCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	usage_showCmd.Flags().String("start", "", "Usage range start date, ex 2012-01-20 (default: 4 weeks ago)")
	usage_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	usageCmd.AddCommand(usage_showCmd)
}
