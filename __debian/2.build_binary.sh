#!/usr/bin/env bash

PKGDIR=vmman3_`cat ../_current_pkg_release`_amd64

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
mv control ${PKGDIR}/DEBIAN/
mv preinst ${PKGDIR}/DEBIAN/

echo "Building binary from source"
cd ../source
go build -o ../__debian/${PKGDIR}/opt/bin/vmman3 .

echo "Binary built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
