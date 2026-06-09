#!/bin/bash

echo -e "\033[0;32mDeploying updates to Github...\033[0m"

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

# Deploy to gh-pages
echo -e "\033[0;32mDeploying to gh-pages branch...\033[0m"
REPO_URL=$(git config --get remote.origin.url)
cd public
git init
git add -A
git commit -m "Deploy site on `date`"
git push -f "$REPO_URL" HEAD:gh-pages
rm -rf .git
cd ..
