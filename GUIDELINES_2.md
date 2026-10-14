# Wenn ein Hypervisor in der Pipeline steckt, muss dieser inkl. seines Betriebssystems auch aktuell gehalten werden.
Warum?
Weil z.B. wenn eine Netzwerkanfrage an die top aktuelle VM kommt, diese erstmal durch die virtuelle Netzwerkschicht des Hypervisors an die VM geleitet werden muss. Findet man an dieser Stelle einen exploit, dann können möglicherweise Datenpackete direkt an den Hypervisor umgeleitet werden. Aus dem Grund muss hier ein Sicherheitsmanagement vorliegen.

Für den Hypervisor gelten die allgemeinen Guidelines im Bezug auf den Supportzeitraum.

Ein Hypervisor ist nicht Pflicht.
Das Deployment findet über Container statt. Allerdings nicht zwingend alles auf einer Hardware. Ein Container Host kann Teil einer vm sein.

# Was nicht umgesetzt werden darf:
Eine Regel besagt, dass die Container Hosts bzw. Nodes unabhängig voneinander bereitgestellt werden können müssen. Also wenn ein Node abgeschaltet wird, darf das nicht bedeuten, dass ein zweiter node mit abgeschaltet wird. Das wäre nicht erlaubt weil dann der innere Node nicht unabhängig ist.

Das bedeutet, wenn z.B. ein node ein Hardwareelement benötigt, welches nicht virtualisiert werden kann (pcie passthrough) und daher in bare metal vorhanden sein muss, dann darf dieser nicht gleichzeitig hypervisor für andere nodes sein, da diese nodes dann nicht erreichbar sind, wenn der andere node abgeschalten wird. In diesem Fall wäre die verstrickung zu eng zwischen den nodes. Das würde die Komplexität zu stark erhöhen und ggf. zu anderen Kompatibilitätsproblemen führen, da z.B. der GPU Node dann zwingend weiter aktualisiert werden müsste.

Zu wartende Konzepte:
- Die Nodes nach Richtlinien von `GUIDELINES.md`
- Himermit hinzukommend, falls mehrere nodes auf einem bare metal rechner deployed werden, ist der hypervisor. Muss auch mit Ansible aktuell gehalten werden. Alternativ kann natürlich direkt auf bare metal deployed werden oder auf einem VPS eines externen Betreibers. Dieser würde sich dann um den Hypervisor kümmern.
