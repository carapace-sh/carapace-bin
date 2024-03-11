package gh

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAuthScopes completes authentication scopes
//
//	admin:gpg_key (Fully manage GPG keys)
//	admin:org (Fully manage the organization and its teams, projects, and memberships)
func ActionAuthScopes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"admin:gpg_key", "Fully manage GPG keys.",
		"admin:org", "Fully manage the organization and its teams, projects, and memberships.",
		"admin:org_hook", "Grants read, write, ping, and delete access to organization hooks.",
		"admin:public_key", "Fully manage public keys.",
		"admin:repo_hook", "Grants read, write, ping, and delete access to repository hooks in public and private repositories.",
		"codespace", "Full control of codespaces",
		"delete:packages", "Grants access to delete packages from GitHub Packages.",
		"delete_repo", "Grants access to delete adminable repositories.",
		"gist", "Grants write access to gists.",
		"notifications", "Grants read access to a user's notifications",
		"project", "Grants read/write access to user and organization projects",
		"public_repo", "Limits access to public repositories.",
		"read:discussion", "Allows read access for team discussions.",
		"read:gpg_key", "List and view details for GPG keys.",
		"read:org", "Read-only access to organization membership, organization projects, and team membership.",
		"read:packages", "Grants access to download or install packages from GitHub Packages.",
		"read:project", "Grants read only access to user and organization projects",
		"read:public_key", "List and view details for public keys.",
		"read:repo_hook", "Grants read and ping access to hooks in public or private repositories.",
		"read:user", "Grants access to read a user's profile data.",
		"repo", "Grants full access to private and public repositories.",
		"repo:invite", "Grants accept/decline abilities for invitations to collaborate on a repository.",
		"repo:status", "Grants read/write access to public and private repository commit statuses.",
		"repo_deployment", "Grants access to deployment statuses for public and private repositories.",
		"security_events", "Grants read and write access to security events in the code scanning API.",
		"user", "Grants read/write access to profile info only.",
		"user:email", "Grants read access to a user's email addresses.",
		"user:follow", "Grants access to follow or unfollow other users.",
		"workflow", "Grants the ability to add and update GitHub Actions workflow files.",
		"write:discussion", "Allows read and write access for team discussions.",
		"write:gpg_key", "Create, list, and view details for GPG keys.",
		"write:org", "Read and write access to organization membership, organization projects, and team membership.",
		"write:packages", "Grants access to upload or publish a package in GitHub Packages.",
		"write:public_key", "Create, list, and view details for public keys.",
		"write:repo_hook", "Grants read, write, and ping access to hooks in public or private repositories.",
	)
}

// ActionCurrentAuthScopes completes current authentication scopes for given hostname
//
//	admin:gpg_key (Fully manage GPG keys)
//	admin:org (Fully manage the organization and its teams, projects, and memberships)
func ActionCurrentAuthScopes(hostname string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"api", "--method", "HEAD", "", "--include"}
		if hostname != "" {
			args = append(args, "--hostname", hostname)
		}

		return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			scopes := []string{}
			for _, line := range lines {
				if strings.HasPrefix(line, "X-Oauth-Scopes: ") {
					scopes = strings.Split(strings.TrimPrefix(line, "X-Oauth-Scopes: "), ", ")
					break
				}
			}
			return ActionAuthScopes().Retain(scopes...)
		})
	})
}
