#!/bin/bash

CODE_VERSION=$(grep -Po 'const LibraryModuleVersion = "v\K[[:digit:]].[[:digit:]].[[:digit:]]' internal/libraryInfo.go)
CODE_MAJOR=$(echo $CODE_VERSION | cut -d '.' -f1)
CODE_MINOR=$(echo $CODE_VERSION | cut -d '.' -f2)
CODE_PATCH=$(echo $CODE_VERSION | cut -d '.' -f3)

TAG_VERSION=$(git describe --tags --abbrev=0 | grep -Po 'v\K[[:digit:]].[[:digit:]].[[:digit:]]')
TAG_MAJOR=$(echo $TAG_VERSION | cut -d '.' -f1)
TAG_MINOR=$(echo $TAG_VERSION | cut -d '.' -f2)
TAG_PATCH=$(echo $TAG_VERSION | cut -d '.' -f3)

ERROR_MSG="\033[31;1m[ Error ] The 'LibraryModuleVersion' variable value (at internal/libraryInfo.go) needs to be newer than the latest one:\n\tLibraryModuleVersion: v$CODE_VERSION\n\tLatest tag: v$TAG_VERSION\033[0m"
if [ $TAG_MAJOR -gt $CODE_MAJOR ]; then
  echo -e >&2 $ERROR_MSG
  exit 1
fi
if [ $TAG_MAJOR -eq $CODE_MAJOR ]; then
  if [ $TAG_MINOR -gt $CODE_MINOR ]; then
    echo -e >&2 $ERROR_MSG
    exit 1
  fi
  if [ $TAG_MINOR -eq $CODE_MINOR ]; then
    if [ $TAG_PATCH -ge $CODE_PATCH ]; then
      echo -e >&2 $ERROR_MSG
      exit 1
    fi
  fi
fi

exit 0
