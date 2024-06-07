#!/bin/bash

# Define the output directory
OUTPUT_DIR="bin"

# Create the output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# Define the platforms you want to cross-compile for
PLATFORMS=(
    "linux/amd64"
    "linux/386"
    "linux/arm"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/arm64"
)

# Loop through each platform and compile the binary
for PLATFORM in "${PLATFORMS[@]}"
do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    
    OUTPUT_NAME="rrm_${GOOS}_${GOARCH}"
    
    if [ $GOOS = "windows" ]; then
        OUTPUT_NAME="${OUTPUT_NAME}.exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_DIR/$OUTPUT_NAME .
done

chmod +x $OUTPUT_DIR/*

echo "Builds complete. Output directory: $OUTPUT_DIR"