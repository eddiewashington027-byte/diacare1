# diacare1

## Getting started (scaffold)

This repository currently includes a minimal Python scaffold, tests, and a CI workflow.

Local quickstart (recommended on your machine):

```bash
# Ensure system venv support is installed (Debian/Ubuntu):
sudo apt update && sudo apt install -y python3-venv

python3 -m venv .venv
.venv/bin/python -m pip install --upgrade pip
.venv/bin/pip install -r requirements-dev.txt
.venv/bin/pytest -q
```

Lint and typecheck locally:

```bash
.venv/bin/flake8 .
.venv/bin/mypy src
```

Pre-commit (format-on-commit):

```bash
pip install -r requirements-dev.txt
pre-commit install
pre-commit run --all-files
```

Go quickstart:

```bash
# Requires Go 1.20+
go test ./...  # run Go tests
go run ./cmd/diacare
```

Note: the dev container used for this workspace may not include `python3-venv`.
If you see errors creating a virtualenv inside the container, run the commands above
on your local machine or enable `python3-venv` in the container image.

CI: see `.github/workflows/ci.yml` for the GitHub Actions configuration; the workflow
installs dependencies and runs `pytest` on Ubuntu runners.

Dia Care is a diabetes management application designed to bridge the gap between patients and healthcare professionals. It primarily functions as a digital logbook and remote monitoring tool, allowing users to track blood glucose levels, medication, and lifestyle factors while enabling doctors to view this data in real-time.
