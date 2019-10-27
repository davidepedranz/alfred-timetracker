#!/bin/sh

set -o errexit
set -o nounset

archive="TimeTracker-${VERSION}.alfredworkflow"

echo "Building go binaries:"
for path in ./cmd/*; do
  d=$(echo "${path}" | sed "s|.*/||")
  echo "  - ${d}"
  GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".alfred/bin/${d}" "./cmd/${d}"
done

echo ""
echo "Crearing archive:"
(
  cd ./.alfred || exit
  envsubst >./info.plist <./info.plist.template
  zip -r "../${archive}" ./*
  zip -d "../${archive}" info.plist.template
)

echo ""
echo "Build completed: \"${archive}\""
