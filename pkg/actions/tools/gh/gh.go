// package gh contains github related actions
package gh

import (
	"github.com/rsteube/carapace/pkg/cache/key"
)

type HostOpts struct {
	Host string
}

func (o HostOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host}
}

type OwnerOpts struct {
	Host  string
	Owner string
}

func (o OwnerOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner}

}

type RepoOpts struct {
	Host  string
	Owner string
	Name  string
}

func (o RepoOpts) cacheKey() key.Key { return key.String(o.Host, o.Owner, o.Name) }
