#!/usr/bin/env bash

FILES=shaders/*

for f in $FILES
do
if [[ "$f" != *\.spv ]]
then
  echo $f
    glslc $f -o $f.spv
fi
done
