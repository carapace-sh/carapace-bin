package action

import (
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rsteube/carapace"
)

func ActionRemotes() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if remotes, err := Remotes(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(remotes...)
		}
	})
}

func Remotes() ([]string, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return []string{}, err
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return []string{}, err
	}

	names := make([]string, len(remotes))
	for i, r := range remotes {
		names[i] = r.Config().Name
	}
	return names, nil
}

func RemoteBranches(remote string) ([]string, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return []string{}, err
	}

	branches, err := repo.References() // TODO verify is a branch Branches didn't seem to work
	if err != nil {
		return []string{}, err
	}
	reg := regexp.MustCompile(`^refs/remotes/[^/]+/`)

	names := []string{}
	branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsRemote() && strings.HasPrefix(ref.Name().String(), "refs/remotes/"+remote) {
			names = append(names, reg.ReplaceAllString(ref.Name().String(), ""))
		}
		return nil
	})
	return names, nil
}
