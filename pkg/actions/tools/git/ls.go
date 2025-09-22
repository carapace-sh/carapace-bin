package git

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
)

type LsRemoteRefOption struct {
	Url      string
	Branches bool
	Tags     bool
}

// ActionLsRemoteRefs completes branches and tags for a remote url
//
//	gh-pages (da4528d0a57ad71417336f0e96fa65ece2fad45a)
//	master (3fbdef3c6a10094812a15cba8e825898b757dfb3)
func ActionLsRemoteRefs(opts LsRemoteRefOption) carapace.Action {
	return carapace.ActionExecCommand("git", "ls-remote", "--refs", "--tags", "--heads", opts.Url)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			if opts.Branches && strings.HasPrefix(fields[1], "refs/heads/") {
				vals = append(vals, strings.TrimPrefix(fields[1], "refs/heads/"), fields[0], styles.Git.Branch)
			} else if opts.Tags && strings.HasPrefix(fields[1], "refs/tags/") {
				vals = append(vals, strings.TrimPrefix(fields[1], "refs/tags/"), fields[0], styles.Git.Tag)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

// ActionRefFiles completes files of given ref
//
//	go.mod
//	pkg/
func ActionRefFiles(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		rootDir, err := rootDir(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		relDir, err := filepath.Rel(rootDir, c.Dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		path := fmt.Sprintf("%v/%v", relDir, filepath.Dir(c.Value))
		if path := filepath.Dir(path); path == "." {
			path = ""
		} else {
			path += "/"
		}

		path = filepath.Clean(path) + "/"
		if path = filepath.Dir(path); path == "." {
			path = ""
		} else {
			path += "/"
		}

		prefix := ""
		if dir := filepath.Dir(c.Value); dir != "." {
			prefix = dir + "/"
		}

		args := []string{"ls-tree", "--name-only", "--full-tree", ref}
		if path != "" {
			args = append(args, path)
		}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			files := lines[:len(lines)-1]
			for index, file := range files {
				files[index] = filepath.Base(file)
			}

			args = append(args, "-d") // only directories
			return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				directories := lines[:len(lines)-1]
				for index, dir := range directories {
					directories[index] = filepath.Base(dir)
				}

				filesA := carapace.ActionValues(files...).Invoke(c).Filter(directories...)

				for index, dir := range directories {
					directories[index] = dir + "/"
				}
				directoriesA := carapace.ActionValues(directories...).Invoke(c)

				return filesA.Merge(directoriesA).Prefix(prefix).ToA().NoSpace().StyleF(style.ForPathExt)
			})
		})
	}).Tag("ref files").UidF(Uid("ref-file", "ref", ref)).QueryF(Uid("ref-file", "ref", ref))
}
