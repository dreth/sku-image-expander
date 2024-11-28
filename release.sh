#!/bin/bash

# Simple GoReleaser bump script
# Usage: ./bump.sh [major|minor|patch]

# Ensure script stops on errors
set -e

# Check if a bump level is passed (major, minor, patch)
if [ -z "$1" ]; then
  echo "Usage: $0 [major|minor|patch]"
  exit 1
fi

# Get the current version from the latest Git tag
CURRENT_VERSION=$(git describe --tags --abbrev=0)

# Strip any leading 'v' (if using vX.X.X style tags)
CURRENT_VERSION=${CURRENT_VERSION#v}

# Bump the version using the specified level
IFS='.' read -r -a VERSION_PARTS <<< "$CURRENT_VERSION"

MAJOR=${VERSION_PARTS[0]}
MINOR=${VERSION_PARTS[1]}
PATCH=${VERSION_PARTS[2]}

case "$1" in
  major)
    MAJOR=$((MAJOR + 1))
    MINOR=0
    PATCH=0
    ;;
  minor)
    MINOR=$((MINOR + 1))
    PATCH=0
    ;;
  patch)
    PATCH=$((PATCH + 1))
    ;;
  *)
    echo "Unknown bump level: $1"
    exit 1
    ;;
esac

# Form the new version
NEW_VERSION="$MAJOR.$MINOR.$PATCH"

# Optionally, prefix the version with 'v'
NEW_TAG="v$NEW_VERSION"

# Update your version in your project files if necessary
# You can modify this section if your project requires updating a version file.
echo "Bumping version to $NEW_TAG"

# Commit and create a Git tag
git commit --allow-empty -m "chore(release): $NEW_TAG"
git tag "$NEW_TAG"

# Push changes and the new tag
git push origin main
git push origin "$NEW_TAG"

# Run GoReleaser
goreleaser release --clean

# Push binaries (GoReleaser handles this part)
echo "Release $NEW_TAG created successfully."
