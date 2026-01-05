package git

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action/run"
	"github.com/carapace-sh/carapace/pkg/execlog"
)

func listRemotes() ([]string, error) {
	remoteCmd, err := GitCommand("remote", "-v")
	if err != nil {
		return nil, err
	}
	output, err := run.PrepareCmd(remoteCmd).Output()
	return outputLines(output), err
}

var GitCommand = func(args ...string) (*execlog.Cmd, error) {
	gitExe, err := execlog.LookPath("git")
	if err != nil {
		programName := "git"
		if runtime.GOOS == "windows" {
			programName = "Git for Windows"
		}
		return nil, fmt.Errorf("unable to find git executable in PATH; please install %s before retrying", programName)
	}
	return execlog.Command(gitExe, args...), nil
}

type Commit struct {
	Sha   string
	Title string
}

func lookupCommit(sha, format string) ([]byte, error) {
	logCmd, err := GitCommand("-c", "log.ShowSignature=false", "show", "-s", "--pretty=format:"+format, sha)
	if err != nil {
		return nil, err
	}
	return run.PrepareCmd(logCmd).Output()
}

func LastCommit() (*Commit, error) {
	output, err := lookupCommit("HEAD", "%H,%s")
	if err != nil {
		return nil, err
	}

	idx := bytes.IndexByte(output, ',')
	return &Commit{
		Sha:   string(output[0:idx]),
		Title: strings.TrimSpace(string(output[idx+1:])),
	}, nil
}

func outputLines(output []byte) []string {
	lines := strings.TrimSuffix(string(output), "\n")
	return strings.Split(lines, "\n")

}
