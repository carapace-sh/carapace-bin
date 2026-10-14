package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_service_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create security service used by project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_service_createCmd).Standalone()

	share_security_service_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_security_service_createCmd.Flags().String("default-ad-site", "", "Default AD site.")
	share_security_service_createCmd.Flags().String("description", "", "Security service description.")
	share_security_service_createCmd.Flags().String("dns-ip", "", "DNS IP address of the security service used inside project's network.")
	share_security_service_createCmd.Flags().String("domain", "", "Security service domain.")
	share_security_service_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_security_service_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_security_service_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_security_service_createCmd.Flags().String("name", "", "Security service name.")
	share_security_service_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_security_service_createCmd.Flags().String("ou", "", "Security service OU (Organizational Unit).")
	share_security_service_createCmd.Flags().String("password", "", "Password used by user.")
	share_security_service_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_security_service_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_security_service_createCmd.Flags().String("server", "", "Security service IP address or hostname.")
	share_security_service_createCmd.Flags().String("user", "", "Security service user or group used by project.")
	share_security_service_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_security_serviceCmd.AddCommand(share_security_service_createCmd)
}
