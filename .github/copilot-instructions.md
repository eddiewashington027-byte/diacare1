<!-- Brief, repo-specific instructions for AI coding agents -->
# Copilot / AI Agent Instructions

Purpose: quickly orient an AI coding agent to this repository so it can be productive without guesswork.

- Repo snapshot: this repository currently contains only the project README: [README.md](README.md).
- Dev environment: development runs in a dev container (Ubuntu 24.04.3 LTS) per workspace configuration.

What to do first
- Inventory files and branches: run `git ls-files` and `git branch -v` to see tracked files and active branches.
- Look for common language indicators: check for `package.json`, `pyproject.toml`, `requirements.txt`, `go.mod`, `Cargo.toml`, `Dockerfile`, `.devcontainer`, `Makefile`, `src/`, and `tests/`.
- If the repo is empty of source (only docs), open [README.md](README.md) and surface project goals or next steps to the human maintainer.

Repository-specific guidance
- This repo currently has no source tree or CI config to inspect. Avoid making assumptions about a build system—ask the user before adding language-specific scaffolding.
- If asked to scaffold code or tests, propose a minimal structure and include a one-command way to run it (e.g., `python -m venv .venv && .venv/bin/pip install -r requirements.txt`).

Search & discovery patterns
- Use these commands to discover hidden or ignored files safely:
  - `git status --porcelain --ignored` (shows ignored files under version control rules)
  - `find . -maxdepth 3 -type f -name "*.md" -o -name "Dockerfile" -o -name "*.yml"`
- Prefer reading manifest files (`package.json`, `pyproject.toml`, `pom.xml`) to infer build/test commands rather than guessing.

How to update this file
- If `.github/copilot-instructions.md` already exists, merge by:
  1. Preserving unique, human-written rationale or project history.
  2. Updating any stale environment commands (validate they run in the dev container).
 3. Adding concrete file references discovered during repo inspection (e.g., `src/` path, `Makefile` targets).

If you need clarification
- Ask the repository owner what the primary language/runtime should be, and whether CI or test scaffolding is desired before creating new files.

Example next questions for the human
- "Should I scaffold a language (Python/Node/Go) for you? If so, which one?"
- "Do you want CI (GitHub Actions) added now, or only local scaffolding?"

Contact / meta
- Repo: eddiewashington027-byte/diacare1 (branch: main). Update this file if repository structure changes.
