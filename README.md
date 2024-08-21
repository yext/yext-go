# yext-go

Go client for the Yext API

[![GoDoc](https://godoc.org/github.com/yext/yext-go?status.svg)](https://godoc.org/github.com/yext/yext-go)
[![Build Status](https://travis-ci.org/yext/yext-go.svg?branch=v2)](https://travis-ci.org/yext/yext-go)

## How to Use Locally With Congo
*  `cd` into this repo's root
*  checkout the latest changes: `git checkout entities && git pull`
*  run `go mod init gopkg.in/yext/yext-go.v2` to initialize the repo as a Go module
*  run `echo "\nreplace gopkg.in/yext/yext-go.v2 => $(pwd)/ // TODO($(whoami)): REMOVE! DO NOT MERGE" >> "$CONGO/go.mod"`
    * This tells Go to use your local changes instead of the fetched package
    * **Don't forget to remove that line before making any commits to Congo**
* You should be able to `cmd + click` on anything that lives in `yext-go` like `yext.Entity` and your IDE should show your local path in the crumbs/file path
    * If you still see `gopkg.in/yext/yext-go.v2@v2.0.0-20240815032338-bcdcdcad0cfb` type of paths, you might need to run run `cd $CONGO && go mod tidy` to fix the issue
    * The correct path will match what is added to the bottom of `$CONGO/go.mod` ex: `/Users/myuser/repo/yext-go`

## Deploying Changes
* Ensure your `$CONGO/go.mod` does NOT have the `replace` line with your local `yext-go` path from the section above
* Open the new commit you want ship in `git log` or Github and copy the commit hash
    * ex: `2e0a9278c9e8d6a354566770512f92e4822a0f6e`
* Then run `cd $CONGO && go get gopkg.in/yext/yext-go.v2@COMMIT_HASH` with `COMMIT_HASH` swapped for hash you just copied
    * ex: `go get gopkg.in/yext/yext-go.v2@2e0a9278c9e8d6a354566770512f92e4822a0f6e` 
* run `make updategomod updatebazeldeps` to update the go mod/sum and bazel files
* make a commit with those changes, wait for `+1` from CongoTestBed and then you should be good to ship it 🥳
    * ex: https://gerrit.yext.com/c/congo/+/259564
