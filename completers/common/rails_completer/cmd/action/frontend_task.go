package action

import "github.com/carapace-sh/carapace"

// ActionCssNamespaces completes css:* namespaces
func ActionCssNamespaces() carapace.Action {
	return carapace.ActionValues("install").Tag("css namespaces").UidF(namespaceUid("css"))
}

// ActionCssInstallTasks completes css:install:* tasks
func ActionCssInstallTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"tailwind", "Install Tailwind CSS",
		"bootstrap", "Install Bootstrap CSS",
		"postcss", "Install PostCSS",
		"sass", "Install Sass",
	).Tag("css install tasks").UidF(taskUid("css.install"))
}

// ActionImportmapTasks completes importmap:* tasks
func ActionImportmapTasks() carapace.Action {
	return carapace.ActionValues("install").Tag("importmap tasks").UidF(taskUid("importmap"))
}

// ActionJavascriptTasks completes javascript:* tasks
func ActionJavascriptTasks() carapace.Action {
	return carapace.ActionValues("install").Tag("javascript tasks").UidF(taskUid("javascript"))
}

// ActionRailtiesNamespaces completes railties:* namespaces
func ActionRailtiesNamespaces() carapace.Action {
	return carapace.ActionValues("install").Tag("railties namespaces").UidF(namespaceUid("railties"))
}

// ActionRailtiesInstallTasks completes railties:install:* tasks
func ActionRailtiesInstallTasks() carapace.Action {
	return carapace.ActionValues("migrations").Tag("railties install tasks").UidF(taskUid("railties.install"))
}

// ActionSecretsTasks completes secrets:* tasks
func ActionSecretsTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"edit", "Edit encrypted secrets",
		"setup", "Setup encrypted secrets",
		"show", "Show decrypted secrets",
	).Tag("secrets tasks").UidF(taskUid("secrets"))
}

// ActionStimulusNamespaces completes stimulus:* namespaces
func ActionStimulusNamespaces() carapace.Action {
	return carapace.ActionValues("install").Tag("stimulus namespaces").UidF(namespaceUid("stimulus"))
}

// ActionStimulusInstallTasks completes stimulus:install:* tasks
func ActionStimulusInstallTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"importmap", "Install with importmap",
		"node", "Install with node",
	).Tag("stimulus install tasks").UidF(taskUid("stimulus.install"))
}

// ActionTurboNamespaces completes turbo:* namespaces
func ActionTurboNamespaces() carapace.Action {
	return carapace.ActionValues("install").Tag("turbo namespaces").UidF(namespaceUid("turbo"))
}

// ActionTurboInstallTasks completes turbo:install:* tasks
func ActionTurboInstallTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"importmap", "Install with importmap",
		"node", "Install with node",
		"redis", "Install with Redis",
		"bun", "Install with bun",
	).Tag("turbo install tasks").UidF(taskUid("turbo.install"))
}

// ActionTailwindcssTasks completes tailwindcss:* tasks
func ActionTailwindcssTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"build", "Build Tailwind CSS",
		"clobber", "Remove built Tailwind CSS",
		"install", "Install Tailwind CSS",
		"watch", "Watch and build on changes",
	).Tag("tailwindcss tasks").UidF(taskUid("tailwindcss"))
}

// ActionYarnTasks completes yarn:* tasks
func ActionYarnTasks() carapace.Action {
	return carapace.ActionValues("install").Tag("yarn tasks").UidF(taskUid("yarn"))
}

// ActionZeitwerkTasks completes zeitwerk:* tasks
func ActionZeitwerkTasks() carapace.Action {
	return carapace.ActionValues("check").Tag("zeitwerk tasks").UidF(taskUid("zeitwerk"))
}
