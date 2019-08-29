#!/bin/#!/usr/bin/env bash

WORKDIR=~/go/src/github.com/insolar
REPOS=( "insolar-api" "insolar-observer-api" "insolar-internal-api" )

for repo in "${REPOS[@]}"
do
  echo checking ${repo}...
  cd ${WORKDIR}/${repo} || exit
  git stash
  git checkout 1.x.x
  git pull origin 1.x.x
done

