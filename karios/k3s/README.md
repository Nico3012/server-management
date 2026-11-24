# Standard image hadron mit k3s:
`quay.io/kairos/hadron:v0.6.3-standard-amd64-generic-v4.7.0-k3s-v1.37.2-k3s1`




docker run --rm --network host \
  -v "$PWD/build:/output" \
  -v "$PWD/cloud-config.yaml:/config.yaml" \
  quay.io/kairos/auroraboot:v4.7.0 \
  build-iso --cloud-config /config.yaml --output /output/ oci:quay.io/kairos/hadron:v0.6.3-standard-amd64-generic-v4.7.0-k3s-v1.37.2-k3s1



# Todo
Updates über Kairos Operator über Kubernetes API steuern
 