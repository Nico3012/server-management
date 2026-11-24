https://www.flatcar.org/

# Laptop
./butane-x86_64-unknown-linux-gnu --strict config.yaml -o config.ign

# ISO:
sudo passwd core

# Laptop
ggf. (rm -r ~/.ssh)
scp config.ign core@192.168.178.101:/home/core/

# Laptop -> ISO:
ssh core@192.168.178.101

# ISO:
sudo flatcar-install -d /dev/sda -i config.ign
reboot
