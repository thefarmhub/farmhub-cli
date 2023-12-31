#!/bin/bash

# Function to download and unpack file
download_and_unpack() {
    url=$1
    filename=$2

    # Download
    if command -v curl > /dev/null; then
        curl -L -o "$filename" "$url"
    elif command -v wget > /dev/null; then
        wget -O "$filename" "$url"
    else
        echo "Error: curl or wget is required to download files."
        exit 1
    fi

    # Unpack
    case $filename in
        *.tar.gz) tar -xzf "$filename" ;;
        *.zip) unzip "$filename" ;;
        *) echo "Cannot unpack $filename, unknown format"; exit 1 ;;
    esac
}

# Function to get the latest release version from GitHub API
get_latest_release() {
    curl --silent "https://api.github.com/repos/thefarmhub/farmhub-cli/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
}

# Determine OS and Architecture
os=""
arch=""
case "$(uname -s)" in
    Darwin) os="Darwin";;
    Linux)  os="Linux";;
    CYGWIN*|MINGW32*|MSYS*|MINGW*) os="Windows";;
    *) echo "Unsupported OS"; exit 1;;
esac

case "$(uname -m)" in
    x86_64) arch="x86_64";;
    arm64) arch="arm64";;
    i386) arch="i386";;
    *) echo "Unsupported architecture"; exit 1;;
esac

# Get the latest version
version=$(get_latest_release)
if [ -z "$version" ]; then
    echo "Error: Could not get latest version."
    exit 1
fi

# Set download URL and file name
file="farmhub_${os}_${arch}"
extension=".tar.gz"
if [ "$os" = "Windows" ]; then
    extension=".zip"
fi
file="${file}${extension}"
url="https://github.com/thefarmhub/farmhub-cli/releases/download/${version}/${file}"

# Download and unpack the file
download_and_unpack "$url" "$file"

echo "Download: $file"
echo "You can now use it by running: ./farmhub"
