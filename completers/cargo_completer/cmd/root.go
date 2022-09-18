package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo",
	Short: "Rust's package manager",
	Long:  "https://doc.rust-lang.org/cargo/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	rootCmd.Flags().String("color", "", "Coloring: auto, always, never")
	rootCmd.Flags().String("explain", "", "Run `rustc --explain CODE`")
	rootCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().Bool("list", false, "List installed commands")
	rootCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	rootCmd.Flags().Bool("offline", false, "Run without accessing the network")
	rootCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.Flags().BoolP("version", "V", false, "Print version info and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if output, err := (carapace.Context{}).Command("cargo", "--list").Output(); err != nil {
			// TODO handle error
		} else {
			re := regexp.MustCompile(`^    (?P<command>\w+)( +(?P<description>.*))?$`)
			installedCommands := make(map[string]string)

			for _, line := range strings.Split(string(output), "\n") {
				if matches := re.FindStringSubmatch(line); matches != nil {
					installedCommands[matches[1]] = matches[2]
				}
			}

			subcommands := make(map[string]bool)
			for _, subcommand := range rootCmd.Commands() {
				subcommands[subcommand.Name()] = true
				if _, exists := installedCommands[subcommand.Name()]; !exists {
					subcommand.Hidden = true // hide from completion
				}
			}

			for name, description := range installedCommands {
				if _, ok := subcommands[name]; !ok {
					pluginCmd := &cobra.Command{
						Use:                name,
						Short:              description,
						Run:                func(cmd *cobra.Command, args []string) {},
						DisableFlagParsing: true,
					}

					carapace.Gen(pluginCmd).PositionalAnyCompletion(
						bridge.ActionCarapaceBin("cargo-" + name),
					)

					rootCmd.AddCommand(pluginCmd)
				}
			}
		}
	})
}
