#!/bin/bash

# Define your GitHub repository
REPO_OWNER="jere-mie"
REPO_NAME="rickroll-me"

# Extract the tag from version.txt
TAG=$(<version.txt)

# Create a release
gh release create $TAG \
    --repo $REPO_OWNER/$REPO_NAME \
    --title "Release $TAG" \
    --notes "Release notes for $TAG"

# Upload built files to the release
for file in $(ls bin/*); do
    gh release upload $TAG $file
done

echo "Release created successfully."
