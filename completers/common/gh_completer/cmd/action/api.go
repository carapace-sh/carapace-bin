package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionApiPreviews https://docs.github.com/en/rest/overview/api-previews
func ActionApiPreviews() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ant-man", "Enhanced deployments",
		"baptiste", "Create and use repository templates",
		"cloak", "Commit search",
		"corsair", "Content attachments",
		"eye-scream", "Pre-receive environments",
		"flash", "Deployment statuses",
		"groot", "List branches or pull requests for a commit",
		"inertia", "Projects",
		"luke-cage", "Require multiple approving reviews",
		"lydian", "Update a pull request branch",
		"mercy", "Repository topics",
		"mockingbird", "Timeline",
		"nebula", "New visibility parameter for the Repositories API",
		"scarlet-witch", "Codes of conduct",
		"squirrel-girl", "Reactions",
		"starfox", "Project card details",
		"superpro", "Global webhooks",
		"surtur", "Repository creation permissions",
		"switcheroo", "Enable and disable Pages",
		"x-ray", "Anonymous Git access to repositories",
		"zzzax", "Require signed commits",
	).Tag("api previews")
}

func ActionApiV3Paths(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		method := "get"
		if f := cmd.Flag("method"); f != nil && f.Changed {
			method = strings.ToLower(f.Value.String())
		}

		return carapace.ActionValuesDescribed(v3Paths[method]...).
			MultiPartsP("/", `{(.*)}`, func(placeholder string, matches map[string]string) carapace.Action {
				repo := gh.RepoOpts{Owner: matches["{owner}"], Name: matches["{repo}"]}

				switch placeholder {
				// TODO completion for other placeholders
				case "{archive_format}":
					return carapace.ActionValues("zip")
				case "{artifact_id}":
					return gh.ActionWorkflowArtifactIds(gh.ArtifactOpts{Owner: repo.Owner, Name: repo.Name})
				case "{cache_id}":
					return gh.ActionCaches(repo)
				case "{artifact_number}":
					issueNumber, _ := strconv.Atoi(matches["{issue_number}"]) // TODO hidden issues
					return gh.ActionIssueArtifacts(gh.IssueArtifactOpts{Owner: repo.Owner, Name: repo.Name, IssueNumber: issueNumber})
				case "{autolink_id}":
					return gh.ActionAutolinks(repo)
				case "{assignee}":
					return gh.ActionAssignableUsers(repo)
				case "{branch}":
					return gh.ActionBranches(repo)
				case "{coc_key}":
					return gh.ActionCocs(gh.HostOpts{Host: repo.Host})
				case "{codespace_name}":
					return gh.ActionCodespaces()
				case "{environment_name}":
					return gh.ActionEnvironments(repo)
				case "{gist_id}":
					return gh.ActionGistIds(gh.GistOpts{Owner: repo.Owner})
				case "{hook_id}":
					if matches["{org}"] != "" {
						return gh.ActionOrganizationHooks(gh.OwnerOpts{Host: repo.Host, Owner: matches["{org}"]})
					}
					return gh.ActionRepositoryHooks(repo)
				case "{gpg_key_id}":
					return gh.ActionGpgKeys(gh.HostOpts{Host: repo.Host})
				case "{gitignore_name}":
					return gh.ActionGitignoreTemplates(gh.HostOpts{Host: repo.Host})
				case "{label}":
					return gh.ActionLabels(gh.RepoOpts{Owner: matches["{owner}"], Name: matches["{repo}"]})
				case "{key_id}":
					return carapace.Batch(
						gh.ActionDeployKeys(repo),                      // repos/{owner}/{repo}/keys/{key_id}
						gh.ActionSshKeys(gh.HostOpts{Host: repo.Host}), // user/keys/{key_id}
					).ToA()
				case "{license}":
					return gh.ActionLicenses(gh.HostOpts{})
				case "{issue_number}":
					return gh.ActionIssues(gh.IssueOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name, Open: true, Closed: true})
				case "{owner}":
					if strings.HasPrefix(c.Value, ":") {
						return carapace.ActionValues(":owner")
					} else {
						return gh.ActionOwners(gh.HostOpts{}) // TODO host
					}
				case "{org}":
					return gh.ActionOrganizations(gh.HostOpts{})
				case "{milestone_number}":
					return gh.ActionMilestoneNumbers(repo)
				case "{package_name}":
					return packageAction(matches, repo.Host, false)
				case "{package_type}":
					return gh.ActionPackageTypes()
				case "{package_version_id}":
					return packageAction(matches, repo.Host, true)
				case "{pull_number}":
					return gh.ActionPullRequests(gh.PullRequestOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name, Open: true, Closed: true, Merged: true})
				case "{repo}":
					if strings.HasPrefix(c.Value, ":") {
						return carapace.ActionValues(":repo")
					} else {
						return gh.ActionRepositories(gh.OwnerOpts{Owner: matches["{owner}"]})
					}
				case "{runner_group_id}":
					if matches["{org}"] != "" {
						return gh.ActionOrganizationRunnerGroups(gh.OwnerOpts{Host: repo.Host, Owner: matches["{org}"]})
					}
					return carapace.ActionValues(placeholder)
				case "{runner_id}":
					if matches["{org}"] != "" {
						return gh.ActionOrganizationRunners(gh.OwnerOpts{Host: repo.Host, Owner: matches["{org}"]})
					}
					return gh.ActionRepositoryRunners(repo)
				case "{ruleset_id}":
					return gh.ActionRulesets(repo)
				case "{run_id}":
					return gh.ActionWorkflowRuns(gh.RunOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name, All: true})
				case "{secret_name}":
					switch {
					case matches["{org}"] != "": // TODO distinguishing app scopes of the path
						return carapace.Batch(
							gh.ActionSecrets(gh.SecretsOpts{Org: matches["{org}"], App: "actions"}),
							gh.ActionSecrets(gh.SecretsOpts{Org: matches["{org}"], App: "codespaces"}),
							gh.ActionSecrets(gh.SecretsOpts{Org: matches["{org}"], App: "dependabot"}),
						).ToA()
					case repo.Owner != "" && repo.Name != "":
						return gh.ActionSecrets(gh.SecretsOpts{Repo: fmt.Sprintf("%v/%v", repo.Owner, repo.Name)})
					default:
						return carapace.ActionValues(placeholder)
					}
				case "{team_slug}":
					return gh.ActionTeams(gh.OwnerOpts{Host: repo.Host, Owner: matches["{org}"]})
				case "{tag}": // only used with releases
					return gh.ActionReleases(gh.ReleaseOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
				case "{template_owner}": // ignore this as it is already provided by `{owner}`
					return carapace.ActionValues()
				case "{template_repo}": // ignore this as it is already provided by `{repo}`
					return carapace.ActionValues()
				case "{username}":
					return gh.ActionUsers(gh.HostOpts{})
				case "{workflow_id}":
					return gh.ActionWorkflows(gh.WorkflowOpts{Enabled: true, Disabled: true, ById: true, Owner: matches["{owner}"], Name: matches["{repo}"]})
				default:
					// static value or placeholder not yet handled
					return carapace.ActionValues(placeholder)
				}
			})
	})
}

func packagesOpts(matches map[string]string, host string) (gh.PackageOpts, bool) {
	opts := gh.PackageOpts{Host: host, Type: matches["{package_type}"]}
	switch {
	case matches["{org}"] != "":
		opts.Owner, opts.Organization = matches["{org}"], true
	case matches["{owner}"] != "" && matches["{repo}"] != "":
		opts.Owner = matches["{owner}"]
	case matches["{username}"] != "":
		opts.Owner = matches["{username}"]
	default:
		return opts, false
	}
	return opts, true
}

func packageAction(matches map[string]string, host string, versions bool) carapace.Action {
	opts, ok := packagesOpts(matches, host)
	if !ok {
		return carapace.ActionValues()
	}
	if versions {
		opts.Name = matches["{package_name}"]
		if opts.Name == "" {
			return carapace.ActionValues()
		}
		return gh.ActionPackageVersions(opts)
	}
	return gh.ActionPackages(opts)
}

func ActionPackageTypes() carapace.Action {
	return carapace.ActionValues("container", "docker", "maven", "npm", "nuget", "rubygems").
		Tag("package types")
}
