# Generated Code Policy

This repository contains **generated Go client code** for the FOSSA REST API. The Go module
is kept at the repo root for a clean import path:

- `github.com/RobertKielty/fossa-go` (current)

## Source of Truth

- Upstream spec: `https://app.fossa.com/api/api-docs/swagger.json`
- Local patched spec: `swagger.json` (downloaded during generation)
- Patch script: `scripts/patch_swagger.py`

We patch the upstream spec to add invitation endpoints that are currently not published in the
Swagger file (but exist in production).

## Patched Endpoints

- `GET /user-invitations` → `listUserInvitations`
- `DELETE /user-invitations/{email}` → `deleteUserInvitation`
- `POST /organizations/{id}/invite` → `sendUserInvitation`

The request/response shapes for these are based on observed UI traffic.

## Regenerating

Local regeneration:

```bash
make generate-go
```

By default we use Podman. If you prefer Docker:

```bash
CONTAINER_ENGINE=docker make generate-go
```

## Warnings

Upstream `swagger.json` contains schema issues that produce warnings in OpenAPI Generator.
We run with `--skip-validate-spec`, so generation succeeds despite these warnings.

## Release Process

Releases should only be cut **after** regeneration. The `Release` GitHub Actions workflow
triggers on `v*` tags, runs `make generate-go`, and fails if the working tree is dirty.
This keeps tagged releases aligned with the generated output.
