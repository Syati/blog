#!/bin/bash

echo -e "\033[0;32mDeploying updates to Github...\033[0m"

# Build search index with Go
go run cmd/indexer/main.go

# Build the project.
hugo

# Add changes to git.
git add -A

# Commit changes.
msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin

# Deploy to gh-pages with force push (local build is source of truth)
git push -f origin $(git subtree split --prefix public):refs/heads/gh-pages
