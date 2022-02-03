package os

import (
	"strconv"

	"github.com/mitchellh/go-ps"
	"github.com/rsteube/carapace"
)

// ActionKillSignals completes linux kill signals
//   ABRT (Abnormal termination)
//   STOP (Stop process, unblockable)
func ActionKillSignals() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ABRT", "Abnormal termination",
		"ALRM", "Virtual alarm clock",
		"BUS", "BUS error",
		"CHLD", "Child status has changed",
		"CONT", "Continue stopped process",
		"FPE", "Floating-point exception",
		"HUP", "Hangup detected on controlling terminal",
		"ILL", "Illegal instruction",
		"INT", "Interrupt from keyboard",
		"KILL", "Kill, unblockable",
		"PIPE", "Broken pipe",
		"POLL", "Pollable event occurred",
		"PROF", "Profiling alarm clock timer expired",
		"PWR", "Power failure restart",
		"QUIT", "Quit from keyboard",
		"SEGV", "Segmentation violation",
		"STKFLT", "Stack fault on coprocessor",
		"STOP", "Stop process, unblockable",
		"SYS", "Bad system call",
		"TERM", "Termination request",
		"TRAP", "Trace/breakpoint trap",
		"TSTP", "Stop typed at keyboard",
		"TTIN", "Background read from tty",
		"TTOU", "Background write to tty",
		"URG", "Urgent condition on socket",
		"USR1", "User-defined signal 1",
		"USR2", "User-defined signal 2",
		"VTALRM", "Virtual alarm clock",
		"WINCH", "Window size change",
		"XCPU", "CPU time limit exceeded",
		"XFSZ", "File size limit exceeded",
	)
}

// ActionProcessExecutables completes executable names of current processes
//   NetworkManager (439)
//   cupsd (454)
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
//   439 (NetworkManager)
//   454 (cupsd)
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
//   I (Idle kernel thread)
//   R (running or runnable on run queue)
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
