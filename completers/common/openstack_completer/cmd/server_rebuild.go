package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_rebuildCmd).Standalone()

	server_rebuildCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_rebuildCmd.Flags().String("description", "", "Set a new description on the rebuilt server (supported by --os-compute-api-version 2.19 or above)")
	server_rebuildCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_rebuildCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_rebuildCmd.Flags().String("hostname", "", "Hostname configured for the server in the metadata service.")
	server_rebuildCmd.Flags().String("image", "", "Recreate server from the specified image (name or ID).")
	server_rebuildCmd.Flags().String("key-name", "", "Set the key name of key pair on the rebuilt server.")
	server_rebuildCmd.Flags().Bool("key-unset", false, "==SUPPRESS==")
	server_rebuildCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_rebuildCmd.Flags().String("name", "", "Set the new name of the rebuilt server")
	server_rebuildCmd.Flags().Bool("no-key-name", false, "Unset the key name of key pair on the rebuilt server.")
	server_rebuildCmd.Flags().Bool("no-preserve-ephemeral", false, "Do not preserve the default ephemeral storage partition on rebuild.")
	server_rebuildCmd.Flags().Bool("no-reimage-boot-volume", false, "Do not rebuild a volume-backed server.")
	server_rebuildCmd.Flags().Bool("no-trusted-image-certs", false, "Remove any existing trusted image certificates from the server.")
	server_rebuildCmd.Flags().Bool("no-user-data", false, "Remove existing user data when rebuilding server.")
	server_rebuildCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_rebuildCmd.Flags().String("password", "", "Set the password on the rebuilt server.")
	server_rebuildCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_rebuildCmd.Flags().Bool("preserve-ephemeral", false, "Preserve the default ephemeral storage partition on rebuild.")
	server_rebuildCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_rebuildCmd.Flags().String("property", "", "Set a new property on the rebuilt server (repeat option to set multiple values)")
	server_rebuildCmd.Flags().Bool("reimage-boot-volume", false, "Rebuild a volume-backed server.")
	server_rebuildCmd.Flags().String("trusted-image-cert", "", "Trusted image certificate IDs used to validate certificates during the image signature verification process.")
	server_rebuildCmd.Flags().String("user-data", "", "Add a new user data file to the rebuilt server.")
	server_rebuildCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_rebuildCmd.Flags().Bool("wait", false, "Wait for rebuild to complete")
	serverCmd.AddCommand(server_rebuildCmd)
}
