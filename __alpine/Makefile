all: 
	cd ../src
	go build . -o ./vmman3.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.20.4 amd64
	sudo apk update
	sudo apk add --no-cache pkgconfig libvirt libvirt-dev gcc musl-dev

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/vmman3.exe $(DESTDIR)$(PREFIX)/bin/vmman3

