#!/usr/bin/env bash

basepath=$(pwd)
cd gopath/src/github.com/cloudfoundry/bosh-init


semver=`cat ${basepath}/version/number`
timestamp=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
git_rev=`git rev-parse --short HEAD`

version="${semver}-${git_rev}-${timestamp}"

echo "version: ${version}"
echo $version > VERSION.txt
