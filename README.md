tlsFileServer
=============

*tlsFileServer* is a simple but effective file server that serves a given directory via https.

Syntax
======

Die generelle syntax lautet:

`tlsFileServer <port> <root_dir>`.


*tlsFileServer* erwartet die beiden Dateien *cert.pem* und *key.pem* im gleichen Verzeichnis.

Schlüssel und Zertifikate
=========================

Die Datei *cert.pem* muss ein valides zu *key.pem* passendes Zertifikat beinhalten. 
Die Datei *key.pem* muss einen validen SSL/TLS key beinhalten.
Beide Dateien müssen PEM kodiert sein.

Die beiden Dateien können mit *generate_cert.go* erzeugt werden.
Etwa so:

`go run generate_cert.go -host=127.0.0.1,localhost`

Weitere Parameter für *generate_cert.go* sind:

- host - eine Liste (Komma separiert) mit IPs und Hostnamen für die ein Zertifikat generiert werden soll.
- validFrom - Datum ab wann das Zertifikat gültig sein soll, im Format: `Jan 1 15:04:05 2011`
- validFor - Zeit (in Sekunden) für die das Zertifikat gültig ist
- isCA - gibt an ob dieses Zertifikat eine CA ist
- rsaBits - Länge (in bit) des RSA Schlüssel der erzeugt werden soll

Um ein 4096bit Key und ein dazugehöriges (selbstunterschriebenes) Zertifikat für scusiblog.org zu erstellen ruft man generate_cert so auf:

`generate_cert.go -host=scusiblog.org,www.scusiblog.org -rsaBits=4096`

Mit dem obigen Befehl erzeugt man einen 4096bit RSA Schlüssel und ein Zertifikat für scusiblog.org. Standardmäßig sind Zertifikate ein Jahr ab dem Zeitpunkt der Erstellung gültig. 

Installation
============

`go get https://github.com/scusi/tlsFileServer`

Holt die sourcen von github.

`cd tlsFileServer`

`go install tlsFileServer`

`tlsFileServer 8443 /usr/share/doc/`

startet einen *tlsFileServer* auf port *8080* mit dem root Verzeichnis */usr/share/doc*


