package git

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionPushOptions completes git push options (`-o`/`--push-option`)
//
//	ci.skip (GitLab: skip CI pipeline)
//	merge_request.create (GitLab: create merge request)
func ActionPushOptions() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValuesDescribed(
					"ci.no_pipeline", "GitLab: prevent creating any pipeline for this push",
					"ci.skip", "GitLab: skip CI pipeline for this push",
					"integrations.skip_ci", "GitLab: skip push events for CI/CD integrations",
					"merge_request.auto_merge", "GitLab: auto-merge when pipeline succeeds",
					"merge_request.create", "GitLab: create merge request for the pushed branch",
					"merge_request.draft", "GitLab: mark merge request as draft",
					"merge_request.remove_source_branch", "GitLab: remove source branch on merge",
					"merge_request.squash", "GitLab: squash commits on merge",
					"secret_push_protection.skip_all", "GitLab: skip secret push protection",
					"skip-validation", "Gerrit: skip commit validation",
					"trace", "Gerrit: enable tracing with auto-generated ID",
				).Tag("push options"),
				carapace.ActionValuesDescribed(
					"ci.input", "GitLab: pass input parameter (NAME=VALUE)",
					"ci.variable", "GitLab: set CI/CD variable (KEY=VALUE)",
					"custom-keyed-value", "Gerrit: attach custom keyed value (KEY:VALUE)",
					"deadline", "Gerrit: set push deadline (e.g. 10m, 30s)",
					"hashtag", "Gerrit: add hashtag to all changes",
					"merge_request.assign", "GitLab: assign user to merge request",
					"merge_request.description", "GitLab: set merge request description",
					"merge_request.label", "GitLab: add label to merge request",
					"merge_request.milestone", "GitLab: set milestone for merge request",
					"merge_request.target", "GitLab: set target branch for merge request",
					"merge_request.target_project", "GitLab: set target project for merge request",
					"merge_request.title", "GitLab: set merge request title",
					"merge_request.unassign", "GitLab: unassign user from merge request",
					"merge_request.unlabel", "GitLab: remove label from merge request",
					"repo.private", "Gitea: set repository visibility",
					"repo.template", "Gitea: set repository as template",
					"security_policy.bypass_reason", "GitLab: set bypass reason for security policy",
					"t", "Gerrit: add hashtag (shorthand)",
					"topic", "Gerrit: set topic name on all changes",
					"trace", "Gerrit: enable tracing with specific ID",
				).Suffix("="),
			).ToA().Tag("push options")
		default:
			switch c.Parts[0] {
			case "repo.private", "repo.template":
				return carapace.ActionValues("false", "true").StyleF(style.ForKeyword)
			default:
				return carapace.ActionValues()
			}
		}
	}).UidF(Uid("push-option"))
}
