#!/bin/bash

baseurl="https://raw.githubusercontent.com/JavaNut13/longrun/master"

os=$(uname)
if [ "$os" == "Darwin" ]; then
  ext="osx"
elif [ "$os" == "Linux" ]; then
  arch=$(lscpu | head -n 1 | cut -f 2 -d ":" | xargs)
  ext="linux_$arch"
else
  echo "Unknown/ unsupported OS: $os"
  exit 1
fi

tmpfile="/tmp/longrun-$$"

curl -s "$baseurl/builds/longrun_$ext" > $tmpfile

notfound=$(head -n 1 $tmpfile | grep -ic "not found")

if [ $notfound -eq 0 ]; then
  sudo cp $tmpfile "/usr/local/bin/push"
  echo "Installed longrun to '/usr/local/bin/push'"
else
  echo "Couldn't find binary for $ext"
  echo "Just install go and build the file yourself"
  exit 2
fi