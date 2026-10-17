package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rex",
	Short: "The multiplexer for all work",
	Long:  "https://www.superlogical.com/rex/docs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("autostart", true, "start a missing local server")
	rootCmd.PersistentFlags().StringP("client", "C", "", "act on client `CLIENT`, given as a label or ID")
	rootCmd.PersistentFlags().String("color", "", "set color WHEN to auto, always, or never")
	rootCmd.Flags().BoolP("help", "h", false, "help for rex")
	rootCmd.PersistentFlags().StringP("server", "S", "", "connect to `SERVER`, a host's label or a server address")
	rootCmd.PersistentFlags().String("timeout", "", "limit each connection and nonstreaming request to `duration`")
	rootCmd.PersistentFlags().CountP("verbose", "v", "increase diagnostic verbosity (repeatable)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"client": action.ActionClients(rootCmd),
		"color":  carapace.ActionValues("auto", "always", "never"),
		"server": rex.ActionHosts(),
	})
}
