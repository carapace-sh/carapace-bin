package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_createCmd).Standalone()

	server_createCmd.Flags().String("auto-network", "", "Automatically allocate a network to the server.")
	server_createCmd.Flags().String("availability-zone", "", "Select an availability zone for the server.")
	server_createCmd.Flags().String("block-device", "", "Create a block device on the server.")
	server_createCmd.Flags().String("block-device-mapping", "", "**Deprecated** Create a block device on the server.")
	server_createCmd.Flags().String("boot-from-volume", "", "When used in conjunction with the ``--image`` or ``--image-property`` option, this option automatically creates a block device mapping with a boot index of 0 and tells the compute service to create a volume of the given size (in GB) from the specified image and use it as the root disk of the server.")
	server_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_createCmd.Flags().String("config-drive", "", "**Deprecated** Use specified volume as the config drive, or 'True' to use an ephemeral drive.")
	server_createCmd.Flags().String("description", "", "Set description for the server (supported by --os-compute-api-version 2.19 or above)")
	server_createCmd.Flags().String("ephemeral", "", "Create and attach a local ephemeral block device of <size> GiB and format it to <format>.")
	server_createCmd.Flags().String("file", "", "File(s) to inject into image before boot (repeat option to set multiple files) (supported by --os-compute-api-version 2.57 or below)")
	server_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_createCmd.Flags().String("flavor", "", "Create server with this flavor (name or ID)")
	server_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_createCmd.Flags().String("hint", "", "Hints for the scheduler")
	server_createCmd.Flags().String("host", "", "Requested host to create servers.")
	server_createCmd.Flags().String("hostname", "", "Hostname configured for the server in the metadata service.")
	server_createCmd.Flags().String("hypervisor-hostname", "", "Requested hypervisor hostname to create servers.")
	server_createCmd.Flags().String("image", "", "Create server boot disk from this image (name or ID)")
	server_createCmd.Flags().String("image-property", "", "Create server using the image that matches the specified property.")
	server_createCmd.Flags().String("key-name", "", "Keypair to inject into this server")
	server_createCmd.Flags().String("max", "", "Maximum number of servers to launch (default=1)")
	server_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_createCmd.Flags().String("min", "", "Minimum number of servers to launch (default=1)")
	server_createCmd.Flags().String("network", "", "Create a NIC on the server and connect it to network.")
	server_createCmd.Flags().String("nic", "", "Create a NIC on the server.")
	server_createCmd.Flags().Bool("no-config-drive", false, "Disable config drive.")
	server_createCmd.Flags().String("no-network", "", "Do not attach a network to the server.")
	server_createCmd.Flags().Bool("no-security-group", false, "Do not associate a security group with ports attached to this server.")
	server_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_createCmd.Flags().String("password", "", "Set the password to this server.")
	server_createCmd.Flags().String("port", "", "Create a NIC on the server and connect it to port.")
	server_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_createCmd.Flags().String("property", "", "Set a property on this server (repeat option to set multiple values)")
	server_createCmd.Flags().String("security-group", "", "Security group to associate with ports attached to this server (name or ID) (repeat option to set multiple groups)")
	server_createCmd.Flags().String("server-group", "", "Server group to create the server within (this is an alias for '--hint group=<server-group-id>')")
	server_createCmd.Flags().String("snapshot", "", "Create server using this snapshot as the boot disk (name or ID)")
	server_createCmd.Flags().String("swap", "", "Create and attach a local swap block device of <swap_size> MiB.")
	server_createCmd.Flags().String("tag", "", "Tags for the server.")
	server_createCmd.Flags().String("trusted-image-cert", "", "Trusted image certificate IDs used to validate certificates during the image signature verification process.")
	server_createCmd.Flags().Bool("use-config-drive", false, "Enable config drive.")
	server_createCmd.Flags().String("user-data", "", "User data file to serve from the metadata server")
	server_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_createCmd.Flags().String("volume", "", "Create server using this volume as the boot disk (name or ID)")
	server_createCmd.Flags().Bool("wait", false, "Wait for build to complete")
	server_createCmd.MarkFlagRequired("flavor")
	serverCmd.AddCommand(server_createCmd)
}
