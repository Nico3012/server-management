# Möglichkeiten
Ziel ist es, dem Nutzer Webanwendungen inkl. aller dafür benötigten Protokolle bereitzustellen.
Das Bedeutet, HTML, STUN, TURN und alles weitere, was heute und in Zukunft möglich ist, um mit dem Browser zu interagieren.

Backendfunktionalitäten sind nicht nur auf klassische CPU/RAM/Storage Komponenten begrenzt sondern ermöglichen auch komplexere Anwendungen wie z.B. GPU / Echtzeitsysteme etc.

### Was nicht möglich ist / aus Sicherheitsgründen ausgeschlossen
Komplexe Berechnungen auf z.B. der GPU sind möglich aber es ist nicht vorgesehen, dem anonymisierten Nutzer Zugriff auf die GPU zu bieten (CUDA injections / GPU miete etc).

### Sicherheitsbarriere (Bastion)
In Kontakt mit dem Nutzer steht ein gehärtetes minimales System, welches nur aus CPU/RAM/Storage besteht um die Angriffsfläche (z.B. Kernel Komplexität) zu minimieren.
Dieses System muss bottom-up Aktualität garantieren. Also von der Firmware bis zur Anwendung muss eine Aktualisierungspipeline sichergestellt werden.
Die Anwendung, die auf diesem System läuft, stellt sicher, dass keine schädlichen Angriffe mehr vom Nutzer an weitere Subsysteme weitergeleitet werden.
[Mehr zur Bastion](bastion/intro.md)

### Subsysteme
Weitere Subsysteme im Verbund können z.B. für GPU compute ein eigener Rechner/VM sein, der eine GPU bereitstellt oder z.B. ein Echtzeitsystem.
Diese Systeme werden in ein privates netzwerk, welches nur mit der gehärteten Hauptmaschine zusammen besteht, isoliert. Die Hauptmaschine kann das Subsystem dann anweisen, gewisse Berechnungen durchzuführen. Hier ist keine prompt injection erlaubt. Also die Nutzeranfrage muss auf der "Bastion" (Hauptrechner) geprüft werden, bevor diese an das Subsystem weitergegeben wird.
Das Subsystem muss nicht bottom-up aktuell gehalten werden. Das ist ein Kompromiss aus geringer Wartung und Sicherheit, die von der Hochaktuellen Bastion maximiert wird.
