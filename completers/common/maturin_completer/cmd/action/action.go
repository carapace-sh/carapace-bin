package action

import "github.com/carapace-sh/carapace"

func ActionAuditwheel() carapace.Action {
	return carapace.ActionValues(
		"repair",
		"check",
		"warn",
		"skip",
	).Tag("auditwheel policies")
}

func ActionBindings() carapace.Action {
	return carapace.ActionValues(
		"pyo3",
		"pyo3-ffi",
		"cffi",
		"uniffi",
		"bin",
	).Tag("bindings")
}

func ActionProjectBindings() carapace.Action {
	return carapace.ActionValues(
		"pyo3",
		"cffi",
		"uniffi",
		"bin",
	).Tag("bindings")
}

func ActionCompatibility() carapace.Action {
	return carapace.ActionValues(
		"pypi",
		"linux",
		"manylinux2014",
		"manylinux_2_17",
		"manylinux_2_24",
		"manylinux_2_28",
		"musllinux_1_1",
		"musllinux_1_2",
	).Tag("compatibility")
}

func ActionCompressionMethods() carapace.Action {
	return carapace.ActionValues(
		"deflated",
		"stored",
		"bzip2",
		"zstd",
	).Tag("compression methods")
}

func ActionRepositories() carapace.Action {
	return carapace.ActionValues(
		"pypi",
		"testpypi",
	).Tag("repositories")
}
