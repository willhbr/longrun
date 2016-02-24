#!/bin/bash

# Install binary

baseurl="https://raw.githubusercontent.com/JavaNut13/longrun/master"

os=$(uname)
if [ "$os" == "Darwin" ]; then
  ext="osx"
elif [ "$os" == "Linux" ]; then
  arch=$(lscpu | head -n 1 | cut -f 2 -d ":" | xargs)
  isarm=$(echo $arch | grep -ic "arm")
  if [ $isarm -eq 1 ]; then
    ext="linux_arm"
  else
    ext="linux_$arch"
  fi
else
  echo "Unknown/ unsupported OS: $os"
  exit 1
fi

tmpfile="/tmp/longrun-$$"

echo "Finding & downloading binary for $ext"

curl -s "$baseurl/builds/longrun_$ext" > $tmpfile

notfound=$(head -n 1 $tmpfile | grep -ic "not found")

if [ $notfound -eq 0 ]; then
  path="/usr/local/bin/lrun"
  sudo cp $tmpfile $path
  if [ $? -eq 0 ]; then
    sudo chmod +x $path
    echo "Installed longrun to '$path'"
  else
    echo "Couldn't copy binary to '$path'"
    echo "You can copy it yourself from '$tmpfile'"
    exit 3
  fi
else
  echo "Couldn't find binary for $ext"
  echo "Just install go and build the file yourself"
  exit 2
fi

# Read key

echo -n "Enter your PushBullet API key (from your settings page) "
read api_key

echo $api_key > "$HOME/.longrun-token"