# Bridge

Bridging completions from another engine for shells not natively supported by it is possible with the corresponding [Macro](./macros.md).

> Even when the command supports your current shell it is still beneficial to bridge it as this enables [embedding](./embed.md) like `sudo [spec.name] <TAB>`.
>
> It also avoids the issue of [shell startup delay] when sourcing the completion in init scripts otherwise circumvented with [lazycomplete].
>
> However, bridging is limited to supported commands/frameworks and how well it actually works.


## Frameworks

### Argcomplete

[kislyuk/argcomplete] based commands can be bridged with the [`bridge.Argcomplete`] macro:

```yaml
name: az
description: Azure Command-Line Interface
completion:
  positionalany: ["$_bridge.Argcomplete([az])"]
```

### Carapace

[rsteube/carapace] based commands can be bridged with the [`bridge.Carapace`] macro:

```yaml
name: freckles
description: simple dotfile manager
completion:
  positionalany: ["$_bridge.Carapace([freckles])"]
```

### CarapaceBin

[Completers](../completers.md) and [Specs](../spec.md) registered in [rsteube/carapace-bin] can be bridged with the [`bridge.CarapaceBin`] macro:

```yaml
name: github-cli
description: Work seamlessly with GitHub from the command line
completion:
  positionalany: ["$_bridge.CarapaceBin([gh])"]
```

### Clap

[clap-rs/clap] based commands can be bridged with the [`bridge.Clap`] macro:

```yaml
name: dyamic
description: clap dynamic example
completion:
  positionalany: ["$_bridge.Clap([dynamic])"]
```

### Click

[pallets/click] based commands can be bridged with the [`bridge.Click`] macro:

```yaml
name: watson
description: Watson is a tool aimed at helping you monitoring your time
completion:
  positionalany: ["$_bridge.Click([watson])"]
```

### Cobra

[spf13/cobra] based commands can be bridged with the [`bridge.Cobra`] macro:

```yaml
name: kubectl
description: kubectl controls the Kubernetes cluster manager
completion:
  positionalany: ["$_bridge.Cobra([kubectl])"]
```

### Complete
[posener/complete] based commands can be bridged with the [`bridge.Complete`] macro:

```yaml
name: vault
description: Manage Secrets & Protect Sensitive Data
completion:
  positionalany: ["$_bridge.Complete([vault])"]
```

### Inshellisense
Commands provided by [microsoft/inshellisense] can be bridged with the [`bridge.Inshellisense`] macro:

```yaml
name: node
description: Run the node interpreter
completion:
  positionalany: ["$_bridge.Inshellisense([node])"]
```

### Kingpin

[alecthomas/kingpin] based commands can be bridged with the [`bridge.Kingpin`] macro:

```yaml
name: tsh
description: Teleport Command Line Client
completion:
  positionalany: ["$_bridge.Kingpin([tsh])"]
```

### Urfavecli

[urfave/cli] based commands can be bridged with the [`bridge.Urfavecli`] macro:

```yaml
name: tea
description: command line tool to interact with Gitea
completion:
  positionalany: ["$_bridge.Urfavecli([tea])"]
```

### Yargs
[yargs/yargs] based commands can be bridged with the [`bridge.Yargs`] macro:

```yaml
name: ng
description: CLI tool for Angular
completion:
  positionalany: ["$_bridge.Yargs([ng])"]
```

## Shells

> For shells custom configurations are loaded from [`${UserConfigDir}/carapace/bridge`].
> Invoking completion in shells is quite tricky though and edge cases are likely to fail.

 Bash

Commands registered in [bash] can be bridged with the [`bridge.Bash`] macro:

```yaml
name:  tail
description: output the last part of files
completion:
  positionalany: ["$_bridge.Bash([tail])"]
```

### Fish

Commands registered in [fish-shell/fish-shell] can be bridged with the [`bridge.Fish`] macro:

```yaml
name:  git
description: the stupid content tracker
completion:
  positionalany: ["$_bridge.Fish([git])"]
```

### Powershell

Commands registered in [powershell] can be bridged with the [`bridge.Powershell`] macro:

```yaml
name:  ConvertTo-Json
description: convert to json
completion:
  positionalany: ["$_bridge.Powershell([ConvertTo-Json])"]
```

### Zsh

Commands registered in [zsh] can be bridged with the [`bridge.Zsh`] macro:

```yaml
name:  git
description: the stupid content tracker
completion:
  positionalany: ["$_bridge.Zsh([git])"]
```

[lazycomplete]:https://github.com/rsteube/lazycomplete
[shell startup delay]:https://jzelinskie.com/posts/dont-recommend-sourcing-shell-completion/

[bash]:https://www.gnu.org/software/bash/
[`bridge.Bash`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionBash

[kislyuk/argcomplete]:https://github.com/kislyuk/argcomplete
[`bridge.Argcomplete`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionArgcomplete

[rsteube/carapace]:https://github.com/rsteube/carapace
[`bridge.Carapace`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionCarapace

[rsteube/carapace-bin]:https://github.com/rsteube/carapace-bin
[`bridge.CarapaceBin`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionCarapaceBin

[clap-rs/clap]:https://github.com/clap-rs/clap
[`bridge.Clap`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionClap

[pallets/click]:https://github.com/pallets/click
[`bridge.Click`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionClick

[spf13/cobra]:https://github.com/spf13/cobra
[`bridge.Cobra`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionCobra

[posener/complete]:https://github.com/posener/complete
[`bridge.Complete`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionComplete

[microsoft/inshellisense]:https://github.com/microsoft/inshellisense
[`bridge.Inshellisense`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionInshellisense

[alecthomas/kingpin]:https://github.com/alecthomas/kingpin
[`bridge.Kingpin`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionKingpin

[powershell]:https://microsoft.com/powershell
[`bridge.Powershell`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionPowershell

[fish-shell/fish-shell]:https://github.com/fish-shell/fish-shell
[`bridge.Fish`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionFish

[urfave/cli]:https://github.com/urfave/cli
[`bridge.Urfavecli`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionUrfavecli

[yargs/yargs]:https://github.com/yargs/yargs
[`bridge.Yargs`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionYargs

[zsh]:https://www.zsh.org/
[`bridge.Zsh`]:https://pkg.go.dev/github.com/rsteube/carapace-bridge/pkg/actions/bridge#ActionZsh

[`${UserConfigDir}/carapace/bridge`]:https://pkg.go.dev/os#UserConfigDir
