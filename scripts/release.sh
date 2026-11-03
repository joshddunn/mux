#!/usr/bin/env bash
# Release a new version of mux and bump the Homebrew formula.
#
# Usage: scripts/release.sh v0.1.12
#
# Expects the tap repo at ../homebrew-tap (override with TAP_DIR).

set -euo pipefail

VERSION="${1:-}"
REPO="joshddunn/mux"
TAP="joshddunn/tap"
MUX_DIR="$(cd "$(dirname "$0")/.." && pwd)"
TAP_DIR="${TAP_DIR:-$MUX_DIR/../homebrew-tap}"
FORMULA="$TAP_DIR/mux.rb"

main() {
  check_version
  check_repo "$MUX_DIR"
  check_repo "$TAP_DIR"
  check_tag_is_new
  confirm

  step "Tagging $VERSION"
  git -C "$MUX_DIR" tag "$VERSION"
  git -C "$MUX_DIR" push origin "$VERSION"

  step "Bumping formula"
  local sha
  sha="$(tarball_sha)"
  update_formula "$sha"
  git -C "$TAP_DIR" commit -am "mux $VERSION"
  git -C "$TAP_DIR" push origin main

  step "Upgrading local install"
  upgrade_local
}

check_version() {
  [[ "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || fail "usage: $0 vX.Y.Z"
  [[ -f "$FORMULA" ]] || fail "formula not found at $FORMULA (set TAP_DIR)"
}

check_repo() {
  local dir="$1"
  [[ -z "$(git -C "$dir" status --porcelain)" ]] || fail "$dir has uncommitted changes"
  [[ "$(git -C "$dir" branch --show-current)" == "main" ]] || fail "$dir is not on main"
  git -C "$dir" fetch --quiet origin
  [[ "$(git -C "$dir" rev-parse HEAD)" == "$(git -C "$dir" rev-parse origin/main)" ]] ||
    fail "$dir is not in sync with origin/main"
}

# Never move a tag: anyone who already downloaded it would get a checksum mismatch.
check_tag_is_new() {
  if git -C "$MUX_DIR" ls-remote --exit-code --tags origin "refs/tags/$VERSION" >/dev/null; then
    fail "tag $VERSION already exists on origin; pick a new version"
  fi
}

confirm() {
  local answer
  read -r -p "Release $VERSION? [y/N] " answer
  [[ "$answer" == [yY] ]] || fail "cancelled"
}

# Hash the tarball GitHub serves, not a local `git archive`, which produces different bytes.
# GitHub can take a few seconds to serve a new tag, so retry.
tarball_sha() {
  local url="https://github.com/$REPO/archive/refs/tags/$VERSION.tar.gz"
  curl -fsSL --retry 5 --retry-all-errors --retry-delay 3 "$url" | shasum -a 256 | cut -d' ' -f1
}

update_formula() {
  local sha="$1"
  sed -i '' -E \
    -e "s/^( *version )\".*\"/\1\"$VERSION\"/" \
    -e "s/^( *sha256 )\".*\"/\1\"$sha\"/" \
    "$FORMULA"
  git -C "$TAP_DIR" diff
}

# Homebrew builds from its own clone of the tap, which `brew upgrade` only pulls once every
# 24 hours. Pull it now so the upgrade sees the new formula.
upgrade_local() {
  git -C "$(brew --repo "$TAP")" pull --ff-only --quiet
  if brew list --formula mux >/dev/null 2>&1; then
    brew upgrade "$TAP/mux"
  else
    brew install "$TAP/mux"
  fi
}

step() {
  echo "==> $*"
}

fail() {
  echo "error: $*" >&2
  exit 1
}

main "$@"
