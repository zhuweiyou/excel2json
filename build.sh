#!/bin/bash

export CGO_ENABLED=0

platforms=(
  windows/amd64
  windows/arm64
  linux/amd64
  linux/arm64
  darwin/amd64
  darwin/arm64
)

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  export GOOS=${platform_split[0]}
  export GOARCH=${platform_split[1]}

  output_name="excel-to-json_${GOOS}_${GOARCH}"

  if [ $GOOS = windows ]; then
    output_name+='.exe'
  fi

  go build -ldflags="-s -w" -o "$output_name"

  sha256sum "$output_name" > "$output_name.sha256"
done
