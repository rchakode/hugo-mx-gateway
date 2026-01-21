# Release v1.0.0

## Security

- Go runtime updated to 1.24.11 — Addresses multiple security vulnerabilities
- Honeypot anti-spam protection — New middleware rejects requests where `website` form fields contain values (bot detection)
- Added SECURITY.md — Vulnerability reporting policy for responsible disclosure
- CI security scanning — Modernized workflows with security checks

## Enhancement

- Honeypot field prefix — Form fields use `website` prefix (see [Migration Section](#migration-guide)).
- Distroless container image — Switched from Alpine to Google's distroless base image for reduced attack surface

## Features

- Improved container security — Distroless image, no shell, minimal footprint
- README key benefits section — Added 4 key benefits highlighting the project's value proposition

## Infrastructure

- Modernized CI pipelines — Updated GitHub Actions workflows
- Simplified Dockerfile — Removed entrypoint.sh, set explicit GOOS/GOARCH build flags
- Helm chart updated — appVersion bumped to 1.0.0

## Documentation

- Added CLAUDE.md — Claude Code guidance for contributors
- Fixed typos and URLs — Corrected documentation errors
- Cleaned up unused docs — Removed deprecated deployment-on-debian.md
- Fixed Helm values.yaml path — Corrected relative path in Kubernetes deployment docs

## Dependencies

- Vendor cleanup — Removed unused vendored dependencies (gorilla/mux, sirupsen/logrus, x/sys/windows, x/crypto/ssh/terminal, etc.)
- Updated go.mod — Refreshed module dependencies

## Migration Guide

Form HTML changes required:
                                                              

```html
<!-- After (new honeypot prefix) -->   

<label style="display:none">
   Website check: <input type="text" name="website-check" value="" autocomplete="off" tabindex="-1" />
</label>                                                                                  
        
```                                                       

See `samples/hugo-partial-contact-form.html` for the updated form template.

---                                                                                                                                                                                                     
# Release v0.4.0

Release Date: May 16, 2022
## Breaking Changes

- Renamed environment variable: SMTP_VERITY_CERT → SMTP_SKIP_VERIFY_CERT
    - Fixed typo in the variable name
    - Renamed for clarity: when set to true, TLS certificate verification is skipped (useful for self-signed certificates)

## Changes

- Updated documentation across all deployment guides to reflect the new variable name:
    - Configuration variables docs
    - Docker deployment guide
    - Helm chart values.yaml
    - Google App Engine app.yaml.sample

## Other

- Updated README headline and added project thumbnail image

## Migration

If upgrading from v0.3.0, rename your environment variable:
### Before
SMTP_VERITY_CERT=true

### After
SMTP_SKIP_VERIFY_CERT=true
