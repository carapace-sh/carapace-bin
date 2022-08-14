// package os contains process related actions
package ps

import (
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/rsteube/carapace/third_party/github.com/mitchellh/go-ps"
)

// ActionKillSignals completes linux kill signals
//
//	ABRT (Abnormal termination)
//	STOP (Stop process, unblockable)
func ActionKillSignals() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionStyledValuesDescribed(
			"ABRT", "Abnormal termination", styles.CarapaceBin.KillSignalCore,
			"ALRM", "Virtual alarm clock", styles.CarapaceBin.KillSignalTerm,
			"BUS", "BUS error", styles.CarapaceBin.KillSignalCore,
			"CHLD", "Child status has changed", styles.CarapaceBin.KillSignalIgn,
			"CONT", "Continue stopped process", styles.CarapaceBin.KillSignalCont,
			"FPE", "Floating-point exception", styles.CarapaceBin.KillSignalCore,
			"HUP", "Hangup detected on controlling terminal", styles.CarapaceBin.KillSignalTerm,
			"ILL", "Illegal instruction", styles.CarapaceBin.KillSignalCore,
			"INT", "Interrupt from keyboard", styles.CarapaceBin.KillSignalTerm,
			"KILL", "Kill, unblockable", styles.CarapaceBin.KillSignalTerm,
			"PIPE", "Broken pipe", styles.CarapaceBin.KillSignalTerm,
			"POLL", "Pollable event occurred", styles.CarapaceBin.KillSignalTerm,
			"PROF", "Profiling alarm clock timer expired", styles.CarapaceBin.KillSignalTerm,
			"PWR", "Power failure restart", styles.CarapaceBin.KillSignalTerm,
			"QUIT", "Quit from keyboard", styles.CarapaceBin.KillSignalCore,
			"SEGV", "Segmentation violation", styles.CarapaceBin.KillSignalCore,
			"STKFLT", "Stack fault on coprocessor", styles.CarapaceBin.KillSignalTerm,
			"STOP", "Stop process, unblockable", styles.CarapaceBin.KillSignalStop,
			"SYS", "Bad system call", styles.CarapaceBin.KillSignalCore,
			"TERM", "Termination request", styles.CarapaceBin.KillSignalTerm,
			"TRAP", "Trace/breakpoint trap", styles.CarapaceBin.KillSignalCore,
			"TSTP", "Stop typed at keyboard", styles.CarapaceBin.KillSignalStop,
			"TTIN", "Background read from tty", styles.CarapaceBin.KillSignalStop,
			"TTOU", "Background write to tty", styles.CarapaceBin.KillSignalStop,
			"URG", "Urgent condition on socket", styles.CarapaceBin.KillSignalIgn,
			"USR1", "User-defined signal 1", styles.CarapaceBin.KillSignalTerm,
			"USR2", "User-defined signal 2", styles.CarapaceBin.KillSignalTerm,
			"VTALRM", "Virtual alarm clock", styles.CarapaceBin.KillSignalTerm,
			"WINCH", "Window size change", styles.CarapaceBin.KillSignalIgn,
			"XCPU", "CPU time limit exceeded", styles.CarapaceBin.KillSignalCore,
			"XFSZ", "File size limit exceeded", styles.CarapaceBin.KillSignalCore,
		)
	})
}

// ActionProcessExecutables completes executable names of current processes
//
//	NetworkManager (439)
//	cupsd (454)
func ActionProcessExecutables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if processes, err := ps.Processes(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			executables := make([]string, 0)
			for _, process := range processes {
				executables = append(executables, process.Executable(), strconv.Itoa(process.Pid()))
			}
			return carapace.ActionValuesDescribed(executables...)
		}
	})
}

// ActionProcessIds completes proces IDs
//
//	439 (NetworkManager)
//	454 (cupsd)
func ActionProcessIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if processes, err := ps.Processes(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			ids := make([]string, 0)
			for _, process := range processes {
				ids = append(ids, strconv.Itoa(process.Pid()), process.Executable())
			}
			return carapace.ActionValuesDescribed(ids...)
		}
	})
}

// ActionProcessStates completes linux process states
//
//	I (Idle kernel thread)
//	R (running or runnable on run queue)
func ActionProcessStates() carapace.Action {
	return carapace.ActionValuesDescribed(
		"D", "uninterruptible sleep (usually IO)",
		"I", "Idle kernel thread",
		"R", "running or runnable (on run queue)",
		"S", "interruptible sleep (waiting for an event to complete)",
		"T", "stopped by job control signal",
		"W", "paging (not valid since the 2.6.xx kernel)",
		"X", "dead (should never be seen)",
		"Z", "defunct (zombie) process, terminated but not reaped by its parent",
		"t", "stopped by debugger during the tracing",
	)
}
