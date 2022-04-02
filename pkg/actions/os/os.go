// package os contains operating system related actions
package os

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

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
		executables := make(map[string]string)

		for _, folder := range strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)) {
			if files, err := ioutil.ReadDir(folder); err == nil {
				for _, f := range files {
					if _, ok := executables[f.Name()]; !ok {
						if !f.IsDir() && isExecAny(f.Mode()) {
							executables[f.Name()] = style.ForPath(folder + "/" + f.Name())
						}
					}
				}
			}
		}

		vals := make([]string, 0)
		for executable, _style := range executables {
			vals = append(vals, executable, _style)
		}
		return carapace.ActionStyledValues(vals...)
	})
}

func isExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
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
