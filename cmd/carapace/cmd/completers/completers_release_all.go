//go:build release && force_all

package completers

import (
	common__acpi "github.com/carapace-sh/carapace-bin/completers_release/common/acpi_completer/cmd"
	common__acpid "github.com/carapace-sh/carapace-bin/completers_release/common/acpid_completer/cmd"
	common__adb "github.com/carapace-sh/carapace-bin/completers_release/common/adb_completer/cmd"
	common__age "github.com/carapace-sh/carapace-bin/completers_release/common/age_completer/cmd"
	common__agg "github.com/carapace-sh/carapace-bin/completers_release/common/agg_completer/cmd"
	common__alsamixer "github.com/carapace-sh/carapace-bin/completers_release/common/alsamixer_completer/cmd"
	common__ant "github.com/carapace-sh/carapace-bin/completers_release/common/ant_completer/cmd"
	common__apk "github.com/carapace-sh/carapace-bin/completers_release/common/apk_completer/cmd"
	common__aplay "github.com/carapace-sh/carapace-bin/completers_release/common/aplay_completer/cmd"
	common__apropos "github.com/carapace-sh/carapace-bin/completers_release/common/apropos_completer/cmd"
	common__apt_cache "github.com/carapace-sh/carapace-bin/completers_release/common/apt-cache_completer/cmd"
	common__apt_get "github.com/carapace-sh/carapace-bin/completers_release/common/apt-get_completer/cmd"
	common__apt "github.com/carapace-sh/carapace-bin/completers_release/common/apt_completer/cmd"
	common__ar "github.com/carapace-sh/carapace-bin/completers_release/common/ar_completer/cmd"
	common__arecord "github.com/carapace-sh/carapace-bin/completers_release/common/arecord_completer/cmd"
	common__aria2c "github.com/carapace-sh/carapace-bin/completers_release/common/aria2c_completer/cmd"
	common__artisan "github.com/carapace-sh/carapace-bin/completers_release/common/artisan_completer/cmd"
	common__asciinema "github.com/carapace-sh/carapace-bin/completers_release/common/asciinema_completer/cmd"
	common__autoconf "github.com/carapace-sh/carapace-bin/completers_release/common/autoconf_completer/cmd"
	common__avdmanager "github.com/carapace-sh/carapace-bin/completers_release/common/avdmanager_completer/cmd"
	common__awk "github.com/carapace-sh/carapace-bin/completers_release/common/awk_completer/cmd"
	common__aws "github.com/carapace-sh/carapace-bin/completers_release/common/aws_completer/cmd"
	common__az "github.com/carapace-sh/carapace-bin/completers_release/common/az_completer/cmd"
	common__baobab "github.com/carapace-sh/carapace-bin/completers_release/common/baobab_completer/cmd"
	common__base32 "github.com/carapace-sh/carapace-bin/completers_release/common/base32_completer/cmd"
	common__base64 "github.com/carapace-sh/carapace-bin/completers_release/common/base64_completer/cmd"
	common__basename "github.com/carapace-sh/carapace-bin/completers_release/common/basename_completer/cmd"
	common__bash_language_server "github.com/carapace-sh/carapace-bin/completers_release/common/bash-language-server_completer/cmd"
	common__bash "github.com/carapace-sh/carapace-bin/completers_release/common/bash_completer/cmd"
	common__bat "github.com/carapace-sh/carapace-bin/completers_release/common/bat_completer/cmd"
	common__batdiff "github.com/carapace-sh/carapace-bin/completers_release/common/batdiff_completer/cmd"
	common__batgrep "github.com/carapace-sh/carapace-bin/completers_release/common/batgrep_completer/cmd"
	common__batman "github.com/carapace-sh/carapace-bin/completers_release/common/batman_completer/cmd"
	common__bats "github.com/carapace-sh/carapace-bin/completers_release/common/bats_completer/cmd"
	common__bc "github.com/carapace-sh/carapace-bin/completers_release/common/bc_completer/cmd"
	common__benthos "github.com/carapace-sh/carapace-bin/completers_release/common/benthos_completer/cmd"
	common__black "github.com/carapace-sh/carapace-bin/completers_release/common/black_completer/cmd"
	common__bloop "github.com/carapace-sh/carapace-bin/completers_release/common/bloop_completer/cmd"
	common__bluetoothctl "github.com/carapace-sh/carapace-bin/completers_release/common/bluetoothctl_completer/cmd"
	common__boundary "github.com/carapace-sh/carapace-bin/completers_release/common/boundary_completer/cmd"
	common__brew "github.com/carapace-sh/carapace-bin/completers_release/common/brew_completer/cmd"
	common__brotli "github.com/carapace-sh/carapace-bin/completers_release/common/brotli_completer/cmd"
	common__bru "github.com/carapace-sh/carapace-bin/completers_release/common/bru_completer/cmd"
	common__btop "github.com/carapace-sh/carapace-bin/completers_release/common/btop_completer/cmd"
	common__bun "github.com/carapace-sh/carapace-bin/completers_release/common/bun_completer/cmd"
	common__bunx "github.com/carapace-sh/carapace-bin/completers_release/common/bunx_completer/cmd"
	common__but "github.com/carapace-sh/carapace-bin/completers_release/common/but_completer/cmd"
	common__cal "github.com/carapace-sh/carapace-bin/completers_release/common/cal_completer/cmd"
	common__calibre "github.com/carapace-sh/carapace-bin/completers_release/common/calibre_completer/cmd"
	common__capslock "github.com/carapace-sh/carapace-bin/completers_release/common/capslock_completer/cmd"
	common__cargo_clippy "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-clippy_completer/cmd"
	common__cargo_fmt "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-fmt_completer/cmd"
	common__cargo_metadata "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-metadata_completer/cmd"
	common__cargo_rm "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-rm_completer/cmd"
	common__cargo_set_version "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-set-version_completer/cmd"
	common__cargo_upgrade "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-upgrade_completer/cmd"
	common__cargo_watch "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-watch_completer/cmd"
	common__cargo "github.com/carapace-sh/carapace-bin/completers_release/common/cargo_completer/cmd"
	common__cat "github.com/carapace-sh/carapace-bin/completers_release/common/cat_completer/cmd"
	common__cfdisk "github.com/carapace-sh/carapace-bin/completers_release/common/cfdisk_completer/cmd"
	common__charm "github.com/carapace-sh/carapace-bin/completers_release/common/charm_completer/cmd"
	common__chcpu "github.com/carapace-sh/carapace-bin/completers_release/common/chcpu_completer/cmd"
	common__chdman "github.com/carapace-sh/carapace-bin/completers_release/common/chdman_completer/cmd"
	common__cheese "github.com/carapace-sh/carapace-bin/completers_release/common/cheese_completer/cmd"
	common__chgrp "github.com/carapace-sh/carapace-bin/completers_release/common/chgrp_completer/cmd"
	common__chmod "github.com/carapace-sh/carapace-bin/completers_release/common/chmod_completer/cmd"
	common__chown "github.com/carapace-sh/carapace-bin/completers_release/common/chown_completer/cmd"
	common__chpasswd "github.com/carapace-sh/carapace-bin/completers_release/common/chpasswd_completer/cmd"
	common__chroma "github.com/carapace-sh/carapace-bin/completers_release/common/chroma_completer/cmd"
	common__chromium "github.com/carapace-sh/carapace-bin/completers_release/common/chromium_completer/cmd"
	common__chroot "github.com/carapace-sh/carapace-bin/completers_release/common/chroot_completer/cmd"
	common__chsh "github.com/carapace-sh/carapace-bin/completers_release/common/chsh_completer/cmd"
	common__circleci "github.com/carapace-sh/carapace-bin/completers_release/common/circleci_completer/cmd"
	common__cksum "github.com/carapace-sh/carapace-bin/completers_release/common/cksum_completer/cmd"
	common__clamav_config "github.com/carapace-sh/carapace-bin/completers_release/common/clamav-config_completer/cmd"
	common__clamav_milter "github.com/carapace-sh/carapace-bin/completers_release/common/clamav-milter_completer/cmd"
	common__clambc "github.com/carapace-sh/carapace-bin/completers_release/common/clambc_completer/cmd"
	common__clamconf "github.com/carapace-sh/carapace-bin/completers_release/common/clamconf_completer/cmd"
	common__clamd "github.com/carapace-sh/carapace-bin/completers_release/common/clamd_completer/cmd"
	common__clamdscan "github.com/carapace-sh/carapace-bin/completers_release/common/clamdscan_completer/cmd"
	common__clamdtop "github.com/carapace-sh/carapace-bin/completers_release/common/clamdtop_completer/cmd"
	common__clamonacc "github.com/carapace-sh/carapace-bin/completers_release/common/clamonacc_completer/cmd"
	common__clamscan "github.com/carapace-sh/carapace-bin/completers_release/common/clamscan_completer/cmd"
	common__clamsubmit "github.com/carapace-sh/carapace-bin/completers_release/common/clamsubmit_completer/cmd"
	common__cmus "github.com/carapace-sh/carapace-bin/completers_release/common/cmus_completer/cmd"
	common__code_insiders "github.com/carapace-sh/carapace-bin/completers_release/common/code-insiders_completer/cmd"
	common__code "github.com/carapace-sh/carapace-bin/completers_release/common/code_completer/cmd"
	common__codecov "github.com/carapace-sh/carapace-bin/completers_release/common/codecov_completer/cmd"
	common__comm "github.com/carapace-sh/carapace-bin/completers_release/common/comm_completer/cmd"
	common__conda_content_trust "github.com/carapace-sh/carapace-bin/completers_release/common/conda-content-trust_completer/cmd"
	common__conda_env "github.com/carapace-sh/carapace-bin/completers_release/common/conda-env_completer/cmd"
	common__conda "github.com/carapace-sh/carapace-bin/completers_release/common/conda_completer/cmd"
	common__conky "github.com/carapace-sh/carapace-bin/completers_release/common/conky_completer/cmd"
	common__consul "github.com/carapace-sh/carapace-bin/completers_release/common/consul_completer/cmd"
	common__coredumpctl "github.com/carapace-sh/carapace-bin/completers_release/common/coredumpctl_completer/cmd"
	common__cp "github.com/carapace-sh/carapace-bin/completers_release/common/cp_completer/cmd"
	common__csplit "github.com/carapace-sh/carapace-bin/completers_release/common/csplit_completer/cmd"
	common__csview "github.com/carapace-sh/carapace-bin/completers_release/common/csview_completer/cmd"
	common__cura "github.com/carapace-sh/carapace-bin/completers_release/common/cura_completer/cmd"
	common__curl "github.com/carapace-sh/carapace-bin/completers_release/common/curl_completer/cmd"
	common__cut "github.com/carapace-sh/carapace-bin/completers_release/common/cut_completer/cmd"
	common__d2 "github.com/carapace-sh/carapace-bin/completers_release/common/d2_completer/cmd"
	common__dagger "github.com/carapace-sh/carapace-bin/completers_release/common/dagger_completer/cmd"
	common__darktable_cli "github.com/carapace-sh/carapace-bin/completers_release/common/darktable-cli_completer/cmd"
	common__darktable "github.com/carapace-sh/carapace-bin/completers_release/common/darktable_completer/cmd"
	common__dart "github.com/carapace-sh/carapace-bin/completers_release/common/dart_completer/cmd"
	common__date "github.com/carapace-sh/carapace-bin/completers_release/common/date_completer/cmd"
	common__dbt "github.com/carapace-sh/carapace-bin/completers_release/common/dbt_completer/cmd"
	common__dc "github.com/carapace-sh/carapace-bin/completers_release/common/dc_completer/cmd"
	common__dd "github.com/carapace-sh/carapace-bin/completers_release/common/dd_completer/cmd"
	common__deadcode "github.com/carapace-sh/carapace-bin/completers_release/common/deadcode_completer/cmd"
	common__delta "github.com/carapace-sh/carapace-bin/completers_release/common/delta_completer/cmd"
	common__deno "github.com/carapace-sh/carapace-bin/completers_release/common/deno_completer/cmd"
	common__devbox "github.com/carapace-sh/carapace-bin/completers_release/common/devbox_completer/cmd"
	common__df "github.com/carapace-sh/carapace-bin/completers_release/common/df_completer/cmd"
	common__dfc "github.com/carapace-sh/carapace-bin/completers_release/common/dfc_completer/cmd"
	common__dict "github.com/carapace-sh/carapace-bin/completers_release/common/dict_completer/cmd"
	common__diff3 "github.com/carapace-sh/carapace-bin/completers_release/common/diff3_completer/cmd"
	common__diff "github.com/carapace-sh/carapace-bin/completers_release/common/diff_completer/cmd"
	common__dig "github.com/carapace-sh/carapace-bin/completers_release/common/dig_completer/cmd"
	common__dir "github.com/carapace-sh/carapace-bin/completers_release/common/dir_completer/cmd"
	common__dircolors "github.com/carapace-sh/carapace-bin/completers_release/common/dircolors_completer/cmd"
	common__direnv "github.com/carapace-sh/carapace-bin/completers_release/common/direnv_completer/cmd"
	common__dirname "github.com/carapace-sh/carapace-bin/completers_release/common/dirname_completer/cmd"
	common__dive "github.com/carapace-sh/carapace-bin/completers_release/common/dive_completer/cmd"
	common__dlv "github.com/carapace-sh/carapace-bin/completers_release/common/dlv_completer/cmd"
	common__dmenu "github.com/carapace-sh/carapace-bin/completers_release/common/dmenu_completer/cmd"
	common__dmesg "github.com/carapace-sh/carapace-bin/completers_release/common/dmesg_completer/cmd"
	common__dms "github.com/carapace-sh/carapace-bin/completers_release/common/dms_completer/cmd"
	common__dngconverter "github.com/carapace-sh/carapace-bin/completers_release/common/dngconverter_completer/cmd"
	common__dnsmasq "github.com/carapace-sh/carapace-bin/completers_release/common/dnsmasq_completer/cmd"
	common__doas "github.com/carapace-sh/carapace-bin/completers_release/common/doas_completer/cmd"
	common__docker_buildx "github.com/carapace-sh/carapace-bin/completers_release/common/docker-buildx_completer/cmd"
	common__docker_compose "github.com/carapace-sh/carapace-bin/completers_release/common/docker-compose_completer/cmd"
	common__docker_scan "github.com/carapace-sh/carapace-bin/completers_release/common/docker-scan_completer/cmd"
	common__docker "github.com/carapace-sh/carapace-bin/completers_release/common/docker_completer/cmd"
	common__dockerd "github.com/carapace-sh/carapace-bin/completers_release/common/dockerd_completer/cmd"
	common__doctl "github.com/carapace-sh/carapace-bin/completers_release/common/doctl_completer/cmd"
	common__doing "github.com/carapace-sh/carapace-bin/completers_release/common/doing_completer/cmd"
	common__dos2unix "github.com/carapace-sh/carapace-bin/completers_release/common/dos2unix_completer/cmd"
	common__downgrade "github.com/carapace-sh/carapace-bin/completers_release/common/downgrade_completer/cmd"
	common__dpkg "github.com/carapace-sh/carapace-bin/completers_release/common/dpkg_completer/cmd"
	common__du "github.com/carapace-sh/carapace-bin/completers_release/common/du_completer/cmd"
	common__ebook_convert "github.com/carapace-sh/carapace-bin/completers_release/common/ebook-convert_completer/cmd"
	common__egrep "github.com/carapace-sh/carapace-bin/completers_release/common/egrep_completer/cmd"
	common__electron "github.com/carapace-sh/carapace-bin/completers_release/common/electron_completer/cmd"
	common__elvish "github.com/carapace-sh/carapace-bin/completers_release/common/elvish_completer/cmd"
	common__env "github.com/carapace-sh/carapace-bin/completers_release/common/env_completer/cmd"
	common__envsubst "github.com/carapace-sh/carapace-bin/completers_release/common/envsubst_completer/cmd"
	common__exa "github.com/carapace-sh/carapace-bin/completers_release/common/exa_completer/cmd"
	common__expand "github.com/carapace-sh/carapace-bin/completers_release/common/expand_completer/cmd"
	common__expr "github.com/carapace-sh/carapace-bin/completers_release/common/expr_completer/cmd"
	common__eza "github.com/carapace-sh/carapace-bin/completers_release/common/eza_completer/cmd"
	common__faas_cli "github.com/carapace-sh/carapace-bin/completers_release/common/faas-cli_completer/cmd"
	common__factor "github.com/carapace-sh/carapace-bin/completers_release/common/factor_completer/cmd"
	common__fakechroot "github.com/carapace-sh/carapace-bin/completers_release/common/fakechroot_completer/cmd"
	common__fakeroot "github.com/carapace-sh/carapace-bin/completers_release/common/fakeroot_completer/cmd"
	common__fastfetch "github.com/carapace-sh/carapace-bin/completers_release/common/fastfetch_completer/cmd"
	common__fc_cache "github.com/carapace-sh/carapace-bin/completers_release/common/fc-cache_completer/cmd"
	common__fc_cat "github.com/carapace-sh/carapace-bin/completers_release/common/fc-cat_completer/cmd"
	common__fc_conflist "github.com/carapace-sh/carapace-bin/completers_release/common/fc-conflist_completer/cmd"
	common__fc_list "github.com/carapace-sh/carapace-bin/completers_release/common/fc-list_completer/cmd"
	common__fd "github.com/carapace-sh/carapace-bin/completers_release/common/fd_completer/cmd"
	common__fdisk "github.com/carapace-sh/carapace-bin/completers_release/common/fdisk_completer/cmd"
	common__ffmpeg "github.com/carapace-sh/carapace-bin/completers_release/common/ffmpeg_completer/cmd"
	common__fgrep "github.com/carapace-sh/carapace-bin/completers_release/common/fgrep_completer/cmd"
	common__file "github.com/carapace-sh/carapace-bin/completers_release/common/file_completer/cmd"
	common__find "github.com/carapace-sh/carapace-bin/completers_release/common/find_completer/cmd"
	common__firefox "github.com/carapace-sh/carapace-bin/completers_release/common/firefox_completer/cmd"
	common__fish "github.com/carapace-sh/carapace-bin/completers_release/common/fish_completer/cmd"
	common__flatpak "github.com/carapace-sh/carapace-bin/completers_release/common/flatpak_completer/cmd"
	common__flutter "github.com/carapace-sh/carapace-bin/completers_release/common/flutter_completer/cmd"
	common__fmt "github.com/carapace-sh/carapace-bin/completers_release/common/fmt_completer/cmd"
	common__fnm "github.com/carapace-sh/carapace-bin/completers_release/common/fnm_completer/cmd"
	common__fold "github.com/carapace-sh/carapace-bin/completers_release/common/fold_completer/cmd"
	common__foot "github.com/carapace-sh/carapace-bin/completers_release/common/foot_completer/cmd"
	common__free "github.com/carapace-sh/carapace-bin/completers_release/common/free_completer/cmd"
	common__freeze "github.com/carapace-sh/carapace-bin/completers_release/common/freeze_completer/cmd"
	common__ftp "github.com/carapace-sh/carapace-bin/completers_release/common/ftp_completer/cmd"
	common__ftpd "github.com/carapace-sh/carapace-bin/completers_release/common/ftpd_completer/cmd"
	common__fury "github.com/carapace-sh/carapace-bin/completers_release/common/fury_completer/cmd"
	common__fzf "github.com/carapace-sh/carapace-bin/completers_release/common/fzf_completer/cmd"
	common__gatsby "github.com/carapace-sh/carapace-bin/completers_release/common/gatsby_completer/cmd"
	common__gcloud "github.com/carapace-sh/carapace-bin/completers_release/common/gcloud_completer/cmd"
	common__gdb "github.com/carapace-sh/carapace-bin/completers_release/common/gdb_completer/cmd"
	common__gdown "github.com/carapace-sh/carapace-bin/completers_release/common/gdown_completer/cmd"
	common__gdu "github.com/carapace-sh/carapace-bin/completers_release/common/gdu_completer/cmd"
	common__get_env "github.com/carapace-sh/carapace-bin/completers_release/common/get-env_completer/cmd"
	common__gftp "github.com/carapace-sh/carapace-bin/completers_release/common/gftp_completer/cmd"
	common__gh_copilot "github.com/carapace-sh/carapace-bin/completers_release/common/gh-copilot_completer/cmd"
	common__gh_dash "github.com/carapace-sh/carapace-bin/completers_release/common/gh-dash_completer/cmd"
	common__gh "github.com/carapace-sh/carapace-bin/completers_release/common/gh_completer/cmd"
	common__ghostty "github.com/carapace-sh/carapace-bin/completers_release/common/ghostty_completer/cmd"
	common__gimp "github.com/carapace-sh/carapace-bin/completers_release/common/gimp_completer/cmd"
	common__git_abort "github.com/carapace-sh/carapace-bin/completers_release/common/git-abort_completer/cmd"
	common__git_alias "github.com/carapace-sh/carapace-bin/completers_release/common/git-alias_completer/cmd"
	common__git_archive_file "github.com/carapace-sh/carapace-bin/completers_release/common/git-archive-file_completer/cmd"
	common__git_authors "github.com/carapace-sh/carapace-bin/completers_release/common/git-authors_completer/cmd"
	common__git_browse_ci "github.com/carapace-sh/carapace-bin/completers_release/common/git-browse-ci_completer/cmd"
	common__git_browse "github.com/carapace-sh/carapace-bin/completers_release/common/git-browse_completer/cmd"
	common__git_clang_format "github.com/carapace-sh/carapace-bin/completers_release/common/git-clang-format_completer/cmd"
	common__git_clear_soft "github.com/carapace-sh/carapace-bin/completers_release/common/git-clear-soft_completer/cmd"
	common__git_clear "github.com/carapace-sh/carapace-bin/completers_release/common/git-clear_completer/cmd"
	common__git_coauthor "github.com/carapace-sh/carapace-bin/completers_release/common/git-coauthor_completer/cmd"
	common__git_extras "github.com/carapace-sh/carapace-bin/completers_release/common/git-extras_completer/cmd"
	common__git_info "github.com/carapace-sh/carapace-bin/completers_release/common/git-info_completer/cmd"
	common__git_standup "github.com/carapace-sh/carapace-bin/completers_release/common/git-standup_completer/cmd"
	common__git_unlock "github.com/carapace-sh/carapace-bin/completers_release/common/git-unlock_completer/cmd"
	common__git_utimes "github.com/carapace-sh/carapace-bin/completers_release/common/git-utimes_completer/cmd"
	common__git "github.com/carapace-sh/carapace-bin/completers_release/common/git_completer/cmd"
	common__gitk "github.com/carapace-sh/carapace-bin/completers_release/common/gitk_completer/cmd"
	common__gitui "github.com/carapace-sh/carapace-bin/completers_release/common/gitui_completer/cmd"
	common__glab "github.com/carapace-sh/carapace-bin/completers_release/common/glab_completer/cmd"
	common__glow "github.com/carapace-sh/carapace-bin/completers_release/common/glow_completer/cmd"
	common__gm "github.com/carapace-sh/carapace-bin/completers_release/common/gm_completer/cmd"
	common__gnome_keyring_daemon "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-keyring-daemon_completer/cmd"
	common__gnome_keyring "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-keyring_completer/cmd"
	common__gnome_maps "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-maps_completer/cmd"
	common__gnome_terminal "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-terminal_completer/cmd"
	common__go_carpet "github.com/carapace-sh/carapace-bin/completers_release/common/go-carpet_completer/cmd"
	common__go_tool_asm "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-asm_completer/cmd"
	common__go_tool_buildid "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-buildid_completer/cmd"
	common__go_tool_cgo "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-cgo_completer/cmd"
	common__go_tool_compile "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-compile_completer/cmd"
	common__go_tool_covdata "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-covdata_completer/cmd"
	common__go_tool_cover "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-cover_completer/cmd"
	common__go_tool_dist "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-dist_completer/cmd"
	common__go_tool_doc "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-doc_completer/cmd"
	common__go_tool_fix "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-fix_completer/cmd"
	common__go_tool_link "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-link_completer/cmd"
	common__go_tool_mockgen "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-mockgen_completer/cmd"
	common__go_tool_nm "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-nm_completer/cmd"
	common__go_tool_objdump "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-objdump_completer/cmd"
	common__go_tool_pack "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-pack_completer/cmd"
	common__go "github.com/carapace-sh/carapace-bin/completers_release/common/go_completer/cmd"
	common__gocyclo "github.com/carapace-sh/carapace-bin/completers_release/common/gocyclo_completer/cmd"
	common__gofmt "github.com/carapace-sh/carapace-bin/completers_release/common/gofmt_completer/cmd"
	common__goimports "github.com/carapace-sh/carapace-bin/completers_release/common/goimports_completer/cmd"
	common__golangci_lint "github.com/carapace-sh/carapace-bin/completers_release/common/golangci-lint_completer/cmd"
	common__gonew "github.com/carapace-sh/carapace-bin/completers_release/common/gonew_completer/cmd"
	common__google_chrome "github.com/carapace-sh/carapace-bin/completers_release/common/google-chrome_completer/cmd"
	common__gopls "github.com/carapace-sh/carapace-bin/completers_release/common/gopls_completer/cmd"
	common__goreleaser "github.com/carapace-sh/carapace-bin/completers_release/common/goreleaser_completer/cmd"
	common__goweight "github.com/carapace-sh/carapace-bin/completers_release/common/goweight_completer/cmd"
	common__gparted "github.com/carapace-sh/carapace-bin/completers_release/common/gparted_completer/cmd"
	common__gpasswd "github.com/carapace-sh/carapace-bin/completers_release/common/gpasswd_completer/cmd"
	common__gpg_agent "github.com/carapace-sh/carapace-bin/completers_release/common/gpg-agent_completer/cmd"
	common__gpg "github.com/carapace-sh/carapace-bin/completers_release/common/gpg_completer/cmd"
	common__gradle "github.com/carapace-sh/carapace-bin/completers_release/common/gradle_completer/cmd"
	common__grep "github.com/carapace-sh/carapace-bin/completers_release/common/grep_completer/cmd"
	common__groupadd "github.com/carapace-sh/carapace-bin/completers_release/common/groupadd_completer/cmd"
	common__groupdel "github.com/carapace-sh/carapace-bin/completers_release/common/groupdel_completer/cmd"
	common__groupmems "github.com/carapace-sh/carapace-bin/completers_release/common/groupmems_completer/cmd"
	common__groupmod "github.com/carapace-sh/carapace-bin/completers_release/common/groupmod_completer/cmd"
	common__groups "github.com/carapace-sh/carapace-bin/completers_release/common/groups_completer/cmd"
	common__grype "github.com/carapace-sh/carapace-bin/completers_release/common/grype_completer/cmd"
	common__gsa "github.com/carapace-sh/carapace-bin/completers_release/common/gsa_completer/cmd"
	common__gulp "github.com/carapace-sh/carapace-bin/completers_release/common/gulp_completer/cmd"
	common__gum "github.com/carapace-sh/carapace-bin/completers_release/common/gum_completer/cmd"
	common__gunzip "github.com/carapace-sh/carapace-bin/completers_release/common/gunzip_completer/cmd"
	common__gzip "github.com/carapace-sh/carapace-bin/completers_release/common/gzip_completer/cmd"
	common__halt "github.com/carapace-sh/carapace-bin/completers_release/common/halt_completer/cmd"
	common__head "github.com/carapace-sh/carapace-bin/completers_release/common/head_completer/cmd"
	common__helix "github.com/carapace-sh/carapace-bin/completers_release/common/helix_completer/cmd"
	common__helm "github.com/carapace-sh/carapace-bin/completers_release/common/helm_completer/cmd"
	common__helmsman "github.com/carapace-sh/carapace-bin/completers_release/common/helmsman_completer/cmd"
	common__hexchat "github.com/carapace-sh/carapace-bin/completers_release/common/hexchat_completer/cmd"
	common__hexdump "github.com/carapace-sh/carapace-bin/completers_release/common/hexdump_completer/cmd"
	common__hostid "github.com/carapace-sh/carapace-bin/completers_release/common/hostid_completer/cmd"
	common__hostname "github.com/carapace-sh/carapace-bin/completers_release/common/hostname_completer/cmd"
	common__htop "github.com/carapace-sh/carapace-bin/completers_release/common/htop_completer/cmd"
	common__http "github.com/carapace-sh/carapace-bin/completers_release/common/http_completer/cmd"
	common__https "github.com/carapace-sh/carapace-bin/completers_release/common/https_completer/cmd"
	common__hugetop "github.com/carapace-sh/carapace-bin/completers_release/common/hugetop_completer/cmd"
	common__hugo "github.com/carapace-sh/carapace-bin/completers_release/common/hugo_completer/cmd"
	common__hurl "github.com/carapace-sh/carapace-bin/completers_release/common/hurl_completer/cmd"
	common__hwinfo "github.com/carapace-sh/carapace-bin/completers_release/common/hwinfo_completer/cmd"
	common__hx "github.com/carapace-sh/carapace-bin/completers_release/common/hx_completer/cmd"
	common__i3_scrot "github.com/carapace-sh/carapace-bin/completers_release/common/i3-scrot_completer/cmd"
	common__i3 "github.com/carapace-sh/carapace-bin/completers_release/common/i3_completer/cmd"
	common__i3exit "github.com/carapace-sh/carapace-bin/completers_release/common/i3exit_completer/cmd"
	common__i3lock "github.com/carapace-sh/carapace-bin/completers_release/common/i3lock_completer/cmd"
	common__i3status_rs "github.com/carapace-sh/carapace-bin/completers_release/common/i3status-rs_completer/cmd"
	common__i3status "github.com/carapace-sh/carapace-bin/completers_release/common/i3status_completer/cmd"
	common__id "github.com/carapace-sh/carapace-bin/completers_release/common/id_completer/cmd"
	common__imv "github.com/carapace-sh/carapace-bin/completers_release/common/imv_completer/cmd"
	common__inkscape "github.com/carapace-sh/carapace-bin/completers_release/common/inkscape_completer/cmd"
	common__inshellisense "github.com/carapace-sh/carapace-bin/completers_release/common/inshellisense_completer/cmd"
	common__install "github.com/carapace-sh/carapace-bin/completers_release/common/install_completer/cmd"
	common__ion "github.com/carapace-sh/carapace-bin/completers_release/common/ion_completer/cmd"
	common__jar "github.com/carapace-sh/carapace-bin/completers_release/common/jar_completer/cmd"
	common__java "github.com/carapace-sh/carapace-bin/completers_release/common/java_completer/cmd"
	common__javac "github.com/carapace-sh/carapace-bin/completers_release/common/javac_completer/cmd"
	common__jj "github.com/carapace-sh/carapace-bin/completers_release/common/jj_completer/cmd"
	common__join "github.com/carapace-sh/carapace-bin/completers_release/common/join_completer/cmd"
	common__journalctl "github.com/carapace-sh/carapace-bin/completers_release/common/journalctl_completer/cmd"
	common__jq "github.com/carapace-sh/carapace-bin/completers_release/common/jq_completer/cmd"
	common__julia "github.com/carapace-sh/carapace-bin/completers_release/common/julia_completer/cmd"
	common__just "github.com/carapace-sh/carapace-bin/completers_release/common/just_completer/cmd"
	common__kak_lsp "github.com/carapace-sh/carapace-bin/completers_release/common/kak-lsp_completer/cmd"
	common__kak "github.com/carapace-sh/carapace-bin/completers_release/common/kak_completer/cmd"
	common__kill "github.com/carapace-sh/carapace-bin/completers_release/common/kill_completer/cmd"
	common__killall "github.com/carapace-sh/carapace-bin/completers_release/common/killall_completer/cmd"
	common__kitten "github.com/carapace-sh/carapace-bin/completers_release/common/kitten_completer/cmd"
	common__kitty "github.com/carapace-sh/carapace-bin/completers_release/common/kitty_completer/cmd"
	common__kmonad "github.com/carapace-sh/carapace-bin/completers_release/common/kmonad_completer/cmd"
	common__kompose "github.com/carapace-sh/carapace-bin/completers_release/common/kompose_completer/cmd"
	common__kotlin "github.com/carapace-sh/carapace-bin/completers_release/common/kotlin_completer/cmd"
	common__kotlinc "github.com/carapace-sh/carapace-bin/completers_release/common/kotlinc_completer/cmd"
	common__ktlint "github.com/carapace-sh/carapace-bin/completers_release/common/ktlint_completer/cmd"
	common__kubeadm "github.com/carapace-sh/carapace-bin/completers_release/common/kubeadm_completer/cmd"
	common__kubectl "github.com/carapace-sh/carapace-bin/completers_release/common/kubectl_completer/cmd"
	common__kubeseal "github.com/carapace-sh/carapace-bin/completers_release/common/kubeseal_completer/cmd"
	common__last "github.com/carapace-sh/carapace-bin/completers_release/common/last_completer/cmd"
	common__lastb "github.com/carapace-sh/carapace-bin/completers_release/common/lastb_completer/cmd"
	common__lastlog "github.com/carapace-sh/carapace-bin/completers_release/common/lastlog_completer/cmd"
	common__lazygit "github.com/carapace-sh/carapace-bin/completers_release/common/lazygit_completer/cmd"
	common__lf "github.com/carapace-sh/carapace-bin/completers_release/common/lf_completer/cmd"
	common__light "github.com/carapace-sh/carapace-bin/completers_release/common/light_completer/cmd"
	common__lightdm "github.com/carapace-sh/carapace-bin/completers_release/common/lightdm_completer/cmd"
	common__link "github.com/carapace-sh/carapace-bin/completers_release/common/link_completer/cmd"
	common__ln "github.com/carapace-sh/carapace-bin/completers_release/common/ln_completer/cmd"
	common__lnav "github.com/carapace-sh/carapace-bin/completers_release/common/lnav_completer/cmd"
	common__lncrawl "github.com/carapace-sh/carapace-bin/completers_release/common/lncrawl_completer/cmd"
	common__locale "github.com/carapace-sh/carapace-bin/completers_release/common/locale_completer/cmd"
	common__localectl "github.com/carapace-sh/carapace-bin/completers_release/common/localectl_completer/cmd"
	common__logname "github.com/carapace-sh/carapace-bin/completers_release/common/logname_completer/cmd"
	common__ls "github.com/carapace-sh/carapace-bin/completers_release/common/ls_completer/cmd"
	common__lsb_release "github.com/carapace-sh/carapace-bin/completers_release/common/lsb_release_completer/cmd"
	common__lsblk "github.com/carapace-sh/carapace-bin/completers_release/common/lsblk_completer/cmd"
	common__lsclocks "github.com/carapace-sh/carapace-bin/completers_release/common/lsclocks_completer/cmd"
	common__lscpu "github.com/carapace-sh/carapace-bin/completers_release/common/lscpu_completer/cmd"
	common__lsfd "github.com/carapace-sh/carapace-bin/completers_release/common/lsfd_completer/cmd"
	common__lsirq "github.com/carapace-sh/carapace-bin/completers_release/common/lsirq_completer/cmd"
	common__lslocks "github.com/carapace-sh/carapace-bin/completers_release/common/lslocks_completer/cmd"
	common__lslogins "github.com/carapace-sh/carapace-bin/completers_release/common/lslogins_completer/cmd"
	common__lsmem "github.com/carapace-sh/carapace-bin/completers_release/common/lsmem_completer/cmd"
	common__lsns "github.com/carapace-sh/carapace-bin/completers_release/common/lsns_completer/cmd"
	common__lsusb "github.com/carapace-sh/carapace-bin/completers_release/common/lsusb_completer/cmd"
	common__lua "github.com/carapace-sh/carapace-bin/completers_release/common/lua_completer/cmd"
	common__lzcat "github.com/carapace-sh/carapace-bin/completers_release/common/lzcat_completer/cmd"
	common__lzma "github.com/carapace-sh/carapace-bin/completers_release/common/lzma_completer/cmd"
	common__magick "github.com/carapace-sh/carapace-bin/completers_release/common/magick_completer/cmd"
	common__make "github.com/carapace-sh/carapace-bin/completers_release/common/make_completer/cmd"
	common__makepkg "github.com/carapace-sh/carapace-bin/completers_release/common/makepkg_completer/cmd"
	common__man "github.com/carapace-sh/carapace-bin/completers_release/common/man_completer/cmd"
	common__marp "github.com/carapace-sh/carapace-bin/completers_release/common/marp_completer/cmd"
	common__mcomix "github.com/carapace-sh/carapace-bin/completers_release/common/mcomix_completer/cmd"
	common__md5sum "github.com/carapace-sh/carapace-bin/completers_release/common/md5sum_completer/cmd"
	common__mdbook "github.com/carapace-sh/carapace-bin/completers_release/common/mdbook_completer/cmd"
	common__meld "github.com/carapace-sh/carapace-bin/completers_release/common/meld_completer/cmd"
	common__melt "github.com/carapace-sh/carapace-bin/completers_release/common/melt_completer/cmd"
	common__micro "github.com/carapace-sh/carapace-bin/completers_release/common/micro_completer/cmd"
	common__minikube "github.com/carapace-sh/carapace-bin/completers_release/common/minikube_completer/cmd"
	common__mitmproxy "github.com/carapace-sh/carapace-bin/completers_release/common/mitmproxy_completer/cmd"
	common__mix "github.com/carapace-sh/carapace-bin/completers_release/common/mix_completer/cmd"
	common__mkcert "github.com/carapace-sh/carapace-bin/completers_release/common/mkcert_completer/cmd"
	common__mkdir "github.com/carapace-sh/carapace-bin/completers_release/common/mkdir_completer/cmd"
	common__mkfifo "github.com/carapace-sh/carapace-bin/completers_release/common/mkfifo_completer/cmd"
	common__mkfs "github.com/carapace-sh/carapace-bin/completers_release/common/mkfs_completer/cmd"
	common__mknod "github.com/carapace-sh/carapace-bin/completers_release/common/mknod_completer/cmd"
	common__mkswap "github.com/carapace-sh/carapace-bin/completers_release/common/mkswap_completer/cmd"
	common__mktemp "github.com/carapace-sh/carapace-bin/completers_release/common/mktemp_completer/cmd"
	common__modinfo "github.com/carapace-sh/carapace-bin/completers_release/common/modinfo_completer/cmd"
	common__modprobe "github.com/carapace-sh/carapace-bin/completers_release/common/modprobe_completer/cmd"
	common__molecule "github.com/carapace-sh/carapace-bin/completers_release/common/molecule_completer/cmd"
	common__more "github.com/carapace-sh/carapace-bin/completers_release/common/more_completer/cmd"
	common__mosh "github.com/carapace-sh/carapace-bin/completers_release/common/mosh_completer/cmd"
	common__mount "github.com/carapace-sh/carapace-bin/completers_release/common/mount_completer/cmd"
	common__mousepad "github.com/carapace-sh/carapace-bin/completers_release/common/mousepad_completer/cmd"
	common__mpv "github.com/carapace-sh/carapace-bin/completers_release/common/mpv_completer/cmd"
	common__mv "github.com/carapace-sh/carapace-bin/completers_release/common/mv_completer/cmd"
	common__mvn "github.com/carapace-sh/carapace-bin/completers_release/common/mvn_completer/cmd"
	common__n_m3u8dl_re "github.com/carapace-sh/carapace-bin/completers_release/common/n-m3u8dl-re_completer/cmd"
	common__nano "github.com/carapace-sh/carapace-bin/completers_release/common/nano_completer/cmd"
	common__nc "github.com/carapace-sh/carapace-bin/completers_release/common/nc_completer/cmd"
	common__ncdu "github.com/carapace-sh/carapace-bin/completers_release/common/ncdu_completer/cmd"
	common__neomutt "github.com/carapace-sh/carapace-bin/completers_release/common/neomutt_completer/cmd"
	common__netcat "github.com/carapace-sh/carapace-bin/completers_release/common/netcat_completer/cmd"
	common__newman "github.com/carapace-sh/carapace-bin/completers_release/common/newman_completer/cmd"
	common__newrelic "github.com/carapace-sh/carapace-bin/completers_release/common/newrelic_completer/cmd"
	common__nfpm "github.com/carapace-sh/carapace-bin/completers_release/common/nfpm_completer/cmd"
	common__ng "github.com/carapace-sh/carapace-bin/completers_release/common/ng_completer/cmd"
	common__nice "github.com/carapace-sh/carapace-bin/completers_release/common/nice_completer/cmd"
	common__nilaway "github.com/carapace-sh/carapace-bin/completers_release/common/nilaway_completer/cmd"
	common__nix_build "github.com/carapace-sh/carapace-bin/completers_release/common/nix-build_completer/cmd"
	common__nix_channel "github.com/carapace-sh/carapace-bin/completers_release/common/nix-channel_completer/cmd"
	common__nix_instantiate "github.com/carapace-sh/carapace-bin/completers_release/common/nix-instantiate_completer/cmd"
	common__nix_shell "github.com/carapace-sh/carapace-bin/completers_release/common/nix-shell_completer/cmd"
	common__nix "github.com/carapace-sh/carapace-bin/completers_release/common/nix_completer/cmd"
	common__nixos_rebuild "github.com/carapace-sh/carapace-bin/completers_release/common/nixos-rebuild_completer/cmd"
	common__nl "github.com/carapace-sh/carapace-bin/completers_release/common/nl_completer/cmd"
	common__nmcli "github.com/carapace-sh/carapace-bin/completers_release/common/nmcli_completer/cmd"
	common__node "github.com/carapace-sh/carapace-bin/completers_release/common/node_completer/cmd"
	common__nohup "github.com/carapace-sh/carapace-bin/completers_release/common/nohup_completer/cmd"
	common__nomad "github.com/carapace-sh/carapace-bin/completers_release/common/nomad_completer/cmd"
	common__npm "github.com/carapace-sh/carapace-bin/completers_release/common/npm_completer/cmd"
	common__nproc "github.com/carapace-sh/carapace-bin/completers_release/common/nproc_completer/cmd"
	common__ntpd "github.com/carapace-sh/carapace-bin/completers_release/common/ntpd_completer/cmd"
	common__nu "github.com/carapace-sh/carapace-bin/completers_release/common/nu_completer/cmd"
	common__numfmt "github.com/carapace-sh/carapace-bin/completers_release/common/numfmt_completer/cmd"
	common__nvim "github.com/carapace-sh/carapace-bin/completers_release/common/nvim_completer/cmd"
	common__od "github.com/carapace-sh/carapace-bin/completers_release/common/od_completer/cmd"
	common__ollama "github.com/carapace-sh/carapace-bin/completers_release/common/ollama_completer/cmd"
	common__openscad "github.com/carapace-sh/carapace-bin/completers_release/common/openscad_completer/cmd"
	common__openssl "github.com/carapace-sh/carapace-bin/completers_release/common/openssl_completer/cmd"
	common__optipng "github.com/carapace-sh/carapace-bin/completers_release/common/optipng_completer/cmd"
	common__packer "github.com/carapace-sh/carapace-bin/completers_release/common/packer_completer/cmd"
	common__pacman_conf "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-conf_completer/cmd"
	common__pacman_db_upgrade "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-db-upgrade_completer/cmd"
	common__pacman_key "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-key_completer/cmd"
	common__pacman_mirrors "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-mirrors_completer/cmd"
	common__pacman "github.com/carapace-sh/carapace-bin/completers_release/common/pacman_completer/cmd"
	common__palemoon "github.com/carapace-sh/carapace-bin/completers_release/common/palemoon_completer/cmd"
	common__pamac "github.com/carapace-sh/carapace-bin/completers_release/common/pamac_completer/cmd"
	common__pandoc "github.com/carapace-sh/carapace-bin/completers_release/common/pandoc_completer/cmd"
	common__paru "github.com/carapace-sh/carapace-bin/completers_release/common/paru_completer/cmd"
	common__pass "github.com/carapace-sh/carapace-bin/completers_release/common/pass_completer/cmd"
	common__passwd "github.com/carapace-sh/carapace-bin/completers_release/common/passwd_completer/cmd"
	common__paste "github.com/carapace-sh/carapace-bin/completers_release/common/paste_completer/cmd"
	common__patch "github.com/carapace-sh/carapace-bin/completers_release/common/patch_completer/cmd"
	common__pathchk "github.com/carapace-sh/carapace-bin/completers_release/common/pathchk_completer/cmd"
	common__patool "github.com/carapace-sh/carapace-bin/completers_release/common/patool_completer/cmd"
	common__pcmanfm "github.com/carapace-sh/carapace-bin/completers_release/common/pcmanfm_completer/cmd"
	common__pgrep "github.com/carapace-sh/carapace-bin/completers_release/common/pgrep_completer/cmd"
	common__php "github.com/carapace-sh/carapace-bin/completers_release/common/php_completer/cmd"
	common__picard "github.com/carapace-sh/carapace-bin/completers_release/common/picard_completer/cmd"
	common__pidof "github.com/carapace-sh/carapace-bin/completers_release/common/pidof_completer/cmd"
	common__pidwait "github.com/carapace-sh/carapace-bin/completers_release/common/pidwait_completer/cmd"
	common__pigz "github.com/carapace-sh/carapace-bin/completers_release/common/pigz_completer/cmd"
	common__ping "github.com/carapace-sh/carapace-bin/completers_release/common/ping_completer/cmd"
	common__pinky "github.com/carapace-sh/carapace-bin/completers_release/common/pinky_completer/cmd"
	common__pip "github.com/carapace-sh/carapace-bin/completers_release/common/pip_completer/cmd"
	common__pkg "github.com/carapace-sh/carapace-bin/completers_release/common/pkg_completer/cmd"
	common__pkgsite "github.com/carapace-sh/carapace-bin/completers_release/common/pkgsite_completer/cmd"
	common__pkill "github.com/carapace-sh/carapace-bin/completers_release/common/pkill_completer/cmd"
	common__pmap "github.com/carapace-sh/carapace-bin/completers_release/common/pmap_completer/cmd"
	common__pngcheck "github.com/carapace-sh/carapace-bin/completers_release/common/pngcheck_completer/cmd"
	common__pnpm "github.com/carapace-sh/carapace-bin/completers_release/common/pnpm_completer/cmd"
	common__podman "github.com/carapace-sh/carapace-bin/completers_release/common/podman_completer/cmd"
	common__poweroff "github.com/carapace-sh/carapace-bin/completers_release/common/poweroff_completer/cmd"
	common__powertop "github.com/carapace-sh/carapace-bin/completers_release/common/powertop_completer/cmd"
	common__pprof "github.com/carapace-sh/carapace-bin/completers_release/common/pprof_completer/cmd"
	common__pr "github.com/carapace-sh/carapace-bin/completers_release/common/pr_completer/cmd"
	common__present "github.com/carapace-sh/carapace-bin/completers_release/common/present_completer/cmd"
	common__prettybat "github.com/carapace-sh/carapace-bin/completers_release/common/prettybat_completer/cmd"
	common__prettyping "github.com/carapace-sh/carapace-bin/completers_release/common/prettyping_completer/cmd"
	common__printenv "github.com/carapace-sh/carapace-bin/completers_release/common/printenv_completer/cmd"
	common__procs "github.com/carapace-sh/carapace-bin/completers_release/common/procs_completer/cmd"
	common__ps "github.com/carapace-sh/carapace-bin/completers_release/common/ps_completer/cmd"
	common__ptx "github.com/carapace-sh/carapace-bin/completers_release/common/ptx_completer/cmd"
	common__pulumi "github.com/carapace-sh/carapace-bin/completers_release/common/pulumi_completer/cmd"
	common__pwd "github.com/carapace-sh/carapace-bin/completers_release/common/pwd_completer/cmd"
	common__pwdx "github.com/carapace-sh/carapace-bin/completers_release/common/pwdx_completer/cmd"
	common__python "github.com/carapace-sh/carapace-bin/completers_release/common/python_completer/cmd"
	common__qmk "github.com/carapace-sh/carapace-bin/completers_release/common/qmk_completer/cmd"
	common__qrencode "github.com/carapace-sh/carapace-bin/completers_release/common/qrencode_completer/cmd"
	common__qutebrowser "github.com/carapace-sh/carapace-bin/completers_release/common/qutebrowser_completer/cmd"
	common__ramalama "github.com/carapace-sh/carapace-bin/completers_release/common/ramalama_completer/cmd"
	common__ranger "github.com/carapace-sh/carapace-bin/completers_release/common/ranger_completer/cmd"
	common__readlink "github.com/carapace-sh/carapace-bin/completers_release/common/readlink_completer/cmd"
	common__reboot "github.com/carapace-sh/carapace-bin/completers_release/common/reboot_completer/cmd"
	common__redis_cli "github.com/carapace-sh/carapace-bin/completers_release/common/redis-cli_completer/cmd"
	common__rename "github.com/carapace-sh/carapace-bin/completers_release/common/rename_completer/cmd"
	common__restic "github.com/carapace-sh/carapace-bin/completers_release/common/restic_completer/cmd"
	common__resume_cli "github.com/carapace-sh/carapace-bin/completers_release/common/resume-cli_completer/cmd"
	common__rg "github.com/carapace-sh/carapace-bin/completers_release/common/rg_completer/cmd"
	common__rifle "github.com/carapace-sh/carapace-bin/completers_release/common/rifle_completer/cmd"
	common__ripsecrets "github.com/carapace-sh/carapace-bin/completers_release/common/ripsecrets_completer/cmd"
	common__rm "github.com/carapace-sh/carapace-bin/completers_release/common/rm_completer/cmd"
	common__rmdir "github.com/carapace-sh/carapace-bin/completers_release/common/rmdir_completer/cmd"
	common__rmmod "github.com/carapace-sh/carapace-bin/completers_release/common/rmmod_completer/cmd"
	common__rsync "github.com/carapace-sh/carapace-bin/completers_release/common/rsync_completer/cmd"
	common__run0 "github.com/carapace-sh/carapace-bin/completers_release/common/run0_completer/cmd"
	common__rust_analyzer "github.com/carapace-sh/carapace-bin/completers_release/common/rust-analyzer_completer/cmd"
	common__rustc "github.com/carapace-sh/carapace-bin/completers_release/common/rustc_completer/cmd"
	common__rustdoc "github.com/carapace-sh/carapace-bin/completers_release/common/rustdoc_completer/cmd"
	common__rustup "github.com/carapace-sh/carapace-bin/completers_release/common/rustup_completer/cmd"
	common__saw "github.com/carapace-sh/carapace-bin/completers_release/common/saw_completer/cmd"
	common__scp "github.com/carapace-sh/carapace-bin/completers_release/common/scp_completer/cmd"
	common__script "github.com/carapace-sh/carapace-bin/completers_release/common/script_completer/cmd"
	common__scriptlive "github.com/carapace-sh/carapace-bin/completers_release/common/scriptlive_completer/cmd"
	common__scriptreplay "github.com/carapace-sh/carapace-bin/completers_release/common/scriptreplay_completer/cmd"
	common__scrot "github.com/carapace-sh/carapace-bin/completers_release/common/scrot_completer/cmd"
	common__sd "github.com/carapace-sh/carapace-bin/completers_release/common/sd_completer/cmd"
	common__sdkmanager "github.com/carapace-sh/carapace-bin/completers_release/common/sdkmanager_completer/cmd"
	common__sed "github.com/carapace-sh/carapace-bin/completers_release/common/sed_completer/cmd"
	common__semver "github.com/carapace-sh/carapace-bin/completers_release/common/semver_completer/cmd"
	common__seq "github.com/carapace-sh/carapace-bin/completers_release/common/seq_completer/cmd"
	common__serie "github.com/carapace-sh/carapace-bin/completers_release/common/serie_completer/cmd"
	common__set_env "github.com/carapace-sh/carapace-bin/completers_release/common/set-env_completer/cmd"
	common__sftp "github.com/carapace-sh/carapace-bin/completers_release/common/sftp_completer/cmd"
	common__sha1sum "github.com/carapace-sh/carapace-bin/completers_release/common/sha1sum_completer/cmd"
	common__sha224sum "github.com/carapace-sh/carapace-bin/completers_release/common/sha224sum_completer/cmd"
	common__sha256sum "github.com/carapace-sh/carapace-bin/completers_release/common/sha256sum_completer/cmd"
	common__sha384sum "github.com/carapace-sh/carapace-bin/completers_release/common/sha384sum_completer/cmd"
	common__sha512sum "github.com/carapace-sh/carapace-bin/completers_release/common/sha512sum_completer/cmd"
	common__showkey "github.com/carapace-sh/carapace-bin/completers_release/common/showkey_completer/cmd"
	common__shred "github.com/carapace-sh/carapace-bin/completers_release/common/shred_completer/cmd"
	common__shutdown "github.com/carapace-sh/carapace-bin/completers_release/common/shutdown_completer/cmd"
	common__slabtop "github.com/carapace-sh/carapace-bin/completers_release/common/slabtop_completer/cmd"
	common__sleep "github.com/carapace-sh/carapace-bin/completers_release/common/sleep_completer/cmd"
	common__slides "github.com/carapace-sh/carapace-bin/completers_release/common/slides_completer/cmd"
	common__soft "github.com/carapace-sh/carapace-bin/completers_release/common/soft_completer/cmd"
	common__sort "github.com/carapace-sh/carapace-bin/completers_release/common/sort_completer/cmd"
	common__speedtest_cli "github.com/carapace-sh/carapace-bin/completers_release/common/speedtest-cli_completer/cmd"
	common__split "github.com/carapace-sh/carapace-bin/completers_release/common/split_completer/cmd"
	common__sqlite3 "github.com/carapace-sh/carapace-bin/completers_release/common/sqlite3_completer/cmd"
	common__ssh_agent "github.com/carapace-sh/carapace-bin/completers_release/common/ssh-agent_completer/cmd"
	common__ssh_copy_id "github.com/carapace-sh/carapace-bin/completers_release/common/ssh-copy-id_completer/cmd"
	common__ssh_keygen "github.com/carapace-sh/carapace-bin/completers_release/common/ssh-keygen_completer/cmd"
	common__ssh "github.com/carapace-sh/carapace-bin/completers_release/common/ssh_completer/cmd"
	common__st "github.com/carapace-sh/carapace-bin/completers_release/common/st_completer/cmd"
	common__starship "github.com/carapace-sh/carapace-bin/completers_release/common/starship_completer/cmd"
	common__stat "github.com/carapace-sh/carapace-bin/completers_release/common/stat_completer/cmd"
	common__staticcheck "github.com/carapace-sh/carapace-bin/completers_release/common/staticcheck_completer/cmd"
	common__strings "github.com/carapace-sh/carapace-bin/completers_release/common/strings_completer/cmd"
	common__stty "github.com/carapace-sh/carapace-bin/completers_release/common/stty_completer/cmd"
	common__su "github.com/carapace-sh/carapace-bin/completers_release/common/su_completer/cmd"
	common__sudo "github.com/carapace-sh/carapace-bin/completers_release/common/sudo_completer/cmd"
	common__sudoedit "github.com/carapace-sh/carapace-bin/completers_release/common/sudoedit_completer/cmd"
	common__sudoreplay "github.com/carapace-sh/carapace-bin/completers_release/common/sudoreplay_completer/cmd"
	common__sulogin "github.com/carapace-sh/carapace-bin/completers_release/common/sulogin_completer/cmd"
	common__sum "github.com/carapace-sh/carapace-bin/completers_release/common/sum_completer/cmd"
	common__supervisorctl "github.com/carapace-sh/carapace-bin/completers_release/common/supervisorctl_completer/cmd"
	common__supervisord "github.com/carapace-sh/carapace-bin/completers_release/common/supervisord_completer/cmd"
	common__svg_term "github.com/carapace-sh/carapace-bin/completers_release/common/svg-term_completer/cmd"
	common__svgcleaner "github.com/carapace-sh/carapace-bin/completers_release/common/svgcleaner_completer/cmd"
	common__sway "github.com/carapace-sh/carapace-bin/completers_release/common/sway_completer/cmd"
	common__swaybar "github.com/carapace-sh/carapace-bin/completers_release/common/swaybar_completer/cmd"
	common__swaybg "github.com/carapace-sh/carapace-bin/completers_release/common/swaybg_completer/cmd"
	common__swayidle "github.com/carapace-sh/carapace-bin/completers_release/common/swayidle_completer/cmd"
	common__swaylock "github.com/carapace-sh/carapace-bin/completers_release/common/swaylock_completer/cmd"
	common__swaymsg "github.com/carapace-sh/carapace-bin/completers_release/common/swaymsg_completer/cmd"
	common__swaynag "github.com/carapace-sh/carapace-bin/completers_release/common/swaynag_completer/cmd"
	common__syft "github.com/carapace-sh/carapace-bin/completers_release/common/syft_completer/cmd"
	common__sync "github.com/carapace-sh/carapace-bin/completers_release/common/sync_completer/cmd"
	common__sysctl "github.com/carapace-sh/carapace-bin/completers_release/common/sysctl_completer/cmd"
	common__systemctl "github.com/carapace-sh/carapace-bin/completers_release/common/systemctl_completer/cmd"
	common__systemd_analyze "github.com/carapace-sh/carapace-bin/completers_release/common/systemd-analyze_completer/cmd"
	common__tac "github.com/carapace-sh/carapace-bin/completers_release/common/tac_completer/cmd"
	common__tail "github.com/carapace-sh/carapace-bin/completers_release/common/tail_completer/cmd"
	common__taplo "github.com/carapace-sh/carapace-bin/completers_release/common/taplo_completer/cmd"
	common__tar "github.com/carapace-sh/carapace-bin/completers_release/common/tar_completer/cmd"
	common__task "github.com/carapace-sh/carapace-bin/completers_release/common/task_completer/cmd"
	common__tea "github.com/carapace-sh/carapace-bin/completers_release/common/tea_completer/cmd"
	common__tee "github.com/carapace-sh/carapace-bin/completers_release/common/tee_completer/cmd"
	common__telnet "github.com/carapace-sh/carapace-bin/completers_release/common/telnet_completer/cmd"
	common__templ "github.com/carapace-sh/carapace-bin/completers_release/common/templ_completer/cmd"
	common__termux_apt_repo "github.com/carapace-sh/carapace-bin/completers_release/common/termux-apt-repo_completer/cmd"
	common__terraform_ls "github.com/carapace-sh/carapace-bin/completers_release/common/terraform-ls_completer/cmd"
	common__terraform "github.com/carapace-sh/carapace-bin/completers_release/common/terraform_completer/cmd"
	common__terragrunt "github.com/carapace-sh/carapace-bin/completers_release/common/terragrunt_completer/cmd"
	common__terramate "github.com/carapace-sh/carapace-bin/completers_release/common/terramate_completer/cmd"
	common__tesseract "github.com/carapace-sh/carapace-bin/completers_release/common/tesseract_completer/cmd"
	common__tig "github.com/carapace-sh/carapace-bin/completers_release/common/tig_completer/cmd"
	common__timeout "github.com/carapace-sh/carapace-bin/completers_release/common/timeout_completer/cmd"
	common__tinygo "github.com/carapace-sh/carapace-bin/completers_release/common/tinygo_completer/cmd"
	common__tldr__tealdeer "github.com/carapace-sh/carapace-bin/completers_release/common/tldr_completer/tealdeer/cmd"
	common__tldr__tldr_python_client "github.com/carapace-sh/carapace-bin/completers_release/common/tldr_completer/tldr-python-client/cmd"
	common__tload "github.com/carapace-sh/carapace-bin/completers_release/common/tload_completer/cmd"
	common__tmate "github.com/carapace-sh/carapace-bin/completers_release/common/tmate_completer/cmd"
	common__tmux "github.com/carapace-sh/carapace-bin/completers_release/common/tmux_completer/cmd"
	common__tofu "github.com/carapace-sh/carapace-bin/completers_release/common/tofu_completer/cmd"
	common__toit_lsp "github.com/carapace-sh/carapace-bin/completers_release/common/toit.lsp_completer/cmd"
	common__toit_pkg "github.com/carapace-sh/carapace-bin/completers_release/common/toit.pkg_completer/cmd"
	common__top "github.com/carapace-sh/carapace-bin/completers_release/common/top_completer/cmd"
	common__tor_browser "github.com/carapace-sh/carapace-bin/completers_release/common/tor-browser_completer/cmd"
	common__tor_gencert "github.com/carapace-sh/carapace-bin/completers_release/common/tor-gencert_completer/cmd"
	common__tor_print_ed_signing_cert "github.com/carapace-sh/carapace-bin/completers_release/common/tor-print-ed-signing-cert_completer/cmd"
	common__tor_resolve "github.com/carapace-sh/carapace-bin/completers_release/common/tor-resolve_completer/cmd"
	common__torsocks "github.com/carapace-sh/carapace-bin/completers_release/common/torsocks_completer/cmd"
	common__touch "github.com/carapace-sh/carapace-bin/completers_release/common/touch_completer/cmd"
	common__tox "github.com/carapace-sh/carapace-bin/completers_release/common/tox_completer/cmd"
	common__tr "github.com/carapace-sh/carapace-bin/completers_release/common/tr_completer/cmd"
	common__traefik "github.com/carapace-sh/carapace-bin/completers_release/common/traefik_completer/cmd"
	common__transmission_cli "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-cli_completer/cmd"
	common__transmission_create "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-create_completer/cmd"
	common__transmission_daemon "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-daemon_completer/cmd"
	common__transmission_edit "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-edit_completer/cmd"
	common__transmission_remote "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-remote_completer/cmd"
	common__transmission_show "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-show_completer/cmd"
	common__tree "github.com/carapace-sh/carapace-bin/completers_release/common/tree_completer/cmd"
	common__truncate "github.com/carapace-sh/carapace-bin/completers_release/common/truncate_completer/cmd"
	common__ts "github.com/carapace-sh/carapace-bin/completers_release/common/ts_completer/cmd"
	common__tsc "github.com/carapace-sh/carapace-bin/completers_release/common/tsc_completer/cmd"
	common__tsh "github.com/carapace-sh/carapace-bin/completers_release/common/tsh_completer/cmd"
	common__tshark "github.com/carapace-sh/carapace-bin/completers_release/common/tshark_completer/cmd"
	common__tsort "github.com/carapace-sh/carapace-bin/completers_release/common/tsort_completer/cmd"
	common__tty "github.com/carapace-sh/carapace-bin/completers_release/common/tty_completer/cmd"
	common__ttyd "github.com/carapace-sh/carapace-bin/completers_release/common/ttyd_completer/cmd"
	common__turbo "github.com/carapace-sh/carapace-bin/completers_release/common/turbo_completer/cmd"
	common__typst "github.com/carapace-sh/carapace-bin/completers_release/common/typst_completer/cmd"
	common__ufw "github.com/carapace-sh/carapace-bin/completers_release/common/ufw_completer/cmd"
	common__umount "github.com/carapace-sh/carapace-bin/completers_release/common/umount_completer/cmd"
	common__uname "github.com/carapace-sh/carapace-bin/completers_release/common/uname_completer/cmd"
	common__unbrotli "github.com/carapace-sh/carapace-bin/completers_release/common/unbrotli_completer/cmd"
	common__unexpand "github.com/carapace-sh/carapace-bin/completers_release/common/unexpand_completer/cmd"
	common__uniq "github.com/carapace-sh/carapace-bin/completers_release/common/uniq_completer/cmd"
	common__unlink "github.com/carapace-sh/carapace-bin/completers_release/common/unlink_completer/cmd"
	common__unlzma "github.com/carapace-sh/carapace-bin/completers_release/common/unlzma_completer/cmd"
	common__unpigz "github.com/carapace-sh/carapace-bin/completers_release/common/unpigz_completer/cmd"
	common__unset_env "github.com/carapace-sh/carapace-bin/completers_release/common/unset-env_completer/cmd"
	common__unxz "github.com/carapace-sh/carapace-bin/completers_release/common/unxz_completer/cmd"
	common__unzip "github.com/carapace-sh/carapace-bin/completers_release/common/unzip_completer/cmd"
	common__upower "github.com/carapace-sh/carapace-bin/completers_release/common/upower_completer/cmd"
	common__uptime "github.com/carapace-sh/carapace-bin/completers_release/common/uptime_completer/cmd"
	common__upx "github.com/carapace-sh/carapace-bin/completers_release/common/upx_completer/cmd"
	common__useradd "github.com/carapace-sh/carapace-bin/completers_release/common/useradd_completer/cmd"
	common__userdel "github.com/carapace-sh/carapace-bin/completers_release/common/userdel_completer/cmd"
	common__usermod "github.com/carapace-sh/carapace-bin/completers_release/common/usermod_completer/cmd"
	common__users "github.com/carapace-sh/carapace-bin/completers_release/common/users_completer/cmd"
	common__vagrant "github.com/carapace-sh/carapace-bin/completers_release/common/vagrant_completer/cmd"
	common__vault "github.com/carapace-sh/carapace-bin/completers_release/common/vault_completer/cmd"
	common__vdir "github.com/carapace-sh/carapace-bin/completers_release/common/vdir_completer/cmd"
	common__vercel "github.com/carapace-sh/carapace-bin/completers_release/common/vercel_completer/cmd"
	common__vhs "github.com/carapace-sh/carapace-bin/completers_release/common/vhs_completer/cmd"
	common__vi "github.com/carapace-sh/carapace-bin/completers_release/common/vi_completer/cmd"
	common__viewnior "github.com/carapace-sh/carapace-bin/completers_release/common/viewnior_completer/cmd"
	common__vim "github.com/carapace-sh/carapace-bin/completers_release/common/vim_completer/cmd"
	common__visudo "github.com/carapace-sh/carapace-bin/completers_release/common/visudo_completer/cmd"
	common__viu "github.com/carapace-sh/carapace-bin/completers_release/common/viu_completer/cmd"
	common__vivid "github.com/carapace-sh/carapace-bin/completers_release/common/vivid_completer/cmd"
	common__vlc "github.com/carapace-sh/carapace-bin/completers_release/common/vlc_completer/cmd"
	common__vmstat "github.com/carapace-sh/carapace-bin/completers_release/common/vmstat_completer/cmd"
	common__volta "github.com/carapace-sh/carapace-bin/completers_release/common/volta_completer/cmd"
	common__w "github.com/carapace-sh/carapace-bin/completers_release/common/w_completer/cmd"
	common__watch "github.com/carapace-sh/carapace-bin/completers_release/common/watch_completer/cmd"
	common__watchexec "github.com/carapace-sh/carapace-bin/completers_release/common/watchexec_completer/cmd"
	common__watchgnupg "github.com/carapace-sh/carapace-bin/completers_release/common/watchgnupg_completer/cmd"
	common__waypoint "github.com/carapace-sh/carapace-bin/completers_release/common/waypoint_completer/cmd"
	common__wc "github.com/carapace-sh/carapace-bin/completers_release/common/wc_completer/cmd"
	common__wezterm "github.com/carapace-sh/carapace-bin/completers_release/common/wezterm_completer/cmd"
	common__wget "github.com/carapace-sh/carapace-bin/completers_release/common/wget_completer/cmd"
	common__whereis "github.com/carapace-sh/carapace-bin/completers_release/common/whereis_completer/cmd"
	common__which "github.com/carapace-sh/carapace-bin/completers_release/common/which_completer/cmd"
	common__who "github.com/carapace-sh/carapace-bin/completers_release/common/who_completer/cmd"
	common__whoami "github.com/carapace-sh/carapace-bin/completers_release/common/whoami_completer/cmd"
	common__wine "github.com/carapace-sh/carapace-bin/completers_release/common/wine_completer/cmd"
	common__wineboot "github.com/carapace-sh/carapace-bin/completers_release/common/wineboot_completer/cmd"
	common__winepath "github.com/carapace-sh/carapace-bin/completers_release/common/winepath_completer/cmd"
	common__wineserver "github.com/carapace-sh/carapace-bin/completers_release/common/wineserver_completer/cmd"
	common__winetricks "github.com/carapace-sh/carapace-bin/completers_release/common/winetricks_completer/cmd"
	common__wire "github.com/carapace-sh/carapace-bin/completers_release/common/wire_completer/cmd"
	common__wireshark "github.com/carapace-sh/carapace-bin/completers_release/common/wireshark_completer/cmd"
	common__wishlist "github.com/carapace-sh/carapace-bin/completers_release/common/wishlist_completer/cmd"
	common__wl_mirror "github.com/carapace-sh/carapace-bin/completers_release/common/wl-mirror_completer/cmd"
	common__woeusb "github.com/carapace-sh/carapace-bin/completers_release/common/woeusb_completer/cmd"
	common__xargs "github.com/carapace-sh/carapace-bin/completers_release/common/xargs_completer/cmd"
	common__xbacklight "github.com/carapace-sh/carapace-bin/completers_release/common/xbacklight_completer/cmd"
	common__xclip "github.com/carapace-sh/carapace-bin/completers_release/common/xclip_completer/cmd"
	common__xdotool "github.com/carapace-sh/carapace-bin/completers_release/common/xdotool_completer/cmd"
	common__xh "github.com/carapace-sh/carapace-bin/completers_release/common/xh_completer/cmd"
	common__xonsh "github.com/carapace-sh/carapace-bin/completers_release/common/xonsh_completer/cmd"
	common__xz "github.com/carapace-sh/carapace-bin/completers_release/common/xz_completer/cmd"
	common__xzcat "github.com/carapace-sh/carapace-bin/completers_release/common/xzcat_completer/cmd"
	common__yarn "github.com/carapace-sh/carapace-bin/completers_release/common/yarn_completer/cmd"
	common__yay "github.com/carapace-sh/carapace-bin/completers_release/common/yay_completer/cmd"
	common__yes "github.com/carapace-sh/carapace-bin/completers_release/common/yes_completer/cmd"
	common__yj "github.com/carapace-sh/carapace-bin/completers_release/common/yj_completer/cmd"
	common__youtube_dl "github.com/carapace-sh/carapace-bin/completers_release/common/youtube-dl_completer/cmd"
	common__yt_dlp "github.com/carapace-sh/carapace-bin/completers_release/common/yt-dlp_completer/cmd"
	common__zathura "github.com/carapace-sh/carapace-bin/completers_release/common/zathura_completer/cmd"
	common__zcat "github.com/carapace-sh/carapace-bin/completers_release/common/zcat_completer/cmd"
	common__zip "github.com/carapace-sh/carapace-bin/completers_release/common/zip_completer/cmd"
	common__zoxide "github.com/carapace-sh/carapace-bin/completers_release/common/zoxide_completer/cmd"
	darwin__fileicon "github.com/carapace-sh/carapace-bin/completers_release/darwin/fileicon_completer/cmd"
	darwin__skhd "github.com/carapace-sh/carapace-bin/completers_release/darwin/skhd_completer/cmd"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
)

var completers = completer.CompleterMap{
	"acpi": {
		{
			Name:        "acpi",
			Description: "Shows information from the /proc filesystem",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/acpi_completer/cmd",
			Variant:     "",
			Url:         "https://www.commandlinux.com/man-page/man1/acpi.1.html",
			Execute:     common__acpi.Execute,
		},
	},
	"acpid": {
		{
			Name:        "acpid",
			Description: "Advanced Configuration and Power Interface event daemon",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/acpid_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/acpid",
			Execute:     common__acpid.Execute,
		},
	},
	"adb": {
		{
			Name:        "adb",
			Description: "Android Debug Bridge",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/adb_completer/cmd",
			Variant:     "",
			Url:         "https://developer.android.com/studio/command-line/adb",
			Execute:     common__adb.Execute,
		},
	},
	"age": {
		{
			Name:        "age",
			Description: "simple, modern, and secure file encryption",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/age_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/FiloSottile/age",
			Execute:     common__age.Execute,
		},
	},
	"agg": {
		{
			Name:        "agg",
			Description: "asciinema gif generator",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/agg_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/asciinema/agg",
			Execute:     common__agg.Execute,
		},
	},
	"alsamixer": {
		{
			Name:        "alsamixer",
			Description: "soundcard mixer for ALSA soundcard driver, with ncurses interface",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/alsamixer_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Alsamixer",
			Execute:     common__alsamixer.Execute,
		},
	},
	"ant": {
		{
			Name:        "ant",
			Description: "software tool for automating software build processes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ant_completer/cmd",
			Variant:     "",
			Url:         "https://ant.apache.org/",
			Execute:     common__ant.Execute,
		},
	},
	"apk": {
		{
			Name:        "apk",
			Description: "Alpine package manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/apk_completer/cmd",
			Variant:     "",
			Url:         "https://gitlab.alpinelinux.org/alpine/apk-tools",
			Execute:     common__apk.Execute,
		},
	},
	"aplay": {
		{
			Name:        "aplay",
			Description: "command-line sound recorder and player for ALSA soundcard driver",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/aplay_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Aplay",
			Execute:     common__aplay.Execute,
		},
	},
	"apropos": {
		{
			Name:        "apropos",
			Description: "search the manual page names and descriptions",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/apropos_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/apropos",
			Execute:     common__apropos.Execute,
		},
	},
	"apt": {
		{
			Name:        "apt",
			Description: "apt is a commandline package manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/apt_completer/cmd",
			Variant:     "",
			Url:         "https://salsa.debian.org/apt-team/apt",
			Execute:     common__apt.Execute,
		},
	},
	"apt-cache": {
		{
			Name:        "apt-cache",
			Description: "query the APT cache",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/apt-cache_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/apt-cache",
			Execute:     common__apt_cache.Execute,
		},
	},
	"apt-get": {
		{
			Name:        "apt-get",
			Description: "APT package handling utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/apt-get_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/apt-get",
			Execute:     common__apt_get.Execute,
		},
	},
	"ar": {
		{
			Name:        "ar",
			Description: "create, modify, and extract from archives",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ar_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ar",
			Execute:     common__ar.Execute,
		},
	},
	"arecord": {
		{
			Name:        "arecord",
			Description: "command-line sound recorder and player for ALSA soundcard driver",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/arecord_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/arecord",
			Execute:     common__arecord.Execute,
		},
	},
	"aria2c": {
		{
			Name:        "aria2c",
			Description: "The ultra fast download utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/aria2c_completer/cmd",
			Variant:     "",
			Url:         "https://aria2.github.io/",
			Execute:     common__aria2c.Execute,
		},
	},
	"artisan": {
		{
			Name:        "artisan",
			Description: "Artisan is the command line interface included with Laravel",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/artisan_completer/cmd",
			Variant:     "",
			Url:         "https://laravel.com/",
			Execute:     common__artisan.Execute,
		},
	},
	"asciinema": {
		{
			Name:        "asciinema",
			Description: "Record and share your terminal sessions, the right way.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/asciinema_completer/cmd",
			Variant:     "",
			Url:         "https://asciinema.org/",
			Execute:     common__asciinema.Execute,
		},
	},
	"autoconf": {
		{
			Name:        "autoconf",
			Description: "Generate a configuration script from a TEMPLATE-FILE",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/autoconf_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/autoconf",
			Execute:     common__autoconf.Execute,
		},
	},
	"avdmanager": {
		{
			Name:        "avdmanager",
			Description: "create and manage Android Virtual Devices",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/avdmanager_completer/cmd",
			Variant:     "",
			Url:         "https://developer.android.com/studio/command-line/avdmanager",
			Execute:     common__avdmanager.Execute,
		},
	},
	"awk": {
		{
			Name:        "awk",
			Description: "pattern scanning and processing language",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/awk_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/AWK",
			Execute:     common__awk.Execute,
		},
	},
	"aws": {
		{
			Name:        "aws",
			Description: "Universal Command Line Interface for Amazon Web Services",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/aws_completer/cmd",
			Variant:     "",
			Url:         "https://aws.amazon.com/cli/",
			Execute:     common__aws.Execute,
		},
	},
	"az": {
		{
			Name:        "az",
			Description: "Azure Command-Line Interface",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/az_completer/cmd",
			Variant:     "",
			Url:         "https://docs.microsoft.com/en-us/cli/azure/",
			Execute:     common__az.Execute,
		},
	},
	"baobab": {
		{
			Name:        "baobab",
			Description: "A graphical disk usage analyzer for the GNOME deskto",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/baobab_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.gnome.org/action/show/Apps/DiskUsageAnalyzer",
			Execute:     common__baobab.Execute,
		},
	},
	"base32": {
		{
			Name:        "base32",
			Description: "Base32 encode or decode FILE, or standard input, to standard output",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/base32_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/base32.1.html",
			Execute:     common__base32.Execute,
		},
	},
	"base64": {
		{
			Name:        "base64",
			Description: "Base64 encode or decode FILE, or standard input, to standard output",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/base64_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/base64.1.html",
			Execute:     common__base64.Execute,
		},
	},
	"basename": {
		{
			Name:        "basename",
			Description: "strip directory and suffix from filenames",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/basename_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Basename",
			Execute:     common__basename.Execute,
		},
	},
	"bash": {
		{
			Name:        "bash",
			Description: "GNU Bourne-Again SHell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bash_completer/cmd",
			Variant:     "",
			Url:         "https://www.gnu.org/software/bash/",
			Execute:     common__bash.Execute,
		},
	},
	"bash-language-server": {
		{
			Name:        "bash-language-server",
			Description: "A language server for Bash",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bash-language-server_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/bash-lsp/bash-language-server",
			Execute:     common__bash_language_server.Execute,
		},
	},
	"bat": {
		{
			Name:        "bat",
			Description: "a cat clone with syntax highlighting and Git integration",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bat_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/sharkdp/bat",
			Execute:     common__bat.Execute,
		},
	},
	"batdiff": {
		{
			Name:        "batdiff",
			Description: "Diff a file against the current git index, or display the diff between two files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/batdiff_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/eth-p/bat-extras/blob/master/doc/batdiff.md",
			Execute:     common__batdiff.Execute,
		},
	},
	"batgrep": {
		{
			Name:        "batgrep",
			Description: "Quickly search through and highlight files using ripgrep",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/batgrep_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/eth-p/bat-extras/blob/master/doc/batgrep.md",
			Execute:     common__batgrep.Execute,
		},
	},
	"batman": {
		{
			Name:        "batman",
			Description: "Read system manual pages (man) using bat",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/batman_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/eth-p/bat-extras/blob/master/doc/batman.md",
			Execute:     common__batman.Execute,
		},
	},
	"bats": {
		{
			Name:        "bats",
			Description: "Bash Automated Testing System",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bats_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/bats-core/bats-core",
			Execute:     common__bats.Execute,
		},
	},
	"bc": {
		{
			Name:        "bc",
			Description: "An arbitrary precision calculator language",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bc_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/bc",
			Execute:     common__bc.Execute,
		},
	},
	"benthos": {
		{
			Name:        "benthos",
			Description: "A stream processor for mundane tasks",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/benthos_completer/cmd",
			Variant:     "",
			Url:         "https://www.benthos.dev",
			Execute:     common__benthos.Execute,
		},
	},
	"black": {
		{
			Name:        "black",
			Description: "The uncompromising code formatter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/black_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/psf/black",
			Execute:     common__black.Execute,
		},
	},
	"bloop": {
		{
			Name:        "bloop",
			Description: "Build and test Scala code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bloop_completer/cmd",
			Variant:     "",
			Url:         "https://scalacenter.github.io/bloop/",
			Execute:     common__bloop.Execute,
		},
	},
	"bluetoothctl": {
		{
			Name:        "bluetoothctl",
			Description: "Bluetooth Control Command Line Tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bluetoothctl_completer/cmd",
			Variant:     "",
			Url:         "https://www.bluez.org/",
			Execute:     common__bluetoothctl.Execute,
		},
	},
	"boundary": {
		{
			Name:        "boundary",
			Description: "Boundary enables identity-based access management for dynamic infrastructure",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/boundary_completer/cmd",
			Variant:     "",
			Url:         "https://www.boundaryproject.io/downloads",
			Execute:     common__boundary.Execute,
		},
	},
	"brew": {
		{
			Name:        "brew",
			Description: "The missing package manager for macOS",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/brew_completer/cmd",
			Variant:     "",
			Url:         "https://brew.sh/",
			Execute:     common__brew.Execute,
		},
	},
	"brotli": {
		{
			Name:        "brotli",
			Description: "compress or decompress files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/brotli_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/google/brotli",
			Execute:     common__brotli.Execute,
		},
	},
	"bru": {
		{
			Name:        "bru",
			Description: "Bruno CLI",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bru_completer/cmd",
			Variant:     "",
			Url:         "https://docs.usebruno.com/bru-cli/overview",
			Execute:     common__bru.Execute,
		},
	},
	"btop": {
		{
			Name:        "btop",
			Description: "A monitor of resources",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/btop_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/aristocratos/btop",
			Execute:     common__btop.Execute,
		},
	},
	"bun": {
		{
			Name:        "bun",
			Description: "a fast bundler, transpiler, JavaScript Runtime and package manager for web software",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bun_completer/cmd",
			Variant:     "",
			Url:         "https://bun.sh/",
			Execute:     common__bun.Execute,
		},
	},
	"bunx": {
		{
			Name:        "bunx",
			Description: "bun package manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/bunx_completer/cmd",
			Variant:     "",
			Url:         "https://bun.sh/docs/cli/bunx",
			Execute:     common__bunx.Execute,
		},
	},
	"but": {
		{
			Name:        "but",
			Description: "A GitButler CLI tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/but_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/gitbutlerapp/gitbutler/tree/master/crates/but",
			Execute:     common__but.Execute,
		},
	},
	"cal": {
		{
			Name:        "cal",
			Description: "display a calendar",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cal_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/cal.1.html",
			Execute:     common__cal.Execute,
		},
	},
	"calibre": {
		{
			Name:        "calibre",
			Description: "Comprehensive e-book software",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/calibre_completer/cmd",
			Variant:     "",
			Url:         "https://calibre-ebook.com/",
			Execute:     common__calibre.Execute,
		},
	},
	"capslock": {
		{
			Name:        "capslock",
			Description: "Capslock is a capability analysis CLI for Go packages ",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/capslock_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/google/capslock",
			Execute:     common__capslock.Execute,
		},
	},
	"cargo": {
		{
			Name:        "cargo",
			Description: "Rust's package manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo_completer/cmd",
			Variant:     "",
			Url:         "https://doc.rust-lang.org/cargo/",
			Execute:     common__cargo.Execute,
		},
	},
	"cargo-clippy": {
		{
			Name:        "cargo-clippy",
			Description: "Checks a package to catch common mistakes and improve your Rust code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-clippy_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/arcnmx/cargo-clippy",
			Execute:     common__cargo_clippy.Execute,
		},
	},
	"cargo-fmt": {
		{
			Name:        "cargo-fmt",
			Description: "format all bin and lib files of the current crate",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-fmt_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/rust-lang/rustfmt",
			Execute:     common__cargo_fmt.Execute,
		},
	},
	"cargo-metadata": {
		{
			Name:        "cargo-metadata",
			Description: "Output the resolved dependencies of a package",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-metadata_completer/cmd",
			Variant:     "",
			Url:         "https://doc.rust-lang.org/cargo/commands/cargo-metadata.html",
			Execute:     common__cargo_metadata.Execute,
		},
	},
	"cargo-rm": {
		{
			Name:        "cargo-rm",
			Description: "Remove a dependency from a Cargo.toml manifest file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-rm_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/killercup/cargo-edit",
			Execute:     common__cargo_rm.Execute,
		},
	},
	"cargo-set-version": {
		{
			Name:        "cargo-set-version",
			Description: "Change a package's version in the local manifest file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-set-version_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/killercup/cargo-edit",
			Execute:     common__cargo_set_version.Execute,
		},
	},
	"cargo-upgrade": {
		{
			Name:        "cargo-upgrade",
			Description: "Update dependencies as recorded in the local lock file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-upgrade_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/killercup/cargo-edit",
			Execute:     common__cargo_upgrade.Execute,
		},
	},
	"cargo-watch": {
		{
			Name:        "cargo-watch",
			Description: "Watches over your Cargo project’s source",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cargo-watch_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/watchexec/cargo-watch",
			Execute:     common__cargo_watch.Execute,
		},
	},
	"cat": {
		{
			Name:        "cat",
			Description: "concatenate files and print on the standard output",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cat_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/cat",
			Execute:     common__cat.Execute,
		},
	},
	"cfdisk": {
		{
			Name:        "cfdisk",
			Description: "display or manipulate a disk partition table",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cfdisk_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/cfdisk.8.html",
			Execute:     common__cfdisk.Execute,
		},
	},
	"charm": {
		{
			Name:        "charm",
			Description: "Do Charm stuff",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/charm_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/charm",
			Execute:     common__charm.Execute,
		},
	},
	"chcpu": {
		{
			Name:        "chcpu",
			Description: "configure CPUs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chcpu_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/chcpu.8.html",
			Execute:     common__chcpu.Execute,
		},
	},
	"chdman": {
		{
			Name:        "chdman",
			Description: "MAME Compressed Hunks of Data (CHD) manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chdman_completer/cmd",
			Variant:     "",
			Url:         "https://docs.mamedev.org/tools/chdman.html",
			Execute:     common__chdman.Execute,
		},
	},
	"cheese": {
		{
			Name:        "cheese",
			Description: "tool to take pictures and videos from your webcam",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cheese_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.gnome.org/Apps/Cheese",
			Execute:     common__cheese.Execute,
		},
	},
	"chgrp": {
		{
			Name:        "chgrp",
			Description: "change group ownership",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chgrp_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/chgrp.1.html",
			Execute:     common__chgrp.Execute,
		},
	},
	"chmod": {
		{
			Name:        "chmod",
			Description: "change file mode bits",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chmod_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/chmod.1.html",
			Execute:     common__chmod.Execute,
		},
	},
	"chown": {
		{
			Name:        "chown",
			Description: "change file owner and group",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chown_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/chown.1.html",
			Execute:     common__chown.Execute,
		},
	},
	"chpasswd": {
		{
			Name:        "chpasswd",
			Description: "update passwords in batch mode",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chpasswd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/chpasswd",
			Execute:     common__chpasswd.Execute,
		},
	},
	"chroma": {
		{
			Name:        "chroma",
			Description: "Chroma is a general purpose syntax highlighter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chroma_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/alecthomas/chroma",
			Execute:     common__chroma.Execute,
		},
	},
	"chromium": {
		{
			Name:        "chromium",
			Description: "chromium browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chromium_completer/cmd",
			Variant:     "",
			Url:         "https://www.chromium.org/Home",
			Execute:     common__chromium.Execute,
		},
	},
	"chroot": {
		{
			Name:        "chroot",
			Description: "run command or interactive shell with special root directory",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chroot_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Chroot",
			Execute:     common__chroot.Execute,
		},
	},
	"chsh": {
		{
			Name:        "chsh",
			Description: "Change your login shell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/chsh_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Chsh",
			Execute:     common__chsh.Execute,
		},
	},
	"circleci": {
		{
			Name:        "circleci",
			Description: "Use CircleCI from the command line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/circleci_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/CircleCI-Public/circleci-cli",
			Execute:     common__circleci.Execute,
		},
	},
	"cksum": {
		{
			Name:        "cksum",
			Description: "compute and verify file checksums",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cksum_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/cksum.1.html",
			Execute:     common__cksum.Execute,
		},
	},
	"clamav-config": {
		{
			Name:        "clamav-config",
			Description: "clamav config",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamav-config_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamav_config.Execute,
		},
	},
	"clamav-milter": {
		{
			Name:        "clamav-milter",
			Description: "milter compatible scanner",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamav-milter_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamav_milter.Execute,
		},
	},
	"clambc": {
		{
			Name:        "clambc",
			Description: "Bytecode Analysis and Testing Tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clambc_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clambc.Execute,
		},
	},
	"clamconf": {
		{
			Name:        "clamconf",
			Description: "Clam AntiVirus configuration utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamconf_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamconf.Execute,
		},
	},
	"clamd": {
		{
			Name:        "clamd",
			Description: "an anti-virus daemon",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamd_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamd.Execute,
		},
	},
	"clamdscan": {
		{
			Name:        "clamdscan",
			Description: "scan files and directories for viruses using Clam AntiVirus Daemon",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamdscan_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamdscan.Execute,
		},
	},
	"clamdtop": {
		{
			Name:        "clamdtop",
			Description: "monitor the Clam AntiVirus Daemon",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamdtop_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamdtop.Execute,
		},
	},
	"clamonacc": {
		{
			Name:        "clamonacc",
			Description: "an anti-virus on-access scanning daemon and clamd client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamonacc_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamonacc.Execute,
		},
	},
	"clamscan": {
		{
			Name:        "clamscan",
			Description: "scan files and directories for viruses",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamscan_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamscan.Execute,
		},
	},
	"clamsubmit": {
		{
			Name:        "clamsubmit",
			Description: "File submission utility for ClamAV",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/clamsubmit_completer/cmd",
			Variant:     "",
			Url:         "http://www.clamav.net/",
			Execute:     common__clamsubmit.Execute,
		},
	},
	"cmus": {
		{
			Name:        "cmus",
			Description: "Curses based music player",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cmus_completer/cmd",
			Variant:     "",
			Url:         "https://cmus.github.io/",
			Execute:     common__cmus.Execute,
		},
	},
	"code": {
		{
			Name:        "code",
			Description: "Visual Studio Code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/code_completer/cmd",
			Variant:     "",
			Url:         "https://code.visualstudio.com/",
			Execute:     common__code.Execute,
		},
	},
	"code-insiders": {
		{
			Name:        "code-insiders",
			Description: "Visual Studio Code Insiders",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/code-insiders_completer/cmd",
			Variant:     "",
			Url:         "https://code.visualstudio.com/insiders/",
			Execute:     common__code_insiders.Execute,
		},
	},
	"codecov": {
		{
			Name:        "codecov",
			Description: "codecov uploader",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/codecov_completer/cmd",
			Variant:     "",
			Url:         "https://docs.codecov.com/docs/codecov-uploader",
			Execute:     common__codecov.Execute,
		},
	},
	"comm": {
		{
			Name:        "comm",
			Description: "compare two sorted files line by line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/comm_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/comm",
			Execute:     common__comm.Execute,
		},
	},
	"conda": {
		{
			Name:        "conda",
			Description: "conda is a tool for managing and deploying applications, environments and packages",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/conda_completer/cmd",
			Variant:     "",
			Url:         "https://docs.conda.io/en/latest/",
			Execute:     common__conda.Execute,
		},
	},
	"conda-content-trust": {
		{
			Name:        "conda-content-trust",
			Description: "Signing and verification tools for Conda",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/conda-content-trust_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/conda/conda-content-trust",
			Execute:     common__conda_content_trust.Execute,
		},
	},
	"conda-env": {
		{
			Name:        "conda-env",
			Description: "Manage conda environments",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/conda-env_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/conda/conda/tree/main/conda_env",
			Execute:     common__conda_env.Execute,
		},
	},
	"conky": {
		{
			Name:        "conky",
			Description: "A system  monitor for X originally based on the torsmo code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/conky_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/brndnmtthws/conky",
			Execute:     common__conky.Execute,
		},
	},
	"consul": {
		{
			Name:        "consul",
			Description: "Consul automates networking for simple and secure application delivery",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/consul_completer/cmd",
			Variant:     "",
			Url:         "https://www.consul.io/",
			Execute:     common__consul.Execute,
		},
	},
	"coredumpctl": {
		{
			Name:        "coredumpctl",
			Description: "List or retrieve coredumps from the journal",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/coredumpctl_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/coredumpctl.html",
			Execute:     common__coredumpctl.Execute,
		},
	},
	"cp": {
		{
			Name:        "cp",
			Description: "copy files and directories",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cp_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/cp.1.html",
			Execute:     common__cp.Execute,
		},
	},
	"csplit": {
		{
			Name:        "csplit",
			Description: "split a file into sections determined by context lines",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/csplit_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/csplit",
			Execute:     common__csplit.Execute,
		},
	},
	"csview": {
		{
			Name:        "csview",
			Description: "A high performance csv viewer with cjk/emoji support",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/csview_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/wfxr/csview",
			Execute:     common__csview.Execute,
		},
	},
	"cura": {
		{
			Name:        "cura",
			Description: "Powerful, easy-to-use 3D printing software",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cura_completer/cmd",
			Variant:     "",
			Url:         "https://ultimaker.com/software/ultimaker-cura",
			Execute:     common__cura.Execute,
		},
	},
	"curl": {
		{
			Name:        "curl",
			Description: "transfer a URL",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/curl_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/curl/curl",
			Execute:     common__curl.Execute,
		},
	},
	"cut": {
		{
			Name:        "cut",
			Description: "remove sections from each line of files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/cut_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/cut.1.html",
			Execute:     common__cut.Execute,
		},
	},
	"d2": {
		{
			Name:        "d2",
			Description: "compiles and renders d2 diagrams into svgs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/d2_completer/cmd",
			Variant:     "",
			Url:         "https://d2lang.com/",
			Execute:     common__d2.Execute,
		},
	},
	"dagger": {
		{
			Name:        "dagger",
			Description: "The Dagger CLI provides a command-line interface to Dagger.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dagger_completer/cmd",
			Variant:     "",
			Url:         "https://dagger.io/",
			Execute:     common__dagger.Execute,
		},
	},
	"darktable": {
		{
			Name:        "darktable",
			Description: "a digital photography workflow application",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/darktable_completer/cmd",
			Variant:     "",
			Url:         "https://darktable-org.github.io/dtdocs/en-us/",
			Execute:     common__darktable.Execute,
		},
	},
	"darktable-cli": {
		{
			Name:        "darktable-cli",
			Description: "a command line darktable variant",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/darktable-cli_completer/cmd",
			Variant:     "",
			Url:         "https://darktable-org.github.io/dtdocs/en-us/special-topics/program-invocation/darktable-cli/",
			Execute:     common__darktable_cli.Execute,
		},
	},
	"dart": {
		{
			Name:        "dart",
			Description: "A command-line utility for Dart development",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dart_completer/cmd",
			Variant:     "",
			Url:         "https://dart.dev/",
			Execute:     common__dart.Execute,
		},
	},
	"date": {
		{
			Name:        "date",
			Description: "print or set the system date and time",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/date_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/date",
			Execute:     common__date.Execute,
		},
	},
	"dbt": {
		{
			Name:        "dbt",
			Description: "An ELT tool for managing your SQL transformations and data models",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dbt_completer/cmd",
			Variant:     "",
			Url:         "https://www.getdbt.com/",
			Execute:     common__dbt.Execute,
		},
	},
	"dc": {
		{
			Name:        "dc",
			Description: "an arbitrary precision calculator",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dc_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dc",
			Execute:     common__dc.Execute,
		},
	},
	"dd": {
		{
			Name:        "dd",
			Description: "convert and copy a file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dd",
			Execute:     common__dd.Execute,
		},
	},
	"deadcode": {
		{
			Name:        "deadcode",
			Description: "The deadcode command reports unreachable functions in Go programs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/deadcode_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/golang.org/x/tools/internal/cmd/deadcode",
			Execute:     common__deadcode.Execute,
		},
	},
	"delta": {
		{
			Name:        "delta",
			Description: "A viewer for git and diff output",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/delta_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/dandavison/delta",
			Execute:     common__delta.Execute,
		},
	},
	"deno": {
		{
			Name:        "deno",
			Description: "A modern JavaScript and TypeScript runtime",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/deno_completer/cmd",
			Variant:     "",
			Url:         "https://deno.land/",
			Execute:     common__deno.Execute,
		},
	},
	"devbox": {
		{
			Name:        "devbox",
			Description: "Instant, easy, predictable development environments",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/devbox_completer/cmd",
			Variant:     "",
			Url:         "https://www.jetpack.io/devbox/",
			Execute:     common__devbox.Execute,
		},
	},
	"df": {
		{
			Name:        "df",
			Description: "report file system disk space usage",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/df_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/df.1.html",
			Execute:     common__df.Execute,
		},
	},
	"dfc": {
		{
			Name:        "dfc",
			Description: "report file system space usage information with style",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dfc_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/rolinh/dfc",
			Execute:     common__dfc.Execute,
		},
	},
	"dict": {
		{
			Name:        "dict",
			Description: "Query a dictd server for the definition of a word",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dict_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dict",
			Execute:     common__dict.Execute,
		},
	},
	"diff": {
		{
			Name:        "diff",
			Description: "compare files line by line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/diff_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/diff",
			Execute:     common__diff.Execute,
		},
	},
	"diff3": {
		{
			Name:        "diff3",
			Description: "compare three files line by line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/diff3_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/diff3",
			Execute:     common__diff3.Execute,
		},
	},
	"dig": {
		{
			Name:        "dig",
			Description: "DNS lookup utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dig_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dig",
			Execute:     common__dig.Execute,
		},
	},
	"dir": {
		{
			Name:        "dir",
			Description: "list directory contents",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dir_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dir",
			Execute:     common__dir.Execute,
		},
	},
	"dircolors": {
		{
			Name:        "dircolors",
			Description: "color setup for ls",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dircolors_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dircolors",
			Execute:     common__dircolors.Execute,
		},
	},
	"direnv": {
		{
			Name:        "direnv",
			Description: "unclutter your .profile",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/direnv_completer/cmd",
			Variant:     "",
			Url:         "https://direnv.net/",
			Execute:     common__direnv.Execute,
		},
	},
	"dirname": {
		{
			Name:        "dirname",
			Description: "strip last component from file name",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dirname_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dirname",
			Execute:     common__dirname.Execute,
		},
	},
	"dive": {
		{
			Name:        "dive",
			Description: "Docker Image Visualizer & Explorer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dive_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/wagoodman/dive",
			Execute:     common__dive.Execute,
		},
	},
	"dlv": {
		{
			Name:        "dlv",
			Description: "Delve is a debugger for the Go programming language.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dlv_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/go-delve/delve",
			Execute:     common__dlv.Execute,
		},
	},
	"dmenu": {
		{
			Name:        "dmenu",
			Description: "dynamic menu",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dmenu_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dmenu",
			Execute:     common__dmenu.Execute,
		},
	},
	"dmesg": {
		{
			Name:        "dmesg",
			Description: "Display or control the kernel ring buffer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dmesg_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/dmesg",
			Execute:     common__dmesg.Execute,
		},
	},
	"dms": {
		{
			Name:        "dms",
			Description: "A UPnP DLNA Digital Media Server",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dms_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/anacrolix/dms",
			Execute:     common__dms.Execute,
		},
	},
	"dngconverter": {
		{
			Name:        "dngconverter",
			Description: "Adobe DNG Converter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dngconverter_completer/cmd",
			Variant:     "",
			Url:         "https://helpx.adobe.com/camera-raw/using/adobe-dng-converter.html",
			Execute:     common__dngconverter.Execute,
		},
	},
	"dnsmasq": {
		{
			Name:        "dnsmasq",
			Description: "A lightweight DHCP and caching DNS server",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dnsmasq_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Dnsmasq",
			Execute:     common__dnsmasq.Execute,
		},
	},
	"doas": {
		{
			Name:        "doas",
			Description: "execute a command as another user",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/doas_completer/cmd",
			Variant:     "",
			Url:         "https://man.openbsd.org/doas",
			Execute:     common__doas.Execute,
		},
	},
	"docker": {
		{
			Name:        "docker",
			Description: "A self-sufficient runtime for containers",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/docker_completer/cmd",
			Variant:     "",
			Url:         "https://docs.docker.com/compose/",
			Execute:     common__docker.Execute,
		},
	},
	"docker-buildx": {
		{
			Name:        "docker-buildx",
			Description: "Docker Buildx",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/docker-buildx_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/docker/buildx",
			Execute:     common__docker_buildx.Execute,
		},
	},
	"docker-compose": {
		{
			Name:        "docker-compose",
			Description: "Docker Compose",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/docker-compose_completer/cmd",
			Variant:     "",
			Url:         "https://docs.docker.com/compose/",
			Execute:     common__docker_compose.Execute,
		},
	},
	"docker-scan": {
		{
			Name:        "docker-scan",
			Description: "A tool to scan your images",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/docker-scan_completer/cmd",
			Variant:     "",
			Url:         "https://docs.docker.com/engine/scan/",
			Execute:     common__docker_scan.Execute,
		},
	},
	"dockerd": {
		{
			Name:        "dockerd",
			Description: "A self-sufficient runtime for containers",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dockerd_completer/cmd",
			Variant:     "",
			Url:         "https://docs.docker.com/engine/reference/commandline/dockerd/",
			Execute:     common__dockerd.Execute,
		},
	},
	"doctl": {
		{
			Name:        "doctl",
			Description: "doctl is a command line interface (CLI) for the DigitalOcean API",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/doctl_completer/cmd",
			Variant:     "",
			Url:         "https://docs.digitalocean.com/reference/doctl/",
			Execute:     common__doctl.Execute,
		},
	},
	"doing": {
		{
			Name:        "doing",
			Description: "CLI for repository/issue workflow on Azure Devops",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/doing_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/ing-bank/doing-cli",
			Execute:     common__doing.Execute,
		},
	},
	"dos2unix": {
		{
			Name:        "dos2unix",
			Description: "DOS/Mac to Unix and vice versa text file format converter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dos2unix_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Unix2dos",
			Execute:     common__dos2unix.Execute,
		},
	},
	"downgrade": {
		{
			Name:        "downgrade",
			Description: "Downgrade Arch Linux packages",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/downgrade_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/pbrisbin/downgrade",
			Execute:     common__downgrade.Execute,
		},
	},
	"dpkg": {
		{
			Name:        "dpkg",
			Description: "package manager for Debian",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/dpkg_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/dpkg",
			Execute:     common__dpkg.Execute,
		},
	},
	"du": {
		{
			Name:        "du",
			Description: "estimate file space usage",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/du_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/du",
			Execute:     common__du.Execute,
		},
	},
	"ebook-convert": {
		{
			Name:        "ebook-convert",
			Description: "Convert an e-book from one format to another",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ebook-convert_completer/cmd",
			Variant:     "",
			Url:         "https://manual.calibre-ebook.com/de/generated/de/ebook-convert.html",
			Execute:     common__ebook_convert.Execute,
		},
	},
	"egrep": {
		{
			Name:        "egrep",
			Description: "print lines that match patterns",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/egrep_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/grep",
			Execute:     common__egrep.Execute,
		},
	},
	"electron": {
		{
			Name:        "electron",
			Description: "Build cross platform desktop apps with JavaScript, HTML, and CSS",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/electron_completer/cmd",
			Variant:     "",
			Url:         "https://www.electronjs.org/",
			Execute:     common__electron.Execute,
		},
	},
	"elvish": {
		{
			Name:        "elvish",
			Description: "expressive programming language and a versatile interactive shell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/elvish_completer/cmd",
			Variant:     "",
			Url:         "https://elv.sh/",
			Execute:     common__elvish.Execute,
		},
	},
	"env": {
		{
			Name:        "env",
			Description: "run a program in a modified environment",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/env_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/env",
			Execute:     common__env.Execute,
		},
	},
	"envsubst": {
		{
			Name:        "envsubst",
			Description: "Substitutes the values of environment variables",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/envsubst_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/envsubst",
			Execute:     common__envsubst.Execute,
		},
	},
	"exa": {
		{
			Name:        "exa",
			Description: "a modern replacement for ls",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/exa_completer/cmd",
			Variant:     "",
			Url:         "https://eza.rocks/",
			Execute:     common__exa.Execute,
		},
	},
	"expand": {
		{
			Name:        "expand",
			Description: "convert tabs to spaces",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/expand_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/expand",
			Execute:     common__expand.Execute,
		},
	},
	"expr": {
		{
			Name:        "expr",
			Description: "evaluate expressions",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/expr_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Expr",
			Execute:     common__expr.Execute,
		},
	},
	"eza": {
		{
			Name:        "eza",
			Description: "a modern replacement for ls",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/eza_completer/cmd",
			Variant:     "",
			Url:         "https://eza.rocks/",
			Execute:     common__eza.Execute,
		},
	},
	"faas-cli": {
		{
			Name:        "faas-cli",
			Description: "Manage your OpenFaaS functions from the command line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/faas-cli_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/openfaas/faas-cli",
			Execute:     common__faas_cli.Execute,
		},
	},
	"factor": {
		{
			Name:        "factor",
			Description: "factor numbers",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/factor_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/factor.1.html",
			Execute:     common__factor.Execute,
		},
	},
	"fakechroot": {
		{
			Name:        "fakechroot",
			Description: "gives a fake chroot environment",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fakechroot_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fakechroot",
			Execute:     common__fakechroot.Execute,
		},
	},
	"fakeroot": {
		{
			Name:        "fakeroot",
			Description: "run a command in an environment faking root privileges for file manipulation",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fakeroot_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fakeroot-sysv",
			Execute:     common__fakeroot.Execute,
		},
	},
	"fastfetch": {
		{
			Name:        "fastfetch",
			Description: "A neofetch-like tool for fetching system information and displaying them in a pretty way",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fastfetch_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/fastfetch-cli/fastfetch",
			Execute:     common__fastfetch.Execute,
		},
	},
	"fc-cache": {
		{
			Name:        "fc-cache",
			Description: "Build font information caches",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fc-cache_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fc-cache",
			Execute:     common__fc_cache.Execute,
		},
	},
	"fc-cat": {
		{
			Name:        "fc-cat",
			Description: "read font information cache files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fc-cat_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fc-cat",
			Execute:     common__fc_cat.Execute,
		},
	},
	"fc-conflist": {
		{
			Name:        "fc-conflist",
			Description: "list the configuration files processed by Fontconfig",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fc-conflist_completer/cmd",
			Variant:     "",
			Url:         "https://manpages.debian.org/testing/fontconfig/fc-conflist.1.en.html",
			Execute:     common__fc_conflist.Execute,
		},
	},
	"fc-list": {
		{
			Name:        "fc-list",
			Description: "list available fonts",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fc-list_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fc-list",
			Execute:     common__fc_list.Execute,
		},
	},
	"fd": {
		{
			Name:        "fd",
			Description: "find entries in the filesystem",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fd_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/sharkdp/fd",
			Execute:     common__fd.Execute,
		},
	},
	"fdisk": {
		{
			Name:        "fdisk",
			Description: "manipulate disk partition table",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fdisk_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Fdisk",
			Execute:     common__fdisk.Execute,
		},
	},
	"ffmpeg": {
		{
			Name:        "ffmpeg",
			Description: "Hyper fast Audio and Video encoder",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ffmpeg_completer/cmd",
			Variant:     "",
			Url:         "https://ffmpeg.org/",
			Execute:     common__ffmpeg.Execute,
		},
	},
	"fgrep": {
		{
			Name:        "fgrep",
			Description: "print lines that match patterns",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fgrep_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Grep",
			Execute:     common__fgrep.Execute,
		},
	},
	"file": {
		{
			Name:        "file",
			Description: "determine file type",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/file_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/file",
			Execute:     common__file.Execute,
		},
	},
	"fileicon": {
		{
			Name:        "fileicon",
			Description: "macOS CLI for managing custom icons for files and folders",
			Group:       "darwin",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/darwin/fileicon_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/mklement0/fileicon",
			Execute:     darwin__fileicon.Execute,
		},
	},
	"find": {
		{
			Name:        "find",
			Description: "search for files in a directory hierarchy",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/find_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/find",
			Execute:     common__find.Execute,
		},
	},
	"firefox": {
		{
			Name:        "firefox",
			Description: "Firefox Browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/firefox_completer/cmd",
			Variant:     "",
			Url:         "https://www.mozilla.org/en-US/firefox/new/",
			Execute:     common__firefox.Execute,
		},
	},
	"fish": {
		{
			Name:        "fish",
			Description: "the friendly interactive shell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fish_completer/cmd",
			Variant:     "",
			Url:         "https://fishshell.com/",
			Execute:     common__fish.Execute,
		},
	},
	"flatpak": {
		{
			Name:        "flatpak",
			Description: "Linux application sandboxing and distribution framework",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/flatpak_completer/cmd",
			Variant:     "",
			Url:         "https://flatpak.org/",
			Execute:     common__flatpak.Execute,
		},
	},
	"flutter": {
		{
			Name:        "flutter",
			Description: "Manage your Flutter app development",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/flutter_completer/cmd",
			Variant:     "",
			Url:         "https://flutter.dev/",
			Execute:     common__flutter.Execute,
		},
	},
	"fmt": {
		{
			Name:        "fmt",
			Description: "simple optimal text formatter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fmt_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fmt",
			Execute:     common__fmt.Execute,
		},
	},
	"fnm": {
		{
			Name:        "fnm",
			Description: "A fast and simple Node.js manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fnm_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/Schniz/fnm",
			Execute:     common__fnm.Execute,
		},
	},
	"fold": {
		{
			Name:        "fold",
			Description: "wrap each input line to fit in specified width",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fold_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/fold",
			Execute:     common__fold.Execute,
		},
	},
	"foot": {
		{
			Name:        "foot",
			Description: "A fast, lightweight and minimalistic Wayland terminal emulator",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/foot_completer/cmd",
			Variant:     "",
			Url:         "https://codeberg.org/dnkl/foot",
			Execute:     common__foot.Execute,
		},
	},
	"free": {
		{
			Name:        "free",
			Description: "Display amount of free and used memory in the system",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/free_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/free.1.html",
			Execute:     common__free.Execute,
		},
	},
	"freeze": {
		{
			Name:        "freeze",
			Description: "Generate images of code and terminal output",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/freeze_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/freeze",
			Execute:     common__freeze.Execute,
		},
	},
	"ftp": {
		{
			Name:        "ftp",
			Description: "File Transfer Protocol client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ftp_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ftp",
			Execute:     common__ftp.Execute,
		},
	},
	"ftpd": {
		{
			Name:        "ftpd",
			Description: "File Transfer Protocol daemon",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ftpd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/ftpd",
			Execute:     common__ftpd.Execute,
		},
	},
	"fury": {
		{
			Name:        "fury",
			Description: "Command line interface to Gemfury API",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fury_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/gemfury/cli",
			Execute:     common__fury.Execute,
		},
	},
	"fzf": {
		{
			Name:        "fzf",
			Description: "a command-line fuzzy finder",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/fzf_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/junegunn/fzf",
			Execute:     common__fzf.Execute,
		},
	},
	"gatsby": {
		{
			Name:        "gatsby",
			Description: "Build blazing fast, modern apps and websites with React",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gatsby_completer/cmd",
			Variant:     "",
			Url:         "https://www.gatsbyjs.com/",
			Execute:     common__gatsby.Execute,
		},
	},
	"gcloud": {
		{
			Name:        "gcloud",
			Description: "manage Google Cloud Platform resources and developer workflow",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gcloud_completer/cmd",
			Variant:     "",
			Url:         "https://cloud.google.com/sdk/gcloud/",
			Execute:     common__gcloud.Execute,
		},
	},
	"gdb": {
		{
			Name:        "gdb",
			Description: "This is the GNU debugger",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gdb_completer/cmd",
			Variant:     "",
			Url:         "https://www.sourceware.org/gdb/",
			Execute:     common__gdb.Execute,
		},
	},
	"gdown": {
		{
			Name:        "gdown",
			Description: "Google Drive Public File Downloader when Curl/Wget Fails",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gdown_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/wkentaro/gdown",
			Execute:     common__gdown.Execute,
		},
	},
	"gdu": {
		{
			Name:        "gdu",
			Description: "Pretty fast disk usage analyzer written in Go",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gdu_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/dundee/gdu",
			Execute:     common__gdu.Execute,
		},
	},
	"get-env": {
		{
			Name:        "get-env",
			Description: "get environment variable",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/get-env_completer/cmd",
			Variant:     "",
			Url:         "https://carapace-sh.github.io/carapace-bin/environment.html",
			Execute:     common__get_env.Execute,
		},
	},
	"gftp": {
		{
			Name:        "gftp",
			Description: "file transfer client for *NIX based machines",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gftp_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/masneyb/gftp",
			Execute:     common__gftp.Execute,
		},
	},
	"gh": {
		{
			Name:        "gh",
			Description: "GitHub CLI",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gh_completer/cmd",
			Variant:     "",
			Url:         "https://cli.github.com/",
			Execute:     common__gh.Execute,
		},
	},
	"gh-copilot": {
		{
			Name:        "gh-copilot",
			Description: "Your AI command line copilot",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gh-copilot_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/github/gh-copilot",
			Execute:     common__gh_copilot.Execute,
		},
	},
	"gh-dash": {
		{
			Name:        "gh-dash",
			Description: "A beautiful CLI dashboard for GitHub",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gh-dash_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/dlvhdr/gh-dash",
			Execute:     common__gh_dash.Execute,
		},
	},
	"ghostty": {
		{
			Name:        "ghostty",
			Description: "a fast, feature-rich, and cross-platform terminal emulator ",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ghostty_completer/cmd",
			Variant:     "",
			Url:         "https://ghostty.org/",
			Execute:     common__ghostty.Execute,
		},
	},
	"gimp": {
		{
			Name:        "gimp",
			Description: "an image manipulation and paint program",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gimp_completer/cmd",
			Variant:     "",
			Url:         "https://www.gimp.org/",
			Execute:     common__gimp.Execute,
		},
	},
	"git": {
		{
			Name:        "git",
			Description: "the stupid content tracker",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git_completer/cmd",
			Variant:     "",
			Url:         "https://git-scm.com/",
			Execute:     common__git.Execute,
		},
	},
	"git-abort": {
		{
			Name:        "git-abort",
			Description: "Abort current rebase, merge or cherry-pick, without the need to find exact command in history",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-abort_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-abort",
			Execute:     common__git_abort.Execute,
		},
	},
	"git-alias": {
		{
			Name:        "git-alias",
			Description: "Define, search and show aliases",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-alias_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-alias",
			Execute:     common__git_alias.Execute,
		},
	},
	"git-archive-file": {
		{
			Name:        "git-archive-file",
			Description: "Export the current HEAD of the git repository to an archive",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-archive-file_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-archive-file",
			Execute:     common__git_archive_file.Execute,
		},
	},
	"git-authors": {
		{
			Name:        "git-authors",
			Description: "Generate authors report",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-authors_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-authors",
			Execute:     common__git_authors.Execute,
		},
	},
	"git-browse": {
		{
			Name:        "git-browse",
			Description: "Opens  the  current  git repository website in your default web browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-browse_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-browse",
			Execute:     common__git_browse.Execute,
		},
	},
	"git-browse-ci": {
		{
			Name:        "git-browse-ci",
			Description: "Opens the current git repository CI page in your default web browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-browse-ci_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-browse-ci",
			Execute:     common__git_browse_ci.Execute,
		},
	},
	"git-clang-format": {
		{
			Name:        "git-clang-format",
			Description: "run clang-format on lines that differ",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-clang-format_completer/cmd",
			Variant:     "",
			Url:         "https://clang.llvm.org/docs/ClangFormat.html",
			Execute:     common__git_clang_format.Execute,
		},
	},
	"git-clear": {
		{
			Name:        "git-clear",
			Description: "Rigorously clean up a repository",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-clear_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-clear",
			Execute:     common__git_clear.Execute,
		},
	},
	"git-clear-soft": {
		{
			Name:        "git-clear-soft",
			Description: "Soft clean up a repository",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-clear-soft_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-clear-soft",
			Execute:     common__git_clear_soft.Execute,
		},
	},
	"git-coauthor": {
		{
			Name:        "git-coauthor",
			Description: "Add a co-author to the last commit",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-coauthor_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-coauthor",
			Execute:     common__git_coauthor.Execute,
		},
	},
	"git-extras": {
		{
			Name:        "git-extras",
			Description: "Awesome GIT utilities",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-extras_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras",
			Execute:     common__git_extras.Execute,
		},
	},
	"git-info": {
		{
			Name:        "git-info",
			Description: "Returns information on current repository",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-info_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-info",
			Execute:     common__git_info.Execute,
		},
	},
	"git-standup": {
		{
			Name:        "git-standup",
			Description: "Recall the commit history",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-standup_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-standup",
			Execute:     common__git_standup.Execute,
		},
	},
	"git-unlock": {
		{
			Name:        "git-unlock",
			Description: "Unlock a file excluded from version control",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-unlock_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-unlock",
			Execute:     common__git_unlock.Execute,
		},
	},
	"git-utimes": {
		{
			Name:        "git-utimes",
			Description: "only touch files that are newer than their last commit date",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/git-utimes_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tj/git-extras/blob/master/Commands.md#git-utimes",
			Execute:     common__git_utimes.Execute,
		},
	},
	"gitk": {
		{
			Name:        "gitk",
			Description: "The Git repository browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gitk_completer/cmd",
			Variant:     "",
			Url:         "https://git-scm.com/docs/gitk",
			Execute:     common__gitk.Execute,
		},
	},
	"gitui": {
		{
			Name:        "gitui",
			Description: "blazing fast terminal-ui for git",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gitui_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/extrawurst/gitui",
			Execute:     common__gitui.Execute,
		},
	},
	"glab": {
		{
			Name:        "glab",
			Description: "A GitLab CLI tool.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/glab_completer/cmd",
			Variant:     "",
			Url:         "https://gitlab.com/gitlab-org/cli",
			Execute:     common__glab.Execute,
		},
	},
	"glow": {
		{
			Name:        "glow",
			Description: "Render markdown on the CLI, with pizzazz!",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/glow_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/glow",
			Execute:     common__glow.Execute,
		},
	},
	"gm": {
		{
			Name:        "gm",
			Description: "command-line utility to create, edit, compare, convert, or display images",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gm_completer/cmd",
			Variant:     "",
			Url:         "http://www.GraphicsMagick.org/",
			Execute:     common__gm.Execute,
		},
	},
	"gnome-keyring": {
		{
			Name:        "gnome-keyring",
			Description: "The gnome-keyring commandline tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-keyring_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.gnome.org/Projects/GnomeKeyring/",
			Execute:     common__gnome_keyring.Execute,
		},
	},
	"gnome-keyring-daemon": {
		{
			Name:        "gnome-keyring-daemon",
			Description: "The Gnome Keyring Daemon",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-keyring-daemon_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.gnome.org/Projects/GnomeKeyring/",
			Execute:     common__gnome_keyring_daemon.Execute,
		},
	},
	"gnome-maps": {
		{
			Name:        "gnome-maps",
			Description: "A map application for GNOME",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-maps_completer/cmd",
			Variant:     "",
			Url:         "https://apps.gnome.org/app/org.gnome.Maps/",
			Execute:     common__gnome_maps.Execute,
		},
	},
	"gnome-terminal": {
		{
			Name:        "gnome-terminal",
			Description: "A terminal emulator for GNOME",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gnome-terminal_completer/cmd",
			Variant:     "",
			Url:         "https://help.gnome.org/users/gnome-terminal/stable/",
			Execute:     common__gnome_terminal.Execute,
		},
	},
	"go": {
		{
			Name:        "go",
			Description: "Go is a tool for managing Go source code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go_completer/cmd",
			Variant:     "",
			Url:         "https://golang.org/",
			Execute:     common__go.Execute,
		},
	},
	"go-carpet": {
		{
			Name:        "go-carpet",
			Description: "show test coverage for Go source files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-carpet_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/msoap/go-carpet",
			Execute:     common__go_carpet.Execute,
		},
	},
	"go-tool-asm": {
		{
			Name:        "go-tool-asm",
			Description: "go assembler",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-asm_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/asm",
			Execute:     common__go_tool_asm.Execute,
		},
	},
	"go-tool-buildid": {
		{
			Name:        "go-tool-buildid",
			Description: "Buildid displays or updates the build ID stored in a Go package or binary.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-buildid_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/buildid",
			Execute:     common__go_tool_buildid.Execute,
		},
	},
	"go-tool-cgo": {
		{
			Name:        "go-tool-cgo",
			Description: "Cgo enables the creation of Go packages that call C code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-cgo_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/cgo",
			Execute:     common__go_tool_cgo.Execute,
		},
	},
	"go-tool-compile": {
		{
			Name:        "go-tool-compile",
			Description: "compiles a single Go packag",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-compile_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/compile",
			Execute:     common__go_tool_compile.Execute,
		},
	},
	"go-tool-covdata": {
		{
			Name:        "go-tool-covdata",
			Description: "read and manipulate coverage data files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-covdata_completer/cmd",
			Variant:     "",
			Url:         "https://go.dev/testing/coverage/",
			Execute:     common__go_tool_covdata.Execute,
		},
	},
	"go-tool-cover": {
		{
			Name:        "go-tool-cover",
			Description: "analyze coverage profiles",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-cover_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/cover",
			Execute:     common__go_tool_cover.Execute,
		},
	},
	"go-tool-dist": {
		{
			Name:        "go-tool-dist",
			Description: "Dist helps bootstrap, build, and test the Go distribution",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-dist_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/dist",
			Execute:     common__go_tool_dist.Execute,
		},
	},
	"go-tool-doc": {
		{
			Name:        "go-tool-doc",
			Description: "show documentation for package or symbol",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-doc_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/doc",
			Execute:     common__go_tool_doc.Execute,
		},
	},
	"go-tool-fix": {
		{
			Name:        "go-tool-fix",
			Description: "Fix finds Go programs that use old APIs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-fix_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/fix",
			Execute:     common__go_tool_fix.Execute,
		},
	},
	"go-tool-link": {
		{
			Name:        "go-tool-link",
			Description: "go linker",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-link_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/link",
			Execute:     common__go_tool_link.Execute,
		},
	},
	"go-tool-mockgen": {
		{
			Name:        "go-tool-mockgen",
			Description: "mock interfaces",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-mockgen_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/uber-go/mock",
			Execute:     common__go_tool_mockgen.Execute,
		},
	},
	"go-tool-nm": {
		{
			Name:        "go-tool-nm",
			Description: "Nm lists the symbols defined or used by an object file, archive, or executabl",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-nm_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/nm",
			Execute:     common__go_tool_nm.Execute,
		},
	},
	"go-tool-objdump": {
		{
			Name:        "go-tool-objdump",
			Description: "Objdump disassembles executable files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-objdump_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/objdump",
			Execute:     common__go_tool_objdump.Execute,
		},
	},
	"go-tool-pack": {
		{
			Name:        "go-tool-pack",
			Description: "Pack is a simple version of the traditional Unix ar tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/go-tool-pack_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/pack",
			Execute:     common__go_tool_pack.Execute,
		},
	},
	"gocyclo": {
		{
			Name:        "gocyclo",
			Description: "Calculate cyclomatic complexities of Go functions",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gocyclo_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/fzipp/gocyclo",
			Execute:     common__gocyclo.Execute,
		},
	},
	"gofmt": {
		{
			Name:        "gofmt",
			Description: "format Go source code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gofmt_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/cmd/gofmt",
			Execute:     common__gofmt.Execute,
		},
	},
	"goimports": {
		{
			Name:        "goimports",
			Description: "updates your Go import lines",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/goimports_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/golang.org/x/tools/cmd/goimports",
			Execute:     common__goimports.Execute,
		},
	},
	"golangci-lint": {
		{
			Name:        "golangci-lint",
			Description: "golangci-lint is a smart linters runner.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/golangci-lint_completer/cmd",
			Variant:     "",
			Url:         "https://golangci-lint.run/",
			Execute:     common__golangci_lint.Execute,
		},
	},
	"gonew": {
		{
			Name:        "gonew",
			Description: "Gonew starts a new Go module by copying a template module",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gonew_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/golang.org/x/tools/cmd/gonew",
			Execute:     common__gonew.Execute,
		},
	},
	"google-chrome": {
		{
			Name:        "google-chrome",
			Description: "chrome browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/google-chrome_completer/cmd",
			Variant:     "",
			Url:         "https://www.google.com/chrome/index.html",
			Execute:     common__google_chrome.Execute,
		},
	},
	"gopls": {
		{
			Name:        "gopls",
			Description: "gopls is a Go language server",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gopls_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/golang.org/x/tools/gopls",
			Execute:     common__gopls.Execute,
		},
	},
	"goreleaser": {
		{
			Name:        "goreleaser",
			Description: "Release engineering, simplified",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/goreleaser_completer/cmd",
			Variant:     "",
			Url:         "https://goreleaser.com/",
			Execute:     common__goreleaser.Execute,
		},
	},
	"goweight": {
		{
			Name:        "goweight",
			Description: "A tool to analyze and troubleshoot a Go binary size",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/goweight_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/jondot/goweight",
			Execute:     common__goweight.Execute,
		},
	},
	"gparted": {
		{
			Name:        "gparted",
			Description: "GNOME Partition Editor for manipulating disk partitions",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gparted_completer/cmd",
			Variant:     "",
			Url:         "https://gparted.org/",
			Execute:     common__gparted.Execute,
		},
	},
	"gpasswd": {
		{
			Name:        "gpasswd",
			Description: "administer /etc/group and /etc/gshadow",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gpasswd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/gpasswd",
			Execute:     common__gpasswd.Execute,
		},
	},
	"gpg": {
		{
			Name:        "gpg",
			Description: "OpenPGP encryption and signing tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gpg_completer/cmd",
			Variant:     "",
			Url:         "https://gnupg.org/",
			Execute:     common__gpg.Execute,
		},
	},
	"gpg-agent": {
		{
			Name:        "gpg-agent",
			Description: "Secret key management for GnuPG",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gpg-agent_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/gpg-agent",
			Execute:     common__gpg_agent.Execute,
		},
	},
	"gradle": {
		{
			Name:        "gradle",
			Description: "Gradle Build Tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gradle_completer/cmd",
			Variant:     "",
			Url:         "https://gradle.org/",
			Execute:     common__gradle.Execute,
		},
	},
	"grep": {
		{
			Name:        "grep",
			Description: "print lines that match patterns",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/grep_completer/cmd",
			Variant:     "",
			Url:         "https://www.gnu.org/software/grep/",
			Execute:     common__grep.Execute,
		},
	},
	"groupadd": {
		{
			Name:        "groupadd",
			Description: "create a new group",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/groupadd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/groupadd",
			Execute:     common__groupadd.Execute,
		},
	},
	"groupdel": {
		{
			Name:        "groupdel",
			Description: "delete a group",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/groupdel_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/groupdel",
			Execute:     common__groupdel.Execute,
		},
	},
	"groupmems": {
		{
			Name:        "groupmems",
			Description: "administer members of a user's primary group",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/groupmems_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/groupmems",
			Execute:     common__groupmems.Execute,
		},
	},
	"groupmod": {
		{
			Name:        "groupmod",
			Description: "modify a group definition on the system",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/groupmod_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/groupmod",
			Execute:     common__groupmod.Execute,
		},
	},
	"groups": {
		{
			Name:        "groups",
			Description: "display current group names",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/groups_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/groups",
			Execute:     common__groups.Execute,
		},
	},
	"grype": {
		{
			Name:        "grype",
			Description: "A vulnerability scanner for container images, filesystems, and SBOMs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/grype_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/anchore/grype",
			Execute:     common__grype.Execute,
		},
	},
	"gsa": {
		{
			Name:        "gsa",
			Description: "A tool for analyzing the size of compiled Go binaries",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gsa_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/Zxilly/go-size-analyzer",
			Execute:     common__gsa.Execute,
		},
	},
	"gulp": {
		{
			Name:        "gulp",
			Description: "Command Line Interface for gulp",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gulp_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/gulpjs/gulp-cli",
			Execute:     common__gulp.Execute,
		},
	},
	"gum": {
		{
			Name:        "gum",
			Description: "A tool for glamorous shell scripts",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gum_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/gum",
			Execute:     common__gum.Execute,
		},
	},
	"gunzip": {
		{
			Name:        "gunzip",
			Description: "Uncompress files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gunzip_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/gunzip",
			Execute:     common__gunzip.Execute,
		},
	},
	"gzip": {
		{
			Name:        "gzip",
			Description: "Compress or uncompress files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/gzip_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/gzip",
			Execute:     common__gzip.Execute,
		},
	},
	"halt": {
		{
			Name:        "halt",
			Description: "halt the machine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/halt_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/halt.html",
			Execute:     common__halt.Execute,
		},
	},
	"head": {
		{
			Name:        "head",
			Description: "output the first part of files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/head_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/head",
			Execute:     common__head.Execute,
		},
	},
	"helix": {
		{
			Name:        "helix",
			Description: "A post-modern text editor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/helix_completer/cmd",
			Variant:     "",
			Url:         "https://helix-editor.com/",
			Execute:     common__helix.Execute,
		},
	},
	"helm": {
		{
			Name:        "helm",
			Description: "The Helm package manager for Kubernetes.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/helm_completer/cmd",
			Variant:     "",
			Url:         "https://helm.sh/",
			Execute:     common__helm.Execute,
		},
	},
	"helmsman": {
		{
			Name:        "helmsman",
			Description: "Helmsman is a Helm Charts as Code tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/helmsman_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/Praqma/Helmsman",
			Execute:     common__helmsman.Execute,
		},
	},
	"hexchat": {
		{
			Name:        "hexchat",
			Description: "IRC Client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hexchat_completer/cmd",
			Variant:     "",
			Url:         "https://hexchat.github.io/",
			Execute:     common__hexchat.Execute,
		},
	},
	"hexdump": {
		{
			Name:        "hexdump",
			Description: "Display file contents in hexadecimal, decimal, octal, or ascii",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hexdump_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/hexdump",
			Execute:     common__hexdump.Execute,
		},
	},
	"hostid": {
		{
			Name:        "hostid",
			Description: "print the numeric identifier for the current host",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hostid_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/hostid",
			Execute:     common__hostid.Execute,
		},
	},
	"hostname": {
		{
			Name:        "hostname",
			Description: "show or set system host name",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hostname_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/hostname",
			Execute:     common__hostname.Execute,
		},
	},
	"htop": {
		{
			Name:        "htop",
			Description: "interactive process viewer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/htop_completer/cmd",
			Variant:     "",
			Url:         "https://htop.dev/",
			Execute:     common__htop.Execute,
		},
	},
	"http": {
		{
			Name:        "http",
			Description: "command-line HTTP client for the API era",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/http_completer/cmd",
			Variant:     "",
			Url:         "https://httpie.io/",
			Execute:     common__http.Execute,
		},
	},
	"https": {
		{
			Name:        "https",
			Description: "command-line HTTP client for the API era",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/https_completer/cmd",
			Variant:     "",
			Url:         "https://httpie.io/",
			Execute:     common__https.Execute,
		},
	},
	"hugetop": {
		{
			Name:        "hugetop",
			Description: "report huge page information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hugetop_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/hugetop.1.html",
			Execute:     common__hugetop.Execute,
		},
	},
	"hugo": {
		{
			Name:        "hugo",
			Description: "hugo builds your site",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hugo_completer/cmd",
			Variant:     "",
			Url:         "https://gohugo.io/",
			Execute:     common__hugo.Execute,
		},
	},
	"hurl": {
		{
			Name:        "hurl",
			Description: "run and test HTTP requests with plain text",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hurl_completer/cmd",
			Variant:     "",
			Url:         "https://hurl.dev/",
			Execute:     common__hurl.Execute,
		},
	},
	"hwinfo": {
		{
			Name:        "hwinfo",
			Description: "Probe for hardware",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hwinfo_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/openSUSE/hwinfo",
			Execute:     common__hwinfo.Execute,
		},
	},
	"hx": {
		{
			Name:        "hx",
			Description: "A post-modern text editor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/hx_completer/cmd",
			Variant:     "",
			Url:         "https://helix-editor.com/",
			Execute:     common__hx.Execute,
		},
	},
	"i3": {
		{
			Name:        "i3",
			Description: "an improved dynamic, tiling window manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/i3_completer/cmd",
			Variant:     "",
			Url:         "https://i3wm.org/",
			Execute:     common__i3.Execute,
		},
	},
	"i3-scrot": {
		{
			Name:        "i3-scrot",
			Description: "simple screenshot script",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/i3-scrot_completer/cmd",
			Variant:     "",
			Url:         "https://gitlab.manjaro.org/packages/community/i3/i3-scrot",
			Execute:     common__i3_scrot.Execute,
		},
	},
	"i3exit": {
		{
			Name:        "i3exit",
			Description: "exit-script for i3",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/i3exit_completer/cmd",
			Variant:     "",
			Url:         "https://aur.archlinux.org/packages/i3exit/",
			Execute:     common__i3exit.Execute,
		},
	},
	"i3lock": {
		{
			Name:        "i3lock",
			Description: "improved screen locker",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/i3lock_completer/cmd",
			Variant:     "",
			Url:         "https://i3wm.org/i3lock/",
			Execute:     common__i3lock.Execute,
		},
	},
	"i3status": {
		{
			Name:        "i3status",
			Description: "Generates a status line for i3bar, dzen2, xmobar or lemonbar",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/i3status_completer/cmd",
			Variant:     "",
			Url:         "https://i3wm.org/i3status/",
			Execute:     common__i3status.Execute,
		},
	},
	"i3status-rs": {
		{
			Name:        "i3status-rs",
			Description: "A feature-rich and resource-friendly replacement for i3status, written in Rust",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/i3status-rs_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/greshake/i3status-rust",
			Execute:     common__i3status_rs.Execute,
		},
	},
	"id": {
		{
			Name:        "id",
			Description: "Print user and group information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/id_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/id",
			Execute:     common__id.Execute,
		},
	},
	"imv": {
		{
			Name:        "imv",
			Description: "Image viewer for X11 and Wayland",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/imv_completer/cmd",
			Variant:     "",
			Url:         "https://git.sr.ht/~exec64/imv/",
			Execute:     common__imv.Execute,
		},
	},
	"inkscape": {
		{
			Name:        "inkscape",
			Description: "an SVG (Scalable Vector Graphics) editing program",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/inkscape_completer/cmd",
			Variant:     "",
			Url:         "https://inkscape.org/",
			Execute:     common__inkscape.Execute,
		},
	},
	"inshellisense": {
		{
			Name:        "inshellisense",
			Description: "IDE style command line auto complete",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/inshellisense_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/microsoft/inshellisense",
			Execute:     common__inshellisense.Execute,
		},
	},
	"install": {
		{
			Name:        "install",
			Description: "copy files and set attributes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/install_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/install",
			Execute:     common__install.Execute,
		},
	},
	"ion": {
		{
			Name:        "ion",
			Description: "The Ion Shell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ion_completer/cmd",
			Variant:     "",
			Url:         "https://gitlab.redox-os.org/redox-os/ion/",
			Execute:     common__ion.Execute,
		},
	},
	"jar": {
		{
			Name:        "jar",
			Description: "create an archive for classes and resources",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/jar_completer/cmd",
			Variant:     "",
			Url:         "https://docs.oracle.com/en/java/javase/16/docs/specs/man/jar.html",
			Execute:     common__jar.Execute,
		},
	},
	"java": {
		{
			Name:        "java",
			Description: "Launches a Java application",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/java_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Java_(programming_language)",
			Execute:     common__java.Execute,
		},
	},
	"javac": {
		{
			Name:        "javac",
			Description: "Reads Java class and interface definitions and compiles them into bytecode and class files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/javac_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Javac",
			Execute:     common__javac.Execute,
		},
	},
	"jj": {
		{
			Name:        "jj",
			Description: "Jujutsu (An experimental VCS)",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/jj_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/martinvonz/jj",
			Execute:     common__jj.Execute,
		},
	},
	"join": {
		{
			Name:        "join",
			Description: "join lines of two files on a common field",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/join_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/join",
			Execute:     common__join.Execute,
		},
	},
	"journalctl": {
		{
			Name:        "journalctl",
			Description: "Query the journal",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/journalctl_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/journalctl.1.html",
			Execute:     common__journalctl.Execute,
		},
	},
	"jq": {
		{
			Name:        "jq",
			Description: "Command-line JSON processor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/jq_completer/cmd",
			Variant:     "",
			Url:         "https://stedolan.github.io/jq/",
			Execute:     common__jq.Execute,
		},
	},
	"julia": {
		{
			Name:        "julia",
			Description: "high-level, high-performance dynamic programming language for technical computing",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/julia_completer/cmd",
			Variant:     "",
			Url:         "https://julialang.org/",
			Execute:     common__julia.Execute,
		},
	},
	"just": {
		{
			Name:        "just",
			Description: "Just a command runner",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/just_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/casey/just",
			Execute:     common__just.Execute,
		},
	},
	"kak": {
		{
			Name:        "kak",
			Description: "a vim-inspired, selection oriented code editor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kak_completer/cmd",
			Variant:     "",
			Url:         "http://kakoune.org/",
			Execute:     common__kak.Execute,
		},
	},
	"kak-lsp": {
		{
			Name:        "kak-lsp",
			Description: "Kakoune Language Server Protocol Client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kak-lsp_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/kak-lsp/kak-lsp",
			Execute:     common__kak_lsp.Execute,
		},
	},
	"kill": {
		{
			Name:        "kill",
			Description: "Forcibly terminate a process",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kill_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/kill",
			Execute:     common__kill.Execute,
		},
	},
	"killall": {
		{
			Name:        "killall",
			Description: "kill processes by name",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/killall_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/killall",
			Execute:     common__killall.Execute,
		},
	},
	"kitten": {
		{
			Name:        "kitten",
			Description: "Fast, statically compiled implementations of various kittens",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kitten_completer/cmd",
			Variant:     "",
			Url:         "https://sw.kovidgoyal.net/kitty/",
			Execute:     common__kitten.Execute,
		},
	},
	"kitty": {
		{
			Name:        "kitty",
			Description: "The fast, feature rich terminal emulator",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kitty_completer/cmd",
			Variant:     "",
			Url:         "https://sw.kovidgoyal.net/kitty/",
			Execute:     common__kitty.Execute,
		},
	},
	"kmonad": {
		{
			Name:        "kmonad",
			Description: "an onion of buttons",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kmonad_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/kmonad/kmonad",
			Execute:     common__kmonad.Execute,
		},
	},
	"kompose": {
		{
			Name:        "kompose",
			Description: "A tool helping Docker Compose users move to Kubernetes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kompose_completer/cmd",
			Variant:     "",
			Url:         "https://kompose.io/",
			Execute:     common__kompose.Execute,
		},
	},
	"kotlin": {
		{
			Name:        "kotlin",
			Description: "run Kotlin programs, scripts or REPL",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kotlin_completer/cmd",
			Variant:     "",
			Url:         "https://kotlinlang.org/docs/command-line.html",
			Execute:     common__kotlin.Execute,
		},
	},
	"kotlinc": {
		{
			Name:        "kotlinc",
			Description: "Kotlin command-line compiler",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kotlinc_completer/cmd",
			Variant:     "",
			Url:         "https://kotlinlang.org/docs/command-line.html#manual-install",
			Execute:     common__kotlinc.Execute,
		},
	},
	"ktlint": {
		{
			Name:        "ktlint",
			Description: "An anti-bikeshedding Kotlin linter with built-in formatter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ktlint_completer/cmd",
			Variant:     "",
			Url:         "https://ktlint.github.io/",
			Execute:     common__ktlint.Execute,
		},
	},
	"kubeadm": {
		{
			Name:        "kubeadm",
			Description: "kubeadm: easily bootstrap a secure Kubernetes cluster",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kubeadm_completer/cmd",
			Variant:     "",
			Url:         "https://kubernetes.io/docs/reference/setup-tools/kubeadm/",
			Execute:     common__kubeadm.Execute,
		},
	},
	"kubectl": {
		{
			Name:        "kubectl",
			Description: "kubectl controls the Kubernetes cluster manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kubectl_completer/cmd",
			Variant:     "",
			Url:         "https://kubernetes.io/docs/reference/kubectl/overview/",
			Execute:     common__kubectl.Execute,
		},
	},
	"kubeseal": {
		{
			Name:        "kubeseal",
			Description: "A Kubernetes controller and tool for one-way encrypted Secrets",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/kubeseal_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/bitnami-labs/sealed-secrets/",
			Execute:     common__kubeseal.Execute,
		},
	},
	"last": {
		{
			Name:        "last",
			Description: "Show a listing of last logged in users",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/last_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/last",
			Execute:     common__last.Execute,
		},
	},
	"lastb": {
		{
			Name:        "lastb",
			Description: "Show a listing of last logged in users",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lastb_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/lastb",
			Execute:     common__lastb.Execute,
		},
	},
	"lastlog": {
		{
			Name:        "lastlog",
			Description: "reports the most recent login of all users or of a given user",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lastlog_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/lastlog",
			Execute:     common__lastlog.Execute,
		},
	},
	"lazygit": {
		{
			Name:        "lazygit",
			Description: "simple terminal UI for git commands",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lazygit_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/jesseduffield/lazygit",
			Execute:     common__lazygit.Execute,
		},
	},
	"lf": {
		{
			Name:        "lf",
			Description: "terminal file manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lf_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/gokcehan/lf",
			Execute:     common__lf.Execute,
		},
	},
	"light": {
		{
			Name:        "light",
			Description: "a program to control backlight controllers",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/light_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/haikarainen/light",
			Execute:     common__light.Execute,
		},
	},
	"lightdm": {
		{
			Name:        "lightdm",
			Description: "a display manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lightdm_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/canonical/lightdm",
			Execute:     common__lightdm.Execute,
		},
	},
	"link": {
		{
			Name:        "link",
			Description: "call the link function to create a link to a file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/link_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/3/link",
			Execute:     common__link.Execute,
		},
	},
	"ln": {
		{
			Name:        "ln",
			Description: "make links between files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ln_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ln",
			Execute:     common__ln.Execute,
		},
	},
	"lnav": {
		{
			Name:        "lnav",
			Description: "ncurses-based log file viewer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lnav_completer/cmd",
			Variant:     "",
			Url:         "https://lnav.org/",
			Execute:     common__lnav.Execute,
		},
	},
	"lncrawl": {
		{
			Name:        "lncrawl",
			Description: "Generate and download e-books from online sources",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lncrawl_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/dipu-bd/lightnovel-crawler",
			Execute:     common__lncrawl.Execute,
		},
	},
	"locale": {
		{
			Name:        "locale",
			Description: "Get locale-specific information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/locale_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/locale",
			Execute:     common__locale.Execute,
		},
	},
	"localectl": {
		{
			Name:        "localectl",
			Description: "Query or change system locale and keyboard settings",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/localectl_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/localectl.1.html",
			Execute:     common__localectl.Execute,
		},
	},
	"logname": {
		{
			Name:        "logname",
			Description: "print user's login name",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/logname_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/logname",
			Execute:     common__logname.Execute,
		},
	},
	"ls": {
		{
			Name:        "ls",
			Description: "list directory contents",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ls_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ls",
			Execute:     common__ls.Execute,
		},
	},
	"lsb_release": {
		{
			Name:        "lsb_release",
			Description: "prints certain LSB (Linux Standard Base) and Distribution information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsb_release_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/lsb_release",
			Execute:     common__lsb_release.Execute,
		},
	},
	"lsblk": {
		{
			Name:        "lsblk",
			Description: "list block devices",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsblk_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/lsblk.8.html",
			Execute:     common__lsblk.Execute,
		},
	},
	"lsclocks": {
		{
			Name:        "lsclocks",
			Description: "display system clocks",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsclocks_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/lsclocks.1.html",
			Execute:     common__lsclocks.Execute,
		},
	},
	"lscpu": {
		{
			Name:        "lscpu",
			Description: "display information about the CPU architecture",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lscpu_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/lscpu.1.html",
			Execute:     common__lscpu.Execute,
		},
	},
	"lsfd": {
		{
			Name:        "lsfd",
			Description: "list file descriptors",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsfd_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/lsfd.1.html",
			Execute:     common__lsfd.Execute,
		},
	},
	"lsirq": {
		{
			Name:        "lsirq",
			Description: "utility to display kernel interrupt information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsirq_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/lsirq.1.html",
			Execute:     common__lsirq.Execute,
		},
	},
	"lslocks": {
		{
			Name:        "lslocks",
			Description: "List local system locks",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lslocks_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/lslocks.8.html",
			Execute:     common__lslocks.Execute,
		},
	},
	"lslogins": {
		{
			Name:        "lslogins",
			Description: "Display information about known users in the system",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lslogins_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/lslogins.1.html",
			Execute:     common__lslogins.Execute,
		},
	},
	"lsmem": {
		{
			Name:        "lsmem",
			Description: "list the ranges of available memory with their online status",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsmem_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/lsmem.1.html",
			Execute:     common__lsmem.Execute,
		},
	},
	"lsns": {
		{
			Name:        "lsns",
			Description: "List system namespaces",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsns_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/lsns.8.html",
			Execute:     common__lsns.Execute,
		},
	},
	"lsusb": {
		{
			Name:        "lsusb",
			Description: "list USB devices",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lsusb_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/lsusb",
			Execute:     common__lsusb.Execute,
		},
	},
	"lua": {
		{
			Name:        "lua",
			Description: "Lua interpreter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lua_completer/cmd",
			Variant:     "",
			Url:         "https://www.lua.org/",
			Execute:     common__lua.Execute,
		},
	},
	"lzcat": {
		{
			Name:        "lzcat",
			Description: "Compress or decompress .xz and .lzma files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lzcat_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xz",
			Execute:     common__lzcat.Execute,
		},
	},
	"lzma": {
		{
			Name:        "lzma",
			Description: "Compress or decompress .xz and .lzma files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/lzma_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xz",
			Execute:     common__lzma.Execute,
		},
	},
	"magick": {
		{
			Name:        "magick",
			Description: "convert between image formats as well as resize an image, blur, crop, despeckle, dither, draw on, flip, join, re-sample, and much more",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/magick_completer/cmd",
			Variant:     "",
			Url:         "https://imagemagick.org/",
			Execute:     common__magick.Execute,
		},
	},
	"make": {
		{
			Name:        "make",
			Description: "GNU make utility to maintain groups of programs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/make_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/make",
			Execute:     common__make.Execute,
		},
	},
	"makepkg": {
		{
			Name:        "makepkg",
			Description: "make packages compatible for use with pacman",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/makepkg_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.archlinux.org/title/Makepkg",
			Execute:     common__makepkg.Execute,
		},
	},
	"man": {
		{
			Name:        "man",
			Description: "an interface to the system reference manuals",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/man_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/man.1.html",
			Execute:     common__man.Execute,
		},
	},
	"marp": {
		{
			Name:        "marp",
			Description: "A CLI interface for Marp and Marpit based converters",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/marp_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/marp-team/marp-cli",
			Execute:     common__marp.Execute,
		},
	},
	"mcomix": {
		{
			Name:        "mcomix",
			Description: "GTK Comic Book Viewer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mcomix_completer/cmd",
			Variant:     "",
			Url:         "https://sourceforge.net/projects/mcomix/",
			Execute:     common__mcomix.Execute,
		},
	},
	"md5sum": {
		{
			Name:        "md5sum",
			Description: "compute and check MD5 message digest",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/md5sum_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/md5sum",
			Execute:     common__md5sum.Execute,
		},
	},
	"mdbook": {
		{
			Name:        "mdbook",
			Description: "Creates a book from markdown files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mdbook_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/rust-lang/mdBook",
			Execute:     common__mdbook.Execute,
		},
	},
	"meld": {
		{
			Name:        "meld",
			Description: "Meld is a file and directory comparison tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/meld_completer/cmd",
			Variant:     "",
			Url:         "https://meldmerge.org/",
			Execute:     common__meld.Execute,
		},
	},
	"melt": {
		{
			Name:        "melt",
			Description: "melt generates a seed phrase from an SSH key",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/melt_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/melt",
			Execute:     common__melt.Execute,
		},
	},
	"micro": {
		{
			Name:        "micro",
			Description: "A modern and intuitive terminal-based text editor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/micro_completer/cmd",
			Variant:     "",
			Url:         "https://micro-editor.github.io/",
			Execute:     common__micro.Execute,
		},
	},
	"minikube": {
		{
			Name:        "minikube",
			Description: "minikube quickly sets up a local Kubernetes cluster",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/minikube_completer/cmd",
			Variant:     "",
			Url:         "https://minikube.sigs.k8s.io/docs/",
			Execute:     common__minikube.Execute,
		},
	},
	"mitmproxy": {
		{
			Name:        "mitmproxy",
			Description: "interactive, SSL/TLS-capable intercepting proxy",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mitmproxy_completer/cmd",
			Variant:     "",
			Url:         "https://mitmproxy.org/",
			Execute:     common__mitmproxy.Execute,
		},
	},
	"mix": {
		{
			Name:        "mix",
			Description: "A build automation tool for working with applications written in the Elixir programming language",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mix_completer/cmd",
			Variant:     "",
			Url:         "https://hexdocs.pm/mix/Mix.html",
			Execute:     common__mix.Execute,
		},
	},
	"mkcert": {
		{
			Name:        "mkcert",
			Description: "simple tool for making locally-trusted development certificates",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mkcert_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/FiloSottile/mkcert",
			Execute:     common__mkcert.Execute,
		},
	},
	"mkdir": {
		{
			Name:        "mkdir",
			Description: "make directories",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mkdir_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/mkdir",
			Execute:     common__mkdir.Execute,
		},
	},
	"mkfifo": {
		{
			Name:        "mkfifo",
			Description: "make FIFOs (named pipes)",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mkfifo_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/3/mkfifo",
			Execute:     common__mkfifo.Execute,
		},
	},
	"mkfs": {
		{
			Name:        "mkfs",
			Description: "Make a Linux filesystem",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mkfs_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/mkfs",
			Execute:     common__mkfs.Execute,
		},
	},
	"mknod": {
		{
			Name:        "mknod",
			Description: "make block or character special files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mknod_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/3/mknod",
			Execute:     common__mknod.Execute,
		},
	},
	"mkswap": {
		{
			Name:        "mkswap",
			Description: "Set up a Linux swap area",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mkswap_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/mkswap",
			Execute:     common__mkswap.Execute,
		},
	},
	"mktemp": {
		{
			Name:        "mktemp",
			Description: "create a temporary file or directory",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mktemp_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/mktemp",
			Execute:     common__mktemp.Execute,
		},
	},
	"modinfo": {
		{
			Name:        "modinfo",
			Description: "Show information about a Linux Kernel module",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/modinfo_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/modinfo",
			Execute:     common__modinfo.Execute,
		},
	},
	"modprobe": {
		{
			Name:        "modprobe",
			Description: "Add and remove modules from the Linux Kernel",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/modprobe_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/modprobe",
			Execute:     common__modprobe.Execute,
		},
	},
	"molecule": {
		{
			Name:        "molecule",
			Description: "Testing framework to aid in the development of Ansible roles",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/molecule_completer/cmd",
			Variant:     "",
			Url:         "https://ansible.readthedocs.io/projects/molecule/",
			Execute:     common__molecule.Execute,
		},
	},
	"more": {
		{
			Name:        "more",
			Description: "display the contents of a file in a terminal",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/more_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/more.1.html",
			Execute:     common__more.Execute,
		},
	},
	"mosh": {
		{
			Name:        "mosh",
			Description: "mobile shell with roaming and intelligent local echo",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mosh_completer/cmd",
			Variant:     "",
			Url:         "https://mosh.org/",
			Execute:     common__mosh.Execute,
		},
	},
	"mount": {
		{
			Name:        "mount",
			Description: "mount a filesystem",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mount_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/mount",
			Execute:     common__mount.Execute,
		},
	},
	"mousepad": {
		{
			Name:        "mousepad",
			Description: "Mousepad is a simple text editor for the Xfce desktop environment",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mousepad_completer/cmd",
			Variant:     "",
			Url:         "http://users.xfce.org/~benny/xfce/apps.html",
			Execute:     common__mousepad.Execute,
		},
	},
	"mpv": {
		{
			Name:        "mpv",
			Description: "a media player",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mpv_completer/cmd",
			Variant:     "",
			Url:         "https://mpv.io/",
			Execute:     common__mpv.Execute,
		},
	},
	"mv": {
		{
			Name:        "mv",
			Description: "move (rename) files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mv_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/mv",
			Execute:     common__mv.Execute,
		},
	},
	"mvn": {
		{
			Name:        "mvn",
			Description: "Apache Maven is a software project management and comprehension tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/mvn_completer/cmd",
			Variant:     "",
			Url:         "https://maven.apache.org/",
			Execute:     common__mvn.Execute,
		},
	},
	"n-m3u8dl-re": {
		{
			Name:        "n-m3u8dl-re",
			Description: "Cross-Platform, modern and powerful stream downloader for MPD/M3U8/ISM",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/n-m3u8dl-re_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/nilaoda/N_m3u8DL-RE",
			Execute:     common__n_m3u8dl_re.Execute,
		},
	},
	"nano": {
		{
			Name:        "nano",
			Description: "Nano's ANOther editor, inspired by Pico",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nano_completer/cmd",
			Variant:     "",
			Url:         "https://www.nano-editor.org/",
			Execute:     common__nano.Execute,
		},
	},
	"nc": {
		{
			Name:        "nc",
			Description: "simple Unix utility which reads and writes data across network connections",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nc_completer/cmd",
			Variant:     "",
			Url:         "https://nc110.sourceforge.io/",
			Execute:     common__nc.Execute,
		},
	},
	"ncdu": {
		{
			Name:        "ncdu",
			Description: "NCurses Disk Usage",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ncdu_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ncdu",
			Execute:     common__ncdu.Execute,
		},
	},
	"neomutt": {
		{
			Name:        "neomutt",
			Description: "The NeoMutt Mail User Agent",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/neomutt_completer/cmd",
			Variant:     "",
			Url:         "https://neomutt.org/",
			Execute:     common__neomutt.Execute,
		},
	},
	"netcat": {
		{
			Name:        "netcat",
			Description: "simple Unix utility which reads and writes data across network connections",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/netcat_completer/cmd",
			Variant:     "",
			Url:         "https://nc110.sourceforge.io/",
			Execute:     common__netcat.Execute,
		},
	},
	"newman": {
		{
			Name:        "newman",
			Description: "Newman is a command-line collection runner for Postman",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/newman_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/postmanlabs/newman",
			Execute:     common__newman.Execute,
		},
	},
	"newrelic": {
		{
			Name:        "newrelic",
			Description: "The New Relic CLI",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/newrelic_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/newrelic/newrelic-cli",
			Execute:     common__newrelic.Execute,
		},
	},
	"nfpm": {
		{
			Name:        "nfpm",
			Description: "Packages apps on RPM, Deb and APK formats based on a YAML configuration file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nfpm_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/goreleaser/nfpm",
			Execute:     common__nfpm.Execute,
		},
	},
	"ng": {
		{
			Name:        "ng",
			Description: "The Angular CLI",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ng_completer/cmd",
			Variant:     "",
			Url:         "https://angular.io/cli",
			Execute:     common__ng.Execute,
		},
	},
	"nice": {
		{
			Name:        "nice",
			Description: "run a program with modified scheduling priority",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nice_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/nice",
			Execute:     common__nice.Execute,
		},
	},
	"nilaway": {
		{
			Name:        "nilaway",
			Description: "Static Analysis tool to detect potential Nil panics in Go code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nilaway_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/uber-go/nilaway",
			Execute:     common__nilaway.Execute,
		},
	},
	"nix": {
		{
			Name:        "nix",
			Description: "a tool for reproducible and declarative configuration management",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nix_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/NixOS/nix",
			Execute:     common__nix.Execute,
		},
	},
	"nix-build": {
		{
			Name:        "nix-build",
			Description: "build a Nix expression",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nix-build_completer/cmd",
			Variant:     "",
			Url:         "https://nixos.org/manual/nix/stable/command-ref/nix-build.html",
			Execute:     common__nix_build.Execute,
		},
	},
	"nix-channel": {
		{
			Name:        "nix-channel",
			Description: "manage Nix channels",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nix-channel_completer/cmd",
			Variant:     "",
			Url:         "https://nixos.org/manual/nix/stable/command-ref/nix-channel.html",
			Execute:     common__nix_channel.Execute,
		},
	},
	"nix-instantiate": {
		{
			Name:        "nix-instantiate",
			Description: "instantiate store derivations from Nix expression",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nix-instantiate_completer/cmd",
			Variant:     "",
			Url:         "https://nixos.org/manual/nix/stable/command-ref/nix-instantiate.html",
			Execute:     common__nix_instantiate.Execute,
		},
	},
	"nix-shell": {
		{
			Name:        "nix-shell",
			Description: "start an interactive shell based on a Nix expression",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nix-shell_completer/cmd",
			Variant:     "",
			Url:         "https://nixos.org/manual/nix/stable/command-ref/nix-shell.html",
			Execute:     common__nix_shell.Execute,
		},
	},
	"nixos-rebuild": {
		{
			Name:        "nixos-rebuild",
			Description: "reconfigure a NixOS machine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nixos-rebuild_completer/cmd",
			Variant:     "",
			Url:         "https://nixos.org",
			Execute:     common__nixos_rebuild.Execute,
		},
	},
	"nl": {
		{
			Name:        "nl",
			Description: "number lines of files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nl_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/nl",
			Execute:     common__nl.Execute,
		},
	},
	"nmcli": {
		{
			Name:        "nmcli",
			Description: "command-line tool for controlling NetworkManager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nmcli_completer/cmd",
			Variant:     "",
			Url:         "https://developer-old.gnome.org/NetworkManager/stable/nmcli.html",
			Execute:     common__nmcli.Execute,
		},
	},
	"node": {
		{
			Name:        "node",
			Description: "server-side JavaScript runtime",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/node_completer/cmd",
			Variant:     "",
			Url:         "https://nodejs.org/en/",
			Execute:     common__node.Execute,
		},
	},
	"nohup": {
		{
			Name:        "nohup",
			Description: "run a command immune to hangups, with output to a non-tty",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nohup_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/nohup",
			Execute:     common__nohup.Execute,
		},
	},
	"nomad": {
		{
			Name:        "nomad",
			Description: "Nomad is an easy-to-use, flexible, and performant workload orchestrator",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nomad_completer/cmd",
			Variant:     "",
			Url:         "https://www.nomadproject.io/",
			Execute:     common__nomad.Execute,
		},
	},
	"npm": {
		{
			Name:        "npm",
			Description: "the package manager for JavaScript",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/npm_completer/cmd",
			Variant:     "",
			Url:         "https://www.npmjs.com/",
			Execute:     common__npm.Execute,
		},
	},
	"nproc": {
		{
			Name:        "nproc",
			Description: "print the number of processing units available",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nproc_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/nproc.1.html",
			Execute:     common__nproc.Execute,
		},
	},
	"ntpd": {
		{
			Name:        "ntpd",
			Description: "NTP daemon program",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ntpd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/ntpd",
			Execute:     common__ntpd.Execute,
		},
	},
	"nu": {
		{
			Name:        "nu",
			Description: "Nushell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nu_completer/cmd",
			Variant:     "",
			Url:         "https://www.nushell.sh/",
			Execute:     common__nu.Execute,
		},
	},
	"numfmt": {
		{
			Name:        "numfmt",
			Description: "Convert numbers from/to human-readable strings",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/numfmt_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/numfmt.1.html",
			Execute:     common__numfmt.Execute,
		},
	},
	"nvim": {
		{
			Name:        "nvim",
			Description: "edit text",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/nvim_completer/cmd",
			Variant:     "",
			Url:         "https://neovim.io/",
			Execute:     common__nvim.Execute,
		},
	},
	"od": {
		{
			Name:        "od",
			Description: "dump files in octal and other formats",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/od_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/od",
			Execute:     common__od.Execute,
		},
	},
	"ollama": {
		{
			Name:        "ollama",
			Description: "Large language model runner",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ollama_completer/cmd",
			Variant:     "",
			Url:         "https://ollama.com/",
			Execute:     common__ollama.Execute,
		},
	},
	"openscad": {
		{
			Name:        "openscad",
			Description: "script file based graphical CAD environment",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/openscad_completer/cmd",
			Variant:     "",
			Url:         "https://openscad.org/",
			Execute:     common__openscad.Execute,
		},
	},
	"openssl": {
		{
			Name:        "openssl",
			Description: "OpenSSL command line program",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/openssl_completer/cmd",
			Variant:     "",
			Url:         "https://www.openssl.org/",
			Execute:     common__openssl.Execute,
		},
	},
	"optipng": {
		{
			Name:        "optipng",
			Description: "Optimize Portable Network Graphics files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/optipng_completer/cmd",
			Variant:     "",
			Url:         "http://optipng.sourceforge.net/",
			Execute:     common__optipng.Execute,
		},
	},
	"packer": {
		{
			Name:        "packer",
			Description: "Create identical machine images for multiple platforms from a single source configuration.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/packer_completer/cmd",
			Variant:     "",
			Url:         "https://www.packer.io/",
			Execute:     common__packer.Execute,
		},
	},
	"pacman": {
		{
			Name:        "pacman",
			Description: "package manager utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pacman_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.archlinux.de/title/Pacman",
			Execute:     common__pacman.Execute,
		},
	},
	"pacman-conf": {
		{
			Name:        "pacman-conf",
			Description: "query pacman's configuration file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-conf_completer/cmd",
			Variant:     "",
			Url:         "https://man.archlinux.org/man/pacman-conf.8.en",
			Execute:     common__pacman_conf.Execute,
		},
	},
	"pacman-db-upgrade": {
		{
			Name:        "pacman-db-upgrade",
			Description: "Upgrade the local pacman database to a newer format",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-db-upgrade_completer/cmd",
			Variant:     "",
			Url:         "https://manpages.debian.org/testing/pacman-package-manager/pacman-db-upgrade.8.en.html",
			Execute:     common__pacman_db_upgrade.Execute,
		},
	},
	"pacman-key": {
		{
			Name:        "pacman-key",
			Description: "Manage pacman's list of trusted keys",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-key_completer/cmd",
			Variant:     "",
			Url:         "https://archlinux.org/pacman/pacman-key.8.html",
			Execute:     common__pacman_key.Execute,
		},
	},
	"pacman-mirrors": {
		{
			Name:        "pacman-mirrors",
			Description: "generate pacman mirrorlist",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pacman-mirrors_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.manjaro.org/index.php?title=Pacman-mirrors",
			Execute:     common__pacman_mirrors.Execute,
		},
	},
	"palemoon": {
		{
			Name:        "palemoon",
			Description: "Pale Moon browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/palemoon_completer/cmd",
			Variant:     "",
			Url:         "https://www.palemoon.org/",
			Execute:     common__palemoon.Execute,
		},
	},
	"pamac": {
		{
			Name:        "pamac",
			Description: "package manager utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pamac_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.manjaro.org/index.php/Pamac",
			Execute:     common__pamac.Execute,
		},
	},
	"pandoc": {
		{
			Name:        "pandoc",
			Description: "general markup converter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pandoc_completer/cmd",
			Variant:     "",
			Url:         "https://pandoc.org/",
			Execute:     common__pandoc.Execute,
		},
	},
	"paru": {
		{
			Name:        "paru",
			Description: "Feature packed AUR helper",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/paru_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/Morganamilo/paru",
			Execute:     common__paru.Execute,
		},
	},
	"pass": {
		{
			Name:        "pass",
			Description: "stores, retrieves, generates, and synchronizes passwords securely",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pass_completer/cmd",
			Variant:     "",
			Url:         "https://www.passwordstore.org/",
			Execute:     common__pass.Execute,
		},
	},
	"passwd": {
		{
			Name:        "passwd",
			Description: "change user password",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/passwd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/passwd",
			Execute:     common__passwd.Execute,
		},
	},
	"paste": {
		{
			Name:        "paste",
			Description: "merge lines of files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/paste_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/paste",
			Execute:     common__paste.Execute,
		},
	},
	"patch": {
		{
			Name:        "patch",
			Description: "appy a diff file to an original",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/patch_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/patch",
			Execute:     common__patch.Execute,
		},
	},
	"pathchk": {
		{
			Name:        "pathchk",
			Description: "check whether file names are valid or portable",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pathchk_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/pathchk",
			Execute:     common__pathchk.Execute,
		},
	},
	"patool": {
		{
			Name:        "patool",
			Description: "An archive file manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/patool_completer/cmd",
			Variant:     "",
			Url:         "http://wummel.github.io/patool/",
			Execute:     common__patool.Execute,
		},
	},
	"pcmanfm": {
		{
			Name:        "pcmanfm",
			Description: "A lightweight Gtk+ based file manager for X Window",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pcmanfm_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.lxde.org/en/PCManFM",
			Execute:     common__pcmanfm.Execute,
		},
	},
	"pgrep": {
		{
			Name:        "pgrep",
			Description: "look up processes based on name and other attributes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pgrep_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/pgrep.1.html",
			Execute:     common__pgrep.Execute,
		},
	},
	"php": {
		{
			Name:        "php",
			Description: "PHP Command Line Interface",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/php_completer/cmd",
			Variant:     "",
			Url:         "https://www.php.net/",
			Execute:     common__php.Execute,
		},
	},
	"picard": {
		{
			Name:        "picard",
			Description: "Picard is a cross-platform music tagger written in Python",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/picard_completer/cmd",
			Variant:     "",
			Url:         "https://picard.musicbrainz.org/",
			Execute:     common__picard.Execute,
		},
	},
	"pidof": {
		{
			Name:        "pidof",
			Description: "find the process ID of a running program",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pidof_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/pidof.1.html",
			Execute:     common__pidof.Execute,
		},
	},
	"pidwait": {
		{
			Name:        "pidwait",
			Description: "wait for processes based on name and other attributes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pidwait_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/pgrep.1.html",
			Execute:     common__pidwait.Execute,
		},
	},
	"pigz": {
		{
			Name:        "pigz",
			Description: "compress or expand files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pigz_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/pigz",
			Execute:     common__pigz.Execute,
		},
	},
	"ping": {
		{
			Name:        "ping",
			Description: "send ICMP ECHO_REQUEST to network hosts",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ping_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/ping",
			Execute:     common__ping.Execute,
		},
	},
	"pinky": {
		{
			Name:        "pinky",
			Description: "lightweight finger",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pinky_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/pinky",
			Execute:     common__pinky.Execute,
		},
	},
	"pip": {
		{
			Name:        "pip",
			Description: "package manager for Python packages",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pip_completer/cmd",
			Variant:     "",
			Url:         "https://pip.pypa.io/en/stable/",
			Execute:     common__pip.Execute,
		},
	},
	"pkg": {
		{
			Name:        "pkg",
			Description: "A tool for managing packages",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pkg_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.termux.com/wiki/Package_Management",
			Execute:     common__pkg.Execute,
		},
	},
	"pkgsite": {
		{
			Name:        "pkgsite",
			Description: "Pkgsite extracts and generates documentation for Go programs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pkgsite_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite",
			Execute:     common__pkgsite.Execute,
		},
	},
	"pkill": {
		{
			Name:        "pkill",
			Description: "look up for processes based on name and other attributes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pkill_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/pgrep.1.html",
			Execute:     common__pkill.Execute,
		},
	},
	"pmap": {
		{
			Name:        "pmap",
			Description: "report memory map of a process",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pmap_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/pmap.1.html",
			Execute:     common__pmap.Execute,
		},
	},
	"pngcheck": {
		{
			Name:        "pngcheck",
			Description: "Test PNG, JNG or MNG image files for corruption",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pngcheck_completer/cmd",
			Variant:     "",
			Url:         "http://www.libpng.org/pub/png/apps/pngcheck.html",
			Execute:     common__pngcheck.Execute,
		},
	},
	"pnpm": {
		{
			Name:        "pnpm",
			Description: "Fast, disk space efficient package manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pnpm_completer/cmd",
			Variant:     "",
			Url:         "https://pnpm.io/",
			Execute:     common__pnpm.Execute,
		},
	},
	"podman": {
		{
			Name:        "podman",
			Description: "Simple management tool for pods, containers and images",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/podman_completer/cmd",
			Variant:     "",
			Url:         "https://podman.io/",
			Execute:     common__podman.Execute,
		},
	},
	"poweroff": {
		{
			Name:        "poweroff",
			Description: "poweroff the machine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/poweroff_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/poweroff.html",
			Execute:     common__poweroff.Execute,
		},
	},
	"powertop": {
		{
			Name:        "powertop",
			Description: "The Linux PowerTOP tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/powertop_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/fenrus75/powertop",
			Execute:     common__powertop.Execute,
		},
	},
	"pprof": {
		{
			Name:        "pprof",
			Description: "pprof is a tool for visualization and analysis of profiling data",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pprof_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/google/pprof",
			Execute:     common__pprof.Execute,
		},
	},
	"pr": {
		{
			Name:        "pr",
			Description: "convert text files for printing",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pr_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/pr",
			Execute:     common__pr.Execute,
		},
	},
	"present": {
		{
			Name:        "present",
			Description: "present implements parsing and rendering of present file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/present_completer/cmd",
			Variant:     "",
			Url:         "https://pkg.go.dev/golang.org/x/tools/present",
			Execute:     common__present.Execute,
		},
	},
	"prettybat": {
		{
			Name:        "prettybat",
			Description: "Pretty-print source code and highlight it with bat",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/prettybat_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/eth-p/bat-extras/blob/master/doc/prettybat.md",
			Execute:     common__prettybat.Execute,
		},
	},
	"prettyping": {
		{
			Name:        "prettyping",
			Description: "This script is a wrapper around the system's \\\"ping\\\" tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/prettyping_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/denilsonsa/prettyping",
			Execute:     common__prettyping.Execute,
		},
	},
	"printenv": {
		{
			Name:        "printenv",
			Description: "print all or part of environment",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/printenv_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/printenv",
			Execute:     common__printenv.Execute,
		},
	},
	"procs": {
		{
			Name:        "procs",
			Description: "a replacement for `ps` written in Rust",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/procs_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/dalance/procs",
			Execute:     common__procs.Execute,
		},
	},
	"ps": {
		{
			Name:        "ps",
			Description: "report a snapshot of the current processes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ps_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/ps.1.html",
			Execute:     common__ps.Execute,
		},
	},
	"ptx": {
		{
			Name:        "ptx",
			Description: "produce a permuted index of file contents",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ptx_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ptx",
			Execute:     common__ptx.Execute,
		},
	},
	"pulumi": {
		{
			Name:        "pulumi",
			Description: "Pulumi command line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pulumi_completer/cmd",
			Variant:     "",
			Url:         "https://www.pulumi.com/",
			Execute:     common__pulumi.Execute,
		},
	},
	"pwd": {
		{
			Name:        "pwd",
			Description: "Print the full filename of the current working directory",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pwd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/pwd",
			Execute:     common__pwd.Execute,
		},
	},
	"pwdx": {
		{
			Name:        "pwdx",
			Description: "report current working directory of a process",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/pwdx_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/pwdx.1.html",
			Execute:     common__pwdx.Execute,
		},
	},
	"python": {
		{
			Name:        "python",
			Description: "an interpreted, interactive, object-oriented programming language",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/python_completer/cmd",
			Variant:     "",
			Url:         "https://www.python.org/",
			Execute:     common__python.Execute,
		},
	},
	"qmk": {
		{
			Name:        "qmk",
			Description: "CLI wrapper for running QMK commands",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/qmk_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/qmk/qmk_cli",
			Execute:     common__qmk.Execute,
		},
	},
	"qrencode": {
		{
			Name:        "qrencode",
			Description: "Encode input data in a QR Code and save as a PNG or EPS image",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/qrencode_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/qrencode",
			Execute:     common__qrencode.Execute,
		},
	},
	"qutebrowser": {
		{
			Name:        "qutebrowser",
			Description: "a keyboard-driven, vim-like browser based on PyQt5",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/qutebrowser_completer/cmd",
			Variant:     "",
			Url:         "https://qutebrowser.org/",
			Execute:     common__qutebrowser.Execute,
		},
	},
	"ramalama": {
		{
			Name:        "ramalama",
			Description: "tool for working with LLM models",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ramalama_completer/cmd",
			Variant:     "",
			Url:         "https://ramalama.ai/",
			Execute:     common__ramalama.Execute,
		},
	},
	"ranger": {
		{
			Name:        "ranger",
			Description: "visual file manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ranger_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/ranger/ranger",
			Execute:     common__ranger.Execute,
		},
	},
	"readlink": {
		{
			Name:        "readlink",
			Description: "print resolved symbolic links or canonical file names",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/readlink_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/readlink",
			Execute:     common__readlink.Execute,
		},
	},
	"reboot": {
		{
			Name:        "reboot",
			Description: "reboot the machine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/reboot_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/reboot.html",
			Execute:     common__reboot.Execute,
		},
	},
	"redis-cli": {
		{
			Name:        "redis-cli",
			Description: "Redis command line interface",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/redis-cli_completer/cmd",
			Variant:     "",
			Url:         "https://redis.io/docs/manual/cli/",
			Execute:     common__redis_cli.Execute,
		},
	},
	"rename": {
		{
			Name:        "rename",
			Description: "rename files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rename_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/rename.1.html",
			Execute:     common__rename.Execute,
		},
	},
	"restic": {
		{
			Name:        "restic",
			Description: "Backup and restore files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/restic_completer/cmd",
			Variant:     "",
			Url:         "https://restic.net/",
			Execute:     common__restic.Execute,
		},
	},
	"resume-cli": {
		{
			Name:        "resume-cli",
			Description: "command line tool for JSON Resume",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/resume-cli_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/jsonresume/resume-cli",
			Execute:     common__resume_cli.Execute,
		},
	},
	"rg": {
		{
			Name:        "rg",
			Description: "recursively search current directory for lines matching a pattern",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rg_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/BurntSushi/ripgrep",
			Execute:     common__rg.Execute,
		},
	},
	"rifle": {
		{
			Name:        "rifle",
			Description: "ranger's file opener",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rifle_completer/cmd",
			Variant:     "",
			Url:         "https://ranger.github.io/",
			Execute:     common__rifle.Execute,
		},
	},
	"ripsecrets": {
		{
			Name:        "ripsecrets",
			Description: "Prevent committing secret keys into your source code",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ripsecrets_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/sirwart/ripsecrets",
			Execute:     common__ripsecrets.Execute,
		},
	},
	"rm": {
		{
			Name:        "rm",
			Description: "remove files or directories",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rm_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/rm",
			Execute:     common__rm.Execute,
		},
	},
	"rmdir": {
		{
			Name:        "rmdir",
			Description: "remove empty directories",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rmdir_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/rmdir",
			Execute:     common__rmdir.Execute,
		},
	},
	"rmmod": {
		{
			Name:        "rmmod",
			Description: "Simple program to remove a module from the Linux Kernel",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rmmod_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/rmmod",
			Execute:     common__rmmod.Execute,
		},
	},
	"rsync": {
		{
			Name:        "rsync",
			Description: "a fast, versatile, remote (and local) file-copying tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rsync_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/WayneD/rsync/",
			Execute:     common__rsync.Execute,
		},
	},
	"run0": {
		{
			Name:        "run0",
			Description: "Elevate privileges interactively",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/run0_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/run0.html",
			Execute:     common__run0.Execute,
		},
	},
	"rust-analyzer": {
		{
			Name:        "rust-analyzer",
			Description: "LSP server for the Rust programming language",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rust-analyzer_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/rust-lang/rust-analyzer",
			Execute:     common__rust_analyzer.Execute,
		},
	},
	"rustc": {
		{
			Name:        "rustc",
			Description: "compiler for the Rust programming language",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rustc_completer/cmd",
			Variant:     "",
			Url:         "https://doc.rust-lang.org/rustc/index.html",
			Execute:     common__rustc.Execute,
		},
	},
	"rustdoc": {
		{
			Name:        "rustdoc",
			Description: "generate documentation for Rust projects",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rustdoc_completer/cmd",
			Variant:     "",
			Url:         "https://doc.rust-lang.org/1.71.0/rustdoc/what-is-rustdoc.html",
			Execute:     common__rustdoc.Execute,
		},
	},
	"rustup": {
		{
			Name:        "rustup",
			Description: "The Rust toolchain installer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/rustup_completer/cmd",
			Variant:     "",
			Url:         "https://rustup.rs/",
			Execute:     common__rustup.Execute,
		},
	},
	"saw": {
		{
			Name:        "saw",
			Description: "A fast, multipurpose tool for AWS CloudWatch Logs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/saw_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/TylerBrock/saw",
			Execute:     common__saw.Execute,
		},
	},
	"scp": {
		{
			Name:        "scp",
			Description: "OpenSSH secure file copy",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/scp_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/scp",
			Execute:     common__scp.Execute,
		},
	},
	"script": {
		{
			Name:        "script",
			Description: "make typescript of terminal session",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/script_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/script.1.html",
			Execute:     common__script.Execute,
		},
	},
	"scriptlive": {
		{
			Name:        "scriptlive",
			Description: "re-run session typescripts, using timing information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/scriptlive_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/scriptlive.1.html",
			Execute:     common__scriptlive.Execute,
		},
	},
	"scriptreplay": {
		{
			Name:        "scriptreplay",
			Description: "play back typescripts, using timing information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/scriptreplay_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/scriptreplay.1.html",
			Execute:     common__scriptreplay.Execute,
		},
	},
	"scrot": {
		{
			Name:        "scrot",
			Description: "command line screen capture utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/scrot_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Scrot",
			Execute:     common__scrot.Execute,
		},
	},
	"sd": {
		{
			Name:        "sd",
			Description: "Intuitive find & replace CLI",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sd_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/chmln/sd",
			Execute:     common__sd.Execute,
		},
	},
	"sdkmanager": {
		{
			Name:        "sdkmanager",
			Description: "Android SDK manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sdkmanager_completer/cmd",
			Variant:     "",
			Url:         "https://developer.android.com/studio/command-line/sdkmanager",
			Execute:     common__sdkmanager.Execute,
		},
	},
	"sed": {
		{
			Name:        "sed",
			Description: "stream editor for filtering and transforming text",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sed_completer/cmd",
			Variant:     "",
			Url:         "https://www.gnu.org/software/sed/manual/sed.html",
			Execute:     common__sed.Execute,
		},
	},
	"semver": {
		{
			Name:        "semver",
			Description: "A JavaScript implementation of the https://semver.org/ specification",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/semver_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/npm/node-semver",
			Execute:     common__semver.Execute,
		},
	},
	"seq": {
		{
			Name:        "seq",
			Description: "print a sequence of numbers",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/seq_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/seq",
			Execute:     common__seq.Execute,
		},
	},
	"serie": {
		{
			Name:        "serie",
			Description: "A rich git commit graph in your terminal, like magic 📚",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/serie_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/lusingander/serie",
			Execute:     common__serie.Execute,
		},
	},
	"set-env": {
		{
			Name:        "set-env",
			Description: "set environment variable",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/set-env_completer/cmd",
			Variant:     "",
			Url:         "https://carapace-sh.github.io/carapace-bin/environment.html",
			Execute:     common__set_env.Execute,
		},
	},
	"sftp": {
		{
			Name:        "sftp",
			Description: "OpenSSH secure file transfer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sftp_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/sftp",
			Execute:     common__sftp.Execute,
		},
	},
	"sha1sum": {
		{
			Name:        "sha1sum",
			Description: "compute and check SHA1 message digest",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sha1sum_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/sha1sum",
			Execute:     common__sha1sum.Execute,
		},
	},
	"sha224sum": {
		{
			Name:        "sha224sum",
			Description: "Print or check SHA224 (224-bit) checksums",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sha224sum_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/sha224sum.1.html",
			Execute:     common__sha224sum.Execute,
		},
	},
	"sha256sum": {
		{
			Name:        "sha256sum",
			Description: "compute and check SHA256 message digest",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sha256sum_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/sha256sum",
			Execute:     common__sha256sum.Execute,
		},
	},
	"sha384sum": {
		{
			Name:        "sha384sum",
			Description: "Print or check SHA384 (384-bit) checksums",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sha384sum_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/sha384sum.1.html",
			Execute:     common__sha384sum.Execute,
		},
	},
	"sha512sum": {
		{
			Name:        "sha512sum",
			Description: "Print or check SHA512 (512-bit) checksums",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sha512sum_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/sha512sum.1.html",
			Execute:     common__sha512sum.Execute,
		},
	},
	"showkey": {
		{
			Name:        "showkey",
			Description: "examine the codes sent by the keyboard",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/showkey_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/showkey",
			Execute:     common__showkey.Execute,
		},
	},
	"shred": {
		{
			Name:        "shred",
			Description: "overwrite a file to hide its contents, and optionally delete it",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/shred_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/shred",
			Execute:     common__shred.Execute,
		},
	},
	"shutdown": {
		{
			Name:        "shutdown",
			Description: "Shut down the system",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/shutdown_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/shutdown.html",
			Execute:     common__shutdown.Execute,
		},
	},
	"skhd": {
		{
			Name:        "skhd",
			Description: "Simple hotkey daemon for macOS",
			Group:       "darwin",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/darwin/skhd_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/koekeishiya/skhd",
			Execute:     darwin__skhd.Execute,
		},
	},
	"slabtop": {
		{
			Name:        "slabtop",
			Description: "display kernel slab cache information in real time",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/slabtop_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/slabtop.1.html",
			Execute:     common__slabtop.Execute,
		},
	},
	"sleep": {
		{
			Name:        "sleep",
			Description: "delay for a specified amount of time",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sleep_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/3/sleep",
			Execute:     common__sleep.Execute,
		},
	},
	"slides": {
		{
			Name:        "slides",
			Description: "Terminal based presentation tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/slides_completer/cmd",
			Variant:     "",
			Url:         "https://maaslalani.com/slides/",
			Execute:     common__slides.Execute,
		},
	},
	"soft": {
		{
			Name:        "soft",
			Description: "A self-hostable Git server for the command line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/soft_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/soft-serve",
			Execute:     common__soft.Execute,
		},
	},
	"sort": {
		{
			Name:        "sort",
			Description: "sort lines of text files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sort_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/sort",
			Execute:     common__sort.Execute,
		},
	},
	"speedtest-cli": {
		{
			Name:        "speedtest-cli",
			Description: "Command line interface for testing internet bandwidth using speedtest.net",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/speedtest-cli_completer/cmd",
			Variant:     "",
			Url:         "https://www.speedtest.net/apps/cli",
			Execute:     common__speedtest_cli.Execute,
		},
	},
	"split": {
		{
			Name:        "split",
			Description: "split a file into pieces",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/split_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/split",
			Execute:     common__split.Execute,
		},
	},
	"sqlite3": {
		{
			Name:        "sqlite3",
			Description: "A command-line interface for SQLite",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sqlite3_completer/cmd",
			Variant:     "",
			Url:         "https://sqlite.org/",
			Execute:     common__sqlite3.Execute,
		},
	},
	"ssh": {
		{
			Name:        "ssh",
			Description: "OpenSSH remote login client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ssh_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ssh",
			Execute:     common__ssh.Execute,
		},
	},
	"ssh-agent": {
		{
			Name:        "ssh-agent",
			Description: "OpenSSH authentication agent",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ssh-agent_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ssh-agent",
			Execute:     common__ssh_agent.Execute,
		},
	},
	"ssh-copy-id": {
		{
			Name:        "ssh-copy-id",
			Description: "use locally available keys to authorise logins on a remote machine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ssh-copy-id_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ssh-copy-id",
			Execute:     common__ssh_copy_id.Execute,
		},
	},
	"ssh-keygen": {
		{
			Name:        "ssh-keygen",
			Description: "OpenSSH authentication key utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ssh-keygen_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ssh-keygen",
			Execute:     common__ssh_keygen.Execute,
		},
	},
	"st": {
		{
			Name:        "st",
			Description: "simple terminal",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/st_completer/cmd",
			Variant:     "",
			Url:         "https://st.suckless.org/",
			Execute:     common__st.Execute,
		},
	},
	"starship": {
		{
			Name:        "starship",
			Description: "The cross-shell prompt for astronauts",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/starship_completer/cmd",
			Variant:     "",
			Url:         "https://starship.rs/",
			Execute:     common__starship.Execute,
		},
	},
	"stat": {
		{
			Name:        "stat",
			Description: "display file or file system status",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/stat_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/2/stat",
			Execute:     common__stat.Execute,
		},
	},
	"staticcheck": {
		{
			Name:        "staticcheck",
			Description: "The advanced Go linter",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/staticcheck_completer/cmd",
			Variant:     "",
			Url:         "https://staticcheck.io/",
			Execute:     common__staticcheck.Execute,
		},
	},
	"strings": {
		{
			Name:        "strings",
			Description: "print the sequences of printable characters in files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/strings_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/strings",
			Execute:     common__strings.Execute,
		},
	},
	"stty": {
		{
			Name:        "stty",
			Description: "change and print terminal line settings",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/stty_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/stty",
			Execute:     common__stty.Execute,
		},
	},
	"su": {
		{
			Name:        "su",
			Description: "run a command with substitute user and group ID",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/su_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/su.1.html",
			Execute:     common__su.Execute,
		},
	},
	"sudo": {
		{
			Name:        "sudo",
			Description: "execute a command as another user",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sudo_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/sudo",
			Execute:     common__sudo.Execute,
		},
	},
	"sudoedit": {
		{
			Name:        "sudoedit",
			Description: "edit files as another user",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sudoedit_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/sudoedit",
			Execute:     common__sudoedit.Execute,
		},
	},
	"sudoreplay": {
		{
			Name:        "sudoreplay",
			Description: "replay sudo session logs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sudoreplay_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/sudoreplay",
			Execute:     common__sudoreplay.Execute,
		},
	},
	"sulogin": {
		{
			Name:        "sulogin",
			Description: "single-user login",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sulogin_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/sulogin",
			Execute:     common__sulogin.Execute,
		},
	},
	"sum": {
		{
			Name:        "sum",
			Description: "checksum and count the blocks in a file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sum_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/sum",
			Execute:     common__sum.Execute,
		},
	},
	"supervisorctl": {
		{
			Name:        "supervisorctl",
			Description: "control applications run by supervisord from the cmd line",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/supervisorctl_completer/cmd",
			Variant:     "",
			Url:         "http://supervisord.org/",
			Execute:     common__supervisorctl.Execute,
		},
	},
	"supervisord": {
		{
			Name:        "supervisord",
			Description: "run a set of applications as daemons",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/supervisord_completer/cmd",
			Variant:     "",
			Url:         "http://supervisord.org/",
			Execute:     common__supervisord.Execute,
		},
	},
	"svg-term": {
		{
			Name:        "svg-term",
			Description: "Share terminal sessions as razor-sharp animated SVG everywhere",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/svg-term_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/marionebl/svg-term-cli",
			Execute:     common__svg_term.Execute,
		},
	},
	"svgcleaner": {
		{
			Name:        "svgcleaner",
			Description: "clean up your SVG files from the unnecessary data",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/svgcleaner_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/RazrFalcon/svgcleaner",
			Execute:     common__svgcleaner.Execute,
		},
	},
	"sway": {
		{
			Name:        "sway",
			Description: "An i3-compatible Wayland compositor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sway_completer/cmd",
			Variant:     "",
			Url:         "https://swaywm.org/",
			Execute:     common__sway.Execute,
		},
	},
	"swaybar": {
		{
			Name:        "swaybar",
			Description: "bar for swaywm",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/swaybar_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/swaywm/sway",
			Execute:     common__swaybar.Execute,
		},
	},
	"swaybg": {
		{
			Name:        "swaybg",
			Description: "Background for Wayland",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/swaybg_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/swaywm/swaybg",
			Execute:     common__swaybg.Execute,
		},
	},
	"swayidle": {
		{
			Name:        "swayidle",
			Description: "Idle manager for Wayland",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/swayidle_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/swaywm/swayidle",
			Execute:     common__swayidle.Execute,
		},
	},
	"swaylock": {
		{
			Name:        "swaylock",
			Description: "Screen locker for Wayland",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/swaylock_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/swaywm/swaylock",
			Execute:     common__swaylock.Execute,
		},
	},
	"swaymsg": {
		{
			Name:        "swaymsg",
			Description: "Send messages to a running instance of sway over the IPC socket",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/swaymsg_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/swaywm/sway",
			Execute:     common__swaymsg.Execute,
		},
	},
	"swaynag": {
		{
			Name:        "swaynag",
			Description: "Show a warning or error message with buttons",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/swaynag_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/swaywm/sway",
			Execute:     common__swaynag.Execute,
		},
	},
	"syft": {
		{
			Name:        "syft",
			Description: "Generate a package SBOM",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/syft_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/anchore/syft",
			Execute:     common__syft.Execute,
		},
	},
	"sync": {
		{
			Name:        "sync",
			Description: "Synchronize cached writes to persistent storage",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sync_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/sync",
			Execute:     common__sync.Execute,
		},
	},
	"sysctl": {
		{
			Name:        "sysctl",
			Description: "configure kernel parameters at runtime",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/sysctl_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/sysctl.8.html",
			Execute:     common__sysctl.Execute,
		},
	},
	"systemctl": {
		{
			Name:        "systemctl",
			Description: "Query or send control commands to the system manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/systemctl_completer/cmd",
			Variant:     "",
			Url:         "https://systemd.io/",
			Execute:     common__systemctl.Execute,
		},
	},
	"systemd-analyze": {
		{
			Name:        "systemd-analyze",
			Description: "Analyze and debug system manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/systemd-analyze_completer/cmd",
			Variant:     "",
			Url:         "https://www.freedesktop.org/software/systemd/man/latest/systemd-analyze.html",
			Execute:     common__systemd_analyze.Execute,
		},
	},
	"tac": {
		{
			Name:        "tac",
			Description: "concatenate and print files in reverse",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tac_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tac",
			Execute:     common__tac.Execute,
		},
	},
	"tail": {
		{
			Name:        "tail",
			Description: "output the last part of files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tail_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tail",
			Execute:     common__tail.Execute,
		},
	},
	"taplo": {
		{
			Name:        "taplo",
			Description: "A TOML toolkit written in Rust",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/taplo_completer/cmd",
			Variant:     "",
			Url:         "https://taplo.tamasfe.dev/",
			Execute:     common__taplo.Execute,
		},
	},
	"tar": {
		{
			Name:        "tar",
			Description: "tar - an archiving utility",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tar_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tar",
			Execute:     common__tar.Execute,
		},
	},
	"task": {
		{
			Name:        "task",
			Description: "A task runner / simpler Make alternative written in Go",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/task_completer/cmd",
			Variant:     "",
			Url:         "https://taskfile.dev/",
			Execute:     common__task.Execute,
		},
	},
	"tea": {
		{
			Name:        "tea",
			Description: "command line tool to interact with Gitea",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tea_completer/cmd",
			Variant:     "",
			Url:         "https://gitea.com/gitea/tea",
			Execute:     common__tea.Execute,
		},
	},
	"tee": {
		{
			Name:        "tee",
			Description: "read from standard input and write to standard output and files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tee_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tee",
			Execute:     common__tee.Execute,
		},
	},
	"telnet": {
		{
			Name:        "telnet",
			Description: "User interface to TELNET",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/telnet_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/telnet",
			Execute:     common__telnet.Execute,
		},
	},
	"templ": {
		{
			Name:        "templ",
			Description: "A language for writing HTML user interfaces in Go",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/templ_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/a-h/templ",
			Execute:     common__templ.Execute,
		},
	},
	"termux-apt-repo": {
		{
			Name:        "termux-apt-repo",
			Description: "Create a repository with deb files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/termux-apt-repo_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/termux/termux-apt-repo",
			Execute:     common__termux_apt_repo.Execute,
		},
	},
	"terraform": {
		{
			Name:        "terraform",
			Description: "infrastructure as code software tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/terraform_completer/cmd",
			Variant:     "",
			Url:         "https://www.terraform.io/",
			Execute:     common__terraform.Execute,
		},
	},
	"terraform-ls": {
		{
			Name:        "terraform-ls",
			Description: "Terraform Language Server",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/terraform-ls_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/hashicorp/terraform-ls",
			Execute:     common__terraform_ls.Execute,
		},
	},
	"terragrunt": {
		{
			Name:        "terragrunt",
			Description: "Terragrunt is a thin wrapper for Terraform",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/terragrunt_completer/cmd",
			Variant:     "",
			Url:         "https://terragrunt.gruntwork.io/",
			Execute:     common__terragrunt.Execute,
		},
	},
	"terramate": {
		{
			Name:        "terramate",
			Description: "A tool for managing terraform stacks",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/terramate_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/mineiros-io/terramate",
			Execute:     common__terramate.Execute,
		},
	},
	"tesseract": {
		{
			Name:        "tesseract",
			Description: "command-line OCR engine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tesseract_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tesseract-ocr/tessdoc",
			Execute:     common__tesseract.Execute,
		},
	},
	"tig": {
		{
			Name:        "tig",
			Description: "text-mode interface for Git",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tig_completer/cmd",
			Variant:     "",
			Url:         "https://jonas.github.io/tig/",
			Execute:     common__tig.Execute,
		},
	},
	"timeout": {
		{
			Name:        "timeout",
			Description: "run a command with a time limit",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/timeout_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/timeout.1.html",
			Execute:     common__timeout.Execute,
		},
	},
	"tinygo": {
		{
			Name:        "tinygo",
			Description: "TinyGo is a Go compiler for small places",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tinygo_completer/cmd",
			Variant:     "",
			Url:         "https://tinygo.org/",
			Execute:     common__tinygo.Execute,
		},
	},
	"tldr": {
		{
			Name:        "tldr",
			Description: "A fast TLDR client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tldr_completer/tealdeer/cmd",
			Variant:     "tealdeer",
			Url:         "https://github.com/dbrgn/tealdeer",
			Execute:     common__tldr__tealdeer.Execute,
		},
		{
			Name:        "tldr",
			Description: "Python command line client for tldr",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tldr_completer/tldr-python-client/cmd",
			Variant:     "tldr-python-client",
			Url:         "https://github.com/tldr-pages/tldr-python-client",
			Execute:     common__tldr__tldr_python_client.Execute,
		},
	},
	"tload": {
		{
			Name:        "tload",
			Description: "graphic representation of system load average",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tload_completer/cmd",
			Variant:     "",
			Url:         "https://www.man7.org/linux/man-pages/man1/tload.1.html",
			Execute:     common__tload.Execute,
		},
	},
	"tmate": {
		{
			Name:        "tmate",
			Description: "Instant terminal sharing",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tmate_completer/cmd",
			Variant:     "",
			Url:         "https://tmate.io/",
			Execute:     common__tmate.Execute,
		},
	},
	"tmux": {
		{
			Name:        "tmux",
			Description: "terminal multiplexer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tmux_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tmux/tmux/wiki",
			Execute:     common__tmux.Execute,
		},
	},
	"tofu": {
		{
			Name:        "tofu",
			Description: "The open source infrastructure as code tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tofu_completer/cmd",
			Variant:     "",
			Url:         "https://opentofu.org/",
			Execute:     common__tofu.Execute,
		},
	},
	"toit.lsp": {
		{
			Name:        "toit.lsp",
			Description: "start the lsp server",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/toit.lsp_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/toitlang/toit",
			Execute:     common__toit_lsp.Execute,
		},
	},
	"toit.pkg": {
		{
			Name:        "toit.pkg",
			Description: "The Toit package manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/toit.pkg_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/toitlang/toit",
			Execute:     common__toit_pkg.Execute,
		},
	},
	"top": {
		{
			Name:        "top",
			Description: "display Linux processes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/top_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/top.1.html",
			Execute:     common__top.Execute,
		},
	},
	"tor-browser": {
		{
			Name:        "tor-browser",
			Description: "Tor Browser",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tor-browser_completer/cmd",
			Variant:     "",
			Url:         "https://www.torproject.org/",
			Execute:     common__tor_browser.Execute,
		},
	},
	"tor-gencert": {
		{
			Name:        "tor-gencert",
			Description: "Generate certs and keys for Tor directory authorities",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tor-gencert_completer/cmd",
			Variant:     "",
			Url:         "https://www.torproject.org/",
			Execute:     common__tor_gencert.Execute,
		},
	},
	"tor-print-ed-signing-cert": {
		{
			Name:        "tor-print-ed-signing-cert",
			Description: "print expiration date of ed25519 signing certificate",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tor-print-ed-signing-cert_completer/cmd",
			Variant:     "",
			Url:         "https://www.torproject.org/",
			Execute:     common__tor_print_ed_signing_cert.Execute,
		},
	},
	"tor-resolve": {
		{
			Name:        "tor-resolve",
			Description: "resolve a hostname to an IP address via tor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tor-resolve_completer/cmd",
			Variant:     "",
			Url:         "https://www.torproject.org/",
			Execute:     common__tor_resolve.Execute,
		},
	},
	"torsocks": {
		{
			Name:        "torsocks",
			Description: "Shell wrapper to simplify the use of the torsocks(8) library to transparently torify an application",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/torsocks_completer/cmd",
			Variant:     "",
			Url:         "https://www.torproject.org/",
			Execute:     common__torsocks.Execute,
		},
	},
	"touch": {
		{
			Name:        "touch",
			Description: "change file timestamps",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/touch_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/touch",
			Execute:     common__touch.Execute,
		},
	},
	"tox": {
		{
			Name:        "tox",
			Description: "automation project",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tox_completer/cmd",
			Variant:     "",
			Url:         "https://tox.wiki/",
			Execute:     common__tox.Execute,
		},
	},
	"tr": {
		{
			Name:        "tr",
			Description: "translate or delete characters",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tr_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tr",
			Execute:     common__tr.Execute,
		},
	},
	"traefik": {
		{
			Name:        "traefik",
			Description: "Traefik is a modern HTTP reverse proxy and load balancer made to deploy microservices with ease",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/traefik_completer/cmd",
			Variant:     "",
			Url:         "https://traefik.io/",
			Execute:     common__traefik.Execute,
		},
	},
	"transmission-cli": {
		{
			Name:        "transmission-cli",
			Description: "A fast and easy BitTorrent client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-cli_completer/cmd",
			Variant:     "",
			Url:         "https://transmissionbt.com/",
			Execute:     common__transmission_cli.Execute,
		},
	},
	"transmission-create": {
		{
			Name:        "transmission-create",
			Description: "A command-line utility to create .torrent files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-create_completer/cmd",
			Variant:     "",
			Url:         "https://transmissionbt.com/",
			Execute:     common__transmission_create.Execute,
		},
	},
	"transmission-daemon": {
		{
			Name:        "transmission-daemon",
			Description: "A daemon-based BitTorrent client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-daemon_completer/cmd",
			Variant:     "",
			Url:         "https://transmissionbt.com/",
			Execute:     common__transmission_daemon.Execute,
		},
	},
	"transmission-edit": {
		{
			Name:        "transmission-edit",
			Description: "A command-line utility to modify .torrent files' announce URLs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-edit_completer/cmd",
			Variant:     "",
			Url:         "https://transmissionbt.com/",
			Execute:     common__transmission_edit.Execute,
		},
	},
	"transmission-remote": {
		{
			Name:        "transmission-remote",
			Description: "A remote control utility for transmission-daemon and transmission",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-remote_completer/cmd",
			Variant:     "",
			Url:         "https://transmissionbt.com/",
			Execute:     common__transmission_remote.Execute,
		},
	},
	"transmission-show": {
		{
			Name:        "transmission-show",
			Description: "A command-line utility to show .torrent file metadata",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/transmission-show_completer/cmd",
			Variant:     "",
			Url:         "https://transmissionbt.com/",
			Execute:     common__transmission_show.Execute,
		},
	},
	"tree": {
		{
			Name:        "tree",
			Description: "list contents of directories in a tree-like format",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tree_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tree",
			Execute:     common__tree.Execute,
		},
	},
	"truncate": {
		{
			Name:        "truncate",
			Description: "Shrink or extend the size of each FILE to the specified size",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/truncate_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/truncate",
			Execute:     common__truncate.Execute,
		},
	},
	"ts": {
		{
			Name:        "ts",
			Description: "timestamp input",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ts_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/ts",
			Execute:     common__ts.Execute,
		},
	},
	"tsc": {
		{
			Name:        "tsc",
			Description: "The TypeScript Compiler",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tsc_completer/cmd",
			Variant:     "",
			Url:         "https://www.typescriptlang.org/docs/handbook/compiler-options.html",
			Execute:     common__tsc.Execute,
		},
	},
	"tsh": {
		{
			Name:        "tsh",
			Description: "Teleport Command Line Client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tsh_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/gravitational/teleport",
			Execute:     common__tsh.Execute,
		},
	},
	"tshark": {
		{
			Name:        "tshark",
			Description: "Dump and analyze network traffic",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tshark_completer/cmd",
			Variant:     "",
			Url:         "https://tshark.dev/",
			Execute:     common__tshark.Execute,
		},
	},
	"tsort": {
		{
			Name:        "tsort",
			Description: "perform topological sort",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tsort_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tsort",
			Execute:     common__tsort.Execute,
		},
	},
	"tty": {
		{
			Name:        "tty",
			Description: "print the file name of the terminal connected to standard input",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/tty_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/tty",
			Execute:     common__tty.Execute,
		},
	},
	"ttyd": {
		{
			Name:        "ttyd",
			Description: "ttyd is a tool for sharing terminal over the web",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ttyd_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/tsl0922/ttyd",
			Execute:     common__ttyd.Execute,
		},
	},
	"turbo": {
		{
			Name:        "turbo",
			Description: "The build system that makes ship happen",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/turbo_completer/cmd",
			Variant:     "",
			Url:         "https://turbo.build/repo",
			Execute:     common__turbo.Execute,
		},
	},
	"typst": {
		{
			Name:        "typst",
			Description: "A new markup-based typesetting system that is powerful and easy to learn.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/typst_completer/cmd",
			Variant:     "",
			Url:         "https://typst.app/",
			Execute:     common__typst.Execute,
		},
	},
	"ufw": {
		{
			Name:        "ufw",
			Description: "program for managing a netfilter firewall",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/ufw_completer/cmd",
			Variant:     "",
			Url:         "https://launchpad.net/ufw",
			Execute:     common__ufw.Execute,
		},
	},
	"umount": {
		{
			Name:        "umount",
			Description: "Unmount filesystems",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/umount_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/umount",
			Execute:     common__umount.Execute,
		},
	},
	"uname": {
		{
			Name:        "uname",
			Description: "print system information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/uname_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/uname",
			Execute:     common__uname.Execute,
		},
	},
	"unbrotli": {
		{
			Name:        "unbrotli",
			Description: "compress or decompress files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unbrotli_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/google/brotli",
			Execute:     common__unbrotli.Execute,
		},
	},
	"unexpand": {
		{
			Name:        "unexpand",
			Description: "convert spaces to tabs",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unexpand_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/unexpand",
			Execute:     common__unexpand.Execute,
		},
	},
	"uniq": {
		{
			Name:        "uniq",
			Description: "report or omit repeated lines",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/uniq_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/uniq",
			Execute:     common__uniq.Execute,
		},
	},
	"unlink": {
		{
			Name:        "unlink",
			Description: "call the unlink function to remove the specified file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unlink_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/2/unlink",
			Execute:     common__unlink.Execute,
		},
	},
	"unlzma": {
		{
			Name:        "unlzma",
			Description: "Compress or decompress .xz and .lzma files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unlzma_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xz",
			Execute:     common__unlzma.Execute,
		},
	},
	"unpigz": {
		{
			Name:        "unpigz",
			Description: "compress or expand files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unpigz_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/pigz",
			Execute:     common__unpigz.Execute,
		},
	},
	"unset-env": {
		{
			Name:        "unset-env",
			Description: "unset environment variable",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unset-env_completer/cmd",
			Variant:     "",
			Url:         "https://carapace-sh.github.io/carapace-bin/environment.html",
			Execute:     common__unset_env.Execute,
		},
	},
	"unxz": {
		{
			Name:        "unxz",
			Description: "Compress or decompress .xz and .lzma files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unxz_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xz",
			Execute:     common__unxz.Execute,
		},
	},
	"unzip": {
		{
			Name:        "unzip",
			Description: "list, test and extract compressed files in a ZIP archive",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/unzip_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/unzip",
			Execute:     common__unzip.Execute,
		},
	},
	"upower": {
		{
			Name:        "upower",
			Description: "UPower command line tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/upower_completer/cmd",
			Variant:     "",
			Url:         "https://upower.freedesktop.org/",
			Execute:     common__upower.Execute,
		},
	},
	"uptime": {
		{
			Name:        "uptime",
			Description: "Tell how long the system has been running",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/uptime_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/uptime.1.html",
			Execute:     common__uptime.Execute,
		},
	},
	"upx": {
		{
			Name:        "upx",
			Description: "compress or expand executable files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/upx_completer/cmd",
			Variant:     "",
			Url:         "https://upx.github.io/",
			Execute:     common__upx.Execute,
		},
	},
	"useradd": {
		{
			Name:        "useradd",
			Description: "create a new user or update default new user information",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/useradd_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/useradd",
			Execute:     common__useradd.Execute,
		},
	},
	"userdel": {
		{
			Name:        "userdel",
			Description: "delete a user account and related files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/userdel_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/userdel",
			Execute:     common__userdel.Execute,
		},
	},
	"usermod": {
		{
			Name:        "usermod",
			Description: "modify a user account",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/usermod_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/usermod",
			Execute:     common__usermod.Execute,
		},
	},
	"users": {
		{
			Name:        "users",
			Description: "print the user names of users currently logged in to the current host",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/users_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/users",
			Execute:     common__users.Execute,
		},
	},
	"vagrant": {
		{
			Name:        "vagrant",
			Description: "tool for building and managing virtual machine environments",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vagrant_completer/cmd",
			Variant:     "",
			Url:         "https://www.vagrantup.com/",
			Execute:     common__vagrant.Execute,
		},
	},
	"vault": {
		{
			Name:        "vault",
			Description: "A tool for secrets management",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vault_completer/cmd",
			Variant:     "",
			Url:         "https://www.vaultproject.io/",
			Execute:     common__vault.Execute,
		},
	},
	"vdir": {
		{
			Name:        "vdir",
			Description: "list directory contents",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vdir_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/vdir",
			Execute:     common__vdir.Execute,
		},
	},
	"vercel": {
		{
			Name:        "vercel",
			Description: "Develop. Preview. Ship.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vercel_completer/cmd",
			Variant:     "",
			Url:         "https://vercel.com/",
			Execute:     common__vercel.Execute,
		},
	},
	"vhs": {
		{
			Name:        "vhs",
			Description: "Run a given tape file and generates its outputs.",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vhs_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/vhs",
			Execute:     common__vhs.Execute,
		},
	},
	"vi": {
		{
			Name:        "vi",
			Description: "screen oriented (visual) display editor based on ex",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vi_completer/cmd",
			Variant:     "",
			Url:         "https://en.wikipedia.org/wiki/Vi",
			Execute:     common__vi.Execute,
		},
	},
	"viewnior": {
		{
			Name:        "viewnior",
			Description: "simple, fast and elegant image viewer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/viewnior_completer/cmd",
			Variant:     "",
			Url:         "https://siyanpanayotov.com/project/viewnior",
			Execute:     common__viewnior.Execute,
		},
	},
	"vim": {
		{
			Name:        "vim",
			Description: "Vi IMproved, a programmer's text editor",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vim_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/vi",
			Execute:     common__vim.Execute,
		},
	},
	"visudo": {
		{
			Name:        "visudo",
			Description: "safely edit the sudoers file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/visudo_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/8/visudo",
			Execute:     common__visudo.Execute,
		},
	},
	"viu": {
		{
			Name:        "viu",
			Description: "View images right from the terminal",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/viu_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/atanunq/viu",
			Execute:     common__viu.Execute,
		},
	},
	"vivid": {
		{
			Name:        "vivid",
			Description: "LS_COLORS manager with multiple themes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vivid_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/sharkdp/vivid",
			Execute:     common__vivid.Execute,
		},
	},
	"vlc": {
		{
			Name:        "vlc",
			Description: "the VLC media player",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vlc_completer/cmd",
			Variant:     "",
			Url:         "https://www.videolan.org/vlc/",
			Execute:     common__vlc.Execute,
		},
	},
	"vmstat": {
		{
			Name:        "vmstat",
			Description: "Report virtual memory statistics",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/vmstat_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man8/vmstat.8.html",
			Execute:     common__vmstat.Execute,
		},
	},
	"volta": {
		{
			Name:        "volta",
			Description: "The JavaScript Launcher",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/volta_completer/cmd",
			Variant:     "",
			Url:         "https://volta.sh/",
			Execute:     common__volta.Execute,
		},
	},
	"w": {
		{
			Name:        "w",
			Description: "Show who is logged on and what they are doing",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/w_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/w.1.html",
			Execute:     common__w.Execute,
		},
	},
	"watch": {
		{
			Name:        "watch",
			Description: "execute a program periodically, showing output fullscreen",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/watch_completer/cmd",
			Variant:     "",
			Url:         "https://man7.org/linux/man-pages/man1/watch.1.html",
			Execute:     common__watch.Execute,
		},
	},
	"watchexec": {
		{
			Name:        "watchexec",
			Description: "Execute commands when watched files change",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/watchexec_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/watchexec/watchexec",
			Execute:     common__watchexec.Execute,
		},
	},
	"watchgnupg": {
		{
			Name:        "watchgnupg",
			Description: "Read and print logs from a socket",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/watchgnupg_completer/cmd",
			Variant:     "",
			Url:         "https://www.gnupg.org/documentation/manuals/gnupg/watchgnupg.html",
			Execute:     common__watchgnupg.Execute,
		},
	},
	"waypoint": {
		{
			Name:        "waypoint",
			Description: "Easy application deployment for Kubernetes and Amazon ECS",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/waypoint_completer/cmd",
			Variant:     "",
			Url:         "https://www.waypointproject.io/",
			Execute:     common__waypoint.Execute,
		},
	},
	"wc": {
		{
			Name:        "wc",
			Description: "print newline, word, and byte counts for each file",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wc_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/wc",
			Execute:     common__wc.Execute,
		},
	},
	"wezterm": {
		{
			Name:        "wezterm",
			Description: "Wez's Terminal Emulator",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wezterm_completer/cmd",
			Variant:     "",
			Url:         "http://github.com/wez/wezterm",
			Execute:     common__wezterm.Execute,
		},
	},
	"wget": {
		{
			Name:        "wget",
			Description: "a non-interactive network retriever",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wget_completer/cmd",
			Variant:     "",
			Url:         "https://www.gnu.org/software/wget/",
			Execute:     common__wget.Execute,
		},
	},
	"whereis": {
		{
			Name:        "whereis",
			Description: "Locate the binary, source, and manual-page files for a command",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/whereis_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/whereis",
			Execute:     common__whereis.Execute,
		},
	},
	"which": {
		{
			Name:        "which",
			Description: "Write the full path of COMMAND(s) to standard output",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/which_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/which",
			Execute:     common__which.Execute,
		},
	},
	"who": {
		{
			Name:        "who",
			Description: "show who is logged on",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/who_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/who",
			Execute:     common__who.Execute,
		},
	},
	"whoami": {
		{
			Name:        "whoami",
			Description: "print effective userid",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/whoami_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/whoami",
			Execute:     common__whoami.Execute,
		},
	},
	"wine": {
		{
			Name:        "wine",
			Description: "run Windows programs on Unix",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wine_completer/cmd",
			Variant:     "",
			Url:         "https://www.winehq.org/",
			Execute:     common__wine.Execute,
		},
	},
	"wineboot": {
		{
			Name:        "wineboot",
			Description: "perform Wine initialization, startup, and shutdown task",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wineboot_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.winehq.org/Wineboot",
			Execute:     common__wineboot.Execute,
		},
	},
	"winepath": {
		{
			Name:        "winepath",
			Description: "Tool to convert Unix paths to/from Win32 paths",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/winepath_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.winehq.org/Winepath",
			Execute:     common__winepath.Execute,
		},
	},
	"wineserver": {
		{
			Name:        "wineserver",
			Description: "the Wine server",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wineserver_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.winehq.org/Wineserver",
			Execute:     common__wineserver.Execute,
		},
	},
	"winetricks": {
		{
			Name:        "winetricks",
			Description: "manage virtual Windows environments using Wine",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/winetricks_completer/cmd",
			Variant:     "",
			Url:         "https://wiki.winehq.org/Winetricks",
			Execute:     common__winetricks.Execute,
		},
	},
	"wire": {
		{
			Name:        "wire",
			Description: "Compile-time Dependency Injection for Go",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wire_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/google/wire",
			Execute:     common__wire.Execute,
		},
	},
	"wireshark": {
		{
			Name:        "wireshark",
			Description: "Interactively dump and analyze network traffic",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wireshark_completer/cmd",
			Variant:     "",
			Url:         "https://www.wireshark.org/",
			Execute:     common__wireshark.Execute,
		},
	},
	"wishlist": {
		{
			Name:        "wishlist",
			Description: "The SSH Directory",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wishlist_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/charmbracelet/soft-serve",
			Execute:     common__wishlist.Execute,
		},
	},
	"wl-mirror": {
		{
			Name:        "wl-mirror",
			Description: "a simple Wayland output mirror client",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/wl-mirror_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/Ferdi265/wl-mirror",
			Execute:     common__wl_mirror.Execute,
		},
	},
	"woeusb": {
		{
			Name:        "woeusb",
			Description: "A Linux program to create a Windows USB stick installer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/woeusb_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/WoeUSB/WoeUSB-ng",
			Execute:     common__woeusb.Execute,
		},
	},
	"xargs": {
		{
			Name:        "xargs",
			Description: "build and execute command lines from standard input",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xargs_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xargs",
			Execute:     common__xargs.Execute,
		},
	},
	"xbacklight": {
		{
			Name:        "xbacklight",
			Description: "adjust backlight brightness using RandR extension",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xbacklight_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xbacklight",
			Execute:     common__xbacklight.Execute,
		},
	},
	"xclip": {
		{
			Name:        "xclip",
			Description: "command line interface to X selections",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xclip_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/astrand/xclip/",
			Execute:     common__xclip.Execute,
		},
	},
	"xdotool": {
		{
			Name:        "xdotool",
			Description: "command-line X11 automation tool",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xdotool_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/jordansissel/xdotool",
			Execute:     common__xdotool.Execute,
		},
	},
	"xh": {
		{
			Name:        "xh",
			Description: "Friendly and fast tool for sending HTTP requests",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xh_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/ducaale/xh",
			Execute:     common__xh.Execute,
		},
	},
	"xonsh": {
		{
			Name:        "xonsh",
			Description: "Python-powered shell",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xonsh_completer/cmd",
			Variant:     "",
			Url:         "https://xon.sh/",
			Execute:     common__xonsh.Execute,
		},
	},
	"xz": {
		{
			Name:        "xz",
			Description: "Compress or decompress .xz and .lzma files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xz_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xz",
			Execute:     common__xz.Execute,
		},
	},
	"xzcat": {
		{
			Name:        "xzcat",
			Description: "Compress or decompress .xz and .lzma files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/xzcat_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/xz",
			Execute:     common__xzcat.Execute,
		},
	},
	"yarn": {
		{
			Name:        "yarn",
			Description: "Yarn is a package manager that doubles down as project manager",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/yarn_completer/cmd",
			Variant:     "",
			Url:         "https://yarnpkg.com/",
			Execute:     common__yarn.Execute,
		},
	},
	"yay": {
		{
			Name:        "yay",
			Description: "An AUR Helper written in Go",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/yay_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/Jguer/yay",
			Execute:     common__yay.Execute,
		},
	},
	"yes": {
		{
			Name:        "yes",
			Description: "output a string repeatedly until killed",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/yes_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/yes",
			Execute:     common__yes.Execute,
		},
	},
	"yj": {
		{
			Name:        "yj",
			Description: "Convert between YAML, TOML, JSON, and HCL",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/yj_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/sclevine/yj",
			Execute:     common__yj.Execute,
		},
	},
	"youtube-dl": {
		{
			Name:        "youtube-dl",
			Description: "download videos from youtube.com or other video platforms",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/youtube-dl_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/ytdl-org/youtube-dl",
			Execute:     common__youtube_dl.Execute,
		},
	},
	"yt-dlp": {
		{
			Name:        "yt-dlp",
			Description: "A youtube-dl fork with additional features and fixes",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/yt-dlp_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/yt-dlp/yt-dlp",
			Execute:     common__yt_dlp.Execute,
		},
	},
	"zathura": {
		{
			Name:        "zathura",
			Description: "a document viewer",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/zathura_completer/cmd",
			Variant:     "",
			Url:         "https://pwmt.org/projects/zathura/",
			Execute:     common__zathura.Execute,
		},
	},
	"zcat": {
		{
			Name:        "zcat",
			Description: "compress or expand files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/zcat_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/zcat",
			Execute:     common__zcat.Execute,
		},
	},
	"zip": {
		{
			Name:        "zip",
			Description: "package and compress (archive) files",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/zip_completer/cmd",
			Variant:     "",
			Url:         "https://linux.die.net/man/1/zip",
			Execute:     common__zip.Execute,
		},
	},
	"zoxide": {
		{
			Name:        "zoxide",
			Description: "A smarter cd command for your terminal",
			Group:       "common",
			Package:     "github.com/carapace-sh/carapace-bin/completers_release/common/zoxide_completer/cmd",
			Variant:     "",
			Url:         "https://github.com/ajeetdsouza/zoxide",
			Execute:     common__zoxide.Execute,
		},
	}}
