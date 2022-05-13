package styles

import "github.com/rsteube/carapace/pkg/style"

var Gh = struct {
	Draft         string `desc:"draft pulls"`
	JobFailed     string `desc:"jobs not finished successfully"`
	JobInProgress string `desc:"jobs in progress"`
	JobSuccess    string `desc:"jobs successfully finished"`
	StateClosed   string `desc:"closed issues/pulls"`
	StateMerged   string `desc:"merged pulls"`
	StateOpen     string `desc:"open issues/pulls"`
	RepoPublic    string `desc:"public repository"`
	RepoPrivate   string `desc:"private repository"`
	RepoFork      string `desc:"forked repository"`
}{
	Draft:         style.Gray,
	JobFailed:     style.Red,
	JobInProgress: style.Yellow,
	JobSuccess:    style.Green,
	StateClosed:   style.Red,
	StateMerged:   style.Magenta,
	StateOpen:     style.Green,
	RepoPublic:    style.Default,
	RepoPrivate:   style.Red,
	RepoFork:      style.Gray,
}

func init() {
	style.Register("gh", &Gh)
}
