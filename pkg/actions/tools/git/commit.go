package git

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
)

type commit struct {
	Ref     string
	Message string
}

func commits(c carapace.Context, refOption RefOption) ([]commit, error) {
	if refOption.Commits <= 0 {
		return []commit{}, nil
	}

	if output, err := c.Command("git", "log", "--pretty=tformat:%h   %<(64,trunc)%s", "--max-count", strconv.Itoa(refOption.Commits)).Output(); err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(output), "\n")
		commits := make([]commit, 0)
		for index, line := range lines[:len(lines)-1] {
			if len(line) > 10 { // TODO duh?
				commits = append(commits, commit{line[:7], strings.TrimSpace(line[10:])})
				if index == 0 {
					commits = append(commits, commit{"HEAD", strings.TrimSpace(line[10:])}) // TOD fix this
				} else {
					commits = append(commits, commit{"HEAD~" + fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(refOption.Commits))-1)+"d", index), strings.TrimSpace(line[10:])}) // TOD fix this
				}
			}
		}
		return commits, nil
	}
}
