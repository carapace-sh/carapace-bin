package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_credentials_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display EC2 credentials details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_credentials_showCmd).Standalone()

	ec2_credentials_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	ec2_credentials_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	ec2_credentials_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	ec2_credentials_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	ec2_credentials_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	ec2_credentials_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	ec2_credentials_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	ec2_credentials_showCmd.Flags().String("user", "", "Show credentials for user (name or ID)")
	ec2_credentials_showCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	ec2_credentials_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	ec2_credentialsCmd.AddCommand(ec2_credentials_showCmd)
}
