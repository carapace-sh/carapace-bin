# CPU Models

CPU model selection and feature configuration for QEMU/KVM. Set with `-cpu MODEL[,FEATURE=on|off[,...]]`.

## Model Approaches

### Host Passthrough

Passes the host CPU exactly as-is to the guest:

```
-cpu host
```

- Maximum performance and features
- **Unsafe for live migration** — host CPUs may differ
- Recommended when live migration is not required

### Named Models

QEMU provides predefined named CPU models representing specific hardware generations.

```
-cpu EPYC
-cpu Cascadelake-Server-v5
-cpu max   # maximum features for this arch
-cpu base  # minimum features for this arch
```

Use `-cpu help` to list models and `-cpu MODEL,help` to list features.

### Host Model (libvirt)

Libvirt's "host model" picks a named model approximating the host, then adds features. Not a QEMU-internal concept — libvirt translates it to QEMU arguments.

## CPU Topology

```
-smp cpus=4,cores=2,threads=2,sockets=1
```

| Parameter | Description |
|-----------|-------------|
| `cpus=N` | Total number of guest CPUs (required) |
| `maxcpus=N` | Maximum CPUs (including hotplug) |
| `sockets=N` | Physical sockets |
| `dies=N` | Dies per socket (x86 since ICX) |
| `clusters=N` | Clusters per die (ARM) |
| `cores=N` | Cores per die/cluster |
| `threads=N` | Threads per core |

The product of sockets × dies × clusters × cores × threads must equal `maxcpus`. `cpus` must be ≤ `maxcpus`.

## CPU Features

Enable/disable features with `+`/`-` prefix or `=on`/`=off`:

```
-cpu host,+svm,-vmx
-cpu Cascadelake-Server-v5,pcid=on,ssbd=on
```

### Common x86 Feature Flags

| Feature | Description |
|---------|-------------|
| `vmx` | Intel VMX (nested virtualization) |
| `svm` | AMD SVM (nested virtualization) |
| `pcid` | Process Context IDentifiers |
| `invpcid` | Invalidate PCID |
| `pdpe1gb` | 1 GiB huge pages |
| `rdrand` | RDRAND instruction |
| `rdseed` | RDSEED instruction |
| `sha-ni` | SHA extensions |
| `avx` / `avx2` / `avx512f` | AVX extensions |
| `xsave` / `xcr0` | XSAVE/XRSTOR |
| `fsgsbase` | FS/GS base instructions |
| `tsc-deadline` | TSC deadline timer |
| `x2apic` | x2APIC |
| `hypervisor` | Hypervisor CPUID bit |
| `kvm` | KVM PV features |
| `kvm-hint-dedicated` | Hint to guest it's on dedicated CPU |
| `migratable` | Only features migratable between hosts |
| `sse4.1` / `sse4.2` | SSE 4.1/4.2 |
| `ssbd` | Speculative Store Bypass Disable |
| `md-clear` | VERW MD_CLEAR (MDS mitigation) |
| `spec-ctrl` | SPEC_CTRL MSR |
| `ibpb` | Indirect Branch Predictor Barrier |
| `stibp` | Single Thread Indirect Branch Predictors |
| `arch-capabilities` | IA32_ARCH_CAPABILITIES MSR |

## NUMA

```
-numa node,cpus=0-1,memdev=mem0
-numa node,cpus=2-3,memdev=mem1
-object memory-backend-ram,id=mem0,size=4G
-object memory-backend-ram,id=mem1,size=4G
```

- Each NUMA node gets a memory backend
- HMAT (Heterogeneous Memory Attribute Table) for performance hints: `-machine hmat=on`
- Use `-numa dist` to set distances between nodes

## Confidential Computing

### AMD SEV, SEV-ES, SEV-SNP

AMD Secure Encrypted Virtualization. Encrypts guest memory, protects from hypervisor.

```
-machine ...,memory-encryption=sev0
-object sev-guest,id=sev0,cbitpos=47,reduced-phys-bits=1,policy=0x01
```

- `sev-guest` — SEV (memory encryption)
- `sev-guest,policy=0x03` — SEV-ES (encryption + register state protection)
- `sev-snp-guest` — SEV-SNP (Secure Nested Paging) (QEMU 9.0+)
- `policy` — bitmask: bit 0=encrypted, bit 1=SEV-ES, bit 2=SEV-SNP, bit 3=debug, bit 4=key sharing

### Intel TDX

Intel Trust Domain Extensions. Hardware-isolated virtual machines.

```
-machine ...,memory-encryption=tdx0
-object tdx-guest,id=tdx0
```

Requires TDX-capable hardware and kernel support.

## CPU Model Recommendations (x86)

| Use Case | -cpu Setting |
|----------|-------------|
| Maximum performance, no migration | `host` |
| Migrate between same-generation Intel | `-cpu Nehalem`, `-cpu Westmere`, `-cpu SandyBridge`, `-cpu Haswell`, `-cpu Broadwell`, `-cpu Skylake-Client`, etc. |
| Migrate between same-generation AMD | `-cpu Opteron_G1` through `-cpu EPYC` |
| Cross-vendor migration | `-cpu qemu64` or `-cpu kvm64` (minimal feature set) |
| VMware compatibility | `-cpu host` with `-overcommit cpu-pm=on` |
| Security-optimized | Add `+ssbd,+md-clear,+spec-ctrl,+stibp,+flush-by-asid` |

## ARM CPU Models

```
-cpu host            # host passthrough (KVM only)
-cpu max             # max QEMU features (TCG)
-cpu cortex-a76      # specific ARM core
-cpu neoverse-n1     # Neoverse N1
-cpu neoverse-v1     # Neoverse V1
```

## RISC-V CPU Models

```
-cpu host            # host passthrough (KVM)
-cpu rv64            # base RISC-V 64-bit
-cpu sifive-u54      # SiFive U54
-cpu shakti-c        # Shakti C
```

## See Also

- `qemu-cpu-models(7)` — canonical reference
- `qemu-system-x86_64 -cpu help` for runtime model/feature listing
