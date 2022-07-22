# DEB (Debian, Ubuntu) PACKAGE :

References:
___
- https://www.internalpointers.com/post/build-binary-deb-package-practical-guide
- https://wiki.debian.org/HowToPackageForDebian
- https://wiki.debian.org/BuildingTutorial

`dpkg-deb -b $DIR`

(where $DIR is the directory containing DEBIAN/ and control files)

# RPM (RedHat, Rocky, Fedora, CentOS, Oracle) PACKAGE :

Basic steps:

- `tito init` (already done, as the presence of the .tito/ directory asserts)
- `tito tag`, to tag an upcoming release (`tito tag --keep-version` if you do not wish to update the release number)
- `git push --follow-tags origin`
- `tito build --rpm` to actually build the package. You will need to manually upload it to Nexus

# APK (Alpine) PACKAGE:

References:
___
- https://wiki.alpinelinux.org/wiki/Creating_an_Alpine_package
- https://www.matthewparris.org/build-an-alpine-package/

1. Everything Alpine-related needs to be in the alpine/ directory
2. Ensure that the file _current_pkg_release corresponds to the version number where the sources are located
3. Customize your APKBUILD file if needed
4. Run `./prepare-tarball.sh`
5. `abuild checksum` once you're confident everything is good
6. `abuild -r [-k -K]` and you're good to go

If you need specific languages to be installed to perform build operations, have a look at `/opt/bin/`
