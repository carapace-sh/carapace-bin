package action

import (
	"fmt"
	"strings"
	"sync"

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
				fakeRepoFlag(cmd, matches["{owner}"], matches["{repo}"])

				switch placeholder {
				// TODO completion for other placeholders
				case "{archive_format}":
					return carapace.ActionValues("zip")
				case "{artifact_id}":
					return ActionWorkflowArtifactIds(cmd, "")
				case "{assignee}":
					return ActionAssignableUsers(cmd)
				case "{branch}":
					return ActionBranches(cmd)
				case "{coc_key}":
					return ActionCocs(cmd)
				case "{environment_name}":
					return ActionEnvironments(cmd)
				case "{gist_id}":
					return ActionGists(cmd)
				case "{gitignore_name}":
					return ActionGitignoreTemplates(cmd)
				case "{label}":
					return gh.ActionLabels(gh.RepoOpts{Owner: matches["{owner}"], Name: matches["{repo}"]})
				case "{license}":
					return gh.ActionLicenses(gh.HostOpts{})
				case "{issue_number}":
					return ActionIssues(cmd, IssueOpts{Open: true, Closed: true})
				case "{owner}":
					if strings.HasPrefix(c.Value, ":") {
						return carapace.ActionValues(":owner")
					} else {
						return gh.ActionOwners(gh.HostOpts{}) // TODO host
					}
				case "{org}":
					return gh.ActionOrganizations(gh.HostOpts{})
				case "{package_type}":
					return ActionPackageTypes()
				case "{pull_number}":
					return ActionPullRequests(cmd, PullRequestOpts{Open: true, Closed: true, Merged: true})
				case "{repo}":
					if strings.HasPrefix(c.Value, ":") {
						return carapace.ActionValues(":repo")
					} else {
						return gh.ActionRepositories(gh.OwnerOpts{Owner: matches["{owner}"]})
					}
				case "{tag}": // only used with releases
					return ActionReleases(cmd)
				case "{template_owner}": // ignore this as it is already provided by `{owner}`
					return carapace.ActionValues()
				case "{template_repo}": // ignore this as it is already provided by `{repo}`
					return carapace.ActionValues()
				case "{username}":
					return gh.ActionUsers(gh.HostOpts{})
				case "{workflow_id}":
					return ActionWorkflows(cmd, WorkflowOpts{Enabled: true, Disabled: true, Id: true})
				default:
					// static value or placeholder not yet handled
					return carapace.ActionValues(placeholder)
				}
			})
	})
}

var fakeRepoFlagMutex sync.Mutex

func fakeRepoFlag(cmd *cobra.Command, owner, repo string) {
	if cmd.Flag("repo") == nil {
		fakeRepoFlagMutex.Lock()
		defer fakeRepoFlagMutex.Unlock()

		if cmd.Flag("repo") == nil {
			cmd.Flags().String("repo", fmt.Sprintf("%v/%v", owner, repo), "fake repo flag")
			cmd.Flag("repo").Changed = true
		}
	}
}

func ActionPackageTypes() carapace.Action {
	return carapace.ActionValues("npm", "maven", "rubygems", "nuget", "docker", "container").
		Tag("package types")
}
