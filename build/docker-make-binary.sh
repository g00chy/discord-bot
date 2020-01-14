#!/usr/bin/env bash
docker-compose run --rm build_arm
docker-compose run --rm build_arm64
docker-compose run --rm build_amd64