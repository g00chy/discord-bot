echo "builded arm-----------------"
docker-compose -f build-docker-compose.yml run --rm build_arm
echo "builded arm64-----------------"
docker-compose -f build-docker-compose.yml run --rm build_arm64
echo "builded amd64-----------------"
docker-compose -f build-docker-compose.yml run --rm build_amd64
