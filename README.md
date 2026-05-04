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

## New Tools (v0.2)

| Tool | Description |
|------|-------------|
| `whois` | WHOIS lookup for domains and IPs |
| `asn-lookup` | ASN/BGP prefix lookup via Team Cymru |
| `mac-lookup` | MAC address vendor identification |
| `entropy` | File entropy calculator for malware analysis |
| `hexdump` | Hex+ASCII file dump |
| `strings-extract` | Extract printable strings from binaries |
| `timestamp` | Unix epoch ↔ human date converter |
| `url-parse` | URL component parser and decoder |
| `ip-range` | IP range expander (CIDR, dash, list) |
| `pass-gen` | Cryptographically secure password generator |
| `file-type` | Magic bytes file type detector |
| `rot` | ROT13/Caesar cipher with brute force mode |
| `email-header` | Email header parser for phishing analysis |
| `netstat-parse` | netstat output parser and summarizer |
| `ua-parse` | User-Agent string parser |
