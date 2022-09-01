all: 
	cd ../source
	go build . -o ./vmman3.exe

prepare:
	/opt/bin/install_golang.sh 1.19 amd64
	#apk add --no-cache gcompat pkgconfig libvirt libvirt-dev gcc musl-dev
	sudo apk add --no-cache pkgconfig libvirt libvirt-dev gcc musl-dev


install:
	test -d $(DESTDIR)$(PREFIX) || mkdir -p $(DESTDIR)$(PREFIX)

	install -pm 644 $(BUILDDIR)$(PREFIX)/etc/$(PKGNM)/README.md $(DESTDIR)$(PREFIX)/etc/$(PKGNM)/README.md


	
