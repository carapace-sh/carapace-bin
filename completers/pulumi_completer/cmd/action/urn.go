package action

import (
	"encoding/json"
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
)

type StackExport struct {
	Deployment struct {
		Resources []struct {
			Urn string `json:"urn"`
		} `json:"resources"`
	} `json:"deployment"`
}

func ActionUrns(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		cwd := cmd.Flag("cwd").Value.String()
		stack := cmd.Flag("stack").Value.String()

		c.Setenv("PULUMI_SKIP_UPDATE_CHECK", "1")
		return carapace.ActionExecCommand("pulumi", "stack", "export", "--cwd", cwd, "--stack", stack)(func(output []byte) carapace.Action {
			var exp StackExport
			if err := json.Unmarshal(output, &exp); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			urns := make([]string, len(exp.Deployment.Resources))
			for i, r := range exp.Deployment.Resources {
				urns[i] = r.Urn
			}
			return carapace.ActionValues(urns...)
		}).Cache(5*time.Second,
			func() (string, error) {
				workdir := c.Dir
				if cmd.Flag("cwd").Changed { // TODO use preinvoke?
					workdir = cmd.Flag("cwd").Value.String()
				}
				stack := cmd.Flag("stack").Value.String()

				absWd, err := filepath.Abs(workdir)
				if err != nil {
					return "", err
				}

				return cache.String(absWd, stack)()
			},
		).Invoke(c).ToA()
	})
}
