# Git Workflow & Branch History

## Main Branch Structure

All development work is done directly on `main`. History starts with independent CI/tooling commits, followed by feature merges.

### Branch Organization

- **CI/Tooling commits** (first): Infrastructure and CI updates that don't depend on any new features (e.g., GitHub Actions workflows, Docker config)
- **Feature merges** (after CI): Integration of feature branches into `main` via `merge --no-ff` commits

### Example History

```
* Merge feature/generalize-auto-sync into main    (feature merge)
|\
| * Fix broken v-btn tag in Settings.vue           (feature branch commits)
| * Generalize auto-sync scheduling for photo sources
* | Merge feature/ha-ingress-auth into main       (feature merge)
|\
| * Add Home Assistant ingress auth bypass        (feature branch commits)
* | Add GHCR publish workflow                      (independent CI/tooling)
|/
* upstream/main base
```

## Integration Pattern

- **Feature branches** (`feature/*`): Contain 1+ commits scoped to that feature only
- **Integration into main**: Use merge commits (`git merge --no-ff`) rather than direct feature commits on `main`
- **CI/Tooling commits**: Direct commits on `main` only (no feature branches for CI)
- **Direct feature commits on main**: Avoid; use feature branches + merge instead

## Rationale

This structure:
- Keeps main organized and intentional
- Makes features easy to revert (entire merge can be reverted)
- Maintains clear separation between tooling/infra and feature work
- Allows standalone CI improvements independent of feature development
