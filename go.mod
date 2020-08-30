module github.com/rsteube/carapace-completers

go 1.12

require (
	github.com/mitchellh/go-ps v1.0.0
	github.com/rsteube/carapace v0.0.17-0.20200830151704-26070a312697
	github.com/spf13/cobra v1.0.0
)

// use the shorthand PR
replace github.com/rsteube/carapace v0.0.16 => github.com/rsteube/carapace v0.0.17-0.20200828122823-e03a70094000

replace github.com/spf13/pflag v1.0.5 => github.com/cornfeedhobo/pflag v1.0.2-0.20200824165833-dd6f6588b61d
