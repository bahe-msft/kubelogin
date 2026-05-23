# Squad Decisions

## Active Decisions

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

## Governance

- All meaningful changes require team consensus
- Document architectural decisions here
- Keep history focused on work, decisions focused on direction
