---
layout:      post
title:      "Document your project with GitHub Pages"
subtitle:   "One Project, one repo, one document"
date:        2020-09-25
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - github
  - devops
---

## GitHub Pages

> [Here is a nice example](https://release-engineering.github.io/exodus-gw/)

#### Use sphinx-build

```
# part of .travis.yml
script: 
  - postmaster --version
  - python setup.py install
  # python scripts/conn.py
  # pytest
  # awslocal s3 mb s3://my-test-bucket
  # python scripts/localstack.py
  - pip install -r requirements.txt
  - sphinx-build -M html docs docs/_build

after_success:
- scripts/push-docs
```

#### Auth with GitHub Token

1. Remember to give `scripts/push-docs` an `executable` permission.
2. Set environment variable in **travis-ci**. like `GITHUB_TOKEN, GITHUB_AUTHOR, GITHUB_EMAIL`

```
# part of scripts/push-docs

cd docs/_build/html

git init
git config user.name "${GITHUB_AUTHOR}"
git config user.email "${GITHUB_EMAIL}"

if enabled; then
  git remote add origin "https://$GITHUB_TOKEN@github.com/$ORG/$REPO.git"
  git fetch origin
  if git rev-parse origin/gh-pages; then
    git reset origin/gh-pages
  fi
fi

git add -A .
git commit -m "Build documentation for ${version} at ${rev}"
```

