package action

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
	"github.com/spf13/cobra"
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
			var stackExport StackExport
			if err := json.Unmarshal(output, &stackExport); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			urns := make([]string, len(stackExport.Deployment.Resources))
			for i, r := range stackExport.Deployment.Resources {
				urns[i] = r.Urn
			}
			return carapace.ActionValues(urns...)
		}).Cache(1*time.Minute,
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

				return key.String(absWd, stack)()
			},
		).Invoke(c).ToA()
	})
}
