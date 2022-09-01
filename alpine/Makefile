all: 
	tar xjf $(SRCDIR)/$(TARBALL)

prepare:
	/opt/bin/install_golang.sh 1.19 amd64
	apk add --no-cache gcompat pkgconfig libvirt libvirt-dev gcc musl-dev

install:
	test -d $(DESTDIR)$(PREFIX)/$(PKGNM) || mkdir -p $(DESTDIR)$(PREFIX)/$(PKGNM)
	test -d $(DESTDIR)$(PREFIX)/etc/$(PKGNM) || mkdir -p $(DESTDIR)$(PREFIX)/etc/$(PKGNM)

	install -pm 644 $(BUILDDIR)$(PREFIX)/$(PKGNM)/* $(DESTDIR)$(PREFIX)/$(PKGNM)/
	install -pm 644 $(BUILDDIR)$(PREFIX)/etc/$(PKGNM)/README.md $(DESTDIR)$(PREFIX)/etc/$(PKGNM)/README.md
	chmod 755 $(DESTDIR)$(PREFIX)/$(PKGNM)/vmman

	
