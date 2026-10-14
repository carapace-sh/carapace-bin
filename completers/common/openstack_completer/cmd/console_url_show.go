package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var console_url_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show server's remote console URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(console_url_showCmd).Standalone()

	console_url_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	console_url_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	console_url_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	console_url_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	console_url_showCmd.Flags().Bool("mks", false, "Show WebMKS console URL")
	console_url_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	console_url_showCmd.Flags().Bool("novnc", false, "Show noVNC console URL (default)")
	console_url_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	console_url_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	console_url_showCmd.Flags().Bool("rdp", false, "Show RDP console URL")
	console_url_showCmd.Flags().Bool("serial", false, "Show serial console URL")
	console_url_showCmd.Flags().Bool("spice", false, "Show SPICE console URL")
	console_url_showCmd.Flags().Bool("spice-direct", false, "Show SPICE direct protocol native console URL")
	console_url_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	console_url_showCmd.Flags().Bool("xvpvnc", false, "Show xvpvnc console URL")
	console_urlCmd.AddCommand(console_url_showCmd)
}
