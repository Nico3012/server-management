Fedora CoreOS + Butane

# Config erstellen
config.bu erstellen

```shell
curl -LO https://github.com/coreos/butane/releases/latest/download/butane-x86_64-unknown-linux-gnu
chmod +x butane-x86_64-unknown-linux-gnu
./butane-x86_64-unknown-linux-gnu --strict config.bu -o config.ign
```

# Config bereitstellen
Auf Laptop ausführen (Vorher auf Windows kopieren): `python3 -m http.server 8000`

# Bare metal
In live iso booten und dann folgenden Befehl ausführen:
```shell
sudo localectl set-keymap de

sudo coreos-installer install /dev/sda \
  --ignition-url http://192.168.178.202:8000/config.ign \
  --insecure-ignition
```

# Podman nutzen
Podman ist auf CoreOS die bessere Wahl
