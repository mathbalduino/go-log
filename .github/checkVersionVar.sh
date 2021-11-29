#!/bin/bash

CODE_VERSION=$(grep -Po 'const LibraryModuleVersion = "\K[^"]*' internal/libraryInfo.go)

if [ "$CODE_VERSION" = "$1" ]; then
  exit 0
else
  echo -e >&2 "\033[31;1m[ Error ] Change the LibraryModuleVersion constant variable value: \n\tCurrent Value: $CODE_VERSION\n\tExpected Value: $1\033[0m"
  exit 1
fi
