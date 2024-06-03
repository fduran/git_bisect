# Git Bisect

## What and Why

Trivial example to learn and practice [git bisect](https://git-scm.com/docs/git-bisect)

## Problem Description

This repository contains a Golang program and a test for it. 

To execute the test, run `go test`. The last commit fails the test. Suppose the first commit passed the test.

Problem: find the commit that first broke the test.

## Solution

```bash
# Find first commit
git log

# Start the bisect process
git bisect start

# Mark the bad commit
git bisect bad HEAD

# Mark the good commit (first one)
git bisect good 9e80a7eb1b09385e93ab4a76cb2c93beec48fd9f

# Run tests automatically
git bisect run go test

# End the bisect process
git bisect reset
```

Solution: `2e44089778e44dcd9b97aa3baacdcff10311841b is the first bad commit` (message `4th`)



