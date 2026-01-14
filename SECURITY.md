# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.4.x   | :white_check_mark: |
| < 0.4   | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please report it responsibly.

**Please do NOT report security vulnerabilities through public GitHub issues.**

### How to Report

Send an email to: **contact@krossboard.app**

Include the following information:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

### Response Timeline

- **Acknowledgment**: within 48 hours
- **Initial assessment**: within 7 days
- **Resolution target**: within 30 days (depending on severity)

### What to Expect

1. We will acknowledge receipt of your report
2. We will investigate and validate the issue
3. We will work on a fix and coordinate disclosure
4. We will credit you in the release notes (unless you prefer anonymity)

## Security Best Practices for Deployment

When deploying hugo-mx-gateway, consider the following:

- Use TLS/HTTPS for all communications
- Store SMTP credentials securely (Kubernetes Secrets, environment variables)
- Restrict network access to the service
- Keep the Docker image updated to the latest version
- Review and restrict CORS origins in production

## Dependency Security

This project uses automated vulnerability scanning:
- **govulncheck** for Go dependencies
- **Trivy** for Docker image scanning

Security updates to dependencies are prioritized and released promptly.
