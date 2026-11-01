package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eslogger",
	Short: "log Endpoint Security events",
	Long:  "https://man.freebsd.org/cgi/man.cgi?eslogger",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("format", "", "Log format to use")
	rootCmd.Flags().Bool("list-events", false, "List supported events and exit")
	rootCmd.Flags().Bool("oslog", false, "Emit event data to unified logging system")
	rootCmd.Flags().String("oslog-category", "", "Log category to use with --oslog")
	rootCmd.Flags().String("oslog-subsystem", "", "Log subsystem to use with --oslog")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
	})
}
