#!/usr/bin/env bash

AppName=tests
BuildPath=build
BuildTags="debug"


mkdir -p $BuildPath
clear
go mod tidy &&
go build -o $BuildPath/$AppName -v -tags $BuildTags && 
./$BuildPath/$AppName


# # go build -o $BuildPath/$AppName -v -gcflags=all="-N -l" -tags "debug hints" && \