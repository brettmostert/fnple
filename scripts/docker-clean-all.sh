#!/usr/bin/env bash

# Removes ALL images for "fnple"
echo
echo "==> Checking fnple images to remove:"
docker image ls -f label="app=fnple"

IMAGES=$(docker image ls -q -f label="app=fnple")

if [[ -n ${IMAGES} ]]; then
	docker image rm -f "${IMAGES}"
fi

# Done!
echo
echo "==> Results:"
docker image ls -f label="app=fnple"
