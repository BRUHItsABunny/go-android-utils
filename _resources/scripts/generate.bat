@echo off
docker build -t protos .
docker run --name="protos_run" protos
docker cp protos_run:proto/github.com/BRUHItsABunny/go-android-utils/. .
docker rm protos_run
docker rmi protos
echo "BAT done"