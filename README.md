# Todo

A simple task manager with YAML configuration, sqlite database tasks, and Waybar integration.

## Description

Tasks are stored in a sqlite database with their unix timestamps, and `todo -t` outputs the current task as JSON for a Waybar custom module.

## Installation

```console
git clone https://github.com/ayuxsys/todo.git --depth=1
cd todo
go build -o .
mv ./todo ~/go/bin
```

## Setup

Create a `config.yaml` with your tasks:

```yaml
tasks:
  - title: "Set up development environment"
    desc: "Install dependencies and configure local tooling."
    duration: "30m"

  - title: "Implement authentication"
    desc: "Add login, logout, and session management."
    duration: "2h"

  - title: "Write unit tests"
    desc: "Add tests for the core application logic."
    duration: "1h"

  - title: "Update documentation"
    desc: "Document the API and configuration options."
    duration: "45m"
```

Load the configuration into the database:

```sh
todo load -rc config.yaml
```

## Configuring Waybar

Add the following module to your Waybar configuration:

```json
{
  "custom/todo": {
    "exec": "\"$HOME\"/go/bin/todo -t",
    "return-type": "json",
    "interval": 5,
    "tooltip": true
  }
}
```

Then add the module to `modules-center`:

```json
"modules-center": ["custom/todo", "clock"]
```

## Hyprland Autostart

To automatically load your configuration when Hyprland starts, add the following to your Hyprland startup configuration:

```lua
-- Loads $HOME/.todo/config.yaml
hl.exec_cmd("\"$HOME\"/go/bin/todo  load -r")
```

This ensures your Todo configuration is loaded automatically when starting Hyprland.