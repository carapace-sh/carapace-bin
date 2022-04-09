package action

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
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
	)
}

//go:embed api.github.com.json
var apiV3Spec string // https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json

func ActionApiV3Paths(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		method := "get"
		if f := cmd.Flag("method"); f != nil && f.Changed {
			method = strings.ToLower(f.Value.String())
		}

		apiV3Spec = strings.NewReplacer( // patch ambiguous placeholders
			"gitignore/templates/{name}", "gitignore/templates/{gitignore_name}",
			"codes_of_conduct/{key}", "codes_of_conduct/{coc_key}",
			"labels/{name}", "labels/{label}",
		).Replace(apiV3Spec)

		return http.ActionOpenApiPaths([]byte(apiV3Spec), method, `{(.*)}`, func(c carapace.Context, matchedData map[string]string, segment string) carapace.Action {
			switch segment {
			// TODO completion for other placeholders
			case "{archive_format}":
				return carapace.ActionValues("zip")
			case "{artifact_id}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionWorkflowArtifactIds(cmd, "")
			case "{assignee}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionAssignableUsers(cmd)
			case "{branch}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionBranches(cmd)
			case "{coc_key}":
				return ActionCocs(cmd)
			case "{environment_name}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionEnvironments(cmd)
			case "{gist_id}":
				return ActionGists(cmd)
			case "{gitignore_name}":
				return ActionGitignoreTemplates(cmd)
			case "{label}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionLabels(cmd)
			case "{license}":
				return ActionLicenses(cmd)
			case "{issue_number}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionIssues(cmd, IssueOpts{Open: true, Closed: true})
			case "{owner}":
				if strings.HasPrefix(c.CallbackValue, ":") {
					return carapace.ActionValues(":owner")
				} else {
					return ActionUsers(cmd, UserOpts{Users: true, Organizations: true})
				}
			case "{org}":
				return ActionUsers(cmd, UserOpts{Organizations: true})
			case "{package_type}":
				return ActionPackageTypes()
			case "{pull_number}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionPullRequests(cmd, PullRequestOpts{Open: true, Closed: true, Merged: true})
			case "{repo}":
				if strings.HasPrefix(c.CallbackValue, ":") {
					return carapace.ActionValues(":repo")
				} else {
					return ActionRepositories(cmd, matchedData["{owner}"], c.CallbackValue)
				}
			case "{tag}": // only used with releases
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionReleases(cmd)
			case "{template_owner}": // ignore this as it is already provided by `{owner}`
				return carapace.ActionValues()
			case "{template_repo}": // ignore this as it is already provided by `{repo}`
				return carapace.ActionValues()
			case "{username}":
				return ActionUsers(cmd, UserOpts{Users: true})
			case "{workflow_id}":
				fakeRepoFlag(cmd, matchedData["{owner}"], matchedData["{repo}"])
				return ActionWorkflows(cmd, WorkflowOpts{Enabled: true, Disabled: true, Id: true})
			default:
				// static value or placeholder not yet handled
				return carapace.ActionValues(segment)
			}
		})
	})
}

func fakeRepoFlag(cmd *cobra.Command, owner, repo string) {
	cmd.Flags().String("repo", fmt.Sprintf("%v/%v", owner, repo), "fake repo flag")
	cmd.Flag("repo").Changed = true
}

func ActionPackageTypes() carapace.Action {
	return carapace.ActionValues("npm", "maven", "rubygems", "nuget", "docker", "container")
}
