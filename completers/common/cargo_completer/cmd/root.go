package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/cargo"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "cargo",
	Short: "Rust's package manager",
	Long:  "https://doc.rust-lang.org/cargo/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

const (
	group_build = iota
	group_manifest
	group_package
	group_publishing
	group_general
)

var groups = []*cobra.Group{
	{ID: "build", Title: "Build Commands"},
	{ID: "manifest", Title: "Manifest Commands"},
	{ID: "package", Title: "Package Commands"},
	{ID: "publishing", Title: "Publishing Commands"},
	{ID: "general", Title: "General Commands"},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(groups...)

	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "Change to DIRECTORY before doing anything (nightly-only)")
	rootCmd.PersistentFlags().StringSliceS("Z", "Z", nil, "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	rootCmd.PersistentFlags().String("color", "", "Coloring")
	rootCmd.PersistentFlags().StringSlice("config", nil, "Override a configuration value")
	rootCmd.Flags().String("explain", "", "Provide a detailed explanation of a rustc error message")
	rootCmd.PersistentFlags().Bool("frozen", false, "Equivalent to specifying both --locked and --offline")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("list", false, "List installed commands")
	rootCmd.PersistentFlags().Bool("locked", false, "Assert that `Cargo.lock` will remain unchanged")
	rootCmd.PersistentFlags().Bool("offline", false, "Run without accessing the network")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.Flags().BoolP("version", "V", false, "Print version info and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C":     carapace.ActionDirectories(),
		"Z":     cargo.ActionNightlyFlags(),
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		c := carapace.Context{Env: os.Environ()}    // TODO how to handle env here?
		c.Setenv("RUSTUP_DIST_SERVER", "localhost") // prevent implicit toolchain synching by the rustup wrapper [#1328]
		if output, err := c.Command("cargo", "--list").Output(); err == nil {
			re := regexp.MustCompile(`^    (?P<command>[^ ]+)( +(?P<description>.*))?$`)
			for _, line := range strings.Split(string(output), "\n") {
				if matches := re.FindStringSubmatch(line); matches != nil {
					pluginCmd := &cobra.Command{
						Use:                matches[1],
						Short:              matches[3],
						Run:                func(cmd *cobra.Command, args []string) {},
						GroupID:            groupFor(matches[1]),
						DisableFlagParsing: true,
					}

					carapace.Gen(pluginCmd).PositionalAnyCompletion(
						bridge.ActionCarapaceBin("cargo-" + matches[1]),
					)

					rootCmd.AddCommand(pluginCmd)
				}
			}
		}
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("C").Value.String())
	})
}

func groupFor(name string) string {
	// https: //doc.rust-lang.org/cargo/commands/index.html
	switch name {
	case "help",
		"version":
		return groups[group_general].ID
	case "bench",
		"build",
		"check",
		"clean",
		"doc",
		"fetch",
		"fix",
		"run",
		"rustc",
		"rustdoc",
		"test",
		"report":
		return groups[group_build].ID
	case "add",
		"generate-lockfile",
		"locate-project",
		"metadata",
		"pkgid",
		"remove",
		"tree",
		"update",
		"vendor",
		"verify-project":
		return groups[group_manifest].ID
	case "init",
		"install",
		"new",
		"search",
		"uninstall":
		return groups[group_package].ID
	case "login",
		"owner",
		"package",
		"publish",
		"yank":
		return groups[group_publishing].ID
	default:
		return ""
	}
}
