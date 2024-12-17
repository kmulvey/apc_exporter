# apc_exporter
[![Build](https://github.com/kmulvey/apc_exporter/actions/workflows/build.yml/badge.svg)](https://github.com/kmulvey/apc_exporter/actions/workflows/build.yml) [![Release](https://github.com/kmulvey/apc_exporter/actions/workflows/release.yml/badge.svg)](https://github.com/kmulvey/apc_exporter/actions/workflows/release.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/apc_exporter)](https://goreportcard.com/report/github.com/kmulvey/apc_exporter) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/apc_exporter.svg)](https://pkg.go.dev/github.com/kmulvey/apc_exporter)

Monitor and store APC UPS statistics.

## Installation and Usage
Several linux package formats are available in the releases. Becasue pwrstat needs to be run as root, this tool needs to be run as root as well.

### Manual linux install:
- `sudo cp apc_exporter /usr/bin/` (this path can be changed if you like, just be sure to change the path in the service file as well)
- `sudo cp apc_exporter.service /etc/systemd/system/`
- `sudo systemctl daemon-reload`
- `sudo systemctl enable apc_exporter`
- `sudo systemctl restart apc_exporter`
- Import grafana-config.json to your grafana instance
- enjoy!

![Screenshot](https://github.com/kmulvey/apc_exporter/blob/main/screenshot.jpg?raw=true)
