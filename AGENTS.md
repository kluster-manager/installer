# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `github.com/kluster-manager/installer` — Helm charts and supporting tooling for installing the AppsCode-curated OCM (Open Cluster Management) stack: the OCM hub/spoke components themselves plus the AppsCode-maintained OCM addons (cluster-auth, cluster-gateway, cluster-profile, cluster-proxy, fluxcd, managed-serviceaccount, multicluster-controlplane, etc.). It also exposes a Go API package describing chart values so other components can consume strongly typed installation parameters.

Charts shipped:
- `charts/cluster-manager-hub`, `charts/cluster-manager-spoke` — OCM hub and klusterlet installs.
- `charts/cluster-auth`, `charts/cluster-auth-manager` — RBAC addon (agent + hub-side manager).
- `charts/cluster-gateway`, `charts/cluster-gateway-manager` — multi-cluster proxy apiserver.
- `charts/cluster-profile-manager` — feature/profile distribution.
- `charts/cluster-proxy-manager` — ANP-based reverse-proxy addon.
- `charts/fluxcd-manager` — Flux CD install addon.
- `charts/managed-serviceaccount-manager` — managed ServiceAccount tokens addon.
- `charts/multicluster-controlplane` — standalone OCM hub control plane.
- `charts/multicluster-ingress-reader` — read aggregated ingress state.
- `charts/hub-cluster-robot`, `charts/spoke-cluster-addons` — install bundles.

The local filesystem path is `open-cluster-management.io/installer`; the **Go module is `github.com/kluster-manager/installer`** (despite `GO_PKG := go.klusters.dev` in the Makefile — that's an inherited value, not the truth). Imports use the GitHub path.

## Architecture

- `charts/` — one subdirectory per Helm chart. Each has `Chart.yaml`, `values.yaml`, `templates/`, plus generated artifacts `doc.yaml`, `README.md`, and (where applicable) `values.openapiv3_schema.yaml` + `crds/`.
- `apis/installer/v1alpha1/` — Go types backing chart values. Used for OpenAPI/values-schema generation and as a typed surface for downstream programs. Single API group: `installer:v1alpha1`.
  - `register.go`, `install/`, `fuzzer/` — standard k8s scheme registration and round-trip fuzz helpers.
- `catalog/imagelist.yaml` — image catalog source of truth; the `kmodules.xyz/image-packer` tooling regenerates the rest of `catalog/` from it. `catalog/README.md` is a generated CVE report.
- `hack/scripts/` — codegen / release helpers:
  - `update-catalog.sh` — regenerate `catalog/` from `imagelist.yaml`.
  - `update-chart-dependencies.sh` — refresh `Chart.lock` / subchart pins.
  - `import-crds.sh` — pull CRDs from the dependent kluster-manager repos into each chart's `crds/`.
  - `ct.sh` — wrap chart-testing.
  - `open-pr.sh`, `trigger.sh`, `update-release-tracker.sh` — release bookkeeping.
- `tests/` — chart-testing configuration.
- `lintconf.yaml` — YAML lint config consumed by `ct`.
- `vendor/` — vendored Go deps.

## Common commands

All Make targets run inside `ghcr.io/appscode/golang-dev` — Docker must be running.

- `make gen` — full regen: `codegen manifests`.
- `make codegen` — regenerate clientset only.
- `make manifests` — `gen-crds patch-crds label-crds gen-schema gen-chart-doc`.
- `make gen-values-schema` (alias `gen-schema`) — regenerate `values.openapiv3_schema.yaml` from `apis/installer/v1alpha1`.
- `make gen-chart-doc` — regenerate per-chart `README.md` (one target per chart subdir under `charts/`).
- `make update-charts` — refresh chart-level metadata across all charts.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make ct` — chart-testing lint+test.
- `make verify` — `verify-gen verify-modules`; `go mod tidy && go mod vendor` must leave the tree clean.
- `make add-license` / `make check-license` — manage license headers.

Auxiliary helpers (invoked directly):

- `./hack/scripts/update-catalog.sh` — regenerate `catalog/` from `imagelist.yaml`.
- `./hack/scripts/import-crds.sh` — pull CRDs from sibling kluster-manager repos into each chart's `crds/` dirs.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/installer/v1alpha1/... -run TestName -v
```

## Conventions

- Module path is `github.com/kluster-manager/installer`. Filesystem location under `open-cluster-management.io/` is for layout only. The `GO_PKG := go.klusters.dev` line in the Makefile is **historical** and does not match the actual module — when bumping the Makefile, leave it unless you can update everything that depends on it.
- Edit `apis/installer/v1alpha1/*_types.go` to change a chart's values surface, then run `make gen` so `values.openapiv3_schema.yaml`, the generated clientset, and per-chart `README.md` stay in sync. Do not hand-edit `zz_generated.*.go`, generated chart `README.md` files, `values.openapiv3_schema.yaml`, or anything under `catalog/` except `imagelist.yaml`.
- License: Apache-2.0 (`LICENSE`); use `make add-license` to apply headers to new Go files.
- Sign off commits (`git commit -s`); contributions follow the DCO (`DCO`).
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- CRDs for charts are **imported** from the upstream operator repo via `import-crds.sh`; do not hand-author `charts/*/crds/*.yaml`.
- Charts come in matched pairs for hub-vs-spoke addons (e.g. `cluster-gateway` + `cluster-gateway-manager`); keep their version bumps coordinated.
