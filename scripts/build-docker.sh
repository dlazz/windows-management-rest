APP_NAME="wrm"
SHA=$(git rev-parse HEAD)
MAIN_VERSION="0.1"
VERSION="${MAIN_VERSION}.${SHA:0:6}"
IMAGE_NAME="${APP_NAME}:${VERSION}"
# Building and tagging image

docker build --rm \
             --build-arg VERSION=${VERSION} \
             -t ${IMAGE_NAME} .


echo "Running container"
docker run --name ${APP_NAME} -dit ${IMAGE_NAME}

echo "copy file from container"
docker cp ${APP_NAME}:/go/build .
echo "archiving file"
zip ./"v.${APP_NAME}-${VERSION}.zip" -r ./build/*

echo "cleaning up"
docker stop ${APP_NAME}
docker rm ${APP_NAME}
docker rmi ${IMAGE_NAME}