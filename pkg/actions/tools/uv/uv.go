// package uv contains uv related actions
package uv

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/uid"
)

// Uid creates a unique identifier for uv completion values
func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		if length := len(opts); length%2 != 0 {
			return nil, fmt.Errorf("invalid amount of arguments [uv.Uid]: %v", length)
		}

		uid := &url.URL{
			Scheme: "uv",
			Host:   host,
			Path:   s,
		}
		values := uid.Query()
		for i := 0; i < len(opts); i += 2 {
			values.Add(opts[i], opts[i+1])
		}
		uid.RawQuery = values.Encode()
		return uid, nil
	}
}

// ActionTools completes installed tools
//
//	ruff (0.5.0)
//	black (24.4.2)
func ActionTools() carapace.Action {
	return carapace.ActionExecCommand("uv", "tool", "list")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`^(?P<name>\S+) \((?P<version>[^)]+)\)`)
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if match := re.FindStringSubmatch(line); match != nil {
				vals = append(vals, match[1], match[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...).
			UidF(Uid("tool")).
			QueryF(Uid("tool"))
	}).Tag("installed tools")
}

type pkg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ActionInstalledPackages completes installed packages
//
//	click (8.1.7)
//	pip (24.0)
func ActionInstalledPackages() carapace.Action {
	return carapace.ActionExecCommand("uv", "pip", "list", "--format", "json")(func(output []byte) carapace.Action {
		var packages []pkg
		if err := json.Unmarshal(output, &packages); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, p := range packages {
			vals = append(vals, p.Name, p.Version)
		}
		return carapace.ActionValuesDescribed(vals...).
			UidF(Uid("installed-package")).
			QueryF(Uid("installed-package"))
	}).Tag("installed packages")
}

type InstallationsOpts struct {
	InstalledOnly bool
}

type installation struct {
	Key     string `json:"key"`
	Path    string `json:"path"`
	Symlink string `json:"symlink"`
	Url     string `json:"url"`
}

// ActionPythonInstallations completes Python installations
//
//	cpython-3.12.3-linux-x86_64-gnu (/home/user/.local/share/uv/python/cpython-3.12.3-linux-x86_64-gnu/bin/python3)
func ActionPythonInstallations(opts InstallationsOpts) carapace.Action {
	args := []string{"python", "list", "--output-format", "json"}
	if opts.InstalledOnly {
		args = append(args, "--only-installed")
	}
	return carapace.ActionExecCommand("uv", args...)(func(output []byte) carapace.Action {
		var installations []installation
		if err := json.Unmarshal(output, &installations); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, i := range installations {
			description := i.Url
			if i.Path != "" {
				description = i.Path
			}
			vals = append(vals, i.Key, description)
		}
		return carapace.ActionValuesDescribed(vals...).
			UidF(Uid("python")).
			QueryF(Uid("python"))
	}).Tag("python installations")
}

// ActionTorchBackends completes torch backends
//
//	auto (Select the appropriate PyTorch index based on the operating system and CUDA driver version)
//	cpu (Use the CPU-only PyTorch index)
func ActionTorchBackends() carapace.Action {
	return carapace.ActionValuesDescribed(
		"auto", "Select the appropriate PyTorch index based on the operating system and CUDA driver version",
		"cpu", "Use the CPU-only PyTorch index",
		"cu132", "Use the PyTorch index for CUDA 13.2",
		"cu130", "Use the PyTorch index for CUDA 13.0",
		"cu129", "Use the PyTorch index for CUDA 12.9",
		"cu128", "Use the PyTorch index for CUDA 12.8",
		"cu126", "Use the PyTorch index for CUDA 12.6",
		"cu125", "Use the PyTorch index for CUDA 12.5",
		"cu124", "Use the PyTorch index for CUDA 12.4",
		"cu123", "Use the PyTorch index for CUDA 12.3",
		"cu122", "Use the PyTorch index for CUDA 12.2",
		"cu121", "Use the PyTorch index for CUDA 12.1",
		"cu120", "Use the PyTorch index for CUDA 12.0",
		"cu118", "Use the PyTorch index for CUDA 11.8",
		"cu117", "Use the PyTorch index for CUDA 11.7",
		"cu116", "Use the PyTorch index for CUDA 11.6",
		"cu115", "Use the PyTorch index for CUDA 11.5",
		"cu114", "Use the PyTorch index for CUDA 11.4",
		"cu113", "Use the PyTorch index for CUDA 11.3",
		"cu112", "Use the PyTorch index for CUDA 11.2",
		"cu111", "Use the PyTorch index for CUDA 11.1",
		"cu110", "Use the PyTorch index for CUDA 11.0",
		"cu102", "Use the PyTorch index for CUDA 10.2",
		"cu101", "Use the PyTorch index for CUDA 10.1",
		"cu100", "Use the PyTorch index for CUDA 10.0",
		"cu92", "Use the PyTorch index for CUDA 9.2",
		"cu91", "Use the PyTorch index for CUDA 9.1",
		"cu90", "Use the PyTorch index for CUDA 9.0",
		"cu80", "Use the PyTorch index for CUDA 8.0",
		"rocm7.2", "Use the PyTorch index for ROCm 7.2",
		"rocm7.1", "Use the PyTorch index for ROCm 7.1",
		"rocm7.0", "Use the PyTorch index for ROCm 7.0",
		"rocm6.4", "Use the PyTorch index for ROCm 6.4",
		"rocm6.3", "Use the PyTorch index for ROCm 6.3",
		"rocm6.2.4", "Use the PyTorch index for ROCm 6.2.4",
		"rocm6.2", "Use the PyTorch index for ROCm 6.2",
		"rocm6.1", "Use the PyTorch index for ROCm 6.1",
		"rocm6.0", "Use the PyTorch index for ROCm 6.0",
		"rocm5.7", "Use the PyTorch index for ROCm 5.7",
		"rocm5.6", "Use the PyTorch index for ROCm 5.6",
		"rocm5.5", "Use the PyTorch index for ROCm 5.5",
		"rocm5.4.2", "Use the PyTorch index for ROCm 5.4.2",
		"rocm5.4", "Use the PyTorch index for ROCm 5.4",
		"rocm5.3", "Use the PyTorch index for ROCm 5.3",
		"rocm5.2", "Use the PyTorch index for ROCm 5.2",
		"rocm5.1.1", "Use the PyTorch index for ROCm 5.1.1",
		"rocm4.2", "Use the PyTorch index for ROCm 4.2",
		"rocm4.1", "Use the PyTorch index for ROCm 4.1",
		"rocm4.0.1", "Use the PyTorch index for ROCm 4.0.1",
		"xpu", "Use the PyTorch index for Intel XPU",
	)
}

// ActionPythonPlatforms completes Python platforms
//
//	windows (An alias for `x86_64-pc-windows-msvc`, the default target for Windows)
//	linux (An alias for `x86_64-unknown-linux-gnu`, the default target for Linux)
func ActionPythonPlatforms() carapace.Action {
	return carapace.ActionValuesDescribed(
		"windows", "An alias for `x86_64-pc-windows-msvc`, the default target for Windows",
		"linux", "An alias for `x86_64-unknown-linux-gnu`, the default target for Linux",
		"macos", "An alias for `aarch64-apple-darwin`, the default target for macOS",
		"x86_64-pc-windows-msvc", "A 64-bit x86 Windows target",
		"aarch64-pc-windows-msvc", "An ARM64 Windows target",
		"i686-pc-windows-msvc", "A 32-bit x86 Windows target",
		"x86_64-unknown-linux-gnu", "An x86 Linux target. Equivalent to `x86_64-manylinux_2_28`",
		"aarch64-apple-darwin", "An ARM-based macOS target, as seen on Apple Silicon devices",
		"x86_64-apple-darwin", "An x86 macOS target",
		"aarch64-unknown-linux-gnu", "An ARM64 Linux target. Equivalent to `aarch64-manylinux_2_28`",
		"aarch64-unknown-linux-musl", "An ARM64 Linux target",
		"x86_64-unknown-linux-musl", "An `x86_64` Linux target",
		"s390x-unknown-linux-gnu", "An s390x Linux target. Equivalent to `s390x-manylinux_2_28`",
		"powerpc64le-unknown-linux-gnu", "A little-endian `PowerPC64` Linux target. Equivalent to `ppc64le-manylinux_2_28`",
		"loongarch64-unknown-linux-gnu", "A `LoongArch64` Linux target. Equivalent to `loongarch64-manylinux_2_36`",
		"riscv64-unknown-linux", "A RISCV64 Linux target",
		"x86_64-manylinux2014", "An `x86_64` target for the `manylinux2014` platform. Equivalent to `x86_64-manylinux_2_17`",
		"x86_64-manylinux_2_17", "An `x86_64` target for the `manylinux_2_17` platform",
		"x86_64-manylinux_2_28", "An `x86_64` target for the `manylinux_2_28` platform",
		"x86_64-manylinux_2_31", "An `x86_64` target for the `manylinux_2_31` platform",
		"x86_64-manylinux_2_32", "An `x86_64` target for the `manylinux_2_32` platform",
		"x86_64-manylinux_2_33", "An `x86_64` target for the `manylinux_2_33` platform",
		"x86_64-manylinux_2_34", "An `x86_64` target for the `manylinux_2_34` platform",
		"x86_64-manylinux_2_35", "An `x86_64` target for the `manylinux_2_35` platform",
		"x86_64-manylinux_2_36", "An `x86_64` target for the `manylinux_2_36` platform",
		"x86_64-manylinux_2_37", "An `x86_64` target for the `manylinux_2_37` platform",
		"x86_64-manylinux_2_38", "An `x86_64` target for the `manylinux_2_38` platform",
		"x86_64-manylinux_2_39", "An `x86_64` target for the `manylinux_2_39` platform",
		"x86_64-manylinux_2_40", "An `x86_64` target for the `manylinux_2_40` platform",
		"aarch64-manylinux2014", "An ARM64 target for the `manylinux2014` platform. Equivalent to `aarch64-manylinux_2_17`",
		"aarch64-manylinux_2_17", "An ARM64 target for the `manylinux_2_17` platform",
		"aarch64-manylinux_2_28", "An ARM64 target for the `manylinux_2_28` platform",
		"aarch64-manylinux_2_31", "An ARM64 target for the `manylinux_2_31` platform",
		"aarch64-manylinux_2_32", "An ARM64 target for the `manylinux_2_32` platform",
		"aarch64-manylinux_2_33", "An ARM64 target for the `manylinux_2_33` platform",
		"aarch64-manylinux_2_34", "An ARM64 target for the `manylinux_2_34` platform",
		"aarch64-manylinux_2_35", "An ARM64 target for the `manylinux_2_35` platform",
		"aarch64-manylinux_2_36", "An ARM64 target for the `manylinux_2_36` platform",
		"aarch64-manylinux_2_37", "An ARM64 target for the `manylinux_2_37` platform",
		"aarch64-manylinux_2_38", "An ARM64 target for the `manylinux_2_38` platform",
		"aarch64-manylinux_2_39", "An ARM64 target for the `manylinux_2_39` platform",
		"aarch64-manylinux_2_40", "An ARM64 target for the `manylinux_2_40` platform",
		"s390x-manylinux2014", "An s390x target for the `manylinux2014` platform. Equivalent to `s390x-manylinux_2_17`",
		"s390x-manylinux_2_17", "An s390x target for the `manylinux_2_17` platform",
		"s390x-manylinux_2_28", "An s390x target for the `manylinux_2_28` platform",
		"s390x-manylinux_2_31", "An s390x target for the `manylinux_2_31` platform",
		"s390x-manylinux_2_32", "An s390x target for the `manylinux_2_32` platform",
		"s390x-manylinux_2_33", "An s390x target for the `manylinux_2_33` platform",
		"s390x-manylinux_2_34", "An s390x target for the `manylinux_2_34` platform",
		"s390x-manylinux_2_35", "An s390x target for the `manylinux_2_35` platform",
		"s390x-manylinux_2_36", "An s390x target for the `manylinux_2_36` platform",
		"s390x-manylinux_2_37", "An s390x target for the `manylinux_2_37` platform",
		"s390x-manylinux_2_38", "An s390x target for the `manylinux_2_38` platform",
		"s390x-manylinux_2_39", "An s390x target for the `manylinux_2_39` platform",
		"s390x-manylinux_2_40", "An s390x target for the `manylinux_2_40` platform",
		"ppc64le-manylinux2014", "A little-endian `PowerPC64` target for the `manylinux2014` platform. Equivalent to `ppc64le-manylinux_2_17`",
		"ppc64le-manylinux_2_17", "A little-endian `PowerPC64` target for the `manylinux_2_17` platform",
		"ppc64le-manylinux_2_28", "A little-endian `PowerPC64` target for the `manylinux_2_28` platform",
		"ppc64le-manylinux_2_31", "A little-endian `PowerPC64` target for the `manylinux_2_31` platform",
		"ppc64le-manylinux_2_32", "A little-endian `PowerPC64` target for the `manylinux_2_32` platform",
		"ppc64le-manylinux_2_33", "A little-endian `PowerPC64` target for the `manylinux_2_33` platform",
		"ppc64le-manylinux_2_34", "A little-endian `PowerPC64` target for the `manylinux_2_34` platform",
		"ppc64le-manylinux_2_35", "A little-endian `PowerPC64` target for the `manylinux_2_35` platform",
		"ppc64le-manylinux_2_36", "A little-endian `PowerPC64` target for the `manylinux_2_36` platform",
		"ppc64le-manylinux_2_37", "A little-endian `PowerPC64` target for the `manylinux_2_37` platform",
		"ppc64le-manylinux_2_38", "A little-endian `PowerPC64` target for the `manylinux_2_38` platform",
		"ppc64le-manylinux_2_39", "A little-endian `PowerPC64` target for the `manylinux_2_39` platform",
		"ppc64le-manylinux_2_40", "A little-endian `PowerPC64` target for the `manylinux_2_40` platform",
		"loongarch64-manylinux_2_36", "A `LoongArch64` target for the `manylinux_2_36` platform",
		"loongarch64-manylinux_2_37", "A `LoongArch64` target for the `manylinux_2_37` platform",
		"loongarch64-manylinux_2_38", "A `LoongArch64` target for the `manylinux_2_38` platform",
		"loongarch64-manylinux_2_39", "A `LoongArch64` target for the `manylinux_2_39` platform",
		"loongarch64-manylinux_2_40", "A `LoongArch64` target for the `manylinux_2_40` platform",
		"aarch64-linux-android", "An ARM64 Android target",
		"x86_64-linux-android", "An `x86_64` Android target",
		"wasm32-pyodide2024", "A wasm32 target using the Pyodide 2024 platform. Meant for use with Python 3.12. See <https://pyodide.org/en/stable/development/abi/312.html>",
		"wasm32-pyodide2025", "A wasm32 target using the Pyodide 2025 platform. Meant for use with Python 3.13. See <https://pyodide.org/en/stable/development/abi/313.html>",
		"arm64-apple-ios", "An ARM64 target for iOS device",
		"arm64-apple-ios-simulator", "An ARM64 target for iOS simulator",
		"x86_64-apple-ios-simulator", "An `x86_64` target for iOS simulator",
	)
}
