package cmd

import (
	bat "github.com/rsteube/carapace-completers/completers/bat_completer/cmd"
	cat "github.com/rsteube/carapace-completers/completers/cat_completer/cmd"
	chgrp "github.com/rsteube/carapace-completers/completers/chgrp_completer/cmd"
	chmod "github.com/rsteube/carapace-completers/completers/chmod_completer/cmd"
	chown "github.com/rsteube/carapace-completers/completers/chown_completer/cmd"
	cksum "github.com/rsteube/carapace-completers/completers/cksum_completer/cmd"
	comm "github.com/rsteube/carapace-completers/completers/comm_completer/cmd"
	cp "github.com/rsteube/carapace-completers/completers/cp_completer/cmd"
	csplit "github.com/rsteube/carapace-completers/completers/csplit_completer/cmd"
	curl "github.com/rsteube/carapace-completers/completers/curl_completer/cmd"
	cut "github.com/rsteube/carapace-completers/completers/cut_completer/cmd"
	df "github.com/rsteube/carapace-completers/completers/df_completer/cmd"
	dircolors "github.com/rsteube/carapace-completers/completers/dircolors_completer/cmd"
	dir "github.com/rsteube/carapace-completers/completers/dir_completer/cmd"
	exa "github.com/rsteube/carapace-completers/completers/exa_completer/cmd"
	expand "github.com/rsteube/carapace-completers/completers/expand_completer/cmd"
	fmt "github.com/rsteube/carapace-completers/completers/fmt_completer/cmd"
	git "github.com/rsteube/carapace-completers/completers/git_completer/cmd"
	gradle "github.com/rsteube/carapace-completers/completers/gradle_completer/cmd"
	install "github.com/rsteube/carapace-completers/completers/install_completer/cmd"
	ln "github.com/rsteube/carapace-completers/completers/ln_completer/cmd"
	ls "github.com/rsteube/carapace-completers/completers/ls_completer/cmd"
	mkdir "github.com/rsteube/carapace-completers/completers/mkdir_completer/cmd"
	mkfifo "github.com/rsteube/carapace-completers/completers/mkfifo_completer/cmd"
	mknod "github.com/rsteube/carapace-completers/completers/mknod_completer/cmd"
	mv "github.com/rsteube/carapace-completers/completers/mv_completer/cmd"
	pkill "github.com/rsteube/carapace-completers/completers/pkill_completer/cmd"
	rm "github.com/rsteube/carapace-completers/completers/rm_completer/cmd"
	rmdir "github.com/rsteube/carapace-completers/completers/rmdir_completer/cmd"
	sha1sum "github.com/rsteube/carapace-completers/completers/sha1sum_completer/cmd"
	shred "github.com/rsteube/carapace-completers/completers/shred_completer/cmd"
	sort "github.com/rsteube/carapace-completers/completers/sort_completer/cmd"
	sync "github.com/rsteube/carapace-completers/completers/sync_completer/cmd"
	tee "github.com/rsteube/carapace-completers/completers/tee_completer/cmd"
	touch "github.com/rsteube/carapace-completers/completers/touch_completer/cmd"
	vdir "github.com/rsteube/carapace-completers/completers/vdir_completer/cmd"
)

var completers = []string{
	"bat",
	"cat",
	"chgrp",
	"chmod",
	"chown",
	"cksum",
	"comm",
	"cp",
	"csplit",
	"curl",
	"cut",
	"df",
	"dircolors",
	"dir",
	"exa",
	"expand",
	"fmt",
	"git",
	"gradle",
	"install",
	"ln",
	"ls",
	"mkdir",
	"mkfifo",
	"mknod",
	"mv",
	"pkill",
	"rm",
	"rmdir",
	"sha1sum",
	"shred",
	"sort",
	"sync",
	"tee",
	"touch",
	"vdir",
}

func executeCompleter(completer string) {
	switch completer {
	case "bat":
		bat.Execute()
	case "cat":
		cat.Execute()
	case "chgrp":
		chgrp.Execute()
	case "chmod":
		chmod.Execute()
	case "chown":
		chown.Execute()
	case "cksum":
		cksum.Execute()
	case "comm":
		comm.Execute()
	case "cp":
		cp.Execute()
	case "csplit":
		csplit.Execute()
	case "curl":
		curl.Execute()
	case "cut":
		cut.Execute()
	case "df":
		df.Execute()
	case "dircolors":
		dircolors.Execute()
	case "dir":
		dir.Execute()
	case "exa":
		exa.Execute()
	case "expand":
		expand.Execute()
	case "fmt":
		fmt.Execute()
	case "git":
		git.Execute()
	case "gradle":
		gradle.Execute()
	case "install":
		install.Execute()
	case "ln":
		ln.Execute()
	case "ls":
		ls.Execute()
	case "mkdir":
		mkdir.Execute()
	case "mkfifo":
		mkfifo.Execute()
	case "mknod":
		mknod.Execute()
	case "mv":
		mv.Execute()
	case "pkill":
		pkill.Execute()
	case "rm":
		rm.Execute()
	case "rmdir":
		rmdir.Execute()
	case "sha1sum":
		sha1sum.Execute()
	case "shred":
		shred.Execute()
	case "sort":
		sort.Execute()
	case "sync":
		sync.Execute()
	case "tee":
		tee.Execute()
	case "touch":
		touch.Execute()
	case "vdir":
		vdir.Execute()
	}
}

