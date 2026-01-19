package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "buildctl",
	Short: "build utility",
	Long:  "https://github.com/moby/buildkit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("addr", "", "buildkitd address")
	rootCmd.Flags().Bool("debug", false, "enable debug output in logs")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().String("log-format", "", "log formatter")
	rootCmd.Flags().String("timeout", "", "timeout backend connection after value seconds")
	rootCmd.Flags().String("tlscacert", "", "CA certificate for validation")
	rootCmd.Flags().String("tlscert", "", "client certificate")
	rootCmd.Flags().String("tlsdir", "", "directory containing CA certificate, client certificate, and client key")
	rootCmd.Flags().String("tlskey", "", "client key")
	rootCmd.Flags().String("tlsservername", "", "buildkitd server name for certificate validation")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")
	rootCmd.Flags().Bool("wait", false, "block RPCs until the connection becomes available")

	rootCmd.MarkFlagsMutuallyExclusive("tlsdir", "tlscacert")
	rootCmd.MarkFlagsMutuallyExclusive("tlsdir", "tlscert")
	rootCmd.MarkFlagsMutuallyExclusive("tlsdir", "tlskey")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-format": carapace.ActionValues("json", "text"),
		"tlscacert":  carapace.ActionFiles(),
		"tlscert":    carapace.ActionFiles(),
		"tlsdir":     carapace.ActionDirectories(),
		"tlskey":     carapace.ActionFiles(),
	})
}
