package cmd

import (
	basename "github.com/rsteube/carapace-completers/completers/basename_completer/cmd"
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
	fold "github.com/rsteube/carapace-completers/completers/fold_completer/cmd"
	git "github.com/rsteube/carapace-completers/completers/git_completer/cmd"
	gradle "github.com/rsteube/carapace-completers/completers/gradle_completer/cmd"
	head "github.com/rsteube/carapace-completers/completers/head_completer/cmd"
	install "github.com/rsteube/carapace-completers/completers/install_completer/cmd"
	join "github.com/rsteube/carapace-completers/completers/join_completer/cmd"
	ln "github.com/rsteube/carapace-completers/completers/ln_completer/cmd"
	ls "github.com/rsteube/carapace-completers/completers/ls_completer/cmd"
	md5sum "github.com/rsteube/carapace-completers/completers/md5sum_completer/cmd"
	mkdir "github.com/rsteube/carapace-completers/completers/mkdir_completer/cmd"
	mkfifo "github.com/rsteube/carapace-completers/completers/mkfifo_completer/cmd"
	mknod "github.com/rsteube/carapace-completers/completers/mknod_completer/cmd"
	mv "github.com/rsteube/carapace-completers/completers/mv_completer/cmd"
	nl "github.com/rsteube/carapace-completers/completers/nl_completer/cmd"
	od "github.com/rsteube/carapace-completers/completers/od_completer/cmd"
	paste "github.com/rsteube/carapace-completers/completers/paste_completer/cmd"
	pkill "github.com/rsteube/carapace-completers/completers/pkill_completer/cmd"
	pr "github.com/rsteube/carapace-completers/completers/pr_completer/cmd"
	ptx "github.com/rsteube/carapace-completers/completers/ptx_completer/cmd"
	rm "github.com/rsteube/carapace-completers/completers/rm_completer/cmd"
	rmdir "github.com/rsteube/carapace-completers/completers/rmdir_completer/cmd"
	sha1sum "github.com/rsteube/carapace-completers/completers/sha1sum_completer/cmd"
	shred "github.com/rsteube/carapace-completers/completers/shred_completer/cmd"
	sort "github.com/rsteube/carapace-completers/completers/sort_completer/cmd"
	split "github.com/rsteube/carapace-completers/completers/split_completer/cmd"
	sum "github.com/rsteube/carapace-completers/completers/sum_completer/cmd"
	sync "github.com/rsteube/carapace-completers/completers/sync_completer/cmd"
	tac "github.com/rsteube/carapace-completers/completers/tac_completer/cmd"
	tail "github.com/rsteube/carapace-completers/completers/tail_completer/cmd"
	tee "github.com/rsteube/carapace-completers/completers/tee_completer/cmd"
	touch "github.com/rsteube/carapace-completers/completers/touch_completer/cmd"
	tr "github.com/rsteube/carapace-completers/completers/tr_completer/cmd"
	tsort "github.com/rsteube/carapace-completers/completers/tsort_completer/cmd"
	unexpand "github.com/rsteube/carapace-completers/completers/unexpand_completer/cmd"
	uniq "github.com/rsteube/carapace-completers/completers/uniq_completer/cmd"
	vdir "github.com/rsteube/carapace-completers/completers/vdir_completer/cmd"
	wc "github.com/rsteube/carapace-completers/completers/wc_completer/cmd"
)

var completers = []string{
	"basename",
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
	"fold",
	"git",
	"gradle",
	"head",
	"install",
	"join",
	"ln",
	"ls",
	"md5sum",
	"mkdir",
	"mkfifo",
	"mknod",
	"mv",
	"nl",
	"od",
	"paste",
	"pkill",
	"pr",
	"ptx",
	"rm",
	"rmdir",
	"sha1sum",
	"shred",
	"sort",
	"split",
	"sum",
	"sync",
	"tac",
	"tail",
	"tee",
	"touch",
	"tr",
	"tsort",
	"unexpand",
	"uniq",
	"vdir",
	"wc",
}

func executeCompleter(completer string) {
	switch completer {
	case "basename":
		basename.Execute()
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
	case "fold":
		fold.Execute()
	case "git":
		git.Execute()
	case "gradle":
		gradle.Execute()
	case "head":
		head.Execute()
	case "install":
		install.Execute()
	case "join":
		join.Execute()
	case "ln":
		ln.Execute()
	case "ls":
		ls.Execute()
	case "md5sum":
		md5sum.Execute()
	case "mkdir":
		mkdir.Execute()
	case "mkfifo":
		mkfifo.Execute()
	case "mknod":
		mknod.Execute()
	case "mv":
		mv.Execute()
	case "nl":
		nl.Execute()
	case "od":
		od.Execute()
	case "paste":
		paste.Execute()
	case "pkill":
		pkill.Execute()
	case "pr":
		pr.Execute()
	case "ptx":
		ptx.Execute()
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
	case "split":
		split.Execute()
	case "sum":
		sum.Execute()
	case "sync":
		sync.Execute()
	case "tac":
		tac.Execute()
	case "tail":
		tail.Execute()
	case "tee":
		tee.Execute()
	case "touch":
		touch.Execute()
	case "tr":
		tr.Execute()
	case "tsort":
		tsort.Execute()
	case "unexpand":
		unexpand.Execute()
	case "uniq":
		uniq.Execute()
	case "vdir":
		vdir.Execute()
	case "wc":
		wc.Execute()
	}
}

