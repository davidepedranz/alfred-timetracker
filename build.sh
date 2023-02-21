#!/bin/sh

set -o errexit
set -o nounset

archive="TimeTracker-${VERSION}.alfredworkflow"

echo "Building go binaries:"
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/tt-amd64" .
GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/tt-arm64" .

echo ""
echo "Crearing archive:"
(
  cd ./.workflow || exit
  envsubst >./info.plist <./info.plist.template
  zip -r "../${archive}" ./*
  zip -d "../${archive}" info.plist.template
)

echo ""
echo "Build completed: \"${archive}\""
