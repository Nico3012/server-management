# Stand jetzt wird hier nur alles reingeschrieben, was beachtet werden muss. Hier ist noch keine Ordnung drin etc.

Mein unternehmen entwickelt software und stellt die VMs etc. selber bereit. Dafür, damit das alles auch sicher ist, habe ich ein Konzept entwickelt, welche Tools verwendet werden und was wie als Abstraktionsschicht zu verstehen ist, damit das deployment system mit möglichst verschiedenster Hardware kompatibel ist aber auch möglichst einheitlich läuft. Das sieht wie folgt aus:

Hardware, BIOS/UEFI, Betriebssystem werden als eine Einheit betrachtet. Bei dieser Einheit wird nicht vorrausgesetzt, dass die Komponenten open source sind oder frei zugänglich. Diese Einheit hat eine Laufzeit von X. Nach dieser Laufzeit darf die Einheit nicht weiter betrieben werden. Diese Laufzeit kann bestimmt werden durch den Supportzeitraum des Betriebssystems oder durch andere Garantien bzw. Beurteilungen, wie lange die Hardware zuverlässig funktioniert. Im Anschluss muss die Einheit erneuert werden (Austauschen einzelner Komponenten dieses Verbunds).
Features dieser möglicherweise propritären Einheit sind folgende:
- Zugriff über Netzwerk
- Root Zugriff. Der Server muss unmanaged sein.
- Tool(s) für Updatemanagement für Sicherheitsupdates und ggf. Funktionsupdates über die Dauer des Supportzeitraums.
- SSH Zugriff
- In der Lage, Linux System Container auszuführen
- Ggf. in der Lage sein, CDI (Container device interface) Dateien anzulegen, um vorhandene Geräte für Container zugänglich zu machen.

Muss die Hardware gekauft werden oder physisch besitzt werden?
- Nein, weder gekauft noch physisch besitzt werden. Wenn die Einheit von Hardware, …, bis OS vollständig irgendwo gemietet wird bzw. virtuell bereitgestellt wird, ist das völlig ausreichend. Statt einem Physischen Start Knopf hat man dann in einer Cloud Console einen Start Knopf. Zugriff/Nutzung passiert immer einfach über Netzwerk.

Wie wird diese Einheit nun in ein standardisiertes Format gebracht, dass dann standardisiert gemanaged werden kann?
- Mittels Ansible wird über ssh auf den Server zugegriffen und der Zustand überwacht bzw. Updates durchgeführt und die benötigte Software installiert um die o.g. Anforderungen zu erfüllen (Ausführen von containern etc.). Dafür muss als Adapter ein entsprechendes Playbook geschrieben werden.
- Natürlich kann keine geforderte Geschwindigkeit bei der Ausführung von Rechenoperationen nicht standardisiert werden oder vorausgesetzt werden, da das stark von verbauten Hardwarebeschleunigern (Z.B. GPU/NPU) abhängt. Falls es da spezielle Anforderungen gibt, muss dies bei der Auswahl der Hardware+OS Einheit beachtet werden. Features können selbstverständlich mittels Software Emulation bereitgestellt werden (Z.B. Falls keine GPU verbaut ist, kann dafür Software Rendering zurückgegriffen werden)



Anmerkungen zu Update typen und Rückwärtskompatibilität (Noch nicht bestätigt):
Für den Linux Kernel ist garantiert, dass dieser zum User space hin eine rückwärtskompatible Api bietet. Das bedeutet, ich kann auch mit dem neusten Kernel alte Container laufen lassen
Beim Einfügen von CDI Geräten in den Container muss das nochmal gesondert evaluiert werden. Wenn der z.B. Grafikkartentreiber, der gemountet wird keine Rückwärtskompatibilität bietet, darf dieser nicht automatisiert mit Funktionsupdates versorgt werden. Unter Umständen kann das bedeuten, dass der Kernel inkl. seiner Treiber, die mittels CDI übergeben werden keine Funktionsupdates erhalten darf sondern nur Sicherheitsupdates. Dabei ist dann zu beachten, den Support Zeitraum klein zu halten weil neuere Container Images auch neuere Kernelfeatures vorraussetzen können.
Ideen um das zu lösen:
- Unterteilung in Up to date branch, der z.B. in VM läuft. Da ist immer der neuste Kernel vorhanden. Z.b. rolling release distro aber keine Hardware Features
- Und bare metal runtime: Da können die Treiber limitieren. Daher nicht zwingend komplett up to date. Dafür auch cdi geräte verfügbar




# Verifizierung gegenüber Industriestandards

KI Prompt / Erklärung des Konzepts (Stand meines Verständnisses 08.10.2026):

Sag mir mal ob so ein typischer large scale tech stack funktioniert:
Server mit Betriebssystem ggf. ist da noch ein hypervisor drunter (muss aber nicht) und das OS hat eine Laufzeit X für die es support gibt. Z.b. ubuntu LTS Zeitraum und in der Zeit wird es von ansible verwaltet. Also ansible führt sicherheitsupdates aus und ggef funktionsupdates und stellt sicher, dass z.B. system container oder vm software installiert ist bzw. die server laufen. Und da drin läuft dann kubernetes, welches sicherstellt, dass mehrere dieser einheiten über wartungszeiträume hinweg voll verfügbar sind. Ist das so richtig?
