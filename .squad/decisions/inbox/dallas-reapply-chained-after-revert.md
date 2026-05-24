### 2026-05-23: Reapply chained login after revert
**By:** Dallas
**What:** Rebased the chained login implementation onto the current post-revert main branch instead of using the stale branch that was already merged and reverted.
**Why:** A PR from the stale branch would compare as already merged on GitHub; a fresh commit on current main is required to reintroduce the intended code and tests for issue #1.
