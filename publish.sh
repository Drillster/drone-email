#!/bin/bash

tags=$(git tag --points-at HEAD)
IFS=$'\n' read -rd '' -a taglist <<<"$tags"

if [ ${#taglist[@]} -gt 0 ]; then
    echo "Tagging Docker images with: ${taglist[@]}"
    for tag in "${taglist[@]}"; do
        docker tag drillster/drone-email:latest drillster/drone-email:$tag
    done
fi

echo "Pushing Docker images..."
docker push drillster/drone-email:latest

if [ "$?" -ne "0" ]; then
    echo "Failed to push image, exiting!"
    exit 1
fi

if [ ${#taglist[@]} -gt 0 ]; then
    for tag in "${taglist[@]}"; do
        docker push drillster/drone-email:$tag
    done
fi