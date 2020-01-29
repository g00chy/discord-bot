docker-compose run --rm build_arm
echo "builded arm-----------------"
docker-compose run --rm build_arm64
echo "builded arm64-----------------"
docker-compose run --rm build_amd64
echo "builded amd64-----------------"