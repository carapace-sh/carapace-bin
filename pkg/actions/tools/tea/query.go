package tea

import (
	"os/exec"
	"strings"

	"github.com/carapace-sh/carapace"
	"gopkg.in/yaml.v3"
)

func fixYamlOutput(b []byte) []byte {
	lines := strings.Split(string(b), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "NOTE:") {
		lines = lines[1:]
	}

	for index, line := range lines {
		line := strings.ReplaceAll(line, `"`, `\"`)
		line = strings.Replace(line, `'`, `"`, 1)
		if last := strings.LastIndex(line, `'`); last != -1 {
			line = line[:last] + strings.Replace(line[last:], `'`, `"`, 1)
		}
		lines[index] = line
	}
	carapace.LOG.Println(strings.Join(lines, "\n"))
	return []byte(strings.Join(lines, "\n"))
}

func actionYamlQuery(opts RepoOpts, result interface{}, args ...string) func(f func() carapace.Action) carapace.Action {
	return func(f func() carapace.Action) carapace.Action {
		args = append(args, "--output", "yaml", "--limit", "100")
		if opts.Login != "" {
			args = append(args, "--login", opts.Login)
		}
		if opts.Remote != "" {
			args = append(args, "--remote", opts.Remote)
		}
		if opts.Repo != "" {
			args = append(args, "--repo", opts.Repo)
		}
		return carapace.ActionExecCommandE("tea", args...)(func(output []byte, err error) carapace.Action {
			output = fixYamlOutput(output)
			if err != nil {
				if _, ok := err.(*exec.ExitError); ok {
					return carapace.ActionMessage(strings.SplitN(string(output), "\n", 2)[0])
				}
				return carapace.ActionMessage(err.Error())
			}

			if err := yaml.Unmarshal(output, result); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f()
		})
	}
}
