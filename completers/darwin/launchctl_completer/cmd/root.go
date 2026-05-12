package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "launchctl",
	Short: "interfaces with launchd",
	Long:  "https://keith.github.io/xcode-man-pages/launchctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(attachCmd)
	rootCmd.AddCommand(blameCmd)
	rootCmd.AddCommand(bootstrapCmd)
	rootCmd.AddCommand(bootoutCmd)
	rootCmd.AddCommand(bootshellCmd)
	rootCmd.AddCommand(bsexecCmd)
	rootCmd.AddCommand(bslistCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(debugCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(dumpstateCmd)
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(errorCmd)
	rootCmd.AddCommand(examineCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(getenvCmd)
	rootCmd.AddCommand(getrusageCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(hostinfoCmd)
	rootCmd.AddCommand(kickstartCmd)
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(limitCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(managernameCmd)
	rootCmd.AddCommand(managerpidCmd)
	rootCmd.AddCommand(manageruidCmd)
	rootCmd.AddCommand(procinfoCmd)
	rootCmd.AddCommand(printCacheCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(printDisabledCmd)
	rootCmd.AddCommand(rebootCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(resolveportCmd)
	rootCmd.AddCommand(runstatsCmd)
	rootCmd.AddCommand(setenvCmd)
	rootCmd.AddCommand(shutdownCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(umaskCmd)
	rootCmd.AddCommand(unloadCmd)
	rootCmd.AddCommand(unsetenvCmd)
	rootCmd.AddCommand(variantCmd)
	rootCmd.AddCommand(versionCmd)
}

var attachCmd = newCommand("attach", "attach a debugger to a service")
var blameCmd = newCommand("blame", "show why a service launched")
var bootstrapCmd = newCommand("bootstrap", "bootstrap domains and services")
var bootoutCmd = newCommand("bootout", "remove domains and services")
var bootshellCmd = newCommand("bootshell", "boot into a shell")
var bsexecCmd = newCommand("bsexec", "execute a command in a bootstrap context")
var bslistCmd = newCommand("bslist", "list Mach bootstrap services")
var configCmd = newCommand("config", "modify launchd configuration")
var debugCmd = newCommand("debug", "configure the next service invocation for debugging")
var disableCmd = newCommand("disable", "disable a service")
var dumpstateCmd = newCommand("dumpstate", "dump launchd state")
var enableCmd = newCommand("enable", "enable a service")
var errorCmd = newCommand("error", "print a launchd error string")
var examineCmd = newCommand("examine", "examine launchd internals")
var exportCmd = newCommand("export", "export launchd environment variables")
var getenvCmd = newCommand("getenv", "get a launchd environment variable")
var getrusageCmd = newCommand("getrusage", "get resource usage statistics")
var helpCmd = newCommand("help", "display help")
var hostinfoCmd = newCommand("hostinfo", "print host information")
var kickstartCmd = newCommand("kickstart", "start a service immediately")
var killCmd = newCommand("kill", "send a signal to a service")
var limitCmd = newCommand("limit", "get or set launchd resource limits")
var listCmd = newCommand("list", "list loaded jobs")
var loadCmd = newCommand("load", "load configuration files")
var logCmd = newCommand("log", "get or set syslog log levels")
var managernameCmd = newCommand("managername", "print manager name")
var managerpidCmd = newCommand("managerpid", "print manager PID")
var manageruidCmd = newCommand("manageruid", "print manager UID")
var procinfoCmd = newCommand("procinfo", "print process information")
var printCacheCmd = newCommand("print-cache", "print service cache")
var printCmd = newCommand("print", "print a domain or service")
var printDisabledCmd = newCommand("print-disabled", "print disabled services")
var rebootCmd = newCommand("reboot", "reboot, halt, or restart userspace")
var removeCmd = newCommand("remove", "remove a job by label")
var resolveportCmd = newCommand("resolveport", "resolve a Mach port")
var runstatsCmd = newCommand("runstats", "print service run statistics")
var setenvCmd = newCommand("setenv", "set a launchd environment variable")
var shutdownCmd = newCommand("shutdown", "prepare launchd for shutdown")
var startCmd = newCommand("start", "start a job by label")
var stopCmd = newCommand("stop", "stop a job by label")
var submitCmd = newCommand("submit", "submit a program without a plist")
var umaskCmd = newCommand("umask", "get or set launchd umask")
var unloadCmd = newCommand("unload", "unload configuration files")
var unsetenvCmd = newCommand("unsetenv", "unset a launchd environment variable")
var variantCmd = newCommand("variant", "print launchd variant")
var versionCmd = newCommand("version", "print launchd version")

func init() {
	addBootstrapFlags(bootstrapCmd)
	addBootstrapFlags(bootoutCmd)
	carapace.Gen(bootstrapCmd).PositionalCompletion(actionDomains())
	carapace.Gen(bootstrapCmd).PositionalAnyCompletion(actionServiceFiles())
	carapace.Gen(bootoutCmd).PositionalCompletion(actionTargets())
	carapace.Gen(bootoutCmd).PositionalAnyCompletion(actionServiceFiles())

	attachCmd.Flags().BoolS("k", "k", false, "kill the running instance")
	attachCmd.Flags().BoolS("s", "s", false, "force the service to start")
	attachCmd.Flags().BoolS("x", "x", false, "attach to xpcproxy before exec")
	carapace.Gen(attachCmd).PositionalCompletion(actionServiceTargets())

	debugCmd.Flags().BoolS("32", "32", false, "run service in a 32-bit architecture")
	debugCmd.Flags().BoolS("NSZombie", "NSZombie", false, "enable NSZombie")
	debugCmd.Flags().BoolS("debug-libraries", "debug-libraries", false, "prefer debug library variants")
	debugCmd.Flags().StringS("environment", "environment", "", "set environment variables")
	debugCmd.Flags().BoolS("guard-malloc", "guard-malloc", false, "enable libgmalloc")
	debugCmd.Flags().BoolS("introspection-libraries", "introspection-libraries", false, "prefer introspection libraries")
	debugCmd.Flags().BoolS("malloc-nano-allocator", "malloc-nano-allocator", false, "enable nano allocator")
	debugCmd.Flags().BoolS("malloc-stack-logging", "malloc-stack-logging", false, "enable malloc stack logging")
	debugCmd.Flags().StringS("program", "program", "", "use a different executable")
	debugCmd.Flags().StringS("stderr", "stderr", "", "set standard error path")
	debugCmd.Flags().StringS("stdin", "stdin", "", "set standard input path")
	debugCmd.Flags().StringS("stdout", "stdout", "", "set standard output path")
	carapace.Gen(debugCmd).FlagCompletion(carapace.ActionMap{
		"program": bridge.ActionCarapaceBin(),
		"stderr":  carapace.ActionFiles(),
		"stdin":   carapace.ActionFiles(),
		"stdout":  carapace.ActionFiles(),
	})
	carapace.Gen(debugCmd).PositionalCompletion(actionServiceTargets())
	carapace.Gen(debugCmd).PositionalAnyCompletion(bridge.ActionCarapaceBin())

	kickstartCmd.Flags().BoolS("k", "k", false, "kill the running instance before restarting")
	kickstartCmd.Flags().BoolS("p", "p", false, "print the PID of the service")
	carapace.Gen(kickstartCmd).PositionalCompletion(actionServiceTargets())

	carapace.Gen(blameCmd).PositionalCompletion(actionServiceTargets())
	carapace.Gen(disableCmd).PositionalCompletion(actionServiceTargets())
	carapace.Gen(enableCmd).PositionalCompletion(actionServiceTargets())
	carapace.Gen(killCmd).PositionalCompletion(ps.ActionKillSignals(), actionServiceTargets())
	carapace.Gen(printCmd).PositionalCompletion(actionTargets())
	carapace.Gen(printCacheCmd).PositionalCompletion(actionDomains())
	carapace.Gen(printDisabledCmd).PositionalCompletion(actionDomains())
	carapace.Gen(runstatsCmd).PositionalCompletion(actionServiceTargets())

	carapace.Gen(configCmd).PositionalCompletion(
		carapace.ActionValues("system", "user"),
		carapace.ActionValuesDescribed(
			"umask", "default file creation mask",
			"path", "default PATH environment variable",
			"maxfiles", "maximum number of open files",
			"maxproc", "maximum number of processes",
		),
	)

	carapace.Gen(dumpstateCmd).PositionalCompletion(actionDomains())
	carapace.Gen(errorCmd).PositionalCompletion(carapace.ActionValues())
	carapace.Gen(getenvCmd).PositionalCompletion(carapace.ActionValues())
	carapace.Gen(getrusageCmd).PositionalCompletion(carapace.ActionValues("self", "children"))
	carapace.Gen(limitCmd).PositionalCompletion(actionLimits())
	carapace.Gen(limitCmd).PositionalAnyCompletion(carapace.ActionValues("both", "hard", "soft", "unlimited"))
	carapace.Gen(logCmd).PositionalCompletion(
		carapace.ActionValues("level", "only", "mask"),
		actionLogLevels(),
	)
	carapace.Gen(logCmd).PositionalAnyCompletion(actionLogLevels())
	carapace.Gen(procinfoCmd).PositionalCompletion(ps.ActionProcessIds())
	carapace.Gen(rebootCmd).PositionalCompletion(carapace.ActionValues("apps", "halt", "logout", "reroot", "system", "userspace"))
	carapace.Gen(resolveportCmd).PositionalCompletion(ps.ActionProcessIds())
	carapace.Gen(umaskCmd).PositionalCompletion(carapace.ActionValues())

	bslistCmd.Flags().BoolS("j", "j", false, "show the job that registered each service")
	carapace.Gen(bslistCmd).PositionalCompletion(ps.ActionProcessIds(), carapace.ActionValues(".."))
	carapace.Gen(bsexecCmd).PositionalCompletion(ps.ActionProcessIds(), bridge.ActionCarapaceBin())
	carapace.Gen(bsexecCmd).PositionalAnyCompletion(bridge.ActionCarapaceBin())

	loadCmd.Flags().StringS("D", "D", "", "look for plists in a domain")
	loadCmd.Flags().BoolS("F", "F", false, "force loading the plist")
	loadCmd.Flags().StringS("S", "S", "", "restrict jobs to a session type")
	loadCmd.Flags().BoolS("w", "w", false, "override the Disabled key")
	carapace.Gen(loadCmd).FlagCompletion(launchdSearchCompletions())
	carapace.Gen(loadCmd).PositionalAnyCompletion(actionServiceFiles())

	unloadCmd.Flags().StringS("D", "D", "", "look for plists in a domain")
	unloadCmd.Flags().StringS("S", "S", "", "restrict jobs to a session type")
	unloadCmd.Flags().BoolS("w", "w", false, "override the Disabled key")
	carapace.Gen(unloadCmd).FlagCompletion(launchdSearchCompletions())
	carapace.Gen(unloadCmd).PositionalAnyCompletion(actionServiceFiles())

	listCmd.Flags().BoolS("x", "x", false, "output as an XML property list")
	carapace.Gen(listCmd).PositionalCompletion(actionLabels())
	carapace.Gen(removeCmd).PositionalCompletion(actionLabels())
	carapace.Gen(startCmd).PositionalCompletion(actionLabels())
	carapace.Gen(stopCmd).PositionalCompletion(actionLabels())

	submitCmd.Flags().StringS("e", "e", "", "stderr path")
	submitCmd.Flags().StringS("l", "l", "", "job label")
	submitCmd.Flags().StringS("o", "o", "", "stdout path")
	submitCmd.Flags().StringS("p", "p", "", "program to execute")
	carapace.Gen(submitCmd).FlagCompletion(carapace.ActionMap{
		"e": carapace.ActionFiles(),
		"o": carapace.ActionFiles(),
		"p": bridge.ActionCarapaceBin(),
	})
	carapace.Gen(submitCmd).PositionalCompletion(bridge.ActionCarapaceBin())
	carapace.Gen(submitCmd).PositionalAnyCompletion(bridge.ActionCarapaceBin())
}

func newCommand(use string, short string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()
	return cmd
}

func addBootstrapFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("D", "D", "", "look for plists in a domain")
	cmd.Flags().BoolS("F", "F", false, "force bootstrapping services")
	cmd.Flags().StringS("S", "S", "", "restrict jobs to a session type")
	carapace.Gen(cmd).FlagCompletion(launchdSearchCompletions())
}

func launchdSearchCompletions() carapace.ActionMap {
	return carapace.ActionMap{
		"D": actionSearchDomains(),
		"S": actionSessionTypes(),
	}
}

func actionDomains() carapace.Action {
	return carapace.ActionValues("system/", "user/", "gui/", "login/", "pid/").NoSpace()
}

func actionServiceTargets() carapace.Action {
	return carapace.Batch(
		actionDomains(),
		actionLabels(),
	).ToA()
}

func actionTargets() carapace.Action {
	return carapace.Batch(
		actionDomains(),
		actionServiceTargets(),
	).ToA()
}

func actionServiceFiles() carapace.Action {
	return carapace.Batch(
		carapace.ActionFiles(".plist"),
		carapace.ActionDirectories(),
	).ToA()
}

func actionLabels() carapace.Action {
	return carapace.ActionExecCommand("launchctl", "list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			fields := strings.Fields(line)
			if len(fields) < 3 || fields[0] == "PID" {
				continue
			}
			vals = append(vals, fields[2])
		}
		return carapace.ActionValues(vals...)
	})
}

func actionLimits() carapace.Action {
	return carapace.ActionValues(
		"core",
		"cpu",
		"data",
		"filesize",
		"maxfiles",
		"maxproc",
		"memlock",
		"rss",
		"stack",
	)
}

func actionLogLevels() carapace.Action {
	return carapace.ActionValues("alert", "critical", "debug", "emergency", "error", "info", "notice", "warning")
}

func actionSearchDomains() carapace.Action {
	return carapace.ActionValues("all", "local", "network", "system", "user")
}

func actionSessionTypes() carapace.Action {
	return carapace.ActionValues("Aqua", "Background", "LoginWindow", "StandardIO", "System")
}
