# Completer backlog

Work queue for commands that have a fish completion but no carapace completer.
This file is the tracker: update a row when you start or finish it.

## How to work this file

1. Pick the first `todo` row in Queue A (or one group of Queue B).
2. Confirm the approach:
   - `scrape:*`: the upstream CLI uses a supported framework. Add a scraper to
     `../scrape` (see `../scrape/AGENTS.md`), then codegen the skeleton.
   - `manual`: the CLI parses arguments by hand. Read its source or man page
     and write the completer directly.
   - `?`: framework not verified. Check the upstream `go.mod` / `Cargo.toml` /
     `pyproject.toml` / `package.json` first; that decides scrape vs manual.
3. Implement in `completers/<group>/<command>_completer`, with every flag and
   subcommand taken from the current upstream release, never from the fish file.
4. Verify (all four must pass):

   ```sh
   go generate ./cmd/...
   go build -tags force_all ./...
   go test -count=1 ./cmd/...
   ./custom-gcl run ./completers/<group>/<command>_completer/...
   ```

5. Update the row: `todo` -> `wip (<branch>)` -> `done (<commit>)`. Add a note
   when the result deviates from the plan (for example a bridge instead).

## Columns

- `size`: fish completion size in KB. Proxy for CLI surface, not popularity;
  use it to order work, not to rank importance.
- `approach`: as described above.
- `source`: upstream repository or documentation to verify the CLI against.
- `status`: `todo`, `wip (<branch>)`, `done (<commit>)`, `skip (<reason>)`.

## Batch 1 (done)

| command | approach | source | status |
|---|---|---|---|
| kind | scrape:cobra | kubernetes-sigs/kind v0.33.0 | done (723a86bcf) |
| skaffold | scrape:cobra | GoogleContainerTools/skaffold v2.25.0 | done (723a86bcf) |
| ruff | scrape:clap | astral-sh/ruff 0.17.3 | done (723a86bcf) |
| pre-commit | scrape:argparse | pre-commit/pre-commit v4.7.0 | done (723a86bcf) |
| psql | manual | PostgreSQL 18 (psql help.c/command.c) | done (0ad345b77) |
| iwctl | manual | iwd 3.12 (client/*.c) | done (0ad345b77) |

Scrapers for batch 1 live in `../scrape` on `feat/completers-batch-1` (da99f55):
`scrapers/{kind,skaffold,ruff,pre-commit}` plus their `package.json` entries.

## Batch 2 (done)

| command | approach | source | status |
|---|---|---|---|
| meson | scrape:argparse | mesonbuild/meson 1.12.1 | done (fa2da5914) |
| llm | scrape:click | simonw/llm 0.36 | done (fa2da5914) |

Scrapers for batch 2 live in `../scrape` on `feat/completers-batch-2` (4e61f8e):
`scrapers/{meson,llm}`. `scrapers/llm` is the first click scraper and can be
reused for other click CLIs by pointing it at another package.

## Suggested next batch

In this order:

1. `btrfs` (manual, mirror `pkg/actions/tools/zfs`). The installed binary can
   dump its own help (`btrfs --help`, `btrfs <group> <cmd> --help`); parse that
   into the command tree and then verify the option names against the current
   upstream `cmds/*.c`.
2. `svn` (manual; subcommands and options from `svn help <cmd>`).
3. `firewall-cmd` (argparse) and `waydroid` (argparse) - both scrapers are
   close to the existing argparse recipe; confirm the parser is reachable.
4. `sfdx` and `heroku` (oclif) - requires the oclif carapace spec plugin, which
   is not wired into `../scrape` yet; verify the plugin still works for these
   versions before adding the scrapers.

## Queue A: fish surface >= 4 KB (127 commands)

| # | command | size | approach | source | status |
|---|---|---|---|---|---|
| 1 | `sfdx` | 114 KB | scrape:oclif | forcedotcom/cli | todo |
| 2 | `virsh` | 113 KB | manual | libvirt (virsh help / docs) | todo |
| 3 | `gcc` | 101 KB | manual | gcc --help, gcc docs (stable flag subset) | todo |
| 4 | `darcs` | 92 KB | manual | darcs/darcs (darcs help) | todo |
| 5 | `zypper` | 52 KB | manual | openSUSE/zypper | todo |
| 6 | `rpm-ostree` | 44 KB | manual | coreos/rpm-ostree | todo |
| 7 | `btrfs` | 41 KB | manual | kdave/btrfs-progs, mirror pkg/actions/tools/zfs | todo |
| 8 | `p4` | 41 KB | manual | Perforce Helix docs | todo |
| 9 | `magento` | 35 KB | manual | magento/magento2 bin/magento | todo |
| 10 | `cvs` | 32 KB | manual | cvs --help-commands | todo |
| 11 | `fossil` | 31 KB | manual | fossil-scm/fossil (fossil help) | todo |
| 12 | `meson` | 27 KB | scrape:argparse | mesonbuild/meson 1.12.1 | done (fa2da5914) |
| 13 | `quilt` | 25 KB | manual | quilt(1) man page | todo |
| 14 | `arc` | 22 KB | manual | phacility/arcanist | todo |
| 15 | `firewall-cmd` | 21 KB | scrape:argparse | firewalld/firewalld | todo |
| 16 | `mdadm` | 18 KB | manual | mdadm(8) man page | todo |
| 17 | `nvme` | 17 KB | manual | linux-nvme/nvme-cli | todo |
| 18 | `cf` | 17 KB | ? | cloudfoundry/cli (check go.mod) | todo |
| 19 | `wajig` | 16 KB | manual | Debian wajig docs | todo |
| 20 | `hledger` | 16 KB | manual | simonmichael/hledger (optparse-applicative) | todo |
| 21 | `emerge` | 16 KB | manual | gentoo/portage | todo |
| 22 | `svn` | 15 KB | manual | apache/subversion (svn help) | todo |
| 23 | `pwn` | 15 KB | ? | pwntools (check entry point) | todo |
| 24 | `sops` | 15 KB | manual | getsops/sops | todo |
| 25 | `apm` | 14 KB | skip | - | skip (Atom is retired) |
| 26 | `keybase` | 14 KB | manual | keybase/client | todo |
| 27 | `iw` | 12 KB | manual | kernel wireless iw | todo |
| 28 | `s3cmd` | 12 KB | scrape:argparse | s3tools/s3cmd | todo |
| 29 | `obnam` | 12 KB | ? | obnam/obnam | todo |
| 30 | `bundle` | 12 KB | manual | rubygems/bundler (bundle help) | todo |
| 31 | `yaourt` | 12 KB | skip | - | skip (AUR helper, retired) |
| 32 | `gem` | 12 KB | manual | rubygems/rubygems (gem help commands) | todo |
| 33 | `aura` | 11 KB | manual | fosskers/aura (Haskell) | todo |
| 34 | `hashcat` | 11 KB | manual | hashcat/hashcat --help | todo |
| 35 | `heroku` | 11 KB | scrape:oclif | heroku/cli | todo |
| 36 | `snap` | 11 KB | manual | canonical/snapd | todo |
| 37 | `prt-get` | 11 KB | manual | CRUX prt-get(8) | todo |
| 38 | `blender` | 11 KB | ? | blender/blender (check argparse) | todo |
| 39 | `pacaur` | 11 KB | skip | - | skip (AUR helper, retired) |
| 40 | `waydroid` | 11 KB | scrape:argparse | waydroid/waydroid | todo |
| 41 | `xterm` | 11 KB | manual | xterm(1) man page | todo |
| 42 | `apkanalyzer` | 11 KB | manual | Android SDK tools | todo |
| 43 | `mono` | 10 KB | manual | mono/mono | todo |
| 44 | `moss` | 10 KB | ? | check upstream (Rust?) | todo |
| 45 | `nmap` | 10 KB | manual | nmap/nmap --help | todo |
| 46 | `sysbench` | 10 KB | manual | akopytov/sysbench | todo |
| 47 | `opam` | 10 KB | manual | ocaml/opam (cmdliner) | todo |
| 48 | `cjpm` | 9 KB | ? | Cangjie tooling | todo |
| 49 | `kitchen` | 9 KB | manual | test-kitchen/test-kitchen | todo |
| 50 | `machinectl` | 9 KB | manual | systemd (reuse pkg/actions/tools/systemd) | todo |
| 51 | `bzr` | 9 KB | skip | - | skip (Bazaar, retired) |
| 52 | `equery` | 9 KB | manual | gentoo/gentoolkit | todo |
| 53 | `rpm` | 9 KB | manual | rpm-software-management/rpm | todo |
| 54 | `csc` | 9 KB | manual | mono/mono csc | todo |
| 55 | `homectl` | 9 KB | manual | systemd | todo |
| 56 | `vbc` | 9 KB | manual | mono/mono vbc | todo |
| 57 | `fail2ban-client` | 8 KB | manual | fail2ban/fail2ban | todo |
| 58 | `mksquashfs` | 8 KB | manual | plougher/squashfs-tools | todo |
| 59 | `clasp` | 8 KB | ? | google/clasp (check framework) | todo |
| 60 | `distrobox` | 8 KB | manual | 89luca89/distrobox | todo |
| 61 | `cryptsetup` | 8 KB | manual | cryptsetup(8) man page | todo |
| 62 | `display` | 8 KB | manual | ImageMagick (sibling of existing convert completer) | todo |
| 63 | `mpc` | 8 KB | manual | Music Player Daemon client | todo |
| 64 | `john` | 8 KB | manual | openwall/john | todo |
| 65 | `zef` | 8 KB | manual | Raku zef | todo |
| 66 | `gio` | 8 KB | manual | GNOME glib | todo |
| 67 | `llm` | 8 KB | scrape:click | simonw/llm | done (fa2da5914) |
| 68 | `qubes-gpg-client` | 7 KB | manual | Qubes OS | todo |
| 69 | `cdrecord` | 7 KB | skip | - | skip (cdrkit, retired) |
| 70 | `yadm` | 7 KB | manual | yadm/yadm | todo |
| 71 | `gnome-extensions` | 7 KB | manual | GNOME Shell | todo |
| 72 | `readelf` | 7 KB | manual | binutils | todo |
| 73 | `phpunit` | 7 KB | manual | sebastianbergmann/phpunit | todo |
| 74 | `chronyc` | 7 KB | manual | chrony | todo |
| 75 | `pactl` | 7 KB | manual | pulseaudio | todo |
| 76 | `roswell` | 7 KB | manual | roswell/roswell | todo |
| 77 | `as` | 7 KB | manual | binutils | todo |
| 78 | `jest` | 7 KB | manual | jestjs/jest (yargs) | todo |
| 79 | `mysql` | 7 KB | manual | mysql/mysql-server | todo |
| 80 | `bosh` | 7 KB | ? | cloudfoundry/bosh-cli (check go.mod) | todo |
| 81 | `busctl` | 6 KB | manual | systemd | todo |
| 82 | `iptables` | 6 KB | manual | netfilter iptables(8) | todo |
| 83 | `cpupower` | 6 KB | manual | linux-tools cpupower | todo |
| 84 | `keepassxc-cli` | 6 KB | manual | keepassxreboot/keepassxc | todo |
| 85 | `mkvextract` | 6 KB | manual | mkvtoolnix | todo |
| 86 | `toot` | 6 KB | ? | ihabunek/toot (check framework) | todo |
| 87 | `valgrind` | 6 KB | manual | valgrind --help | todo |
| 88 | `aptitude` | 6 KB | manual | Debian aptitude | todo |
| 89 | `flac` | 6 KB | manual | xiph/flac | todo |
| 90 | `yubico-piv-tool` | 6 KB | manual | Yubico | todo |
| 91 | `import` | 6 KB | manual | ImageMagick | todo |
| 92 | `fsharpc` | 6 KB | manual | mono/mono fsharpc | todo |
| 93 | `opkg` | 6 KB | manual | openwrt opkg | todo |
| 94 | `strace` | 6 KB | manual | strace/strace --help | todo |
| 95 | `src` | 6 KB | ? | verify upstream | todo |
| 96 | `scss` | 6 KB | manual | sass/dart-sass | todo |
| 97 | `sass` | 6 KB | manual | sass/dart-sass | todo |
| 98 | `udevadm` | 5 KB | manual | systemd | todo |
| 99 | `climate` | 5 KB | ? | verify upstream | todo |
| 100 | `animate` | 5 KB | manual | ImageMagick | todo |
| 101 | `fastboot` | 5 KB | manual | Android platform tools | todo |
| 102 | `ack` | 5 KB | manual | beyondgrep/ack | todo |
| 103 | `latexmk` | 5 KB | manual | latexmk(1) man page | todo |
| 104 | `gphoto2` | 5 KB | manual | gphoto/gphoto2 | todo |
| 105 | `dhcpcd` | 5 KB | manual | dhcpcd(8) man page | todo |
| 106 | `checkinstall` | 5 KB | manual | checkinstall | todo |
| 107 | `objdump` | 5 KB | manual | binutils | todo |
| 108 | `xrandr` | 5 KB | manual | xrandr(1) man page | todo |
| 109 | `lxc` | 5 KB | manual | lxc/lxc | todo |
| 110 | `cwebp` | 5 KB | manual | webmproject/libwebp | todo |
| 111 | `connmanctl` | 5 KB | manual | connman | todo |
| 112 | `pg_dump` | 5 KB | manual | PostgreSQL (extend psql_completer pattern) | todo |
| 113 | `travis` | 5 KB | skip | - | skip (Travis CLI, retired) |
| 114 | `rbenv` | 4 KB | manual | rbenv/rbenv | todo |
| 115 | `configure` | 4 KB | manual | autoconf (generic configure flags) | todo |
| 116 | `btrbk` | 4 KB | manual | digint/btrbk | todo |
| 117 | `gsettings` | 4 KB | manual | GNOME glib | todo |
| 118 | `udisksctl` | 4 KB | manual | storaged-project/udisks | todo |
| 119 | `pyenv` | 4 KB | manual | pyenv/pyenv | todo |
| 120 | `portmaster` | 4 KB | manual | FreeBSD ports | todo |
| 121 | `guile` | 4 KB | manual | guile scheme | todo |
| 122 | `msfvenom` | 4 KB | manual | rapid7/metasploit-framework | todo |
| 123 | `tmuxp` | 4 KB | manual | tmux-python/tmuxp | todo |
| 124 | `wpctl` | 4 KB | manual | pipewire | todo |
| 125 | `fab` | 4 KB | manual | fabric/fabric | todo |
| 126 | `ncat` | 4 KB | manual | nmap/nmap | todo |
| 127 | `resolvectl` | 4 KB | manual | systemd | todo |

## Queue B: smaller surface (515 commands)

Grouped by domain so one session can take a whole category. All rows are
`todo`; append ` -> wip (<branch>)` or ` -> done (<commit>)` when you touch one.

### Version control (7)

- [ ] `sv` (4 KB)
- [ ] `git-subtree` (3 KB)
- [ ] `git-sizer` (1 KB)
- [ ] `smerge` (1 KB)
- [ ] `rejmerge` (1 KB)
- [ ] `git-lfs` (1 KB)
- [ ] `colorsvn` (1 KB)

### Package managers (25)

- [ ] `sbt` (4 KB)
- [ ] `lein` (3 KB)
- [ ] `yum` (3 KB)
- [ ] `eix` (3 KB)
- [ ] `cpan` (3 KB)
- [ ] `eselect` (3 KB)
- [ ] `emaint` (3 KB)
- [ ] `ebuild` (2 KB)
- [ ] `bower` (2 KB)
- [ ] `topgrade` (2 KB)
- [ ] `rakudo` (2 KB)
- [ ] `apt-mark` (2 KB)
- [ ] `stow` (2 KB)
- [ ] `apt-file` (1 KB)
- [ ] `pkgfile` (1 KB)
- [ ] `pkg_add` (1 KB)
- [ ] `dlocate` (1 KB)
- [ ] `cabal` (1 KB)
- [ ] `pkg_delete` (1 KB)
- [ ] `stack` (1 KB)
- [ ] `pip2` (1 KB)
- [ ] `pip3` (1 KB)
- [ ] `pkg_info` (1 KB)
- [ ] `composer.phar` (1 KB)
- [ ] `gradlew` (1 KB)

### Build and compiler tooling (20)

- [ ] `cjxl` (4 KB)
- [ ] `scons` (4 KB)
- [ ] `stream` (3 KB)
- [ ] `ttx` (2 KB)
- [ ] `avifenc` (2 KB)
- [ ] `gprof` (2 KB)
- [ ] `oggenc` (2 KB)
- [ ] `nm` (2 KB)
- [ ] `m4` (2 KB)
- [ ] `ninja` (1 KB)
- [ ] `vips` (1 KB)
- [ ] `bison` (1 KB)
- [ ] `djxl` (1 KB)
- [ ] `conjure` (1 KB)
- [ ] `avifdec` (1 KB)
- [ ] `clang` (1 KB)
- [ ] `makedepend` (1 KB)
- [ ] `clang++` (1 KB)
- [ ] `gdbserver` (1 KB)
- [ ] `g++` (1 KB)

### Virtualization, containers, orchestration (6)

- [ ] `systemd-nspawn` (4 KB)
- [ ] `flux` (1 KB)
- [ ] `devspace` (1 KB)
- [ ] `cilium` (1 KB)
- [ ] `hubble` (1 KB)
- [ ] `istioctl` (1 KB)

### System, hardware, storage (23)

- [ ] `ethtool` (4 KB)
- [ ] `networkctl` (4 KB)
- [ ] `duply` (3 KB)
- [ ] `loginctl` (3 KB)
- [ ] `systemd-cryptenroll` (3 KB)
- [ ] `alsactl` (3 KB)
- [ ] `losetup` (3 KB)
- [ ] `setxkbmap` (2 KB)
- [ ] `pv` (2 KB)
- [ ] `timedatectl` (2 KB)
- [ ] `hostnamectl` (2 KB)
- [ ] `mkdosfs` (2 KB)
- [ ] `userdbctl` (2 KB)
- [ ] `xinput` (2 KB)
- [ ] `brightnessctl` (1 KB)
- [ ] `dmidecode` (1 KB)
- [ ] `pacmd` (1 KB)
- [ ] `badblocks` (1 KB)
- [ ] `tuned-adm` (1 KB)
- [ ] `amixer` (1 KB)
- [ ] `rfkill` (1 KB)
- [ ] `synclient` (1 KB)
- [ ] `mkfs.fat` (1 KB)

### Networking, firewall, remote access (14)

- [ ] `arping` (4 KB)
- [ ] `arp-scan` (3 KB)
- [ ] `dhclient` (3 KB)
- [ ] `mtr` (3 KB)
- [ ] `create_ap` (3 KB)
- [ ] `ipset` (2 KB)
- [ ] `caddy` (2 KB)
- [ ] `iftop` (2 KB)
- [ ] `wg-quick` (1 KB)
- [ ] `dumpcap` (1 KB)
- [ ] `sshfs` (1 KB)
- [ ] `tracepath` (1 KB)
- [ ] `proxychains` (1 KB)
- [ ] `ngrok` (1 KB)

### Desktop, media, productivity (21)

- [ ] `emacs` (4 KB)
- [ ] `mplayer` (3 KB)
- [ ] `scanimage` (3 KB)
- [ ] `dconf` (2 KB)
- [ ] `kdeconnect-cli` (2 KB)
- [ ] `gdbus` (2 KB)
- [ ] `xdg-mime` (2 KB)
- [ ] `subl` (2 KB)
- [ ] `atom` (1 KB)
- [ ] `notify-send` (1 KB)
- [ ] `qdbus` (1 KB)
- [ ] `evince` (1 KB)
- [ ] `gapplication` (1 KB)
- [ ] `tmuxinator` (1 KB)
- [ ] `i3-msg` (1 KB)
- [ ] `mupdf` (1 KB)
- [ ] `gvim` (1 KB)
- [ ] `okular` (1 KB)
- [ ] `djview4` (1 KB)
- [ ] `djview` (1 KB)
- [ ] `vimdiff` (1 KB)

### Security (9)

- [ ] `firejail` (3 KB)
- [ ] `signify` (3 KB)
- [ ] `msfdb` (2 KB)
- [ ] `setfacl` (2 KB)
- [ ] `msfconsole` (2 KB)
- [ ] `searchsploit` (2 KB)
- [ ] `age-keygen` (1 KB)
- [ ] `gpg1` (1 KB)
- [ ] `gpg2` (1 KB)

### Languages and runtimes (27)

- [ ] `fsharpi` (4 KB)
- [ ] `mocha` (3 KB)
- [ ] `clj` (3 KB)
- [ ] `monodis` (3 KB)
- [ ] `python3` (3 KB)
- [ ] `perl` (3 KB)
- [ ] `jhipster` (3 KB)
- [ ] `octave` (3 KB)
- [ ] `matlab` (3 KB)
- [ ] `java_home` (2 KB)
- [ ] `elm` (2 KB)
- [ ] `godoc` (2 KB)
- [ ] `irb` (2 KB)
- [ ] `python2` (2 KB)
- [ ] `grunt` (1 KB)
- [ ] `sbcl` (1 KB)
- [ ] `ruby` (1 KB)
- [ ] `iex` (1 KB)
- [ ] `elixir` (1 KB)
- [ ] `monop` (1 KB)
- [ ] `ruby-build` (1 KB)
- [ ] `gorename` (1 KB)
- [ ] `yo` (1 KB)
- [ ] `spago` (1 KB)
- [ ] `golint` (1 KB)
- [ ] `uvx` (1 KB)
- [ ] `clojure` (1 KB)

### Documentation, data and publishing (37)

- [ ] `mkdocs` (4 KB)
- [ ] `pg_restore` (4 KB)
- [ ] `pg_dumpall` (4 KB)
- [ ] `sphinx-quickstart` (2 KB)
- [ ] `asciidoctor` (2 KB)
- [ ] `gojq` (2 KB)
- [ ] `sphinx-apidoc` (2 KB)
- [ ] `sphinx-build` (2 KB)
- [ ] `createdb` (1 KB)
- [ ] `rst2odt` (1 KB)
- [ ] `whatis` (1 KB)
- [ ] `hjson` (1 KB)
- [ ] `tex` (1 KB)
- [ ] `dropdb` (1 KB)
- [ ] `csvlens` (1 KB)
- [ ] `rst2s5` (1 KB)
- [ ] `ijq` (1 KB)
- [ ] `cmark` (1 KB)
- [ ] `md-to-clip` (1 KB)
- [ ] `rst2xml` (1 KB)
- [ ] `sphinx-autogen` (1 KB)
- [ ] `rstpep2html` (1 KB)
- [ ] `rst2latex` (1 KB)
- [ ] `rst2xetex` (1 KB)
- [ ] `rst2html5` (1 KB)
- [ ] `rst2html` (1 KB)
- [ ] `rst2man` (1 KB)
- [ ] `rst2pseudoxml` (1 KB)
- [ ] `tex-fmt` (1 KB)
- [ ] `rst2html4` (1 KB)
- [ ] `hexyl` (1 KB)
- [ ] `mysqldump` (1 KB)
- [ ] `xelatex` (1 KB)
- [ ] `lualatex` (1 KB)
- [ ] `pdflatex` (1 KB)
- [ ] `pdftex` (1 KB)
- [ ] `latex` (1 KB)

### Text editors and shells (4)

- [ ] `tcsh` (1 KB)
- [ ] `csh` (1 KB)
- [ ] `ksh` (1 KB)
- [ ] `oksh` (1 KB)

### Other (322)

- [ ] `netctl-auto` (4 KB)
- [ ] `mkbundle` (3 KB)
- [ ] `sass-convert` (3 KB)
- [ ] `gv` (3 KB)
- [ ] `lpadmin` (3 KB)
- [ ] `dotslash` (3 KB)
- [ ] `xgettext` (3 KB)
- [ ] `ldapsearch` (3 KB)
- [ ] `zabbix_server` (3 KB)
- [ ] `subfinder` (3 KB)
- [ ] `efibootmgr` (3 KB)
- [ ] `grub-install` (3 KB)
- [ ] `ezjail-admin` (3 KB)
- [ ] `vim-addons` (3 KB)
- [ ] `grub-file` (3 KB)
- [ ] `usbip` (3 KB)
- [ ] `bootctl` (3 KB)
- [ ] `mariner` (3 KB)
- [ ] `qjsc` (3 KB)
- [ ] `wpa_cli` (3 KB)
- [ ] `asp` (2 KB)
- [ ] `cygport` (2 KB)
- [ ] `mocp` (2 KB)
- [ ] `wlr-randr` (2 KB)
- [ ] `zstd` (2 KB)
- [ ] `pkg-config` (2 KB)
- [ ] `msgfmt` (2 KB)
- [ ] `VBoxSDL` (2 KB)
- [ ] `zstdmt` (2 KB)
- [ ] `reflector` (2 KB)
- [ ] `navi` (2 KB)
- [ ] `cygpath` (2 KB)
- [ ] `wesnoth` (2 KB)
- [ ] `xpdf` (2 KB)
- [ ] `alternatives` (2 KB)
- [ ] `castnow` (2 KB)
- [ ] `zabbix_sender` (2 KB)
- [ ] `adduser` (2 KB)
- [ ] `guild` (2 KB)
- [ ] `nodeenv` (2 KB)
- [ ] `mkinitcpio` (2 KB)
- [ ] `xsp` (2 KB)
- [ ] `grub-mkrescue` (2 KB)
- [ ] `exif` (2 KB)
- [ ] `mutt` (2 KB)
- [ ] `netctl` (2 KB)
- [ ] `xsv` (2 KB)
- [ ] `duf` (2 KB)
- [ ] `kldload` (2 KB)
- [ ] `konsole` (2 KB)
- [ ] `atool` (2 KB)
- [ ] `winemaker` (2 KB)
- [ ] `tflint` (2 KB)
- [ ] `fcitx5-remote` (2 KB)
- [ ] `protontricks` (2 KB)
- [ ] `zabbix_proxy` (2 KB)
- [ ] `windscribe` (2 KB)
- [ ] `scrypt` (2 KB)
- [ ] `cygstart` (2 KB)
- [ ] `apt-listbugs` (2 KB)
- [ ] `figlet` (2 KB)
- [ ] `lz4` (2 KB)
- [ ] `apt-build` (2 KB)
- [ ] `krita` (2 KB)
- [ ] `lunchy` (2 KB)
- [ ] `locate` (2 KB)
- [ ] `makensis` (2 KB)
- [ ] `attributes` (2 KB)
- [ ] `t-rec` (2 KB)
- [ ] `abook` (2 KB)
- [ ] `ifdata` (2 KB)
- [ ] `pkginfo` (2 KB)
- [ ] `eix-sync` (2 KB)
- [ ] `tokei` (2 KB)
- [ ] `unzstd` (2 KB)
- [ ] `pstree` (2 KB)
- [ ] `pygmentize` (2 KB)
- [ ] `mc` (2 KB)
- [ ] `zabbix_get` (2 KB)
- [ ] `openocd` (1 KB)
- [ ] `loadkeys` (1 KB)
- [ ] `icdiff` (1 KB)
- [ ] `service` (1 KB)
- [ ] `apt-src` (1 KB)
- [ ] `ouch` (1 KB)
- [ ] `setx` (1 KB)
- [ ] `gacutil` (1 KB)
- [ ] `cmdkey` (1 KB)
- [ ] `xviewer` (1 KB)
- [ ] `apt-move` (1 KB)
- [ ] `totem` (1 KB)
- [ ] `zabbix_agentd` (1 KB)
- [ ] `dput` (1 KB)
- [ ] `unrar` (1 KB)
- [ ] `lpstat` (1 KB)
- [ ] `canto` (1 KB)
- [ ] `xrdb` (1 KB)
- [ ] `nethack` (1 KB)
- [ ] `optimus-manager` (1 KB)
- [ ] `web-ext` (1 KB)
- [ ] `pkgmk` (1 KB)
- [ ] `xsel` (1 KB)
- [ ] `wvdial` (1 KB)
- [ ] `zstdcat` (1 KB)
- [ ] `apt-ftparchive` (1 KB)
- [ ] `xprop` (1 KB)
- [ ] `pzstd` (1 KB)
- [ ] `pydf` (1 KB)
- [ ] `apt-show-versions` (1 KB)
- [ ] `ibus` (1 KB)
- [ ] `actionlint` (1 KB)
- [ ] `qjs` (1 KB)
- [ ] `xplayer` (1 KB)
- [ ] `gendarme` (1 KB)
- [ ] `encfs` (1 KB)
- [ ] `efivar` (1 KB)
- [ ] `pine` (1 KB)
- [ ] `trash` (1 KB)
- [ ] `xed` (1 KB)
- [ ] `apt-spy` (1 KB)
- [ ] `invoke-rc.d` (1 KB)
- [ ] `wicd-cli` (1 KB)
- [ ] `VBoxHeadless` (1 KB)
- [ ] `unlz4` (1 KB)
- [ ] `apt-rdepends` (1 KB)
- [ ] `apt-listchanges` (1 KB)
- [ ] `termite` (1 KB)
- [ ] `lpinfo` (1 KB)
- [ ] `cascadia` (1 KB)
- [ ] `dm-tool` (1 KB)
- [ ] `update-eix-remote` (1 KB)
- [ ] `update-eix` (1 KB)
- [ ] `protontricks-launch` (1 KB)
- [ ] `sylpheed` (1 KB)
- [ ] `udevil` (1 KB)
- [ ] `zopflipng` (1 KB)
- [ ] `ark` (1 KB)
- [ ] `trash-empty` (1 KB)
- [ ] `apt-zip-list` (1 KB)
- [ ] `cleanmgr` (1 KB)
- [ ] `trash-put` (1 KB)
- [ ] `logkeys` (1 KB)
- [ ] `choose` (1 KB)
- [ ] `dpkg-reconfigure` (1 KB)
- [ ] `pactree` (1 KB)
- [ ] `gresource` (1 KB)
- [ ] `epkginfo` (1 KB)
- [ ] `xmms` (1 KB)
- [ ] `apt-cdrom` (1 KB)
- [ ] `bzip2` (1 KB)
- [ ] `jv` (1 KB)
- [ ] `kcmshell5` (1 KB)
- [ ] `apt-show-source` (1 KB)
- [ ] `sad` (1 KB)
- [ ] `cowthink` (1 KB)
- [ ] `mddiagnose` (1 KB)
- [ ] `mkpasswd` (1 KB)
- [ ] `lpoptions` (1 KB)
- [ ] `cowsay` (1 KB)
- [ ] `xxh64sum` (1 KB)
- [ ] `fortune` (1 KB)
- [ ] `lp` (1 KB)
- [ ] `pabcnetcclear` (1 KB)
- [ ] `zed` (1 KB)
- [ ] `jnv` (1 KB)
- [ ] `eg` (1 KB)
- [ ] `apt-proxy-import` (1 KB)
- [ ] `expect` (1 KB)
- [ ] `ilasm` (1 KB)
- [ ] `lz4cat` (1 KB)
- [ ] `zabbix_agent2` (1 KB)
- [ ] `rc` (1 KB)
- [ ] `setsid` (1 KB)
- [ ] `csi` (1 KB)
- [ ] `gping` (1 KB)
- [ ] `yash` (1 KB)
- [ ] `yast2` (1 KB)
- [ ] `entr` (1 KB)
- [ ] `trash-list` (1 KB)
- [ ] `serve` (1 KB)
- [ ] `es` (1 KB)
- [ ] `namei` (1 KB)
- [ ] `horcrux` (1 KB)
- [ ] `zabbix_js` (1 KB)
- [ ] `vmctl` (1 KB)
- [ ] `apt-zip-inst` (1 KB)
- [ ] `bd` (1 KB)
- [ ] `xreader` (1 KB)
- [ ] `calendar` (1 KB)
- [ ] `poff` (1 KB)
- [ ] `root` (1 KB)
- [ ] `trash-restore` (1 KB)
- [ ] `bunzip2` (1 KB)
- [ ] `combine` (1 KB)
- [ ] `ikdasm` (1 KB)
- [ ] `dolphin` (1 KB)
- [ ] `pbget` (1 KB)
- [ ] `jbake` (1 KB)
- [ ] `wish` (1 KB)
- [ ] `zopfli` (1 KB)
- [ ] `yajsv` (1 KB)
- [ ] `fusermount` (1 KB)
- [ ] `camcontrol` (1 KB)
- [ ] `j` (1 KB)
- [ ] `ports` (1 KB)
- [ ] `wicd-client` (1 KB)
- [ ] `apt-config` (1 KB)
- [ ] `wicd-gtk` (1 KB)
- [ ] `kldunload` (1 KB)
- [ ] `apt-sortpkgs` (1 KB)
- [ ] `lxpanel` (1 KB)
- [ ] `rcctl` (1 KB)
- [ ] `wslpath` (1 KB)
- [ ] `kb` (1 KB)
- [ ] `airmon-ng` (1 KB)
- [ ] `pkgadd` (1 KB)
- [ ] `block` (1 KB)
- [ ] `pix` (1 KB)
- [ ] `batsh` (1 KB)
- [ ] `lppasswd` (1 KB)
- [ ] `apt-extracttemplates` (1 KB)
- [ ] `hikari` (1 KB)
- [ ] `pkgrm` (1 KB)
- [ ] `zabbix_web_service` (1 KB)
- [ ] `csharp` (1 KB)
- [ ] `native2ascii` (1 KB)
- [ ] `isatty` (1 KB)
- [ ] `atd` (1 KB)
- [ ] `a2disconf` (1 KB)
- [ ] `a2dissite` (1 KB)
- [ ] `prime-run` (1 KB)
- [ ] `a2dismod` (1 KB)
- [ ] `test-spice` (1 KB)
- [ ] `sqlsharp` (1 KB)
- [ ] `apt-key` (1 KB)
- [ ] `container` (1 KB)
- [ ] `unbuffer` (1 KB)
- [ ] `a2enconf` (1 KB)
- [ ] `a2ensite` (1 KB)
- [ ] `a2enmod` (1 KB)
- [ ] `bzcat` (1 KB)
- [ ] `apt-setup` (1 KB)
- [ ] `aunpack` (1 KB)
- [ ] `cancel` (1 KB)
- [ ] `ifdown` (1 KB)
- [ ] `asd` (1 KB)
- [ ] `acat` (1 KB)
- [ ] `ifup` (1 KB)
- [ ] `als` (1 KB)
- [ ] `atq` (1 KB)
- [ ] `ibmcloud` (1 KB)
- [ ] `clean` (1 KB)
- [ ] `bzip2recover` (1 KB)
- [ ] `evi` (1 KB)
- [ ] `accept` (1 KB)
- [ ] `qshell` (1 KB)
- [ ] `reject` (1 KB)
- [ ] `pon` (1 KB)
- [ ] `gsutil` (1 KB)
- [ ] `atrm` (1 KB)
- [ ] `dvipdfm` (1 KB)
- [ ] `dvipdf` (1 KB)
- [ ] `zstdless` (1 KB)
- [ ] `ps2pdf` (1 KB)
- [ ] `validate-spice` (1 KB)
- [ ] `fluxbox-remote` (1 KB)
- [ ] `xdvi` (1 KB)
- [ ] `zstdgrep` (1 KB)
- [ ] `firefox-developer-edition` (1 KB)
- [ ] `prek` (1 KB)
- [ ] `xcodes` (1 KB)
- [ ] `btdownloadheadless.py` (1 KB)
- [ ] `lz4c` (1 KB)
- [ ] `pstack` (1 KB)
- [ ] `7za` (1 KB)
- [ ] `7zr` (1 KB)
- [ ] `btdownloadcurses.py` (1 KB)
- [ ] `warp-cli` (1 KB)
- [ ] `kops` (1 KB)
- [ ] `speedtest` (1 KB)
- [ ] `folderify` (1 KB)
- [ ] `nc.traditional` (1 KB)
- [ ] `code-exploration` (1 KB)
- [ ] `ps2pdf12` (1 KB)
- [ ] `ps2pdf13` (1 KB)
- [ ] `ps2pdf14` (1 KB)
- [ ] `ps2pdfwr` (1 KB)
- [ ] `cobra-cli` (1 KB)
- [ ] `pacman-color` (1 KB)
- [ ] `zizmor` (1 KB)
- [ ] `nc.openbsd` (1 KB)
- [ ] `xxh128sum` (1 KB)
- [ ] `bb-wrapper` (1 KB)
- [ ] `mkfs.vfat` (1 KB)
- [ ] `stackit` (1 KB)
- [ ] `xxh32sum` (1 KB)
- [ ] `powerpill` (1 KB)
- [ ] `rumdl` (1 KB)
- [ ] `cabal-dev` (1 KB)
- [ ] `niri` (1 KB)
- [ ] `pacmatic` (1 KB)
- [ ] `zed-preview` (1 KB)
- [ ] `bws` (1 KB)
- [ ] `colordiff` (1 KB)
- [ ] `dua` (1 KB)
- [ ] `arepack` (1 KB)
- [ ] `buf` (1 KB)
- [ ] `pdfelatex` (1 KB)
- [ ] `sudo-rs` (1 KB)
- [ ] `gvimdiff` (1 KB)
- [ ] `oc` (1 KB)
- [ ] `adiff` (1 KB)
- [ ] `apack` (1 KB)
- [ ] `lpmove` (1 KB)
- [ ] `pdfetex` (1 KB)
- [ ] `elatex` (1 KB)
- [ ] `rgrep` (1 KB)
- [ ] `gw` (1 KB)
- [ ] `lprm` (1 KB)
- [ ] `omega` (1 KB)
- [ ] `pftp` (1 KB)
- [ ] `etex` (1 KB)

## Not in scope

- fish builtins and fish-only tooling (`alias`, `abbr`, `function`, `psub`,
  `commandline`, `fish_*`, `vared`, `status`, ...). The shell handles these and
  `CARAPACE_BRIDGES=fish` already routes to the fish completions.
- Commands served by a known bridge (`cmd/carapace/cmd/completers/bridges.go`)
  are not listed; promote one to a native completer only if the bridge is poor.
- Retired tools are marked `skip` in Queue A; do not resurrect them.
- Punctuation-only fish files (`!`, `.`, `[`) are not commands and are omitted.

## Workflow notes (verified in batch 1)

- Scraper usage: add the tool to `../scrape/package.json`, add
  `scrapers/<tool>/Dockerfile` plus a patch, run `./init.sh`, then
  `docker compose build <tool>` and
  `docker compose run --quiet <tool> > /tmp/<tool>.yaml`.
- Codegen: `carapace --codegen /tmp/<tool>.yaml` prints the generated files;
  copy them into `completers/<group>/<tool>_completer/cmd/`.
- The release optimizer assumes every `.go` file directly inside `cmd/` has an
  `init()`; put helpers in `cmd/action/` or in `root.go`.
- Inherited flags must be registered on the parent command: the storage check
  rejects `FlagCompletion` keys that are not in `cmd.LocalFlags()`.
- The argparse converter emits a phantom `- name: ""` command for the root
  parser; `scrapers/pre-commit/clean.py` strips it.
- Verify completion output without a shell:
  `./carapace <tool> export <tool> <args...> "prefix"`.
- Group commands in a spec (for example `llm keys`) do not carry the flags of
  their children; register flag completion per leaf command. Where a shared map
  is convenient, filter it with `cmd.LocalFlags().Lookup(name)`.
- Follow-ups for batch 2: `llm` cannot complete model names, keys, schemas,
  templates, fragments or conversation ids yet (they live in the llm database
  or config directory); `meson` leaves `-D` project options and `test` names
  alone because both need a configured build directory.
