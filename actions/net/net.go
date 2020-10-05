package net

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/rsteube/carapace"
)

func ActionHosts() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		hosts := []string{}
		if file, err := homedir.Expand("~/.ssh/known_hosts"); err == nil {
			if content, err := ioutil.ReadFile(file); err == nil {
				r := regexp.MustCompile(`^(?P<host>[^ ,#]+)`)
				for _, entry := range strings.Split(string(content), "\n") {
					if r.MatchString(entry) {
						hosts = append(hosts, r.FindStringSubmatch(entry)[0])
					}
				}
			} else {
				return carapace.ActionMessage(err.Error())
			}
		}
		return carapace.ActionValues(hosts...)
	})
}

func ActionNetInterfaces() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("ifconfig").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			interfaces := []string{}
			for _, line := range strings.Split(string(output), "\n") {
				if matches, _ := regexp.MatchString("^[0-9a-zA-Z]", line); matches {
					interfaces = append(interfaces, strings.Split(line, ":")[0])
				}
			}
			return carapace.ActionValues(interfaces...)
		}
	})
}
