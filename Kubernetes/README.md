# Talos
Talos ist ein Betriebssystem welches für Kubernetes gedacht ist und möglichst wenig Angriffsfläche bietet.

### Talos image builder
Über den Talos image builder: https://factory.talos.dev/
Wird ein image gebaut, welches dann mit Modulen für z.B. Nvidia gpu konfiguriert wird
Zudem wird die Architektur und das image Format (iso für bare metal oder cloud image für vm etc.) ausgewählt.
Sobald die z.B. iso für die spezifische hardware erstellt ist, wird diese installiert und alle Konfiguration im Bezug auf Kubernetes passiert remote über Konfigurationslisten

### Beispielconfig Talos image builder
Nvidia Pascal GPU Modul:
- siderolabs/nonfree-kmod-nvidia-lts
- siderolabs/nvidia-container-toolkit-lts




# Setup
1. talosctl installieren
2. Ordner erstellen z.B. ~/talos-config
3. talosctl gen config mein-cluster https://192.168.178.96:6443
3. talosctl gen secrets -o secrets.yaml
3. talosctl gen config mein-cluster https://192.168.178.96:6443 --with-secrets secrets.yaml -f


talosctl config merge ./talosconfig
talosctl config endpoint 192.168.178.96
talosctl config node 192.168.178.96


