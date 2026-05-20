package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/darwin/launchctl_completer/cmd/action"
	actionos "github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "launchctl",
	Short: "Manage launchd agents and daemons",
	Long:  "https://www.unix.com/man_page/osx/1/launchctl/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "domain", Title: "Domain Commands"},
		&cobra.Group{ID: "service", Title: "Service Commands"},
		&cobra.Group{ID: "legacy", Title: "Legacy Commands"},
		&cobra.Group{ID: "manager", Title: "Manager Commands"},
		&cobra.Group{ID: "other", Title: "Other Commands"},
	)

	addDomainCommand("bootstrap", "Bootstraps a domain or service into a domain")
	addDomainCommand("bootout", "Tears down a domain or removes a service from a domain")
	addServiceCommand("enable", "Enables an existing service")
	addServiceCommand("disable", "Disables an existing service")
	addServiceCommand("kickstart", "Forces an existing service to start")
	addServiceCommand("attach", "Attach the system's debugger to a service")
	addServiceCommand("debug", "Configures the next invocation of a service for debugging")
	addServiceCommand("blame", "Prints the reason a service is running")
	addServiceCommand("print", "Prints a description of a domain or service")

	addCommand("kill", "Sends a signal to the service instance", "service", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(
			ps.ActionKillSignals(),
			action.ActionServiceTargets(),
		)
	})
	addCommand("print-cache", "Prints information about the service cache", "service", nil)
	addCommand("print-disabled", "Prints which services are disabled", "service", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(action.ActionDomains())
	})
	addCommand("print-token", "Prints service identifier given an xpc event token", "service", nil)

	addCommand("plist", "Prints a property list embedded in a binary", "other", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(carapace.ActionFiles())
	})
	addCommand("procinfo", "Prints port information about a process", "other", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(ps.ActionProcessIds())
	})
	addCommand("hostinfo", "Prints port information about the host", "other", nil)
	addCommand("resolveport", "Resolves a port name from a process to an endpoint in launchd", "other", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(ps.ActionProcessIds())
	})
	addCommand("limit", "Reads or modifies launchd's resource limits", "other", nil)
	addCommand("examine", "Runs the specified analysis tool against launchd", "other", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(carapace.ActionExecutables())
	})
	addCommand("config", "Modifies persistent launchd domain configuration", "domain", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(
			carapace.ActionValues("system", "user"),
			carapace.ActionValues("umask", "path"),
		)
	})
	addCommand("dumpstate", "Dumps launchd state to stdout", "manager", nil)
	addCommand("dumpjpcategory", "Dumps jetsam properties category for all services", "manager", nil)
	addCommand("reboot", "Initiates a system reboot of the specified type", "manager", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(carapace.ActionValues("system", "userspace", "halt", "logout", "apps"))
	})
	addCommand("bootshell", "Brings the system up from single-user mode with a console shell", "manager", nil)

	addLegacyFileCommand("load", "Bootstraps a service or directory of services")
	addLegacyFileCommand("unload", "Unloads a service or directory of services")
	addLegacyServiceCommand("remove", "Unloads the specified service name")
	addLegacyServiceCommand("list", "Lists information about services")
	addLegacyServiceCommand("start", "Starts the specified service")
	addLegacyServiceCommand("stop", "Stops the specified service if it is running")

	addCommand("setenv", "Sets an environment variable for all services within the domain", "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(actionos.ActionEnvironmentVariables())
	})
	addCommand("unsetenv", "Unsets an environment variable from within launchd", "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(actionos.ActionEnvironmentVariables())
	})
	addCommand("getenv", "Gets the value of an environment variable from within launchd", "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(actionos.ActionEnvironmentVariables())
	})
	addCommand("bsexec", "Execute a program in another process' bootstrap context", "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(ps.ActionProcessIds(), carapace.ActionExecutables())
	})
	addCommand("asuser", "Execute a program in the bootstrap context of a given user", "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(action.ActionUserIds(), carapace.ActionExecutables())
	})
	addCommand("reload-atf", "Reloads the active trial factors", "legacy", nil)
	addCommand("submit", "Submit a basic job from the command line", "legacy", nil)
	addCommand("managerpid", "Prints the PID of the launchd controlling the session", "manager", nil)
	addCommand("manageruid", "Prints the UID of the current launchd session", "manager", nil)
	addCommand("managername", "Prints the name of the current launchd session", "manager", nil)
	addCommand("error", "Prints a description of an error", "other", nil)
	addCommand("variant", "Prints the launchd variant", "manager", nil)
	addCommand("version", "Prints the launchd version", "manager", nil)
	addCommand("help", "Prints the usage for a given subcommand", "other", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(actionSubcommands())
	})
}

func addCommand(name string, short string, groupID string, configure func(*cobra.Command)) {
	cmd := &cobra.Command{
		Use:     name,
		Short:   short,
		GroupID: groupID,
		Run:     func(cmd *cobra.Command, args []string) {},
	}
	rootCmd.AddCommand(cmd)
	if configure != nil {
		configure(cmd)
	}
}

func addDomainCommand(name string, short string) {
	addCommand(name, short, "domain", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(action.ActionDomains())
		carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles(".plist"))
	})
}

func addServiceCommand(name string, short string) {
	addCommand(name, short, "service", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(action.ActionServiceTargets())
	})
}

func addLegacyFileCommand(name string, short string) {
	addCommand(name, short, "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles(".plist"))
	})
}

func addLegacyServiceCommand(name string, short string) {
	addCommand(name, short, "legacy", func(cmd *cobra.Command) {
		carapace.Gen(cmd).PositionalCompletion(action.ActionCurrentServices())
	})
}

func actionSubcommands() carapace.Action {
	values := make([]string, 0)
	for _, cmd := range rootCmd.Commands() {
		values = append(values, cmd.Name(), cmd.Short)
	}
	return carapace.ActionValuesDescribed(values...)
}
