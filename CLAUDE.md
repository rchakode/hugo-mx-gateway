# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Hugo MX Gateway is a lightweight Go application providing a RESTful API backend for handling contact and demo request forms on static websites (Hugo-based sites). It accepts form submissions via POST and sends templated emails via SMTP.

## Build Commands

```bash
make all         # Run deps, test, and build
make build       # Build Linux/amd64 binary to ./bin/hugo-mx-gateway
make build-ci    # Build inside Docker container (reproducible builds)
make test        # Run all tests (go test -v ./...)
make container   # Build Docker image (runs build-ci first)
make vendor      # Vendor Go dependencies
make clean       # Remove build artifacts
make run         # Build and run locally
```

## Architecture

### Request Processing Pipeline

The `/sendmail` endpoint uses nested middleware for security:

```
HTTP POST /sendmail
    ↓
MuxSecAllowedDomainsHandler  → CORS origin validation against ALLOWED_ORIGINS
    ↓
MuxSecHoneypotHandler        → Anti-spam: rejects if website fields are non-empty
    ↓
MuxSecReCaptchaHandler       → Optional Google reCaptcha v3 verification
    ↓
SendMail                     → Core email logic
```

### Key Files

- `main.go` - HTTP server setup, route definitions, Viper configuration initialization
- `sendmail.go` - Email sending logic, security middleware (CORS, honeypot, reCaptcha), form handling
- `healthz.go` - Health check endpoint (`GET /`)
- `templates/` - Go html/template files for email responses

### Form Types

The `target` form field determines behavior:
- `demo` - Sends email to submitter + BCC recipient
- `contact` - Sends email only to BCC recipient (internal notification)

## Configuration

All configuration via environment variables (Viper):

**Required:**
- `SMTP_SERVER_ADDR` - SMTP host:port (TLS on port 465)
- `SMTP_CLIENT_USERNAME` / `SMTP_CLIENT_PASSWORD` - SMTP auth credentials
- `CONTACT_REPLY_EMAIL` - From address for emails
- `CONTACT_REPLY_BCC_EMAIL` - BCC recipient for tracking
- `ALLOWED_ORIGINS` - Comma-separated allowed domain origins

**Optional:**
- `SMTP_AUTHENTICATION_ENABLED` - Boolean, default true
- `SMTP_SKIP_VERIFY_CERT` - Boolean, for self-signed certs
- `RECAPTCHA_PRIVATE_KEY` - Enables reCaptcha validation
- `DEMO_URL` - URL included in demo request emails
- `PORT` - Listen port, default 8080

## Code Conventions

- Apache 2.0 License header required on all .go files
- Copyright: "Copyright 2020 Rodrigue Chakode and contributors"
- Configuration via environment variables only (no config files)
- Security middleware pattern: wrap handlers for validation layers

## Deployment

Deployment options: Docker, Kubernetes (via Helm), Google App Engine, or any cloud VM.

- Container runs as non-root user `mxgateway` (UID 4583)
- Base image: Alpine Linux
- Helm chart in `helm/` for Kubernetes deployment
- Read-only root filesystem in production
