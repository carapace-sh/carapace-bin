package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("config", "", "Location of client config files (default \"/home/rsteube/.docker\")")
	rootCmd.Flags().StringP("context", "c", "", "Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with \"docker context use\")")
	rootCmd.Flags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket(s) to connect to")
	rootCmd.Flags().StringP("log-level", "l", "", "Set the logging level (\"debug\"|\"info\"|\"warn\"|\"error\"|\"fatal\") (default \"info\")")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "", "Trust certs signed only by this CA (default \"/home/rsteube/.docker/ca.pem\")")
	rootCmd.Flags().String("tlscert", "", "Path to TLS certificate file (default \"/home/rsteube/.docker/cert.pem\")")
	rootCmd.Flags().String("tlskey", "", "Path to TLS key file (default \"/home/rsteube/.docker/key.pem\")")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":    carapace.ActionFiles(""),
		"log-level": carapace.ActionValues("debug", "info", "warn", "error", "fatal"),
		"tlscacert": carapace.ActionFiles(""),
		"tlscert":   carapace.ActionFiles(""),
		"tlskey":    carapace.ActionFiles(""),
	})

	rootCmd.AddCommand(container_attachCmd)
	rootCmd.AddCommand(image_buildCmd)
	rootCmd.AddCommand(container_commitCmd)
	rootCmd.AddCommand(container_cpCmd)
	rootCmd.AddCommand(container_createCmd)
	rootCmd.AddCommand(container_diffCmd)
	rootCmd.AddCommand(system_eventsCmd)
	rootCmd.AddCommand(container_execCmd)
	rootCmd.AddCommand(container_exportCmd)
	rootCmd.AddCommand(image_historyCmd)
	rootCmd.AddCommand(image_importCmd)
	rootCmd.AddCommand(system_infoCmd)
	// TODO docker inspect command
	rootCmd.AddCommand(container_killCmd)
	rootCmd.AddCommand(image_loadCmd)
	// TODO docker login command
	// TODO docker logout command
	rootCmd.AddCommand(container_logsCmd)
	rootCmd.AddCommand(container_pauseCmd)
	rootCmd.AddCommand(container_portCmd)
	psCmd := container_lsCmd
	psCmd.Use = "ps"
	rootCmd.AddCommand(psCmd)
	rootCmd.AddCommand(image_pullCmd)
	rootCmd.AddCommand(image_pushCmd)
	rootCmd.AddCommand(container_renameCmd)
	rootCmd.AddCommand(container_restartCmd)
	rootCmd.AddCommand(container_rmCmd)
	rmiCmd := image_rmCmd
	rmiCmd.Use = "rmi"
	rootCmd.AddCommand(rmiCmd)
	rootCmd.AddCommand(container_runCmd)
	rootCmd.AddCommand(image_saveCmd)
	// TODO docker search command
	rootCmd.AddCommand(container_startCmd)
	rootCmd.AddCommand(container_statsCmd)
	rootCmd.AddCommand(container_stopCmd)
	rootCmd.AddCommand(image_tagCmd)
	rootCmd.AddCommand(container_topCmd)
	rootCmd.AddCommand(container_unpauseCmd)
	rootCmd.AddCommand(container_updateCmd)
	rootCmd.AddCommand(container_waitCmd)
}
