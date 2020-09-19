module github.com/rsteube/carapace-bin

go 1.12

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-ps v1.0.0
	github.com/rsteube/carapace v0.0.23
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cobra v1.0.0
	gopkg.in/ini.v1 v1.60.2
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v0.0.0-20200824165833-dd6f6588b61d
