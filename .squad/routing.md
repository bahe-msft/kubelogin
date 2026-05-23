# Work Routing

How to decide who handles what.

## Routing Table

| Work Type | Route To | Examples |
|-----------|----------|----------|
| Architecture, scope, design, review | Ripley | Cross-package changes, compatibility decisions, reviewer gates |
| Go implementation | Dallas | CLI commands, token providers, cache behavior, kubeconfig conversion |
| Security and auth review | Lambert | Secret handling, credential flows, cache protection, threat-focused review |
| Testing and validation | Parker | Unit tests, VCR fixtures, regression tests, edge cases |
| Documentation and release workflow | Ash | Docs site, CLI reference, changelog, release notes |
| Session logging | Scribe | Automatic — never needs routing |
| Work queue monitoring | Ralph | Backlog scan, issue follow-up, keep-alive work loops |

## Issue Routing

| Label | Action | Who |
|-------|--------|-----|
| `squad` | Triage: analyze issue, assign `squad:{member}` label | Ripley |
| `squad:ripley` | Pick up architecture/scope/review work | Ripley |
| `squad:dallas` | Pick up Go implementation work | Dallas |
| `squad:lambert` | Pick up security/auth review work | Lambert |
| `squad:parker` | Pick up test/quality work | Parker |
| `squad:ash` | Pick up docs/release work | Ash |

### How Issue Assignment Works

1. When a GitHub issue gets the `squad` label, **Ripley** triages it — analyzing content, assigning the right `squad:{member}` label, and commenting with triage notes.
2. When a `squad:{member}` label is applied, that member picks up the issue in their next session.
3. Members can reassign by removing their label and adding another member's label.
4. The `squad` label is the inbox for untriaged issues.

## Rules

1. **Eager by default** — spawn all agents who could usefully start work, including anticipatory downstream work.
2. **Scribe always runs** after substantial work, always as `mode: "background"`. Never blocks.
3. **Quick facts -> coordinator answers directly.** Don't spawn an agent for simple repo-status questions.
4. **When two agents could handle it**, pick the one whose domain is the primary concern.
5. **"Team, ..." -> fan-out.** Spawn all relevant agents in parallel as `mode: "background"`.
6. **Anticipate downstream work.** If a feature is being built, spawn Parker to write test cases from requirements simultaneously.
7. **Issue-labeled work** — when a `squad:{member}` label is applied to an issue, route to that member. Ripley handles the base `squad` label.
