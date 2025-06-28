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

	rootCmd.Flags().StringS("C", "C", "", "Parse and check configuration file, do not execute command")
	rootCmd.Flags().BoolS("L", "L", false, "Clear any persisted authentications and exit")
	rootCmd.Flags().BoolS("n", "n", false, "Non interactive mode, fail if password entry is needed")
	rootCmd.Flags().BoolS("s", "s", false, "Execute shell from $SHELL or /etc/passwd")
	rootCmd.Flags().StringS("u", "u", "", "Execute command as specified user")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionFiles(),
		"u": os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
