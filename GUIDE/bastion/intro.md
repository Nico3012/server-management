# Bastion server
Ein gehärtetes System mit minimaler Angriffsfläche.
Das ist der Hauptserver, der die Gesamtanwendung bereitstellt. Hier laufen jedoch keine speziellen hardwarebeschleuniger sondern nur organisatorische Software, die CPU/RAM/Storage nutzt.
Die einzelne Hardware muss nicht hochverfügbar sein, hat eine maximale Laufzeit und wird in dieser Zeit durch regelmäßige Wartungszeiträume aus dem Produktivbetrieb genommen.

## Software
Ein immutable OS, welches open source ist und möglcihst unabhängig finanziert ist. Sicherheitsupdates werden regelmäßig bereitgestellt.

## Hardware
Die Hardware wird für die Software ausgewählt. Also die Software ist die Anforderung und die Hardware muss dafür funktionieren.
Die Hardware ist so minimal wie möglich um die Angriffsfläche zu reduzieren.
Die Hardware hat einen vom Hersteller vorgesehenen Wartungszeitraum, in der diese Updates erhällt, falls z.B. Firmware fehlerhaft ist.
Die aktualität muss dauerhaft geprüft werden. Die Hardwaresicherheitsprüfung umfasst zwei Level:

1. Physische Hardware:

   In diesem Fall wird überprüft, ob die hardware Sicherheitslücken im SIlizium beinhaltet, die durch Software nicht korrigierbar sind. Sehr selten. Beispiel: Spectre, Meltdown sind solche Lücken, die jedoch durch Firmwareanpassungen (Software) geschlossen werden konnten.

2. Firmware / BIOS (Software):

   Der Hersteller stellt Firmwareupdates im Wartungszeitraum bereit. Diese müssen installiert werden.
   Dafür MUSS die Hardware eine Remoteschnittstelle bereitstellen, über diese das durchgeführt werden kann.
   Diese Remoteschnittstelle darf nur in einem privaten und kontrollierten Netzwerk exposed werden!

## Hardwareauswahl
Beispiele für Bereitstellungshardware und was dann speziell an Wartung anfällt:

1. Basic Hetzner VPS:

   Bei einem Cloud VPS wird der hypervisor vom provider gewartet. Also nur die wartung für die Software (OS + application layer) fällt an

2. Industriestandard Server:

   Z.B. Dell PowerEdge oder vergleichbare Server haben ein eigenes BMC System, welches das Steuern inkl. Firmware / BIOS flashing von remote über das private Netzwerk ermöglicht.

   Sind oft Teuer weil z.B. auch ECC RAM verbaut wird und erfordern oft Lizenzkosten für das BMC System.

3. Selbstbau System:

   Z.B. Minipc wie Lenovo ThinkCentre und dann ein JetKVM Mini oder GL.iNet Comet (GL-RM1) davor geschaltet. Der kvm over ip ermöglicht den Zugriff auf das BIOS über das private Netzwerk. Zudem können USB devices für BIOS files oder ISO images über dieses Gerät über das private Netzwerk gemountet werden. Der JetKVM Mini oder GL.iNet Comet (GL-RM1) ist auch gefährdet, da er über usb, hdmi, ... mit dem Server verbunden ist und muss daher auch regelmäßig Sicherheitsupdates durchführen. Das ist über das Webinterface von diesem möglich.

   https://www.gl-inet.com/de-de/products/gl-rm1pe

   Keine klassische Serverhardware aber deutlich günstiger in Gesamt als ein out of the Box System.
