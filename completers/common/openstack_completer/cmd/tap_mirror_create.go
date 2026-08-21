package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_mirror_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new tap mirror.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_mirror_createCmd).Standalone()

	tap_mirror_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_mirror_createCmd.Flags().String("description", "", "Description of the tap service.")
	tap_mirror_createCmd.Flags().String("directions", "", "Dictionary of direction and tunnel_id.")
	tap_mirror_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_mirror_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_mirror_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_mirror_createCmd.Flags().String("mirror-type", "", "Mirror type.")
	tap_mirror_createCmd.Flags().String("name", "", "Name of the tap service.")
	tap_mirror_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_mirror_createCmd.Flags().String("port", "", "Port (name or ID) to which the Tap Mirror is connected.")
	tap_mirror_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	tap_mirror_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_mirror_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	tap_mirror_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	tap_mirror_createCmd.Flags().String("remote-ip", "", "Remote IP address for the tap mirror (remote end of the GRE or ERSPAN v1 tunnel).")
	tap_mirror_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	tap_mirror_createCmd.MarkFlagRequired("directions")
	tap_mirror_createCmd.MarkFlagRequired("mirror-type")
	tap_mirror_createCmd.MarkFlagRequired("port")
	tap_mirror_createCmd.MarkFlagRequired("remote-ip")
	tap_mirrorCmd.AddCommand(tap_mirror_createCmd)
}
