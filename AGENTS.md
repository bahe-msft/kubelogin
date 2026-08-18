# Agent Guidelines

## Testing

- New table-driven Go tests must call `t.Parallel()` in the parent test and in each subtest when the cases do not share mutable state.
- Keep parallel test cases isolated: create flags, options, temporary files, and other mutable dependencies per subtest.
