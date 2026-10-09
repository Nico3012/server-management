# Referenzen
`https://www.youtube.com/watch?v=Ctg7rWGhqbk&t=313s`

# Begriffserklärungen
1. `Server`: Das ist das Gerät wo ansible drauf läuft. Also z.B. WSL auf einem Laptop
2. `Client`: Das ist der von ansible verwaltete Server. Also z.B. Ubuntu auf bare metal

# Ansible Installation auf server
```shell
sudo apt update
sudo apt install ansible sshpass
```

# Ansible Installation auf client
Nicht nötig. Nur ssh ist nötig

# Ausführen eines Playbooks
```shell
ansible-playbook --user root --ask-pass -i ./playbooks/hosts.ini ./playbooks/ubuntu-nvidia-docker.yml
```




# Server auf Englisch umstellen
ssh nico@192.168.178.68
sudo locale-gen en_US.UTF-8
sudo update-locale LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8
