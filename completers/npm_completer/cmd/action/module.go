package action

import (
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func nodeModulesPath() (path string, err error) {
	var wd string
	if wd, err = os.Getwd(); err == nil {
		path, err = util.FindReverse(wd, "node_modules")
	}
	return
}

func ActionModules() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		path, err := nodeModulesPath()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		fullSegments := make([]string, 0)
		for _, part := range c.Parts {
			if strings.HasPrefix(part, "@") {
				fullSegments = append(fullSegments, part)
			} else {
				fullSegments = append(fullSegments, part, "node_modules")
			}
		}

		contents, err := os.ReadDir(fmt.Sprintf("%v/%v", path, strings.Join(fullSegments, "/")))
		if err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		for _, content := range contents {
			if content.IsDir() && !strings.HasPrefix(content.Name(), ".") {
				vals = append(vals, content.Name())
			}
		}
		return carapace.ActionValues(vals...)
	})
}
