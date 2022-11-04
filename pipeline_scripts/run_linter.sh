#!/bin/bash
set -x
printf "\nRunning hadolint on Dockerfile ...\n"

MSYS_NO_PATHCONV=1 docker run --rm -i \
                              -v $(pwd):/lint \
                              -w /lint \
                              hadolint/hadolint:v1.17.5-2-g44d0caa-alpine \
                              hadolint --config hadolint_config.yml Dockerfile