# UI Installation und Build Anleitung

Diese Version von NZB Monkey Go enthält jetzt eine moderne grafische Benutzeroberfläche (GUI) mit Fyne.

## Schnellstart

### 1. Build-Skript ausführbar machen (Linux):
```bash
chmod +x build.sh
```

### 2. Für Linux bauen (Hauptplattform):
```bash
./build.sh linux
```

Die ausführbare Datei wird erstellt unter: `bin/nzb-monkey-go-linux-amd64`

### 3. Für alle Plattformen bauen:
```bash
# Linux/macOS
./build.sh all

# Windows PowerShell
.\build.ps1 -All

# Mit Make
make build-all
```

## Installation (Linux)

### Systemweite Installation:
```bash
./build.sh install
```

### Desktop-Integration (optional):
```bash
# Desktop-Datei installieren
sudo cp nzb-monkey-go.desktop /usr/share/applications/

# Desktop-Datenbank aktualisieren
sudo update-desktop-database
```

## Verwendung

### GUI-Modus (Standard):
```bash
# GUI startet automatisch wenn keine Parameter übergeben werden
nzb-monkey-go

# Oder explizit mit --gui Flag
nzb-monkey-go --gui
```

### CLI-Modus (wie vorher):
```bash
# NZBLNK verwenden
nzb-monkey-go "nzblnk://?h=..."

# Manuelle Suche
nzb-monkey-go --subject "suchbegriff" --title "Mein Download"
```

## GUI Funktionen

- **Search Tab**: NZBLNK eingeben oder manuelle Suchparameter
- **Settings Tab**: Wichtigste Konfigurationseinstellungen
- **Log Tab**: Echtzeit-Logs der Suchvorgänge

## Abhängigkeiten (Linux)

Für die GUI werden zusätzliche System-Bibliotheken benötigt:

```bash
# Debian/Ubuntu
sudo apt-get install libgl1-mesa-dev xorg-dev

# Fedora/RHEL
sudo dnf install mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel

# Arch Linux
sudo pacman -S libgl libxcursor libxrandr libxinerama libxi
```

## Build-Optionen

### Einzelne Plattformen:
```bash
./build.sh linux          # Linux AMD64
./build.sh linux-arm64    # Linux ARM64 (z.B. Raspberry Pi)
./build.sh windows        # Windows
./build.sh darwin         # macOS (Intel + Apple Silicon)
```

### Mit Make:
```bash
make build-linux
make build-windows
make build-darwin
make build-all
```

### Mit PowerShell (Windows):
```powershell
.\build.ps1 -Target linux
.\build.ps1 -Target windows
.\build.ps1 -All
```

## Aufräumen

```bash
# Build-Artefakte entfernen
./build.sh clean

# oder
make clean

# oder Windows
.\build.ps1 -Clean
```

## Hinweise

- Die GUI und CLI teilen sich die gleiche Konfigurationsdatei
- Konfigurationspfad Linux: `~/.config/nzb-monkey-go/nzb-monkey-go.conf`
- Konfigurationspfad Windows: `%APPDATA%\nzb-monkey-go\nzb-monkey-go.conf`

## Cross-Compilation

Die Build-Skripte unterstützen Cross-Compilation out-of-the-box:
- Von Windows können Linux/macOS Binaries gebaut werden
- Von Linux können Windows/macOS Binaries gebaut werden
- Von macOS können Windows/Linux Binaries gebaut werden

Das funktioniert, weil Go nativ Cross-Compilation unterstützt und Fyne ebenfalls cross-platform ist!

## Weitere Informationen

Siehe [BUILD_UI.md](BUILD_UI.md) für detaillierte Build-Anweisungen und Troubleshooting.
