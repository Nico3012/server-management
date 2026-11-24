# Einleitung
Dieser guide startet sobald der k3s server läuft und möchte den k3s Server konfigurieren, so dass ein hallo welt image läuft

# Cmds
```shell
ssh -t nico@192.168.178.103 "sudo cat /etc/rancher/k3s/k3s.yaml" | sed 's/127.0.0.1/192.168.178.103/g' > ./k3s.yaml

chmod 600 ./k3s.yaml

kubectl --kubeconfig ./k3s.yaml get nodes

kubectl --kubeconfig ./k3s.yaml apply -f hello-world.yaml
```
