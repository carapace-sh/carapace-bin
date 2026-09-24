package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a BGP speaker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_createCmd).Standalone()

	bgp_speaker_createCmd.Flags().Bool("advertise-floating-ip-host-routes", false, "Enable the advertisement of floating IP host routes by the BGP speaker.")
	bgp_speaker_createCmd.Flags().Bool("advertise-tenant-networks", false, "Enable the advertisement of tenant network routes by the BGP speaker.")
	bgp_speaker_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgp_speaker_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgp_speaker_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgp_speaker_createCmd.Flags().String("ip-version", "", "IP version for the BGP speaker (default is 4)")
	bgp_speaker_createCmd.Flags().String("local-as", "", "Local AS number.")
	bgp_speaker_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgp_speaker_createCmd.Flags().Bool("no-advertise-floating-ip-host-routes", false, "Disable the advertisement of floating IP host routes by the BGP speaker.")
	bgp_speaker_createCmd.Flags().Bool("no-advertise-tenant-networks", false, "Disable the advertisement of tenant network routes by the BGP speaker.")
	bgp_speaker_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgp_speaker_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgp_speaker_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgp_speaker_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	bgp_speaker_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	bgp_speaker_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgp_speaker_createCmd.MarkFlagRequired("local-as")
	bgp_speakerCmd.AddCommand(bgp_speaker_createCmd)
}
