package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker image and container command line interface",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("config", "", "Location of client config files (default \"/home/user/.docker\")")
	rootCmd.Flags().StringP("context", "c", "", "Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with \"docker context use\")")
	rootCmd.Flags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket(s) to connect to")
	rootCmd.Flags().StringP("log-level", "l", "", "Set the logging level (\"debug\"|\"info\"|\"warn\"|\"error\"|\"fatal\") (default \"info\")")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "", "Trust certs signed only by this CA (default \"/home/user/.docker/ca.pem\")")
	rootCmd.Flags().String("tlscert", "", "Path to TLS certificate file (default \"/home/user/.docker/cert.pem\")")
	rootCmd.Flags().String("tlskey", "", "Path to TLS key file (default \"/home/user/.docker/key.pem\")")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":    carapace.ActionFiles(),
		"log-level": carapace.ActionValues("debug", "info", "warn", "error", "fatal"),
		"tlscacert": carapace.ActionFiles(),
		"tlscert":   carapace.ActionFiles(),
		"tlskey":    carapace.ActionFiles(),
	})
}

func rootAlias(cmd *cobra.Command, initCompletion func(cmd *cobra.Command, isAlias bool)) {
	aliasCmd := cmd
	rootCmd.AddCommand(aliasCmd)

	for i, c := range []*cobra.Command{cmd, aliasCmd} {
		initCompletion(c, i == 1)
	}
}
