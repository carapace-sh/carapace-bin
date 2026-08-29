package openrc

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionServices completes available services
//
//	hostname
//	networking
func ActionServices() carapace.Action {
	return carapace.ActionExecCommand("rc-service", "-l")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("services")
}

// ActionRunlevels completes available runlevels
//
//	default
//	boot
func ActionRunlevels() carapace.Action {
	return carapace.ActionExecCommand("rc-status", "-l")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("runlevels")
}

// ActionServiceStates completes service states accepted by rc-status --in-state
//
//	started
//	stopped
func ActionServiceStates() carapace.Action {
	return carapace.ActionValues(
		"started",
		"stopped",
		"starting",
		"stopping",
		"inactive",
		"hotplugged",
		"failed",
		"crashed",
		"scheduled",
	).Tag("service states")
}

// ActionCommands completes commands accepted by openrc-run service scripts
//
//	start
//	stop
func ActionCommands() carapace.Action {
	return carapace.ActionValues(
		"describe",
		"help",
		"depend",
		"status",
		"ineed",
		"iuse",
		"iwant",
		"needsme",
		"usesme",
		"wantsme",
		"iafter",
		"ibefore",
		"iprovide",
		"start",
		"stop",
		"restart",
		"condrestart",
		"conditionalrestart",
		"zap",
	).Tag("service commands")
}

// ActionDependencyTypes completes dependency types accepted by rc-depend --type
//
//	ineed
//	iuse
func ActionDependencyTypes() carapace.Action {
	return carapace.ActionValues(
		"ineed",
		"iuse",
		"iwant",
		"needsme",
		"usesme",
		"wantsme",
		"iafter",
		"ibefore",
		"iprovide",
		"providedby",
		"keyword",
		"config",
		"reexport",
		"broken",
	).Tag("dependency types")
}
