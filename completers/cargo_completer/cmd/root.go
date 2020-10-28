package cmd

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo",
	Short: "Rust's package manager",
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
		"color": carapace.ActionValues("auto", "always", "never"),
	})

	c, _, _ := rootCmd.Find([]string{"_carapace"})
	c.PreRun = func(cmd *cobra.Command, args []string) {
		if len(args) == 1 { // non-lazy completion script generation
			if output, err := exec.Command("cargo", "--list").Output(); err != nil {
				// TODO handle error
			} else {
				re := regexp.MustCompile(`^    (?P<command>\w+) +(?P<description>.*)$`)
				installedCommands := make(map[string]bool)

				for _, line := range strings.Split(string(output), "\n") {
					if re.MatchString(line) {
						installedCommands[re.FindStringSubmatch(line)[1]] = true
					}
				}

				for _, subcommand := range rootCmd.Commands() {
					if _, exists := installedCommands[subcommand.Name()]; !exists {
						subcommand.Hidden = true // hide from completion
					}
				}
			}
		}
	}
}
