#!/bin/bash

CODE_VERSION=$(grep -Po 'const LibraryModuleVersion = "\K[^"]*' internal/libraryInfo.go)
GIT_TAG_VERSION=$(git describe --tags --abbrev=0)

if [ "$CODE_VERSION" = "$GIT_TAG_VERSION" ]; then
  exit 0
else
  echo -e >&2 "\033[31;1m[ Error ] Change the LibraryModuleVersion constant variable value: \n\tCurrent Value: $CODE_VERSION\n\tExpected Value: $GIT_TAG_VERSION\033[0m"
  exit 1
fi
