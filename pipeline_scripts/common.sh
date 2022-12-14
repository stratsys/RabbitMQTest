#!/bin/bash
source ./build_config.sh
if [[ "$COMPLETE_RELEASE_BRANCH" =~ refs/tags/(.+)$ ]]; then
    TAG_NAME=${BASH_REMATCH[1]}
else
    TAG_NAME=preview-${RELEASE_BRANCH}-${TIME_STAMP}
fi
IMAGE_TAG="$CONTAINER_REGISTRY/$CONTAINER_REGISTRY_REPOSITORY:${TAG_NAME}"
[[ "$PUSH_LATEST" == "1" ]] && IMAGE_TAG_LATEST="$CONTAINER_REGISTRY/$CONTAINER_REGISTRY_REPOSITORY:latest"