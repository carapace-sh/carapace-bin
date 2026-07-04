package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnscmd",
	Short: "manage DNS servers",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/dnscmd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"info", "display server information",
			"config", "display or modify server configuration",
			"statistics", "display or clear server statistics",
			"clear", "clear cache or statistics",
			"restart", "restart the DNS server",
			"zone", "manage zones",
			"record", "manage resource records",
			"directory", "manage directory partitions",
		),
	)
}
