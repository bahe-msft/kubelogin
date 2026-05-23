# Squad Decisions

## Active Decisions

### PR #2 CI Failure Resolution

**Date:** 2026-05-23T16:02:33.592-07:00 | **From:** livingston | **Issue:** PR #2

The `dependency-review` check in PR #2 failed with 403 Forbidden. **Root Cause:** Missing workflow permission in `.github/workflows/dependency-review.yml`.

The workflow lacked the `pull-requests: read` permission required by the GitHub dependency-review action to fetch PR metadata and analyze dependency changes.

**Fix Applied:** Added `pull-requests: read` to workflow permissions (commit 8671f8e on branch pr707).

**Impact:** This is a DevOps configuration gap unrelated to Rusty's chained credential implementation. Once merged, dependency-review will function correctly.

### Issue #1 Routing: Reimplement Azure/kubelogin PR #707

**Date:** 2026-05-23T16:02:33.592-07:00 | **From:** danny | **Status:** proposed_handoff

**Route to:** Rusty (Go/auth/Kubernetes specialist)

The chained credential login mode implementation is a straightforward feature addition following established architectural patterns. No architectural changes are needed. Rusty should own the implementation.

#### Reasoning

- **Scope:** Adding a new login mode (`chained`) that uses Azure SDK's `DefaultAzureCredential`
- **Architectural status:** No new architectural decisions required. This follows the existing credential provider factory pattern used by all 9 existing login modes
- **Pattern precedent:** PR #703 established the pattern; PR #707 applies it again. The architecture is proven
- **Implementation type:** Pure feature work—no cross-cutting concerns, no component boundary changes
- **Domain:** Azure authentication, credential management, Go SDK integration
- **Complexity:** Medium—involves adding login mode constant, creating `chainedcredential.go`, unit tests, and documentation updates

#### Key Files

- `pkg/internal/token/chainedcredential.go` (49 lines)
- `pkg/internal/token/chainedcredential_test.go` (52 lines)
- Updates to `options.go`, `provider.go`, documentation

#### Notes

- This is a direct port from Azure/kubelogin#707—leverage the upstream implementation
- Maintains backward compatibility
- The `DefaultAzureCredential` chain order is non-negotiable per Azure SDK design
- Branch `pr707` ready at repo root

## Completed Decisions

### Linus Review Verdict — Issue #1 Chained Credential

**Date:** 2026-05-23T16:02:33.592-07:00 | **From:** linus | **Status:** approved | **Issue:** #1

APPROVE Rusty's issue #1 implementation.

- The new `chained` login mode is registered in the internal provider factory, supported login list, and public `pkg/token` constants.
- The implementation delegates to Azure SDK `DefaultAzureCredential` and preserves kubelogin cloud, tenant, instance discovery, and custom HTTP transport options.
- Added test-only coverage to assert internal factory wiring returns `*ChainedCredential` and public `GetTokenProvider` can construct chained login.
- Validation passed for `go test ./pkg/internal/token ./pkg/token` and `go build ./...`.
- Full `go test ./...` is blocked by unrelated `pkg/internal/pop/cache` keychain state failures on this machine; no chained credential regression observed.

No rejection; no lockout or reassignment required.

### Chained credential order follows installed Azure SDK

**Date:** 2026-05-23T16:02:33.592-07:00 | **From:** rusty | **Status:** implemented | **Issue:** #1

The new `chained` login mode delegates to Azure SDK for Go `azidentity.NewDefaultAzureCredential` rather than building a custom chain.

Acceptance requires respecting Azure SDK `DefaultAzureCredential` order. In the repository's installed `github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.0`, `DefaultAzureCredential` documents this order: Environment Credential, Workload Identity Credential, Managed Identity Credential, Azure CLI Credential, Azure Developer CLI Credential. Documentation for chained mode reflects that SDK order, including Azure Developer CLI.

Future SDK upgrades may change the chain order. If that happens, keep kubelogin docs aligned with the SDK rather than hard-coding a local chain unless the team explicitly decides to diverge.

## Governance

- All meaningful changes require team consensus
- Document architectural decisions here
- Keep history focused on work, decisions focused on direction
