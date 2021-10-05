// package os contains operating system related actions
package os

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
	"github.com/rsteube/carapace"
)

// ActionEnvironmentVariables completes environment values
//   SHELL (/bin/elvish)
//   LANG (en_US.utf8)
func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		env := os.Environ()
		vars := make([]string, len(env)*2)
		for index, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			vars[index*2] = pair[0]
			vars[(index*2)+1] = pair[1]
		}
		return carapace.ActionValuesDescribed(vars...)
	})
}

// ActionGroups completes system group names
//    root (0)
//    ssh (101)
func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
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

// ActionGroupMembers completes system group members
//   root
//   daemon
func ActionGroupMembers(group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := []string{}
		if content, err := ioutil.ReadFile("/etc/group"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 3 &&
					splitted[0] == group {
					if len(strings.TrimSpace(group)) > 0 &&
						len(strings.TrimSpace(splitted[3])) > 0 {
						vals = append(vals, strings.Split(splitted[3], ",")...)
					}
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

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

// ActionUsers completes system user names
//   root (0)
//   daemon (1)
func ActionUsers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
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
//   bin:audio
//   lp:list
func ActionUserGroup() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionUsers().Invoke(c).Suffix(":").ToA()
		case 1:
			return ActionGroups()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionShells completes available terminal shells
//   /bin/elvish
//   /bin/bash
func ActionShells() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("chsh", "--list-shells")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

// ActionPathExecutables completes executable files from PATH
//   nvim
//   chmod
func ActionPathExecutables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		executables := make(map[string]bool)

		for _, folder := range strings.Split(os.Getenv("PATH"), ":") {
			if files, err := ioutil.ReadDir(folder); err == nil {
				for _, f := range files {
					if f.Mode().IsRegular() && isExecAny(f.Mode()) {
						executables[f.Name()] = true
					}
				}
			}
		}

		vals := make([]string, 0)
		for executable := range executables {
			vals = append(vals, executable)
		}
		return carapace.ActionValues(vals...)
	})
}

func isExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

// ActionDisplays completes x displays
//   :0
//   :1
func ActionDisplays() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("w", "-hsu")(func(output []byte) carapace.Action {
			vals := make([]string, 0)
			r := regexp.MustCompile("/usr/lib/Xorg (?P<display>:[0-9]+)")
			for _, line := range strings.Split(string(output), "\n") {
				if r.MatchString(line) {
					vals = append(vals, r.FindStringSubmatch(line)[1])
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

// ActionCgroups completes cgroup names
//   blkio
//   cpu
func ActionCgroups() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cpu", `Cgroups  can  be  guaranteed  a minimum number of "CPU shares" when a system is busy.`,
		"cpuacct", "This provides accounting for CPU usage by groups of processes.",
		"cpuset", "This cgroup can be used to bind the processes in a cgroup to a specified set of  CPUs and NUMA nodes.",
		"memory", "The  memory controller supports reporting and limiting of process memory, kernel memory, and swap used by cgroups.",
		"devices", "This supports controlling which processes may create",
		"freezer", "The  freezer  cgroup  can  suspend  and  restore",
		"net_cls", "This  places  a  classid,  specified  for the cgroup, on network packets created by a cgroup.",
		"blkio", "The blkio cgroup controls and limits access to specified block devices.",
		"perf_event", "This controller allows perf monitoring of the set of processes grouped in a cgroup.",
		"net_prio", "This allows priorities to be specified, per network interface, for cgroups.",
		"hugetlb", "This supports limiting the use of huge pages by cgroups.",
		"pids", "This controller permits limiting the number of process  that  may  be  created  in  a cgroup",
		"rdma", "The  RDMA  controller  permits  limiting  the  use  of RDMA/IB-specific resources per cgroup.",
	)
}
