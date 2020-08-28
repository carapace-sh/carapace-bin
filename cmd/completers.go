package cmd

import (
	bat "github.com/rsteube/carapace-completers/completers/bat_completer/cmd"
	chgrp "github.com/rsteube/carapace-completers/completers/chgrp_completer/cmd"
	chmod "github.com/rsteube/carapace-completers/completers/chmod_completer/cmd"
	chown "github.com/rsteube/carapace-completers/completers/chown_completer/cmd"
	cksum "github.com/rsteube/carapace-completers/completers/cksum_completer/cmd"
	comm "github.com/rsteube/carapace-completers/completers/comm_completer/cmd"
	cp "github.com/rsteube/carapace-completers/completers/cp_completer/cmd"
	curl "github.com/rsteube/carapace-completers/completers/curl_completer/cmd"
	df "github.com/rsteube/carapace-completers/completers/df_completer/cmd"
	exa "github.com/rsteube/carapace-completers/completers/exa_completer/cmd"
	git "github.com/rsteube/carapace-completers/completers/git_completer/cmd"
	gradle "github.com/rsteube/carapace-completers/completers/gradle_completer/cmd"
	ln "github.com/rsteube/carapace-completers/completers/ln_completer/cmd"
	mkdir "github.com/rsteube/carapace-completers/completers/mkdir_completer/cmd"
	mv "github.com/rsteube/carapace-completers/completers/mv_completer/cmd"
	pkill "github.com/rsteube/carapace-completers/completers/pkill_completer/cmd"
	rm "github.com/rsteube/carapace-completers/completers/rm_completer/cmd"
	rmdir "github.com/rsteube/carapace-completers/completers/rmdir_completer/cmd"
	sha1sum "github.com/rsteube/carapace-completers/completers/sha1sum_completer/cmd"
	sort "github.com/rsteube/carapace-completers/completers/sort_completer/cmd"
	tee "github.com/rsteube/carapace-completers/completers/tee_completer/cmd"
	touch "github.com/rsteube/carapace-completers/completers/touch_completer/cmd"
)

var completers = []string{
	"bat",
	"chgrp",
	"chmod",
	"chown",
	"cksum",
	"comm",
	"cp",
	"curl",
	"df",
	"exa",
	"git",
	"gradle",
	"ln",
	"mkdir",
	"mv",
	"pkill",
	"rm",
	"rmdir",
	"sha1sum",
	"sort",
	"tee",
	"touch",
}

func executeCompleter(completer string) {
	switch completer {
	case "bat":
		bat.Execute()
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
	case "curl":
		curl.Execute()
	case "df":
		df.Execute()
	case "exa":
		exa.Execute()
	case "git":
		git.Execute()
	case "gradle":
		gradle.Execute()
	case "ln":
		ln.Execute()
	case "mkdir":
		mkdir.Execute()
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
	case "sort":
		sort.Execute()
	case "tee":
		tee.Execute()
	case "touch":
		touch.Execute()
	}
}

