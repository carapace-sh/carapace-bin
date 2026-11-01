package flatpak

import "github.com/carapace-sh/carapace"

// ActionApplicationColumns completes application columns
//
//	name (Show the name)
//	description (Show the description)
func ActionApplicationColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"active", "Show the active commit",
		"all", "Show all columns",
		"application", "Show the application ID",
		"arch", "Show the architecture",
		"branch", "Show the branch",
		"description", "Show the description",
		"help", "Show available columns",
		"installation", "Show the installation",
		"latest", "Show the latest commit",
		"name", "Show the name",
		"options", "Show options",
		"origin", "Show the origin remote",
		"ref", "Show the ref",
		"runtime", "Show the used runtime",
		"size", "Show the installed size",
		"version", "Show the version",
	)
}

// ActionDocumentColums completes document columns
//
//	id (Show the document ID)
//	path (Show the document path)
func ActionDocumentColums() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Show all columns",
		"application", "Show applications with permission",
		"help", "Show available columns",
		"id", "Show the document ID",
		"origin", "Show the document path",
		"path", "Show the document path",
		"permissions", "Show permissions for applications",
	)
}

// ActionHistoryColums completes history columns
//
//	time (Show when the change happened)
//	change (Show the kind of change)
func ActionHistoryColums() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Show all columns",
		"application", "Show the application/runtime ID",
		"arch", "Show the architecture",
		"branch", "Show the branch",
		"change", "Show the kind of change",
		"commit", "Show the current commit",
		"help", "Show available columns",
		"installation", "Show the affected installation",
		"old-commit", "Show the previous commit",
		"ref", "Show the ref",
		"remote", "Show the remote",
		"time", "Show when the change happened",
		"tool", "Show the tool that was used",
		"url", "Show the remote URL",
		"user", "Show the user doing the change",
		"version", "Show the Flatpak version",
	)
}

// ActionProcessColums completes process columns
//
//	instance (Show the instance ID)
//	pid (Show the PID of the wrapper process)
func ActionProcessColums() carapace.Action {
	return carapace.ActionValuesDescribed(
		"active", "Show whether the app is active",
		"all", "Show all columns",
		"application", "Show the application ID",
		"arch", "Show the architecture",
		"background", "Show whether the app is background",
		"branch", "Show the application branch",
		"child-pid", "Show the PID of the sandbox process",
		"commit", "Show the application commit",
		"help", "Show available columns",
		"instance", "Show the instance ID",
		"pid", "Show the PID of the wrapper process",
		"runtime", "Show the runtime ID",
		"runtime-branch", "Show the runtime branch",
		"runtime-commit", "Show the runtime commit",
	)
}

// ActionRemoteColumns completes remote columns
//
//	name (Show the name)
//	title (Show the title)
func ActionRemoteColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Show all columns",
		"collection", "Show the collection ID",
		"comment", "Show comment",
		"description", "Show description",
		"filter", "Show filter file",
		"help", "Show available columns",
		"homepage", "Show homepage",
		"icon", "Show icon",
		"name", "Show the name",
		"options", "Show options",
		"priority", "Show the priority",
		"subset", "Show the subset",
		"title", "Show the title",
		"url", "Show the URL",
	)
}

// ActionRemoteContentColumns completes remote content columns
//
//	name (Show the name)
//	description (Show the description)
func ActionRemoteContentColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Show all columns",
		"application", "Show the application ID",
		"arch", "Show the architecture",
		"branch", "Show the branch",
		"commit", "Show the active commit",
		"description", "Show the description",
		"download-size", "Show the download size",
		"help", "Show available columns",
		"installed-size", "Show the installed size",
		"name", "Show the name",
		"options", "Show options",
		"origin", "Show the origin remote",
		"ref", "Show the ref",
		"runtime", "Show the runtime",
		"version", "Show the version",
	)
}

// ActionSearchColumns completes search columns
//
//	name (Show the name)
//	description (Show the description)
func ActionSearchColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Show all columns",
		"application", "Show the application ID",
		"branch", "Show the application branch",
		"description", "Show the description",
		"help", "Show available columns",
		"name", "Show the name",
		"remotes", "Show the remotes",
		"version", "Show the version",
	)
}

// ActionEllipsizations completes elipsizations
//
//	s (start)
//	m (middle)
func ActionEllipsizations() carapace.Action {
	return carapace.ActionValuesDescribed(
		"e", "end",
		"f", "full",
		"m", "middle",
		"s", "start",
	)
}
