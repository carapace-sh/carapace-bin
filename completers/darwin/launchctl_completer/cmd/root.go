package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/darwin/launchctl_completer/cmd/actions"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "launchctl",
	Short: "Interfaces with launchd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.AddCommand(
		simpleCommand("bootstrap", "Bootstraps a domain or service", domainFilesCompletion),
		simpleCommand("bootout", "Tears down a domain or service", domainServiceCompletion),
		simpleCommand("enable", "Enables a service", serviceTargetCompletion),
		simpleCommand("disable", "Disables a service", serviceTargetCompletion),
		simpleCommand("kickstart", "Forces an existing service to start", serviceTargetCompletion),
		simpleCommand("kill", "Sends a signal to a service", signalServiceCompletion),
		simpleCommand("print", "Prints a domain or service", domainServiceCompletion),
		simpleCommand("print-cache", "Prints launchd cache information", nil),
		simpleCommand("print-disabled", "Prints disabled services", actions.ActionDomains),
		simpleCommand("procinfo", "Prints launchd information for a process", ps.ActionProcessIds),
		simpleCommand("blame", "Prints service blame information", serviceTargetCompletion),
		simpleCommand("list", "Lists loaded services", nil),
		simpleCommand("start", "Starts a loaded service", actions.ActionServices),
		simpleCommand("stop", "Stops a loaded service", actions.ActionServices),
		simpleCommand("remove", "Removes a loaded service", actions.ActionServices),
		simpleCommand("load", "Loads configuration files", launchdPlistCompletion),
		simpleCommand("unload", "Unloads configuration files", launchdPlistCompletion),
	)
}

func simpleCommand(use, short string, completion func() carapace.Action) *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: func(cmd *cobra.Command, args []string) {}}
	carapace.Gen(cmd).Standalone()
	if completion != nil {
		carapace.Gen(cmd).PositionalAnyCompletion(completion())
	}
	return cmd
}

func launchdPlistCompletion() carapace.Action {
	return carapace.ActionFiles(".plist")
}

func domainFilesCompletion() carapace.Action {
	return carapace.Batch(
		actions.ActionDomains(),
		carapace.ActionFiles(".plist"),
	).ToA()
}

func domainServiceCompletion() carapace.Action {
	return carapace.Batch(
		actions.ActionDomains(),
		actions.ActionServices(),
	).ToA()
}

func serviceTargetCompletion() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return actions.ActionServices().Invoke(c).Prefix("gui/").ToA()
	})
}

func signalServiceCompletion() carapace.Action {
	return carapace.Batch(
		carapace.ActionValues("SIGTERM", "SIGKILL", "SIGHUP", "SIGINT", "TERM", "KILL", "HUP", "INT"),
		serviceTargetCompletion(),
	).ToA()
}
