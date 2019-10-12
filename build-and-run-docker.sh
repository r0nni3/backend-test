#!/bin/bash

# Builds and creates container image with just binary
# it uses Docker mulstage builds to simplify process
# for more info see the Dockerfile
docker build -t backend-test .

# This executes the binary in a docker container.
# - passes exports context
# - runs command with arguments
docker run -v $PWD/feed-exports:/feed-exports --rm backend-test import /feed-exports/
