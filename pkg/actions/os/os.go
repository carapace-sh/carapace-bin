package os

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
	"github.com/rsteube/carapace"
)

// ActionEnvironmentVariables completes environment values
func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		env := os.Environ()
		vars := make([]string, len(env)*2)
		for index, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			vars[index*2] = pair[0]
			if len(pair[1]) > 40 {
				vars[(index*2)+1] = pair[1][:37] + "..."
			} else {
				vars[(index*2)+1] = pair[1]
			}
		}
		return carapace.ActionValuesDescribed(vars...)
	})
}

// ActionGroups completes system group names
func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		groups := []string{}
		if content, err := ioutil.ReadFile("/etc/group"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 2 {
					group := splitted[0]
					id := splitted[2]
					if len(strings.TrimSpace(group)) > 0 {
						groups = append(groups, group, id)
					}
				}
			}
		}
		return carapace.ActionValuesDescribed(groups...)
	})
}

// ActionKillSignals completes linux kill signals
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
func ActionProcessExecutables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
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

// ActionProcessStates completes linux process states
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

// ActionUsers completes system user names
func ActionUsers() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		users := []string{}
		if content, err := ioutil.ReadFile("/etc/passwd"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 2 {
					user := splitted[0]
					id := splitted[2]
					if len(strings.TrimSpace(user)) > 0 {
						users = append(users, user, id)
					}
				}
			}
		}
		return carapace.ActionValuesDescribed(users...)
	})
}

// ActionUserGroup completes system user:group separately
func ActionUserGroup() carapace.Action {
	return carapace.ActionMultiParts(":", func(args []string, parts []string) carapace.Action {
		switch len(parts) {
		case 0:
			return ActionUsers().Invoke(args).Suffix(":").ToA()
		case 1:
			return ActionGroups()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionShells completes available terminal shells
func ActionShells() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("chsh", "--list-shells").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")...)
		}
	})
}
