package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "systemd-analyze",
	Short: "Analyze and debug system manager",
	Long:  "https://www.freedesktop.org/software/systemd/man/latest/systemd-analyze.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("base-time", "", "Calculate calendar times relative to specified time")
	rootCmd.PersistentFlags().String("debugger", "", "Use the given debugger")
	rootCmd.PersistentFlags().StringP("debugger-arguments", "A", "", "Pass the given arguments to the debugger")
	rootCmd.PersistentFlags().Bool("detailed", false, "Add more details to SVG plot, e.g. show activation timestamps")
	rootCmd.PersistentFlags().String("drm-device", "", "Use this DRM device sysfs path to get EDID")
	rootCmd.PersistentFlags().String("from-pattern", "", "Show only origins in the graph")
	rootCmd.PersistentFlags().String("fuzz", "", "Also print services which finished SECONDS earlier than the latest in the branch")
	rootCmd.PersistentFlags().String("generators", "", "Do [not] run unit generators (requires privileges)")
	rootCmd.PersistentFlags().Bool("global", false, "Operate on global user configuration")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show this help")
	rootCmd.PersistentFlags().StringP("host", "H", "", "Operate on remote host")
	rootCmd.PersistentFlags().String("image", "", "Operate on disk image as filesystem root")
	rootCmd.PersistentFlags().String("image-policy", "", "Specify disk image dissection policy")
	rootCmd.PersistentFlags().String("instance", "", "Specify fallback instance name for template units")
	rootCmd.PersistentFlags().String("iterations", "", "Show the specified number of iterations")
	rootCmd.PersistentFlags().String("json", "", "Generate JSON output of the security analysis table, or plot's raw time data")
	rootCmd.PersistentFlags().StringP("machine", "M", "", "Operate on local container")
	rootCmd.PersistentFlags().String("man", "", "Do [not] check for existence of man pages")
	rootCmd.PersistentFlags().BoolP("mask", "m", false, "Parse parameter as numeric capability mask")
	rootCmd.PersistentFlags().Bool("no-legend", false, "Disable column headers and hints in plot with either --table or --json=")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.PersistentFlags().String("offline", "", "Perform a security review on unit file(s)")
	rootCmd.PersistentFlags().Bool("order", false, "Show only order in the graph")
	rootCmd.PersistentFlags().String("profile", "", "Include the specified profile in the security review of the unit(s)")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Do not emit hints")
	rootCmd.PersistentFlags().String("recursive-errors", "", "Control which units are verified")
	rootCmd.PersistentFlags().Bool("require", false, "Show only requirement in the graph")
	rootCmd.PersistentFlags().String("root", "", "Operate on an alternate filesystem root")
	rootCmd.PersistentFlags().String("scale-svg", "", "Stretch x-axis of plot by FACTOR")
	rootCmd.PersistentFlags().String("security-policy", "", "Use custom JSON security policy instead of built-in one")
	rootCmd.PersistentFlags().Bool("system", false, "Operate on system systemd instance")
	rootCmd.PersistentFlags().Bool("table", false, "Output plot's raw time data as a table")
	rootCmd.PersistentFlags().String("threshold", "", "Exit with a non-zero status when overall exposure level is over threshold value")
	rootCmd.PersistentFlags().Bool("tldr", false, "Skip comments and empty lines")
	rootCmd.PersistentFlags().String("to-pattern", "", "Show only destinations in the graph")
	rootCmd.PersistentFlags().String("unit", "", "Evaluate conditions and asserts of unit")
	rootCmd.PersistentFlags().Bool("user", false, "Operate on user systemd instance")
	rootCmd.PersistentFlags().Bool("version", false, "Show package version")

	rootCmd.Flag("generators").NoOptDefVal = " "
	rootCmd.Flag("man").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"drm-device":       carapace.ActionDirectories(),
		"generators":       carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"image":            carapace.ActionFiles(),
		"json":             carapace.ActionValues("pretty", "short", "off"),
		"man":              carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"offline":          carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"profile":          carapace.ActionFiles(),
		"recursive-errors": carapace.ActionValues("yes", "no", "one").StyleF(style.ForKeyword),
		"root":             carapace.ActionDirectories(),
		"security-policy":  carapace.ActionFiles(),
		"unit": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return systemctl.ActionUnits(systemctl.UnitOpts{User: rootCmd.Flag("user").Changed, Active: true, Inactive: true})
		}),
	})
}
