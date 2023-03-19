#!/bin/bash

# get mode (update or add)
mode=$1

# get filename and file extension
name=$(git status -s | sed -nE 's/.*([0-9]+.*)\..*/\1/p')
extension=$(git status -s | sed -nE 's/.*\.([a-z]+)"/\1/p')

if [[ mode == "update" ]]
then
  git add :
  git cm -m "$name in $extension"
else
  git add :
  git cm -m "update $name"
fi

git push origin master
