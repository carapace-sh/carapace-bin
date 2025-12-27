package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "run0",
	Short: "Elevate privileges interactively",
	Long:  "https://www.freedesktop.org/software/systemd/man/latest/run0.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().String("area", "", "Home area to log into")
	rootCmd.Flags().String("background", "", "Set ANSI color for background")
	rootCmd.Flags().StringP("chdir", "D", "", "Set working directory")
	rootCmd.Flags().String("description", "", "Description for unit")
	rootCmd.Flags().Bool("empower", false, "Give privileges to selected or current user")
	rootCmd.Flags().StringP("group", "g", "", "Run as system group")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolS("i", "i", false, "Shortcut for --via-shell --chdir='~'")
	rootCmd.Flags().Bool("lightweight", false, "Control whether to register a session with service manager or without")
	rootCmd.Flags().String("machine", "", "Operate on local container")
	rootCmd.Flags().String("nice", "", "Nice level")
	rootCmd.Flags().Bool("no-ask-password", false, "Do not prompt for password")
	rootCmd.Flags().Bool("pipe", false, "Request direct pipe for stdio")
	rootCmd.Flags().StringArray("property", nil, "Set service or scope unit property")
	rootCmd.Flags().Bool("pty", false, "Request allocation of a pseudo TTY for stdio")
	rootCmd.Flags().Bool("pty-late", false, "Just like --pty, but leave TTY access to agents until unit is started up")
	rootCmd.Flags().StringArray("setenv", nil, "Set environment variable")
	rootCmd.Flags().String("shell-prompt-prefix", "", "Set $SHELL_PROMPT_PREFIX")
	rootCmd.Flags().String("slice", "", "Run in the specified slice")
	rootCmd.Flags().Bool("slice-inherit", false, "Inherit the slice")
	rootCmd.Flags().String("unit", "", "Run under the specified unit name")
	rootCmd.Flags().StringP("user", "u", "", "Run as system user")
	rootCmd.Flags().BoolP("version", "V", false, "Show package version")
	rootCmd.Flags().Bool("via-shell", false, "Invoke command via target user's login shell")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"background": color.ActionAnsiBackgroundColors(false),
		"chdir":      carapace.ActionDirectories(),
		"group":      os.ActionGroups(),
		"setenv":     env.ActionNameValues(false),
		"user":       os.ActionUsers(),
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

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			envs, err := rootCmd.Flags().GetStringArray("setenv")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			for _, e := range envs {
				if name, value, ok := strings.Cut(e, "="); ok {
					c.Setenv(name, value)
				}
			}
			return action.Chdir(cmd.Flag("chdir").Value.String()).
				Invoke(c).ToA()
		})
	})
}
