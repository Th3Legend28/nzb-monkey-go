# NZB Monkey Go - UI Edition

This version includes a modern graphical user interface (GUI) built with Fyne.

## Building

### Prerequisites

- Go 1.25 or higher
- For Linux GUI: `libgl1-mesa-dev xorg-dev` (Debian/Ubuntu) or equivalent packages

### Quick Build

#### Linux (Primary Platform):
```bash
chmod +x build.sh
./build.sh linux
```

#### Cross-Platform Build:
```bash
# Build for all platforms
./build.sh all

# Or use PowerShell on Windows
.\build.ps1 -All

# Or use Make
make build-all
```

### Platform-Specific Builds

#### Linux:
```bash
./build.sh linux          # AMD64
./build.sh linux-arm64    # ARM64 (e.g., Raspberry Pi)
make build-linux
```

#### Windows:
```powershell
.\build.ps1 -Target windows
```
or
```bash
make build-windows
```

#### macOS:
```bash
./build.sh darwin
make build-darwin
```

## Installation (Linux)

### System-wide installation:
```bash
./build.sh install
```

This will install the binary to `/usr/local/bin/nzb-monkey-go`.

### Desktop Integration:
```bash
# Copy desktop file
sudo cp nzb-monkey-go.desktop /usr/share/applications/

# Update desktop database
sudo update-desktop-database
```

## Usage

### Graphical User Interface (GUI):
```bash
nzb-monkey-go --gui
# or simply:
nzb-monkey-go
```

The GUI will start automatically if no command-line arguments are provided.

### Command Line Interface (CLI):
```bash
# Use NZBLNK URI
nzb-monkey-go "nzblnk://?h=..."

# Manual search
nzb-monkey-go --subject "search term" --title "My Download"

# With all parameters
nzb-monkey-go --subject "ubuntu" --title "Ubuntu ISO" --group alt.binaries.linux
```

### Register NZBLNK Protocol:
```bash
nzb-monkey-go --register
```

## UI Features

The GUI includes three main tabs:

### 1. Search Tab
- NZBLNK URI input field
- Manual search parameters (header, title, password, groups, date, category)
- Real-time results display with file/segment completeness information

### 2. Settings Tab
- General settings (target, debug mode)
- NZB check options (skip failed, find best NZB)
- SABnzbd configuration

### 3. Log Tab
- Real-time logging of search operations
- View and clear logs

## Dependencies

The UI build requires additional dependencies:
- `fyne.io/fyne/v2` - Modern cross-platform GUI framework

All dependencies are automatically installed when running the build scripts.

## Build Outputs

Binaries are created in the `bin/` directory:
- `nzb-monkey-go-linux-amd64` - Linux 64-bit
- `nzb-monkey-go-linux-arm64` - Linux ARM64
- `nzb-monkey-go-windows-amd64.exe` - Windows 64-bit
- `nzb-monkey-go-darwin-amd64` - macOS Intel
- `nzb-monkey-go-darwin-arm64` - macOS Apple Silicon

## Clean Build

```bash
./build.sh clean
# or
make clean
# or on Windows
.\build.ps1 -Clean
```

## Development

### Run in Development Mode:
```bash
# GUI mode
go run . --gui

# CLI mode
go run . --help
```

### Update Dependencies:
```bash
go mod tidy
```

## Notes

- The GUI and CLI modes share the same configuration file
- Configuration file location:
  - Linux: `~/.config/nzb-monkey-go/nzb-monkey-go.conf`
  - Windows: `%APPDATA%\nzb-monkey-go\nzb-monkey-go.conf`
  - macOS: `~/Library/Application Support/nzb-monkey-go/nzb-monkey-go.conf`

## Troubleshooting

### Linux: Missing libraries
If you get errors about missing libraries, install required packages:
```bash
# Debian/Ubuntu
sudo apt-get install libgl1-mesa-dev xorg-dev

# Fedora/RHEL
sudo dnf install mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel

# Arch
sudo pacman -S libgl libxcursor libxrandr libxinerama libxi
```

### Windows: Build fails
Make sure Go is installed and in your PATH:
```powershell
go version
```

### macOS: Unsigned binary warning
The macOS binaries are not signed. You may need to allow them in System Preferences > Security & Privacy.

## Original Documentation

For more information about the original NZB Monkey Go functionality, configuration, and usage, see the main [README.md](README.md) and the [Wiki](https://github.com/Tensai75/nzb-monkey-go/wiki).
