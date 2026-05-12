package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zfs",
	Short: "configure ZFS datasets",
	Long:  "https://openzfs.github.io/openzfs-docs/man/v2.4/8/zfs.8.html",
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

	listCmd := newCommand("list", "list properties of datasets")
	boolFlag(listCmd, "H", "scripted mode")
	boolFlag(listCmd, "p", "display exact numeric values")
	boolFlag(listCmd, "r", "recurse into children")
	boolFlag(listCmd, "j", "display JSON")
	listCmd.Flags().Bool("json-int", false, "display JSON numbers as integers")
	stringFlag(listCmd, "d", "limit recursion depth")
	stringFlag(listCmd, "o", "properties to display")
	stringFlag(listCmd, "s", "sort by property")
	stringFlag(listCmd, "S", "sort by property in reverse")
	stringFlag(listCmd, "t", "dataset types to display")
	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"S": zfs.ActionDatasetProperties(),
		"o": propertyListAction(zfs.ActionDatasetProperties()),
		"s": zfs.ActionDatasetProperties(),
		"t": zfs.ActionDatasetTypes().UniqueList(","),
	})
	carapace.Gen(listCmd).PositionalAnyCompletion(zfs.ActionDatasets("all"))

	createCmd := newCommand("create", "create a new ZFS file system or volume")
	boolFlag(createCmd, "p", "create parent datasets")
	boolFlag(createCmd, "s", "create sparse volume")
	boolFlag(createCmd, "u", "do not mount the created file system")
	stringFlag(createCmd, "o", "set a dataset property")
	stringFlag(createCmd, "V", "create a volume of the given size")
	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(createCmd).PositionalCompletion(
		zfs.ActionDatasets("filesystem", "volume").NoSpace(),
	)

	destroyCmd := newCommand("destroy", "destroy datasets, snapshots, or bookmarks")
	boolFlag(destroyCmd, "d", "defer snapshot deletion")
	boolFlag(destroyCmd, "f", "force unmount")
	boolFlag(destroyCmd, "n", "dry run")
	boolFlag(destroyCmd, "p", "print machine-parsable output")
	boolFlag(destroyCmd, "r", "destroy descendants")
	boolFlag(destroyCmd, "R", "destroy dependents")
	boolFlag(destroyCmd, "v", "print verbose information")
	carapace.Gen(destroyCmd).PositionalAnyCompletion(zfs.ActionDatasets("all"))

	renameCmd := newCommand("rename", "rename a dataset or snapshot")
	boolFlag(renameCmd, "f", "force unmount")
	boolFlag(renameCmd, "p", "create parent datasets")
	boolFlag(renameCmd, "r", "recursively rename snapshots")
	boolFlag(renameCmd, "u", "do not remount")
	carapace.Gen(renameCmd).PositionalCompletion(
		zfs.ActionDatasets("all"),
		zfs.ActionDatasets("filesystem", "volume").NoSpace(),
	)

	upgradeCmd := newCommand("upgrade", "manage upgrading the on-disk dataset version")
	boolFlag(upgradeCmd, "a", "upgrade all datasets")
	boolFlag(upgradeCmd, "r", "upgrade descendants")
	boolFlag(upgradeCmd, "v", "display supported versions")
	carapace.Gen(upgradeCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))

	snapshotCmd := newCommand("snapshot", "create snapshots")
	boolFlag(snapshotCmd, "r", "recursively create snapshots")
	stringFlag(snapshotCmd, "o", "set a snapshot property")
	carapace.Gen(snapshotCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(snapshotCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume").NoSpace().Suffix("@"))

	rollbackCmd := newCommand("rollback", "roll back to a previous snapshot")
	boolFlag(rollbackCmd, "f", "force unmount")
	boolFlag(rollbackCmd, "R", "destroy later snapshots and clones")
	boolFlag(rollbackCmd, "r", "destroy later snapshots")
	carapace.Gen(rollbackCmd).PositionalCompletion(zfs.ActionSnapshots())

	holdCmd := newCommand("hold", "add a hold reference to snapshots")
	boolFlag(holdCmd, "r", "recursively hold snapshots")
	carapace.Gen(holdCmd).PositionalAnyCompletion(positionalByIndex(
		carapace.ActionValues(),
		zfs.ActionSnapshots(),
	))

	holdsCmd := newCommand("holds", "list hold references on snapshots")
	boolFlag(holdsCmd, "H", "scripted mode")
	boolFlag(holdsCmd, "r", "recursively list holds")
	carapace.Gen(holdsCmd).PositionalAnyCompletion(zfs.ActionSnapshots())

	releaseCmd := newCommand("release", "remove a hold reference from snapshots")
	boolFlag(releaseCmd, "r", "recursively release holds")
	carapace.Gen(releaseCmd).PositionalAnyCompletion(positionalByIndex(
		carapace.ActionValues(),
		zfs.ActionSnapshots(),
	))

	diffCmd := newCommand("diff", "display differences between snapshots")
	boolFlag(diffCmd, "F", "display file type")
	boolFlag(diffCmd, "H", "scripted mode")
	boolFlag(diffCmd, "t", "display path change time")
	carapace.Gen(diffCmd).PositionalAnyCompletion(zfs.ActionSnapshots())

	cloneCmd := newCommand("clone", "create a writable clone of a snapshot")
	boolFlag(cloneCmd, "p", "create parent datasets")
	stringFlag(cloneCmd, "o", "set a dataset property")
	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(cloneCmd).PositionalCompletion(
		zfs.ActionSnapshots(),
		zfs.ActionDatasets("filesystem", "volume").NoSpace(),
	)

	promoteCmd := newCommand("promote", "promote a clone dataset")
	carapace.Gen(promoteCmd).PositionalCompletion(zfs.ActionDatasets("filesystem", "volume"))

	sendCmd := newCommand("send", "generate a send stream")
	boolFlag(sendCmd, "D", "deduplicate the stream")
	boolFlag(sendCmd, "L", "allow large blocks")
	boolFlag(sendCmd, "P", "print parsable stream package size")
	boolFlag(sendCmd, "R", "replicate descendants")
	boolFlag(sendCmd, "c", "write compressed data")
	boolFlag(sendCmd, "e", "embed blocks")
	boolFlag(sendCmd, "h", "include holds")
	boolFlag(sendCmd, "n", "dry run")
	boolFlag(sendCmd, "p", "include properties")
	boolFlag(sendCmd, "s", "resume interrupted send")
	boolFlag(sendCmd, "v", "verbose output")
	boolFlag(sendCmd, "w", "raw encrypted send")
	stringFlag(sendCmd, "i", "incremental source")
	stringFlag(sendCmd, "I", "incremental source with intermediates")
	stringFlag(sendCmd, "t", "resume token")
	carapace.Gen(sendCmd).FlagCompletion(carapace.ActionMap{
		"I": carapace.Batch(zfs.ActionSnapshots(), zfs.ActionBookmarks()).ToA(),
		"i": carapace.Batch(zfs.ActionSnapshots(), zfs.ActionBookmarks()).ToA(),
	})
	carapace.Gen(sendCmd).PositionalCompletion(carapace.Batch(zfs.ActionSnapshots(), zfs.ActionBookmarks()).ToA())

	receiveCmd := newCommand("receive", "receive a ZFS send stream", "recv")
	boolFlag(receiveCmd, "F", "force rollback")
	boolFlag(receiveCmd, "M", "force unmount")
	boolFlag(receiveCmd, "d", "discard leading path components")
	boolFlag(receiveCmd, "e", "discard all but the last path component")
	boolFlag(receiveCmd, "n", "dry run")
	boolFlag(receiveCmd, "s", "save partially received state")
	boolFlag(receiveCmd, "u", "do not mount")
	boolFlag(receiveCmd, "v", "verbose output")
	stringFlag(receiveCmd, "o", "override a property")
	stringFlag(receiveCmd, "x", "exclude a property")
	carapace.Gen(receiveCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
		"x": zfs.ActionDatasetProperties(),
	})
	carapace.Gen(receiveCmd).PositionalCompletion(zfs.ActionDatasets("filesystem", "volume").NoSpace())

	bookmarkCmd := newCommand("bookmark", "create a bookmark of a snapshot or bookmark")
	carapace.Gen(bookmarkCmd).PositionalCompletion(
		carapace.Batch(zfs.ActionSnapshots(), zfs.ActionBookmarks()).ToA(),
		zfs.ActionDatasets("filesystem", "volume").NoSpace().Suffix("#"),
	)

	redactCmd := newCommand("redact", "generate a redaction bookmark")
	carapace.Gen(redactCmd).PositionalAnyCompletion(carapace.Batch(zfs.ActionSnapshots(), zfs.ActionBookmarks()).ToA())

	getCmd := newCommand("get", "display dataset properties")
	boolFlag(getCmd, "H", "scripted mode")
	boolFlag(getCmd, "p", "display exact numeric values")
	boolFlag(getCmd, "r", "recurse into children")
	stringFlag(getCmd, "d", "limit recursion depth")
	stringFlag(getCmd, "o", "output fields")
	stringFlag(getCmd, "s", "property sources")
	stringFlag(getCmd, "t", "dataset types")
	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionGetFields().UniqueList(","),
		"s": zfs.ActionPropertySources().UniqueList(","),
		"t": zfs.ActionDatasetTypes().UniqueList(","),
	})
	carapace.Gen(getCmd).PositionalAnyCompletion(positionalByIndex(
		propertyListAction(zfs.ActionDatasetProperties()),
		zfs.ActionDatasets("all"),
	))

	setCmd := newCommand("set", "set dataset properties")
	carapace.Gen(setCmd).PositionalAnyCompletion(positionalByIndex(
		propertyAssignmentAction(),
		zfs.ActionDatasets("filesystem", "volume"),
	))

	inheritCmd := newCommand("inherit", "inherit a dataset property")
	boolFlag(inheritCmd, "S", "revert to received value")
	boolFlag(inheritCmd, "r", "recurse into children")
	carapace.Gen(inheritCmd).PositionalAnyCompletion(positionalByIndex(
		zfs.ActionDatasetProperties(),
		zfs.ActionDatasets("filesystem", "volume"),
	))

	userspaceCmd := newSpaceCommand("userspace", "display user space accounting")
	groupspaceCmd := newSpaceCommand("groupspace", "display group space accounting")
	projectspaceCmd := newSpaceCommand("projectspace", "display project space accounting")
	carapace.Gen(userspaceCmd).PositionalCompletion(zfs.ActionDatasets("filesystem", "snapshot"))
	carapace.Gen(groupspaceCmd).PositionalCompletion(zfs.ActionDatasets("filesystem", "snapshot"))
	carapace.Gen(projectspaceCmd).PositionalCompletion(zfs.ActionDatasets("filesystem", "snapshot"))

	mountCmd := newCommand("mount", "mount ZFS file systems")
	boolFlag(mountCmd, "O", "overlay mount")
	boolFlag(mountCmd, "a", "mount all available file systems")
	boolFlag(mountCmd, "l", "load keys for encrypted datasets")
	boolFlag(mountCmd, "v", "verbose output")
	stringFlag(mountCmd, "o", "temporary mount options")
	carapace.Gen(mountCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem"))

	unmountCmd := newCommand("unmount", "unmount ZFS file systems", "umount")
	boolFlag(unmountCmd, "a", "unmount all ZFS file systems")
	boolFlag(unmountCmd, "f", "force unmount")
	boolFlag(unmountCmd, "u", "do not unload keys")
	carapace.Gen(unmountCmd).PositionalAnyCompletion(carapace.Batch(
		zfs.ActionDatasets("filesystem"),
		carapace.ActionDirectories(),
	).ToA())

	shareCmd := newCommand("share", "share ZFS file systems")
	boolFlag(shareCmd, "a", "share all file systems")
	carapace.Gen(shareCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem"))

	unshareCmd := newCommand("unshare", "unshare ZFS file systems")
	boolFlag(unshareCmd, "a", "unshare all file systems")
	carapace.Gen(unshareCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem"))

	allowCmd := newCommand("allow", "delegate permissions on a dataset")
	boolFlag(allowCmd, "c", "set create-time permissions")
	boolFlag(allowCmd, "d", "set descendant permissions")
	boolFlag(allowCmd, "e", "set everyone permissions")
	boolFlag(allowCmd, "g", "set group permissions")
	boolFlag(allowCmd, "l", "set local permissions")
	boolFlag(allowCmd, "s", "define a permission set")
	boolFlag(allowCmd, "u", "set user permissions")
	carapace.Gen(allowCmd).PositionalAnyCompletion(carapace.Batch(
		zfs.ActionPermissions(),
		zfs.ActionDatasets("filesystem", "volume"),
	).ToA())

	unallowCmd := newCommand("unallow", "remove delegated permissions")
	boolFlag(unallowCmd, "d", "remove descendant permissions")
	boolFlag(unallowCmd, "e", "remove everyone permissions")
	boolFlag(unallowCmd, "g", "remove group permissions")
	boolFlag(unallowCmd, "l", "remove local permissions")
	boolFlag(unallowCmd, "r", "remove permissions recursively")
	boolFlag(unallowCmd, "u", "remove user permissions")
	carapace.Gen(unallowCmd).PositionalAnyCompletion(carapace.Batch(
		zfs.ActionPermissions(),
		zfs.ActionDatasets("filesystem", "volume"),
	).ToA())

	changeKeyCmd := newCommand("change-key", "add or change an encryption key")
	boolFlag(changeKeyCmd, "i", "inherit key location")
	boolFlag(changeKeyCmd, "l", "load the key after changing it")
	stringFlag(changeKeyCmd, "o", "set an encryption property")
	carapace.Gen(changeKeyCmd).FlagCompletion(carapace.ActionMap{
		"o": propertyAssignmentAction(),
	})
	carapace.Gen(changeKeyCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))

	loadKeyCmd := newCommand("load-key", "load encryption keys")
	boolFlag(loadKeyCmd, "a", "load all keys")
	boolFlag(loadKeyCmd, "n", "dry run")
	boolFlag(loadKeyCmd, "r", "recursively load keys")
	stringFlag(loadKeyCmd, "L", "key location")
	carapace.Gen(loadKeyCmd).FlagCompletion(carapace.ActionMap{
		"L": carapace.ActionFiles(),
	})
	carapace.Gen(loadKeyCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))

	unloadKeyCmd := newCommand("unload-key", "unload encryption keys")
	boolFlag(unloadKeyCmd, "a", "unload all keys")
	boolFlag(unloadKeyCmd, "r", "recursively unload keys")
	carapace.Gen(unloadKeyCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))

	programCmd := newCommand("program", "execute a channel program")
	boolFlag(programCmd, "j", "display JSON output")
	boolFlag(programCmd, "n", "dry run")
	stringFlag(programCmd, "t", "terminate after the given timeout")
	carapace.Gen(programCmd).PositionalCompletion(
		zfs.ActionDatasets("filesystem", "volume"),
		carapace.ActionFiles(),
	)

	projectCmd := newCommand("project", "manage project IDs")
	boolFlag(projectCmd, "d", "clear project ID")
	boolFlag(projectCmd, "r", "recurse into directories")
	boolFlag(projectCmd, "s", "set project ID")
	boolFlag(projectCmd, "C", "clear inherit flag")
	boolFlag(projectCmd, "c", "check project ID")
	carapace.Gen(projectCmd).PositionalAnyCompletion(carapace.ActionFiles())

	rewriteCmd := newCommand("rewrite", "rewrite files without modification")
	carapace.Gen(rewriteCmd).PositionalAnyCompletion(carapace.ActionFiles())

	jailCmd := newCommand("jail", "attach a dataset to a jail", "zone")
	carapace.Gen(jailCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))

	unjailCmd := newCommand("unjail", "detach a dataset from a jail", "unzone")
	carapace.Gen(unjailCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))

	waitCmd := newCommand("wait", "wait for background activity to complete")
	boolFlag(waitCmd, "H", "scripted mode")
	boolFlag(waitCmd, "p", "display exact numeric values")
	stringFlag(waitCmd, "t", "activity type")
	carapace.Gen(waitCmd).FlagCompletion(carapace.ActionMap{
		"t": zfs.ActionWaitActivities(),
	})
	carapace.Gen(waitCmd).PositionalAnyCompletion(zfs.ActionDatasets("filesystem", "volume"))
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

func newSpaceCommand(use string, short string) *cobra.Command {
	cmd := newCommand(use, short)
	boolFlag(cmd, "H", "scripted mode")
	boolFlag(cmd, "i", "translate SID to POSIX ID")
	boolFlag(cmd, "n", "display numeric IDs")
	boolFlag(cmd, "p", "display exact numeric values")
	stringFlag(cmd, "o", "output fields")
	stringFlag(cmd, "s", "sort by field")
	stringFlag(cmd, "S", "sort by field in reverse")
	stringFlag(cmd, "t", "types to display")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"S": zfs.ActionUserSpaceFields(),
		"o": propertyListAction(zfs.ActionUserSpaceFields()),
		"s": zfs.ActionUserSpaceFields(),
		"t": zfs.ActionUserSpaceTypes().UniqueList(","),
	})
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
