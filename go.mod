module github.com/rsteube/carapace-bin

go 1.12

require (
	github.com/mitchellh/go-ps v1.0.0
	github.com/rsteube/carapace v0.0.17
	github.com/spf13/cobra v1.0.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

// use the shorthand PR
replace github.com/rsteube/carapace v0.0.17 => github.com/rsteube/carapace v0.0.17-0.20200902095507-4a5937d65250

replace github.com/spf13/pflag v1.0.5 => github.com/cornfeedhobo/pflag v1.0.2-0.20200824165833-dd6f6588b61d
