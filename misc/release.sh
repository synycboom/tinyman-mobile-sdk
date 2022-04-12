#!/bin/bash

if [ "$1" = "" ] || [ "$2" = "" ] || [ "$3" = "" ] || [ "$4" = "" ] || [ "$5" = "" ]
then
  echo "Usage: $0 <version> <android_home> <android_ndk_home> <go_mobile> <module>"
  exit
fi

mkdir -p release
mkdir -p release/$1
rm -rf release/$1/*

ANDROID_HOME=$2 ANDROID_NDK_HOME=$3 go run $4 bind -target=android -o release/$1/tinyman.aar $5
go run $4 bind -target=ios -o release/$1/Tinyman.xcframework $5

zip -r release/$1/tinyman-mobile-ios-$1.zip release/$1/Tinyman.xcframework/
zip -r release/$1/tinyman-mobile-android-$1.zip release/$1/tinyman.aar release/$1/tinyman-sources.jar