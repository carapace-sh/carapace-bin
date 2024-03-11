package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go code from *.templ files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().StringS("cmd", "cmd", "", "Set the command to run after generating code")
	generateCmd.Flags().StringS("f", "f", "", "Optionally generates code for a single file")
	generateCmd.Flags().BoolS("help", "help", false, "Print help and exit.")
	generateCmd.Flags().StringS("path", "path", "", "Generates code for all files in path")
	generateCmd.Flags().StringS("pprof", "pprof", "", "Port to start pprof web server on")
	generateCmd.Flags().StringS("proxy", "proxy", "", "Set the URL to proxy")
	generateCmd.Flags().StringS("proxyport", "proxyport", "", "The port the proxy will listen on")
	generateCmd.Flags().BoolS("sourceMapVisualisations", "sourceMapVisualisations", false, "Generate HTML files to visualise the templ code")
	generateCmd.Flags().StringS("w", "w", "", "Number of workers to run in parallel")
	generateCmd.Flags().BoolS("watch", "watch", false, "Watch the path for changes and regenerate code")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).FlagCompletion(carapace.ActionMap{
		"cmd":       bridge.ActionCarapaceBin().Split(),
		"f":         carapace.ActionFiles(),
		"path":      carapace.ActionDirectories(),
		"pprof":     net.ActionPorts(),
		"proxyport": net.ActionPorts(),
	})
}
