package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var docs_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docs_serveCmd).Standalone()

	docs_serveCmd.Flags().Bool("no-browser", false, "")
	docs_serveCmd.Flags().String("port", "", "Specify the port number for the docs server.")
	docs_serveCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	docs_serveCmd.Flags().String("vars", "", "Supply variables to the project.")
	addProjectDirFlag(docs_serveCmd)
	addProfileFlag(docs_serveCmd)
	docsCmd.AddCommand(docs_serveCmd)

	carapace.Gen(docs_serveCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})
}
