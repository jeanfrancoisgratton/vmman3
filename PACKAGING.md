**Each package assumes that you're using the latest pkgbuilder container: debbuilder, rpmbuilder, apkbuilder**<br><br><br>

# DEB (Debian, Ubuntu) PACKAGE :

References:
___
- https://www.internalpointers.com/post/build-binary-deb-package-practical-guide
- https://wiki.debian.org/HowToPackageForDebian
- https://wiki.debian.org/BuildingTutorial

cd into `__debian` and then run all scripts in their numerical order<br>
**DO NOT `GIT COMMIT` THE CHANGES IN `__DEBIAN` ONCE COMPLETED**



# RPM (RedHat, Rocky, Fedora, CentOS, Oracle) PACKAGE :

Basic steps:

- from the projetc directory, run `tito init` . This is unneeded if there already is a `.tito` directory.
- Specfiles with `BuildRequires` need those packages to be installed ahead of tito build; run `./rpm-install-build-deps.sh`
- `tito tag`, to tag an upcoming release (`tito tag --keep-version` if you do not wish to update the release number)
- `git push --follow-tags origin`
- `tito build --rpm [--verbose] [--no-cleanup]` to actually build the package. You will need to manually upload it to Nexus

# APK (Alpine) PACKAGE:

References:
___
- https://wiki.alpinelinux.org/wiki/Creating_an_Alpine_package
- https://www.matthewparris.org/build-an-alpine-package/

1. Everything Alpine-related needs to be in the `__alpine/' directory
2. Customize your APKBUILD file if needed
3. Run `abuild -r [-k -K]` and you're good to go

If you need specific languages to be installed to perform build operations, have a look at `/opt/bin/`
