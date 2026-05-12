package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zpool"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zpool",
	Short: "configure ZFS storage pools",
	Long:  "https://openzfs.github.io/openzfs-docs/man/v2.4/8/zpool.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "display help")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	versionCmd := newCommand("version", "display ZFS version")
	versionCmd.Flags().BoolS("j", "j", false, "display JSON")

	listCmd := newCommand("list", "list storage pools")
	boolFlag(listCmd, "H", "scripted mode")
	boolFlag(listCmd, "L", "display real paths for vdevs")
	boolFlag(listCmd, "P", "display full paths for vdevs")
	boolFlag(listCmd, "g", "display vdev GUIDs")
	boolFlag(listCmd, "j", "display JSON")
	boolFlag(listCmd, "p", "display exact numeric values")
	boolFlag(listCmd, "v", "display vdev statistics")
	listCmd.Flags().Bool("json-int", false, "display JSON numbers as integers")
	listCmd.Flags().Bool("json-pool-key-guid", false, "key JSON pools by GUID")
	stringFlag(listCmd, "T", "timestamp style")
	stringFlag(listCmd, "o", "properties to display")
	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"T": zpool.ActionTimestampStyles(),
		"o": propertyListAction(zpool.ActionListProperties()),
	})
	carapace.Gen(listCmd).PositionalAnyCompletion(zpool.ActionPools())

	createCmd := newCommand("create", "create a new storage pool")
	boolFlag(createCmd, "d", "do not enable features automatically")
	boolFlag(createCmd, "f", "force use of devices")
	boolFlag(createCmd, "n", "dry run")
	stringFlag(createCmd, "R", "alternate root directory")
	stringFlag(createCmd, "m", "mount point")
	stringFlag(createCmd, "O", "set a root dataset property")
	stringFlag(createCmd, "o", "set a pool property")
	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"O": datasetPropertyAssignmentAction(),
		"R": carapace.ActionDirectories(),
		"m": carapace.ActionDirectories(),
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(createCmd).PositionalAnyCompletion(positionalByIndex(
		carapace.ActionValues(),
		vdevAction(),
	))

	addCmd := newCommand("add", "add virtual devices to a pool")
	boolFlag(addCmd, "L", "display real paths for vdevs")
	boolFlag(addCmd, "f", "force use of devices")
	boolFlag(addCmd, "g", "display vdev GUIDs")
	boolFlag(addCmd, "n", "dry run")
	stringFlag(addCmd, "o", "set a pool property")
	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(addCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		vdevAction(),
	))

	attachCmd := newCommand("attach", "attach a device to a pool")
	boolFlag(attachCmd, "f", "force use of device")
	stringFlag(attachCmd, "o", "set a pool property")
	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(attachCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	checkpointCmd := newCommand("checkpoint", "checkpoint a pool")
	checkpointCmd.Flags().Bool("discard", false, "discard an existing checkpoint")
	carapace.Gen(checkpointCmd).PositionalAnyCompletion(zpool.ActionPools())

	clearCmd := newCommand("clear", "clear device errors in a pool")
	boolFlag(clearCmd, "F", "rewind to checkpoint")
	boolFlag(clearCmd, "n", "dry run")
	carapace.Gen(clearCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	destroyCmd := newCommand("destroy", "destroy a storage pool")
	boolFlag(destroyCmd, "f", "force destroy")
	carapace.Gen(destroyCmd).PositionalAnyCompletion(zpool.ActionPools())

	detachCmd := newCommand("detach", "detach a device from a pool")
	carapace.Gen(detachCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	eventsCmd := newCommand("events", "list recent ZFS events")
	boolFlag(eventsCmd, "H", "scripted mode")
	boolFlag(eventsCmd, "c", "clear events")
	boolFlag(eventsCmd, "f", "follow event stream")
	boolFlag(eventsCmd, "v", "verbose output")

	exportCmd := newCommand("export", "export pools from the system")
	boolFlag(exportCmd, "a", "export all pools")
	boolFlag(exportCmd, "f", "force export")
	carapace.Gen(exportCmd).PositionalAnyCompletion(zpool.ActionPools())

	getCmd := newCommand("get", "retrieve pool properties")
	boolFlag(getCmd, "H", "scripted mode")
	boolFlag(getCmd, "p", "display exact numeric values")
	stringFlag(getCmd, "o", "output fields")
	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"o": zpool.ActionGetFields().UniqueList(","),
	})
	carapace.Gen(getCmd).PositionalAnyCompletion(positionalByIndex(
		propertyListAction(zpool.ActionPoolProperties()),
		zpool.ActionPools(),
	))

	historyCmd := newCommand("history", "display pool command history")
	boolFlag(historyCmd, "i", "display internal events")
	boolFlag(historyCmd, "l", "display long format")
	carapace.Gen(historyCmd).PositionalAnyCompletion(zpool.ActionPools())

	importCmd := newCommand("import", "import pools into the system")
	boolFlag(importCmd, "D", "include destroyed pools")
	boolFlag(importCmd, "F", "rewind if needed")
	boolFlag(importCmd, "N", "do not mount datasets")
	boolFlag(importCmd, "a", "import all pools")
	boolFlag(importCmd, "f", "force import")
	boolFlag(importCmd, "l", "load keys")
	boolFlag(importCmd, "m", "allow missing log devices")
	boolFlag(importCmd, "n", "dry run")
	stringFlag(importCmd, "R", "alternate root directory")
	stringFlag(importCmd, "c", "cache file")
	stringFlag(importCmd, "d", "device search directory")
	stringFlag(importCmd, "o", "set a pool property")
	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"R": carapace.ActionDirectories(),
		"c": carapace.ActionFiles(),
		"d": carapace.ActionDirectories(),
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(importCmd).PositionalAnyCompletion(zpool.ActionPools())

	initializeCmd := newCommand("initialize", "initialize pool devices")
	boolFlag(initializeCmd, "c", "cancel initialization")
	boolFlag(initializeCmd, "s", "suspend initialization")
	boolFlag(initializeCmd, "u", "uninitialize")
	boolFlag(initializeCmd, "w", "wait for completion")
	carapace.Gen(initializeCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	iostatCmd := newCommand("iostat", "display pool I/O statistics")
	boolFlag(iostatCmd, "H", "scripted mode")
	boolFlag(iostatCmd, "L", "display real paths for vdevs")
	boolFlag(iostatCmd, "P", "display full paths for vdevs")
	boolFlag(iostatCmd, "p", "display exact numeric values")
	boolFlag(iostatCmd, "r", "display request size histograms")
	boolFlag(iostatCmd, "v", "display vdev statistics")
	boolFlag(iostatCmd, "w", "display latency histograms")
	boolFlag(iostatCmd, "y", "omit the first report")
	stringFlag(iostatCmd, "T", "timestamp style")
	stringFlag(iostatCmd, "c", "run a command after each report")
	carapace.Gen(iostatCmd).FlagCompletion(carapace.ActionMap{
		"T": zpool.ActionTimestampStyles(),
	})
	carapace.Gen(iostatCmd).PositionalAnyCompletion(zpool.ActionPools())

	labelclearCmd := newCommand("labelclear", "remove ZFS labels from a device")
	boolFlag(labelclearCmd, "f", "force label removal")
	carapace.Gen(labelclearCmd).PositionalAnyCompletion(deviceAction())

	offlineCmd := newCommand("offline", "take devices offline")
	boolFlag(offlineCmd, "f", "force fault")
	boolFlag(offlineCmd, "t", "temporarily offline")
	carapace.Gen(offlineCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	onlineCmd := newCommand("online", "bring devices online")
	boolFlag(onlineCmd, "e", "expand devices")
	carapace.Gen(onlineCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	prefetchCmd := newCommand("prefetch", "prefetch pool data")
	carapace.Gen(prefetchCmd).PositionalAnyCompletion(zpool.ActionPools())

	reguidCmd := newCommand("reguid", "generate a new pool GUID")
	carapace.Gen(reguidCmd).PositionalAnyCompletion(zpool.ActionPools())

	removeCmd := newCommand("remove", "remove devices from a pool")
	boolFlag(removeCmd, "n", "dry run")
	boolFlag(removeCmd, "p", "display exact numeric values")
	boolFlag(removeCmd, "w", "wait for removal")
	carapace.Gen(removeCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	reopenCmd := newCommand("reopen", "reopen pool vdevs")
	boolFlag(reopenCmd, "n", "dry run")
	carapace.Gen(reopenCmd).PositionalAnyCompletion(zpool.ActionPools())

	replaceCmd := newCommand("replace", "replace a device in a pool")
	boolFlag(replaceCmd, "f", "force use of device")
	stringFlag(replaceCmd, "o", "set a pool property")
	carapace.Gen(replaceCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(replaceCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	resilverCmd := newCommand("resilver", "start a resilver")
	carapace.Gen(resilverCmd).PositionalAnyCompletion(zpool.ActionPools())

	scrubCmd := newCommand("scrub", "start or resume a scrub")
	boolFlag(scrubCmd, "p", "pause scrub")
	boolFlag(scrubCmd, "s", "stop scrub")
	boolFlag(scrubCmd, "w", "wait for completion")
	carapace.Gen(scrubCmd).PositionalAnyCompletion(zpool.ActionPools())

	setCmd := newCommand("set", "set pool properties")
	carapace.Gen(setCmd).PositionalAnyCompletion(positionalByIndex(
		propertyAssignmentAction(),
		zpool.ActionPools(),
	))

	splitCmd := newCommand("split", "split a mirrored pool")
	boolFlag(splitCmd, "g", "display vdev GUIDs")
	boolFlag(splitCmd, "L", "display real paths for vdevs")
	boolFlag(splitCmd, "n", "dry run")
	stringFlag(splitCmd, "R", "alternate root directory")
	stringFlag(splitCmd, "o", "set a pool property")
	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"R": carapace.ActionDirectories(),
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(splitCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		carapace.ActionValues(),
		deviceAction(),
	))

	statusCmd := newCommand("status", "display pool health status")
	boolFlag(statusCmd, "D", "display dedup statistics")
	boolFlag(statusCmd, "L", "display real paths for vdevs")
	boolFlag(statusCmd, "P", "display full paths for vdevs")
	boolFlag(statusCmd, "e", "display slow I/O statistics")
	boolFlag(statusCmd, "g", "display vdev GUIDs")
	boolFlag(statusCmd, "j", "display JSON")
	boolFlag(statusCmd, "s", "display slow vdevs")
	boolFlag(statusCmd, "v", "verbose output")
	boolFlag(statusCmd, "x", "display only unhealthy pools")
	stringFlag(statusCmd, "c", "run vdev script")
	stringFlag(statusCmd, "t", "health status filter")
	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"t": carapace.ActionValues("ONLINE", "DEGRADED", "FAULTED", "OFFLINE", "REMOVED", "UNAVAIL"),
	})
	carapace.Gen(statusCmd).PositionalAnyCompletion(zpool.ActionPools())

	syncCmd := newCommand("sync", "force dirty data to disk")
	carapace.Gen(syncCmd).PositionalAnyCompletion(zpool.ActionPools())

	trimCmd := newCommand("trim", "initiate TRIM on a pool")
	boolFlag(trimCmd, "d", "secure trim")
	boolFlag(trimCmd, "r", "trim assigned ranges")
	boolFlag(trimCmd, "w", "wait for completion")
	carapace.Gen(trimCmd).PositionalAnyCompletion(positionalByIndex(
		zpool.ActionPools(),
		deviceAction(),
	))

	upgradeCmd := newCommand("upgrade", "manage pool on-disk format versions")
	boolFlag(upgradeCmd, "a", "upgrade all pools")
	boolFlag(upgradeCmd, "v", "display supported versions")
	stringFlag(upgradeCmd, "V", "upgrade to a specific version")
	carapace.Gen(upgradeCmd).PositionalAnyCompletion(zpool.ActionPools())

	waitCmd := newCommand("wait", "wait for pool background activity")
	boolFlag(waitCmd, "H", "scripted mode")
	boolFlag(waitCmd, "p", "display exact numeric values")
	stringFlag(waitCmd, "t", "activity type")
	carapace.Gen(waitCmd).FlagCompletion(carapace.ActionMap{
		"t": zpool.ActionWaitActivities(),
	})
	carapace.Gen(waitCmd).PositionalAnyCompletion(zpool.ActionPools())

	ddtpruneCmd := newCommand("ddtprune", "prune deduplication tables")
	carapace.Gen(ddtpruneCmd).PositionalAnyCompletion(zpool.ActionPools())
}

func newCommand(use string, short string, aliases ...string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Aliases: aliases,
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()
	rootCmd.AddCommand(cmd)
	return cmd
}

func boolFlag(cmd *cobra.Command, name string, description string) {
	cmd.Flags().BoolS(name, name, false, description)
}

func stringFlag(cmd *cobra.Command, name string, description string) {
	cmd.Flags().StringS(name, name, "", description)
}

func propertyListAction(action carapace.Action) carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		return action.Invoke(c).Filter(c.Parts...).ToA().NoSpace()
	})
}

func propertyAssignmentAction() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return zpool.ActionPoolProperties().Suffix("=")
		default:
			return carapace.ActionValues()
		}
	})
}

func datasetPropertyAssignmentAction() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return zfs.ActionDatasetProperties().Suffix("=")
		default:
			return carapace.ActionValues()
		}
	})
}

func positionalByIndex(actions ...carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) < len(actions) {
			return actions[len(c.Args)]
		}
		return actions[len(actions)-1]
	})
}

func deviceAction() carapace.Action {
	return carapace.Batch(
		fs.ActionBlockDevices(),
		carapace.ActionFiles(),
	).ToA()
}

func vdevAction() carapace.Action {
	return carapace.Batch(
		zpool.ActionVdevTypes(),
		fs.ActionBlockDevices(),
		carapace.ActionFiles(),
	).ToA()
}
