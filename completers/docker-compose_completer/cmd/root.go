package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "Define and run multi-container applications with Docker",
	Long:  "https://docs.docker.com/compose/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("compatibility", false, "If set, Compose will attempt to convert keys")
	rootCmd.Flags().StringP("context", "c", "", "Specify a context name")
	rootCmd.Flags().String("env-file", "", "Specify an alternate environment file")
	rootCmd.Flags().StringP("file", "f", "", "Specify an alternate compose file")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket to connect to")
	rootCmd.Flags().String("log-level", "", "Set log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)")
	rootCmd.Flags().Bool("no-ansi", false, "Do not print ANSI control characters")
	rootCmd.Flags().String("project-directory", "", "Specify an alternate working directory")
	rootCmd.Flags().StringP("project-name", "p", "", "Specify an alternate project name")
	rootCmd.Flags().Bool("skip-hostname-check", false, "Don't check the daemon's hostname against the")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "", "Trust certs signed only by this CA")
	rootCmd.Flags().String("tlscert", "", "Path to TLS certificate file")
	rootCmd.Flags().String("tlskey", "", "Path to TLS key file")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().Bool("verbose", false, "Show more output")
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"env-file":          carapace.ActionFiles(),
		"file":              carapace.ActionFiles(),
		"log-level":         carapace.ActionValues("DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"),
		"project-directory": carapace.ActionDirectories(),
		"tlscert":           carapace.ActionFiles(".crt"),
		"tlskey":            carapace.ActionFiles(".key"),
	})
}
