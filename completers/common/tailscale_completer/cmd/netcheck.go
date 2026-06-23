package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netcheckCmd = &cobra.Command{
	Use:   "netcheck",
	Short: "Print an analysis of local network conditions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netcheckCmd).Standalone()

	netcheckCmd.Flags().String("bind-address", "", "send and receive connectivity probes using this locally bound IP address")
	netcheckCmd.Flags().String("bind-port", "", "send and receive connectivity probes using this UDP port")
	netcheckCmd.Flags().String("every", "", "do an incremental report with the given frequency")
	netcheckCmd.Flags().String("format", "", "output format")
	netcheckCmd.Flags().Bool("verbose", false, "verbose logs")
	rootCmd.AddCommand(netcheckCmd)

	carapace.Gen(netcheckCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "json-line"),
	})
}
