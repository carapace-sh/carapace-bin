package styles

import "github.com/rsteube/carapace/pkg/style"

var Gh = struct {
	Draft             string `desc:"draft pulls"`
	JobFailed         string `desc:"jobs not finished successfully"`
	JobInProgress     string `desc:"jobs in progress"`
	JobSuccess        string `desc:"jobs successfully finished"`
	OwnerOrganization string `desc:"organization"`
	OwnerUser         string `desc:"user"`
	RepoArchived      string `desc:"archived repository"`
	RepoFork          string `desc:"forked repository"`
	RepoLocked        string `desc:"locked repository"`
	RepoMirror        string `desc:"mirror repository"`
	RepoPrivate       string `desc:"private repository"`
	RepoPublic        string `desc:"public repository"`
	RepoTemplate      string `desc:"template repository"`
	StateClosed       string `desc:"closed issues/pulls"`
	StateMerged       string `desc:"merged pulls"`
	StateOpen         string `desc:"open issues/pulls"`
}{
	Draft:             style.Gray,
	JobFailed:         style.Red,
	JobInProgress:     style.Yellow,
	JobSuccess:        style.Green,
	OwnerOrganization: style.Blue,
	OwnerUser:         style.Default,
	RepoArchived:      style.Magenta,
	RepoFork:          style.Gray,
	RepoLocked:        style.Bold,
	RepoMirror:        style.Cyan,
	RepoPrivate:       style.Red,
	RepoPublic:        style.Default,
	RepoTemplate:      style.Yellow,
	StateClosed:       style.Red,
	StateMerged:       style.Magenta,
	StateOpen:         style.Green,
}

func init() {
	style.Register("gh", &Gh)
}
