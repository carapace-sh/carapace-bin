# Bridge

## Argcomplete

[kislyuk/argcomplete] based commands can be bridged with the [bridge.Argcomplete(\<COMMAND\>)] macro:

```yaml
name: az
description: Azure Command-Line Interface
completion:
  positionalany: ["$_bridge.Argcomplete(az)"]
````

## Carapace

[rsteube/carapace] based commands can be bridged with the [bridge.Carapace(\<COMMAND\>)] macro:

```yaml
name: freckles
description: simple dotfile manager
completion:
  positionalany: ["$_bridge.Carapace(freckles)"]
````

## CarapaceBin

[Completers](../completers.md) and [Specs](../spec.md) registered in [rsteube/carapace-bin] can be bridged with the [bridge.CarapaceBin(\<COMMAND\>)] macro:

```yaml
name: gh
description: Work seamlessly with GitHub from the command line
completion:
  positionalany: ["$_bridge.CarapaceBin(gh)"]
````

## Click

[pallets/click] based commands can be bridged with the [bridge.Click(\<COMMAND\>)] macro:

```yaml
name: watson
description: Watson is a tool aimed at helping you monitoring your time
completion:
  positionalany: ["$_bridge.Click(watson)"]
````

## Cobra

[spf13/cobra] based commands can be bridged with the [bridge.Cobra(\<COMMAND\>)] macro:

```yaml
name: kubectl
description: kubectl controls the Kubernetes cluster manager
completion:
  positionalany: ["$_bridge.Cobra(kubectl)"]
````

## Fish

Commands registered in [fish-shell/fish-shell] can be bridged with the [bridge.Fish(\<COMMAND\>)] macro:

```yaml
name:  git
description: the stupid content tracker
completion:
  positionalany: ["$_bridge.Fish(git)"]
````

## Posener (TODO: rename to `Complete`)
[posener/complete] based commands can be bridged with the [bridge.Posener(\<COMMAND\>)] macro:

```yaml
name: vault
description: Manage Secrets & Protect Sensitive Data
completion:
  positionalany: ["$_bridge.Posener(vault)"]
````


[kislyuk/argcomplete]:https://github.com/kislyuk/argcomplete
[bridge.Argcomplete(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionArgcomplete

[rsteube/carapace]:https://github.com/rsteube/carapace
[bridge.Carapace(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionCarapace

[rsteube/carapace-bin]:https://github.com/rsteube/carapace-bin
[bridge.CarapaceBin(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionCarapaceBin

[pallets/click]:https://github.com/pallets/click
[bridge.Click(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionClick

[spf13/cobra]:https://github.com/spf13/cobra
[bridge.Cobra(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionCobra

[fish-shell/fish-shell]:https://github.com/fish-shell/fish-shell
[bridge.Fish(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionFish

[posener/complete]:https://github.com/posener/complete
[bridge.Posener(\<COMMAND\>)]:https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/bridge#ActionPosener
