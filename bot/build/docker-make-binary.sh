#!/usr/bin/env bash
export UID=${UID}
export GID=${GID}
docker-compose run --rm build_arm
docker-compose run --rm build_arm64
docker-compose run --rm build_amd64