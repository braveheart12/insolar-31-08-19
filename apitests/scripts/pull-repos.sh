#!/bin/#!/usr/bin/env bash

WORKDIR=~/go/src/github.com/insolar
REPOS=( "insolar-api" "insolar-observer-api" "insolar-internal-api" )

for repo in "${REPOS[@]}"
do
  echo checking ${repo}...
  cd ${WORKDIR}/${repo} || exit
  git stash
  git checkout tags/1.0.0
  git pull origin tags/1.0.0
done

