package action

import "github.com/carapace-sh/carapace"

func ActionAuditwheel() carapace.Action {
	return carapace.ActionValues(
		"check",
		"repair",
		"skip",
		"warn",
	).Tag("auditwheel policies")
}

func ActionBindings() carapace.Action {
	return carapace.ActionValues(
		"bin",
		"cffi",
		"pyo3",
		"pyo3-ffi",
		"uniffi",
	).Tag("bindings")
}

func ActionProjectBindings() carapace.Action {
	return carapace.ActionValues(
		"bin",
		"cffi",
		"pyo3",
		"uniffi",
	).Tag("bindings")
}

func ActionCompatibility() carapace.Action {
	return carapace.ActionValues(
		"linux",
		"manylinux2014",
		"manylinux_2_17",
		"manylinux_2_24",
		"manylinux_2_28",
		"musllinux_1_1",
		"musllinux_1_2",
		"pypi",
	).Tag("compatibility")
}

func ActionCompressionMethods() carapace.Action {
	return carapace.ActionValues(
		"bzip2",
		"deflated",
		"stored",
		"zstd",
	).Tag("compression methods")
}

func ActionRepositories() carapace.Action {
	return carapace.ActionValues(
		"pypi",
		"testpypi",
	).Tag("repositories")
}
