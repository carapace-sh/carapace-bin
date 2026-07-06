# Display and Graphics

Configure how the VM's graphical output is presented, both locally and remotely.

## Display Backends

| Display | Option | Description |
|---------|--------|-------------|
| GTK | `-display gtk` | Native GTK3 window with zoom, grab, and Ctrl+Alt+ shortcuts |
| SDL | `-display sdl` | SDL2 window |
| curses | `-display curses` | Text-mode terminal output (no VGA) |
| VNC | `-vnc HOST:PORT` | Remote framebuffer protocol (via `-vnc`, not `-display`) |
| spice-app | `-display spice-app` | SPICE remote desktop protocol (standalone GTK window) |
| egl-headless | `-display egl-headless` | Offscreen EGL rendering, remote-only |
| dbus | `-display dbus` | Wayland/D-Bus display sharing |
| cocoa | `-display cocoa` | Native macOS window (Cocoa API, macOS only) |
| none | `-display none` | No display (use QMP/HMP only) |

## VGA Interface

```
-vga MODEL
-vga std      # Standard VGA (bochs-display, default for most machines)
-vga cirrus   # Cirrus CLGD 5446 (legacy)
-vga vmware   # VMware SVGA II (for VMware drivers)
-vga qxl      # SPICE QXL (for SPICE acceleration)
-vga virtio   # virtio-gpu (virgl/3D acceleration)
-vga none     # No VGA (use -device for custom graphics)
```

## VNC

VNC is provided via `-vnc`, not `-display`:

```
-vnc HOST:PORT[,option=value[,...]]
```

| Option | Description |
|--------|-------------|
| `HOST:PORT` | e.g. `0.0.0.0:0` (display 0 on all interfaces) |
| `unix:PATH` | UNIX socket |
| `password` | Require password (set via QMP `set_password`) |
| `lossy` | Use lossy compression methods |
| `non-adaptive` | Disable adaptive updates |
| `acl` | Enable access control list |
| `reverse=HOST:PORT` | Reverse/VNC server connection |
| `websocket=PORT` | WebSocket port for noVNC |

VNC key mapping: Ctrl+Alt+[1-9] to switch monitors; Ctrl+Alt+Shift to release grab.

### Connecting
```
vncviewer HOST:DISPLAY
```

## SPICE

```
-spice port=PORT,addr=ADDR,password=SECRET,disable-ticketing
```

| Option | Description |
|--------|-------------|
| `port=N` | TCP port |
| `tls-port=N` | TLS-encrypted port |
| `addr=IP` | Bind address |
| `password=VAL` | Plaintext password (mutually exclusive with `disable-ticketing`) |
| `disable-ticketing` | No authentication |
| `image-compression=auto|quic|glz|lz|off` | Image compression |
| `jpeg-wan-compression=auto|never|always` | JPEG compression over WAN |
| `zlib-glz-wan-compression=auto|never|always` | Zlib compression |
| `streaming-video=off|all|filter` | Video streaming detection |
| `agent-mouse=on|off` | Guest mouse via spice-vdagent |
| `playback-compression=on|off` | Audio playback compression |
| `seamless-migration=on|off` | Seamless migration support |

### Connecting
```
spicy -h HOST -p PORT
remote-viewer spice://HOST:PORT
```

## VirGL / 3D Acceleration

```
-display gtk,gl=on
-vga virtio
```

or

```
-display egl-headless,gl=on
-vga virtio
```

Requires:
- `virglrenderer` library
- Paravirtualized GPU: `-device virtio-vga` or `-device virtio-gpu-pci`
- `-display egl-headless,gl=on` for remote-only

VirGL provides OpenGL acceleration to the guest without emulating a physical GPU.

## Multi-Monitor

```
-device virtio-vga,max_outputs=2
-device virtio-gpu-pci,max_outputs=2
```

Each output appears as a separate monitor in the guest. Switch between heads with `-display gtk` using Ctrl+Alt+[1-9].

## Full Screen and Window Options

| Option | Description |
|--------|-------------|
| `-full-screen` | Start full screen |
| `-g WxH[xDEPTH]` | Guest resolution (e.g. `-g 1024x768x32`) |
| `-no-frame` | No window decorations (SDL) |
| `-alt-grab` | Use Ctrl+Alt+Shift for grab (was Ctrl+Alt) |
| `-ctrl-grab` | Use right Ctrl for grab |

## Audio

Audio is configured via `-audiodev` (not a display option, but related):

```
-audiodev pa,id=audio0       # PulseAudio
-audiodev alsa,id=audio0     # ALSA
-audiodev sdl,id=audio0      # SDL audio
-audiodev spice,id=audio0    # SPICE audio tunnel
-audiodev none,id=audio0     # No host audio
-audiodev oss,id=audio0      # OSS
-audiodev jack,id=audio0     # JACK
-audiodev dbus,id=audio0     # PipeWire/D-Bus
-audiodev pipewire,id=audio0 # PipeWire
```

Then attach to a device:
```
-device intel-hda -device hda-duplex,audiodev=audio0
-device AC97,audiodev=audio0
-device sb16,audiodev=audio0
```
