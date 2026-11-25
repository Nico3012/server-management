# Standard image hadron mit k3s:
`quay.io/kairos/hadron:v0.6.3-standard-amd64-generic-v4.7.0-k3s-v1.37.2-k3s1`



```shell
docker run --rm --network host \
  -v "$PWD/build:/output" \
  -v "$PWD/cloud-config.yaml:/config.yaml" \
  quay.io/kairos/auroraboot:v4.7.0 \
  build-iso --cloud-config /config.yaml --output /output/ oci:quay.io/kairos/hadron:v0.5.1-standard-amd64-generic-v4.3.0-k3s-v1.36.4-k3s1
```




# Updates manuell
### Beispielupdate mit neuerem Tag
```shell
sudo kairos-agent state
sudo kubectl get nodes

# Bei Tag auswahl, muss NUR drauf geachtet werden, dass k3s genau einen minor step vorwärts macht. Hadron oder Kairos version sind egal.
sudo kairos-agent upgrade --source oci:quay.io/kairos/hadron:v0.6.3-standard-amd64-generic-v4.7.0-k3s-v1.37.2-k3s1

sudo reboot

sudo kairos-agent state
sudo kubectl get nodes
```

# Todo
Updates über Kairos Operator über Kubernetes API steuern
