package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doas",
	Short: "execute a command as another user",
	Long:  "https://man.openbsd.org/doas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("clear", "L", false, "Clear any persisted authentications and exit")
	rootCmd.Flags().StringS("config", "C", "", "Parse and check configuration file, do not execute command")
	rootCmd.Flags().BoolS("shell", "s", false, "Execute shell from $SHELL or /etc/passwd")
	rootCmd.Flags().BoolS("silent", "n", false, "Non interactive mode, fail if password entry is needed")
	rootCmd.Flags().StringS("user", "u", "root", "Execute command as specified user")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"user":   os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
