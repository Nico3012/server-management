# Image bauen
docker build -t mein-hadron-webserver:v1.0.0 .


# Image speichern
mkdir -p build
docker save mein-hadron-webserver:v1.0.0 -o build/image.tar



# Image hochladen
docker run -d --name local-reg -p 5000:5000 registry:2
docker tag mein-hadron-webserver:v1.0.0 localhost:5000/mein-hadron-webserver:v1.0.0
docker push localhost:5000/mein-hadron-webserver:v1.0.0


# bauen
docker run --rm --network host \
  -v "$PWD/build:/output" \
  -v "$PWD/cloud-config.yaml:/config.yaml" \
  quay.io/kairos/auroraboot:v0.28.0 \
  build-iso --cloud-config /config.yaml --output /output/ oci:localhost:5000/mein-hadron-webserver:v1.0.0



# cleanup
docker rm -f local-reg
rm -f build/image.tar
