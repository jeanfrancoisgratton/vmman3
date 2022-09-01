all: 
	cd ../source
	go build . -o ./vmman3.exe

prepare:
	/opt/bin/install_golang.sh 1.19 amd64
	apk add --no-cache gcompat pkgconfig libvirt libvirt-dev gcc musl-dev
	sudo apk add --no-cache pkgconfig libvirt libvirt-dev gcc musl-dev


install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin

	install -pm 755 $(BUILDDIR)/vmman3.exe $(DESTDIR)$(PREFIX)/bin/vmman3


	
