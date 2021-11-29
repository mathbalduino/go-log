#!/bin/bash

grep -Po 'const LibraryModuleVersion = "v\K[[:digit:]].[[:digit:]].[[:digit:]]' internal/libraryInfo.go
