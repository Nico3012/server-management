# Prompt
Ich möchte ein Kubernetes cluster aufbauen, welches mit ingress controller im Internet steht. Erstmal wird es ein Kairos mit hadron und k3s single node cluster für den Start. Ich brauche jetzt hardware. Wichtig wäre mir, dass es viele flexible Möglichkeiten für BIOS und alle Firmware updates gibt. Zudem brauche ich möglicht viel unabhängigkeit. Denn es könnte auch sein, dass ich irgendwann einfach auf Debian oder ubuntu oder Windows Server oder Windows 11 pro oder Flatcar Linux oder Talos oder ein anderes OS wechseln möchte. Es sollte zudem die Möglichkeit bestehen, dass z.B. die Firmware und Bios und so viel wie möglich über das LVFS mit fwupd aktualisiert werden kann und alternativ auch das von Microsoft Windows genutzt System für BIOS updates soll auch unterstützt werden. Zudem muss auch ein manueller Weg bereitstehen, wie ich mit möglichst wenig Aufwand z.B. alle Firmware flashen kann mit einem USB Stick oder direkt aus Bios vom hersteller herunterladen oder sowas. Also mir ist wichtig, dass Updates recht einfach möglich sind. Ggf. auch gerne mit einem remote bmc oder so. Auch da gerne so viel Industriestandard wie möglich. Ich habe keine großen Anforderungen an die Hardware. 16gb RAM oder mehr wären ok. CPU keine speziellen Anforderungen. min. 1 ssd. Lan Anschluss. Wenn möglich alles unter 1000€ Und wenn möglich mit relativ geringem Idle Stromverbrauch. Offizieller Linux Support wäre mir auch wichtig.
Wenn möglich wäre auch cool, wenn ich eine ISO hochladen kann. Es wäre auch ok wenn das über ein extra Baustein möglich ist. Also z.B. mit einem extra IP-KVM.
ECC Ram und so brauche ich nicht zwingend. Das wäre cool aber muss nicht sein.

# Sicherstellen (Fragen)
Werden in einem Update z.B. das Bios aktualisiert, die Netzwerkkarte, Festplatte, ... oder muss ich all diese Updates unabhängig voneinander installieren?




# Beispiel
Dell, Lenovo, HP scheinen die ANforderungen am besten zu unterstützen

### Dell OptiPlex Micro
Bios und Intel Ethernet werden zusammen im Bios Update ausgeliefert.
