package git

import (
	exec "golang.org/x/sys/execabs"
	"strings"
)

type tag struct {
	Name    string
	Message string
}

func tags(refOption RefOption) ([]tag, error) {
	if !refOption.Tags {
		return []tag{}, nil
	}

	if output, err := exec.Command("git", "tag", "--format", "%(refname)\n%(subject)").Output(); err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(output), "\n")
		tags := make([]tag, len(lines)/2)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 {
				tags[index/2] = tag{strings.TrimPrefix(line, "refs/tags/"), lines[index+1]}
			}
		}
		return tags, err
	}
}
