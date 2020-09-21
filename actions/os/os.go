package os

import (
	"io/ioutil"
	"os"
	"strings"

	ps "github.com/mitchellh/go-ps"
	"github.com/rsteube/carapace"
)

func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		env := os.Environ()
		vars := make([]string, len(env))
		for index, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			vars[index] = pair[0]
		}
		return carapace.ActionValues(vars...)
	})
}

func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		groups := []string{}
		if content, err := ioutil.ReadFile("/etc/group"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				group := strings.Split(entry, ":")[0]
				if len(strings.TrimSpace(group)) > 0 {
					groups = append(groups, group)
				}
			}
		}
		return carapace.ActionValues(groups...)
	})
}

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

func ActionProcessExecutables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if processes, err := ps.Processes(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			executables := make([]string, 0)
			for _, process := range processes {
				executables = append(executables, process.Executable())
			}
			return carapace.ActionValues(executables...)
		}
	})
}

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

func ActionUsers() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		users := []string{}
		if content, err := ioutil.ReadFile("/etc/passwd"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				user := strings.Split(entry, ":")[0]
				if len(strings.TrimSpace(user)) > 0 {
					users = append(users, user)
				}
			}
		}
		return carapace.ActionValues(users...)
	})
}
