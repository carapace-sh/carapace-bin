package gh

import (
	"fmt"
	"os"
	"sort"

	"github.com/carapace-sh/carapace"
	"go.yaml.in/yaml/v3"
)

// gistUser resolves host and owner for gist access (the authenticated user),
// following gh's own host precedence: GH_HOST env, else the hosts.yml entries
// (github.com first), else no host.
func gistUser(c carapace.Context) (host string, user string, err error) {
	content, err := os.ReadFile(configDir(c) + "/hosts.yml")
	if err != nil {
		return "", "", err
	}

	var config hostConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		return "", "", err
	}

	hosts := make([]string, 0, len(config))
	for h := range config {
		hosts = append(hosts, h)
	}
	sort.SliceStable(hosts, func(i, j int) bool { return hosts[i] == "github.com" })

	host = c.Getenv("GH_HOST")
	if host == "" && len(hosts) > 0 {
		host = hosts[0]
	}
	conf, ok := config[host]
	if !ok {
		return "", "", fmt.Errorf("unknown host %q", host)
	}
	if conf.User == "" {
		return "", "", fmt.Errorf("could not retrieve default user")
	}
	user = conf.User
	return host, user, nil
}

type gist struct {
	Name        string
	Description string
	Files       []struct {
		Name string
		Size int
	}
}

func gistQueryAction(host string, owner string, withFiles bool, transform func(gists []gist) carapace.Action) carapace.Action {
	flags := "name, description"
	if withFiles {
		flags += ", files { name, size }"
	}

	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if owner == "" {
			h, u, err := gistUser(c)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			if host == "" {
				host = h
			}
			owner = u
		}
		var queryResult struct {
			Data struct {
				User struct {
					Gists struct {
						Edges []struct {
							Node gist
						}
					}
				}
			}
		}
		return graphQlAction(RepoOpts{Host: host},
			fmt.Sprintf(`user(login: %q) { gists(first: 100, privacy: ALL, orderBy: {field: UPDATED_AT, direction: DESC}) { edges { node { %v } } } }`, owner, flags),
			&queryResult, func() carapace.Action {
				gists := make([]gist, 0)
				for _, edge := range queryResult.Data.User.Gists.Edges {
					gists = append(gists, edge.Node)
				}
				return transform(gists)
			})
	})
}

type GistOpts struct {
	Host string
	// Owner completes gists of this user (default: authenticated user of the resolved host)
	Owner string
	// Gist completes files of this gist
	Gist string
}

func (o GistOpts) Default() GistOpts {
	return o
}

// ActionGistIds completes gist ids
//
//	1a2b3c (fix contract bug)
//	e5f6d7 (personal snippet)
func ActionGistIds(opts GistOpts) carapace.Action {
	return gistQueryAction(opts.Host, opts.Owner, false, func(gists []gist) carapace.Action {
		vals := make([]string, 0, len(gists)*2)
		for _, gist := range gists {
			vals = append(vals, gist.Name, gist.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("gist ids")
}

// ActionGistFiles completes files of a gist
//
//	contract.txt (1.1 kB)
func ActionGistFiles(opts GistOpts) carapace.Action {
	return gistQueryAction(opts.Host, opts.Owner, true, func(gists []gist) carapace.Action {
		vals := make([]string, 0)
		for _, gist := range gists {
			if gist.Name == opts.Gist {
				for _, file := range gist.Files {
					vals = append(vals, file.Name, byteCountSI(file.Size))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("gist files")
}
