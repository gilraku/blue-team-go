# blue-team-go

A collection of command-line security tools for blue team operations, written in Go.

## Tools

| Tool | Description |
|------|-------------|
| `hash-checker` | Compute and compare file/string hashes (MD5, SHA1, SHA256, SHA512) |
| `b64` | Base64 encode and decode |
| `cidr-calc` | CIDR range calculator and IP subnet analyzer |
| `dns-lookup` | DNS query tool supporting multiple record types |
| `jwt-decode` | Decode and inspect JWT tokens without verification |
| `ioc-extract` | Extract Indicators of Compromise from text (IPs, domains, hashes, URLs) |
| `log-parser` | Parse and filter common log formats (syslog, Apache, Nginx) |
| `port-scanner` | Fast TCP port scanner with banner grabbing |
| `tls-check` | TLS/SSL certificate inspector |
| `pwcheck` | Password entropy and strength estimator |
| `http-inspect` | HTTP response header inspector |
| `report-gen` | Generate Markdown security reports from JSON findings |

## Installation

```bash
git clone https://github.com/gilraku/blue-team-go.git
cd blue-team-go
go build ./...
```

## Usage

Each tool is a standalone binary located in `cmd/<tool-name>/`.

```bash
# Example
go run ./cmd/hash-checker -algo sha256 -input "hello world"
go run ./cmd/b64 -d -input "aGVsbG8gd29ybGQ="
go run ./cmd/cidr-calc -cidr 192.168.1.0/24
```

## Requirements

- Go 1.21+

## License

MIT
