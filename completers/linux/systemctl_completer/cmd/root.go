package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/journalctl"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "systemctl",
	Short: "Query or send control commands to the system manager",
	Long:  "https://systemd.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "unit", Title: "Unit Commands"},
		&cobra.Group{ID: "unit file", Title: "Unit File Commands"},
		&cobra.Group{ID: "machine", Title: "Machine Commands"},
		&cobra.Group{ID: "job", Title: "Job Commands"},
		&cobra.Group{ID: "environment", Title: "Environment Commands"},
		&cobra.Group{ID: "manager state", Title: "Manager State Commands"},
		&cobra.Group{ID: "system", Title: "System Commands"},
	)

	rootCmd.PersistentFlags().StringS("P", "P", "", "Equivalent to --value --property=NAME")
	rootCmd.PersistentFlags().Bool("after", false, "Show units ordered after with 'list-dependencies'")
	rootCmd.PersistentFlags().BoolP("all", "a", false, "Show all properties/all units currently in memory, including dead/empty ones")
	rootCmd.PersistentFlags().Bool("before", false, "Show units ordered before with 'list-dependencies'")
	rootCmd.PersistentFlags().String("boot-loader-entry", "", "Boot into a specific boot loader entry on next boot")
	rootCmd.PersistentFlags().String("boot-loader-menu", "", "Boot into boot loader menu on next boot")
	rootCmd.PersistentFlags().StringP("capsule", "C", "", "Connect to service manager of specified capsule")
	rootCmd.PersistentFlags().String("check-inhibitors", "", "Specify if checking inhibitors before shutting down, sleeping or hibernating")
	rootCmd.PersistentFlags().String("drop-in", "", "Edit unit files using the specified drop-in file name")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Only print what would be done")
	rootCmd.PersistentFlags().Bool("failed", false, "Shortcut for --state=failed")
	rootCmd.PersistentFlags().Bool("firmware-setup", false, "Tell the firmware to show the setup menu on next boot")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "When enabling unit files, override existing symlinks When shutting down, execute action immediately")
	rootCmd.PersistentFlags().BoolP("full", "l", false, "Don't ellipsize unit names on output")
	rootCmd.PersistentFlags().Bool("global", false, "Enable/disable/mask default user unit files globally")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show this help")
	rootCmd.PersistentFlags().StringP("host", "H", "", "Operate on remote host")
	rootCmd.PersistentFlags().BoolS("i", "i", false, "Shortcut for --check-inhibitors=no")
	rootCmd.PersistentFlags().String("image", "", "Edit/enable/disable/mask unit files in the specified disk image")
	rootCmd.PersistentFlags().String("image-policy", "", "Specify disk image dissection policy")
	rootCmd.PersistentFlags().String("job-mode", "", "Specify how to deal with already queued jobs, when queueing a new job")
	rootCmd.PersistentFlags().String("kill-subgroup", "", "Send signal to sub-control group only")
	rootCmd.PersistentFlags().String("kill-value", "", "Signal value to enqueue")
	rootCmd.PersistentFlags().String("kill-whom", "", "Whom to send signal to")
	rootCmd.PersistentFlags().String("legend", "", "Enable/disable the legend (column headers and hints)")
	rootCmd.PersistentFlags().StringP("lines", "n", "", "Number of journal entries to show")
	rootCmd.PersistentFlags().StringP("machine", "M", "", "Operate on a local container")
	rootCmd.PersistentFlags().Bool("marked", false, "Restart/reload previously marked units")
	rootCmd.PersistentFlags().String("message", "", "Specify human-readable reason for system shutdown")
	rootCmd.PersistentFlags().Bool("mkdir", false, "Create directory before mounting, if missing")
	rootCmd.PersistentFlags().Bool("no-ask-password", false, "Do not ask for system passwords")
	rootCmd.PersistentFlags().Bool("no-block", false, "Do not wait until operation finished")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.PersistentFlags().Bool("no-reload", false, "Don't reload daemon after en-/dis-abling unit files")
	rootCmd.PersistentFlags().Bool("no-wall", false, "Don't send wall message before halt/power-off/reboot")
	rootCmd.PersistentFlags().Bool("no-warn", false, "Suppress several warnings shown by default")
	rootCmd.PersistentFlags().Bool("now", false, "Start or stop unit after enabling or disabling it")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Change journal output mode")
	rootCmd.PersistentFlags().Bool("plain", false, "Print unit dependencies as a list instead of a tree")
	rootCmd.PersistentFlags().String("preset-mode", "", "Apply only enable, only disable, or all presets")
	rootCmd.PersistentFlags().StringP("property", "p", "", "Show only properties by this name")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress output")
	rootCmd.PersistentFlags().Bool("read-only", false, "Create read-only bind mount")
	rootCmd.PersistentFlags().String("reboot-argument", "", "Specify argument string to pass to reboot()")
	rootCmd.PersistentFlags().BoolP("recursive", "r", false, "Show unit list of host and local containers")
	rootCmd.PersistentFlags().Bool("reverse", false, "Show reverse dependencies with 'list-dependencies'")
	rootCmd.PersistentFlags().String("root", "", "Enable/disable/mask unit files in the specified root directory")
	rootCmd.PersistentFlags().Bool("runtime", false, "Enable/disable/mask unit files temporarily until next reboot")
	rootCmd.PersistentFlags().BoolP("show-transaction", "T", false, "When enqueuing a unit job, show full transaction")
	rootCmd.PersistentFlags().Bool("show-types", false, "When showing sockets, explicitly show their type")
	rootCmd.PersistentFlags().StringP("signal", "s", "", "Which signal to send")
	rootCmd.PersistentFlags().String("state", "", "List units with particular LOAD or SUB or ACTIVE state")
	rootCmd.PersistentFlags().Bool("stdin", false, "Read new contents of edited file from stdin")
	rootCmd.PersistentFlags().Bool("system", false, "Connect to system manager")
	rootCmd.PersistentFlags().String("timestamp", "", "Change format of printed timestamps.")
	rootCmd.PersistentFlags().StringP("type", "t", "", "List units of a particular type")
	rootCmd.PersistentFlags().Bool("user", false, "Connect to user service manager")
	rootCmd.PersistentFlags().Bool("value", false, "When showing properties, only print the value")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Show unit logs while executing operation")
	rootCmd.PersistentFlags().Bool("version", false, "Show package version")
	rootCmd.PersistentFlags().Bool("wait", false, "wait until service stopped or startup completed")
	rootCmd.PersistentFlags().String("what", "", "Which types of resources to remove")
	rootCmd.PersistentFlags().String("when", "", "Schedule halt/power-off/reboot/kexec action after a certain timestamp")
	rootCmd.PersistentFlags().Bool("with-dependencies", false, "Show unit dependencies with 'status', 'cat', 'list-units', and 'list-unit-files'.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"P":                action.ActionProperties(rootCmd),
		"check-inhibitors": carapace.ActionValues("auto", "yes", "no").StyleF(style.ForKeyword),
		"host":             net.ActionHosts(),
		"job-mode": carapace.ActionValuesDescribed(
			"fail", "cause fail on conflict",
			"replace", "replace on conflict",
			"replace-irreversibly", "replace on conflict irreversibly",
			"isolate", "cause all other units to be stopped when starting unit",
			"ignore-dependencies", "ignore all unit dependencies",
			"ignore-requirements", "ignore all unit requirements",
			"flush", "cancel all queued jobs",
			"triggering", "stop unit and any active unit that may trigger it",
		),
		"kill-whom": carapace.ActionValuesDescribed(
			"main", "main process",
			"control", "control process",
			"cgroup", "all processes in the unit's control group",
			"all", "all processes",
		),
		"output":      journalctl.ActionOutputs(),
		"preset-mode": carapace.ActionValues("full", "enable-only", "disable-only"),
		"property":    action.ActionProperties(rootCmd),
		"root":        carapace.ActionDirectories(),
		"signal":      ps.ActionKillSignals(),
		"state":       systemctl.ActionStates(),
		"timestamp": carapace.ActionValuesDescribed(
			"pretty", "Day YYYY-MM-DD HH:MM:SS TZ",
			"us", "Day YYYY-MM-DD HH:MM:SS.UUUUUU TZ",
			"utc", "Day YYYY-MM-DD HH:MM:SS UTC",
			"us+utc", "Day YYYY-MM-DD HH:MM:SS.UUUUUU UTC",
		),
		"type": systemctl.ActionUnitTypes(),
		"what": carapace.ActionValues("configuration", "state", "cache", "logs", "runtime"),
	})
}
