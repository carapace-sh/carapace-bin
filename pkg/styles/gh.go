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
	RepoArchived  string `desc:"archived repository"`
	RepoFork      string `desc:"forked repository"`
	RepoLocked    string `desc:"locked repository"`
	RepoMirror    string `desc:"mirror repository"`
	RepoPrivate   string `desc:"private repository"`
	RepoPublic    string `desc:"public repository"`
	RepoTemplate  string `desc:"template repository"`
}{
	Draft:         style.Gray,
	JobFailed:     style.Red,
	JobInProgress: style.Yellow,
	JobSuccess:    style.Green,
	StateClosed:   style.Red,
	StateMerged:   style.Magenta,
	StateOpen:     style.Green,
	RepoArchived:  style.Magenta,
	RepoFork:      style.Gray,
	RepoLocked:    style.Bold,
	RepoMirror:    style.Cyan,
	RepoPrivate:   style.Red,
	RepoPublic:    style.Default,
	RepoTemplate:  style.Yellow,
}

func init() {
	style.Register("gh", &Gh)
}
