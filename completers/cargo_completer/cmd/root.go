package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(groups...)

	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringSliceP("Z", "Z", []string{}, "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	rootCmd.PersistentFlags().String("color", "", "Coloring: auto, always, never")
	rootCmd.PersistentFlags().StringSlice("config", []string{}, "Override a configuration value")
	rootCmd.Flags().String("explain", "", "Run `rustc --explain CODE`")
	rootCmd.PersistentFlags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("list", false, "List installed commands")
	rootCmd.PersistentFlags().Bool("locked", false, "Require Cargo.lock is up to date")
	rootCmd.PersistentFlags().Bool("offline", false, "Run without accessing the network")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.Flags().BoolP("version", "V", false, "Print version info and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
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
