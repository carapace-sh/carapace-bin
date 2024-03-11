package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:     "listen",
	Short:   "Starts the engine server",
	Aliases: []string{"l"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listenCmd).Standalone()

	listenCmd.Flags().Bool("allow-cors", false, "allow Cross-Origin Resource Sharing (CORS) requests")
	listenCmd.Flags().Bool("disable-host-read-write", false, "disable host read/write access")
	listenCmd.PersistentFlags().Bool("focus", false, "Only show output for focused commands")
	listenCmd.Flags().String("listen", "", "Listen on network address ADDR")
	listenCmd.PersistentFlags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	rootCmd.AddCommand(listenCmd)

	carapace.Gen(listenCmd).FlagCompletion(carapace.ActionMap{
		"listen": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return net.ActionPorts()
			}
		}),
		"mod": dagger.ActionMods(),
	})
}
