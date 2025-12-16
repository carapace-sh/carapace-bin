package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sudo",
	Short: "execute a command as another user",
	Long:  "https://linux.die.net/man/8/sudo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolP("askpass", "A", false, "use a helper program for password prompting")
	rootCmd.Flags().BoolP("background", "b", false, "run command in the background")
	rootCmd.Flags().BoolP("bell", "B", false, "ring bell when prompting")
	rootCmd.Flags().StringP("close-from", "C", "", "close all file descriptors >= num")
	rootCmd.Flags().BoolP("edit", "e", false, "edit files instead of running a command")
	rootCmd.Flags().StringP("group", "g", "", "run command as the specified group name or ID")
	rootCmd.Flags().BoolP("help", "h", false, "display help message and exit")
	rootCmd.Flags().String("host", "", "run command on host (if supported by plugin)")
	rootCmd.Flags().BoolP("list", "l", false, "list user's privileges or check a specific command; use twice for longer format")
	rootCmd.Flags().BoolP("login", "i", false, "run login shell as the target user; a command may also be specified")
	rootCmd.Flags().BoolP("non-interactive", "n", false, "non-interactive mode, no prompts are used")
	rootCmd.Flags().StringP("other-user", "U", "", "in list mode, display privileges for user")
	rootCmd.Flags().StringP("preserve-env", "E", "", "preserve user environment when running command")
	rootCmd.Flags().BoolP("preserve-groups", "P", false, "preserve group vector instead of setting to target's")
	rootCmd.Flags().StringP("prompt", "p", "", "use the specified password prompt")
	rootCmd.Flags().BoolP("remove-timestamp", "K", false, "remove timestamp file completely")
	rootCmd.Flags().BoolP("reset-timestamp", "k", false, "invalidate timestamp file")
	rootCmd.Flags().BoolP("set-home", "H", false, "set HOME variable to target user's home dir")
	rootCmd.Flags().BoolP("shell", "s", false, "run shell as the target user; a command may also be specified")
	rootCmd.Flags().BoolP("stdin", "S", false, "read password from standard input")
	rootCmd.Flags().StringP("user", "u", "", "run command (or edit file) as specified user name or ID")
	rootCmd.Flags().BoolP("validate", "v", false, "update user's timestamp without running a command")
	rootCmd.Flags().BoolP("version", "V", false, "display version information and exit")

	rootCmd.Flag("preserve-env").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"group":        os.ActionGroups(),
		"other-user":   os.ActionUsers(),
		"preserve-env": os.ActionEnvironmentVariables().UniqueList(","),
		"user":         os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
