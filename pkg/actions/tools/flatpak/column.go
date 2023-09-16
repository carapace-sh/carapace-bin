package flatpak

import "github.com/rsteube/carapace"

// ActionApplicationColumns completes application columns
//
//	name (Show the name)
//	description (Show the description)
func ActionApplicationColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "Show the name",
		"description", "Show the description",
		"application", "Show the application ID",
		"version", "Show the version",
		"branch", "Show the branch",
		"arch", "Show the architecture",
		"runtime", "Show the used runtime",
		"origin", "Show the origin remote",
		"installation", "Show the installation",
		"ref", "Show the ref",
		"active", "Show the active commit",
		"latest", "Show the latest commit",
		"size", "Show the installed size",
		"options", "Show options",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionDocumentColums completes document columns
//
//	id (Show the document ID)
//	path (Show the document path)
func ActionDocumentColums() carapace.Action {
	return carapace.ActionValuesDescribed(
		"id", "Show the document ID",
		"path", "Show the document path",
		"origin", "Show the document path",
		"application", "Show applications with permission",
		"permissions", "Show permissions for applications",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionHistoryColums completes history columns
//
//	time (Show when the change happened)
//	change (Show the kind of change)
func ActionHistoryColums() carapace.Action {
	return carapace.ActionValuesDescribed(
		"time", "Show when the change happened",
		"change", "Show the kind of change",
		"ref", "Show the ref",
		"application", "Show the application/runtime ID",
		"arch", "Show the architecture",
		"branch", "Show the branch",
		"installation", "Show the affected installation",
		"remote", "Show the remote",
		"commit", "Show the current commit",
		"old-commit", "Show the previous commit",
		"url", "Show the remote URL",
		"user", "Show the user doing the change",
		"tool", "Show the tool that was used",
		"version", "Show the Flatpak version",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionProcessColums completes process columns
//
//	instance (Show the instance ID)
//	pid (Show the PID of the wrapper process)
func ActionProcessColums() carapace.Action {
	return carapace.ActionValuesDescribed(
		"instance", "Show the instance ID",
		"pid", "Show the PID of the wrapper process",
		"child-pid", "Show the PID of the sandbox process",
		"application", "Show the application ID",
		"arch", "Show the architecture",
		"branch", "Show the application branch",
		"commit", "Show the application commit",
		"runtime", "Show the runtime ID",
		"runtime-branch", "Show the runtime branch",
		"runtime-commit", "Show the runtime commit",
		"active", "Show whether the app is active",
		"background", "Show whether the app is background",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionRemoteColumns completes remote columns
//
//	name (Show the name)
//	title (Show the title)
func ActionRemoteColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "Show the name",
		"title", "Show the title",
		"url", "Show the URL",
		"collection", "Show the collection ID",
		"subset", "Show the subset",
		"filter", "Show filter file",
		"priority", "Show the priority",
		"options", "Show options",
		"comment", "Show comment",
		"description", "Show description",
		"homepage", "Show homepage",
		"icon", "Show icon",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionRemoteContentColumns completes remote content columns
//
//	name (Show the name)
//	description (Show the description)
func ActionRemoteContentColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "Show the name",
		"description", "Show the description",
		"application", "Show the application ID",
		"version", "Show the version",
		"branch", "Show the branch",
		"arch", "Show the architecture",
		"origin", "Show the origin remote",
		"ref", "Show the ref",
		"commit", "Show the active commit",
		"runtime", "Show the runtime",
		"installed-size", "Show the installed size",
		"download-size", "Show the download size",
		"options", "Show options",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionSearchColumns completes search columns
//
//	name (Show the name)
//	description (Show the description)
func ActionSearchColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "Show the name",
		"description", "Show the description",
		"application", "Show the application ID",
		"version", "Show the version",
		"branch", "Show the application branch",
		"remotes", "Show the remotes",
		"all", "Show all columns",
		"help", "Show available columns",
	)
}

// ActionEllipsizations completes elipsizations
//
//	s (start)
//	m (middle)
func ActionEllipsizations() carapace.Action {
	return carapace.ActionValuesDescribed(
		"s", "start",
		"m", "middle",
		"e", "end",
		"f", "full",
	)
}
