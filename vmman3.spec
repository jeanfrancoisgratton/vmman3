%define debug_package   %{nil}
%define _build_id_links none
%define _name   vmman3
%define _prefix /opt
%define _version 0.700
%define _rel 0
%define _arch x86_64

Name:       vmman3
Version:    %{_version}
Release:    %{_rel}
Summary:    vmman

Group:      virtualMachines/orchestration
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/devops/vmman3

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: libvirt-devel,wget,gcc
Requires: libvirt-devel,libvirt,virt-clone,sudo
#Obsoletes: vmman1 > 1.140

%description
GoLang-based libvirt client

%prep
#%setup -q
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_name} .

%clean
rm -rf $RPM_BUILD_ROOT

%pre
/usr/sbin/groupadd kvm 2> /dev/null
/usr/sbin/groupadd -g 2500 devops 2> /dev/null
exit 0

%install
#%{__mkdir_p} "$RPM_BUILD_ROOT%{_prefix}/bin"
#install -Dpm 0755 %{buildroot}/%{_name} "$RPM_BUILD_ROOT%{_prefix}/bin/"
install -Dpm 0755 %{_sourcedir}/%{name} %{buildroot}%{_bindir}/%{name}

%post
strip %{_prefix}/bin/%{name}

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{name}


%changelog
* Sat Nov 12 2022 builder <builder@famillegratton.net> 0.650-0
- Completed hypervisor subcommand (jean-francois@famillegratton.net)
- hypervisor stub (jean-francois@famillegratton.net)
- Added extra field to dbCreds (now renamed envCreds), version bump (jean-
  francois@famillegratton.net)
- hypervisor subcommand stub (jean-francois@famillegratton.net)
- Helpers refactoring, version bump (missing a file) (jean-
  francois@famillegratton.net)
- Version bump, doc update (jean-francois@famillegratton.net)
- Updates to README, TODO and FIXME; fix in args handling of vmman vm xmldump
  (jean-francois@famillegratton.net)

* Fri Nov 11 2022 builder <builder@famillegratton.net> 0.600-0
- vmman snap subcommands completed; version bump (jean-
  francois@famillegratton.net)
- DomainLookupByName() is now in the helpers package (...and it works this
  time) (jean-francois@famillegratton.net)
- CreateSnapshot() stub, sync zenika-> (jean-francois@famillegratton.net)
- Completed snap rm (jean-francois@famillegratton.net)
- Refactoring: moved back Wait4Shutdown() in the helpers package (jean-
  francois@famillegratton.net)
- sync bergen-> (jean-francois@famillegratton.net)
- Completed snap ls (jean-francois@famillegratton.net)
- fixed vmman reset, near-completed vmman snap ls (jean-
  francois@famillegratton.net)
- new samples (jean-francois@famillegratton.net)
- vm_management package refactoring and db export format change (jean-
  francois@famillegratton.net)
- sync bergen -> (jean-francois@famillegratton.net)
- corrected variable mismatch (jean-francois@famillegratton.net)

* Sat Nov 05 2022 builder <builder@famillegratton.net> 0.550-0
- Completed vm console (jean-francois@famillegratton.net)
- libvirt.Connect has been moved in the helpers package; fixed issue with vm
  stopall -a not working (jean-francois@famillegratton.net)
- Minor fixes in vm start-stop (jean-francois@famillegratton.net)
- Completed virsh xmldump (jean-francois@famillegratton.net)
- Plugged a memory leak with libvirt (jean-francois@famillegratton.net)
- Added FIXME file (jean-francois@famillegratton.net)
- moved vmHelpers in the proper package (jean-francois@famillegratton.net)
- Sync between branch chg (jean-francois@famillegratton.net)
- First vm create stub... this one is the biggie.... (jean-
  francois@famillegratton.net)
- Completed DB updates in vm rename (jean-francois@famillegratton.net)
- Removed the clusters table (jean-francois@famillegratton.net)

* Sun Oct 30 2022 builder <builder@famillegratton.net> 0.525-1
- vm-rename was not attached in cmd/ (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)

* Sun Oct 30 2022 builder <builder@famillegratton.net> 0.525-0
- Completed vm rm (jean-francois@famillegratton.net)
- Completed vmRename (jean-francois@famillegratton.net)
- New samples, new subcommand: vm rename (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- sync across devices (jean-francois@famillegratton.net)

* Sat Oct 22 2022 builder <builder@famillegratton.net> 0.510-0
- Tag bump, db subcommands completed for this release, new table (jean-
  francois@famillegratton.net)
- Completed all tables (jean-francois@famillegratton.net)
- package refactoring, new table, new db subcommand (jean-
  francois@famillegratton.net)
- Completed "uptime" in ls subcommand (jean-francois@famillegratton.net)
- sync across devices (jean-francois@famillegratton.net)

* Wed Oct 19 2022 builder <builder@famillegratton.net> 0.500-0
- Completed stop-start-reset (jean-francois@famillegratton.net)
- before 0.420 (jean-francois@famillegratton.net)
- minor refactoring (jean-francois@famillegratton.net)
- Really fixed VMstateChange(), re-wrote most PGSQL connection strings (jean-
  francois@famillegratton.net)
- Updated PGX module to 5.0.3, even if it's a dev branch (jean-
  francois@famillegratton.net)

* Sun Oct 16 2022 builder <builder@famillegratton.net> 0.400-0
- Added file needed to build (jean-francois@famillegratton.net)
  (builder@famillegratton.net)
- Needed to compile 0.400 (jean-francois@famillegratton.net)
  (builder@famillegratton.net)
- Finalized files for tag 0.400 (jean-francois@famillegratton.net)
- ls completed (jean-francois@famillegratton.net)
- completed vmman ls (jean-francois@famillegratton.net)
- sync across devices (jean-francois@famillegratton.net)
- completed URIbuilder (jean-francois@famillegratton.net)
- another round of packages refactoring (jean-francois@famillegratton.net)
- minor linting in the db package (jean-francois@famillegratton.net)
- sync across devices (jean-francois@famillegratton.net)
- cosmetic change (jean-francois@famillegratton.net)
- Commenting format uniformizations (jean-francois@famillegratton.net)
- Third attempt at fixing EnvironmentFile (jean-francois@famillegratton.net)
- environmentfile improvements (jean-francois@famillegratton.net)
- Moved tables from the "config" schema to the "public" one (jean-
  francois@famillegratton.net)
- Fixed unresolved var (jean-francois@famillegratton.net)
- Enabled default environment file (jean-francois@famillegratton.net)
- onliner to prettify the json file (jean-francois@famillegratton.net)
- Roadmap update (jean-francois@famillegratton.net)
- Packaging update for new release (jean-francois@famillegratton.net)
- Rewriting stuff; reverting to old tags (jean-francois@famillegratton.net)
- Packaging update for new release (jean-francois@famillegratton.net)
- Roadmap update (jean-francois@famillegratton.net)
- onliner to prettify the json file (jean-francois@famillegratton.net)
- Enabled default environment file (jean-francois@famillegratton.net)
- Moved tables from the "config" schema to the "public" one (jean-
  francois@famillegratton.net)
- fixes to vmman ls (jean-francois@famillegratton.net)
- Completed vmstop (jean-francois@famillegratton.net)
- renamed all pgsql-related "conn" vars to "dbconn" (jean-
  francois@famillegratton.net)
- Hopefully a logic fix to cli flags (jean-francois@famillegratton.net)
- First stab at vm_management (jean-francois@famillegratton.net)
- update FIXME (jean-francois@famillegratton.net)
- Removed helpers.SurroundText() (jean-francois@famillegratton.net)
- config dir change on OSX (jean-francois@famillegratton.net)
- commit before switching (jean-francois@famillegratton.net)
- Automatic commit of package [vmman3] release [0.300-1].
  (builder@famillegratton.net)
- specfile fix (jean-francois@famillegratton.net)
- specfile fix (jean-francois@famillegratton.net)
- GetConn() is not working, removed (jean-francois@famillegratton.net)
- Packaging updates (jean-francois@famillegratton.net)

* Sun Sep 25 2022 builder <builder@famillegratton.net> 0.300-1
- specfile fix (jean-francois@famillegratton.net)
- GetConn() is not working, removed (jean-francois@famillegratton.net)
- Packaging updates (jean-francois@famillegratton.net)

* Sun Sep 25 2022 builder <builder@famillegratton.net> 0.300-0
- Version bump (even if it will never be packaged) (jean-
  francois@famillegratton.net)
- inventory near-completion; issue with -c -1 -a flags (jean-
  francois@famillegratton.net)
- some data changes, added postinst tasks to all packages (jean-
  francois@famillegratton.net)
- ALTERed config.vmstates (jean-francois@famillegratton.net)
- sync across devices (jean-francois@famillegratton.net)
- Completed vm ls refactoring (untested yet) (jean-francois@famillegratton.net)
- doc update, script cosmectic change (jean-francois@famillegratton.net)
- vmman ls refactoring in progress....sync across devices (jean-
  francois@famillegratton.net)
- Added column in the templates table (jean-francois@famillegratton.net)
- sync across devices (jean-francois@famillegratton.net)
- More todos.... (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- New subcommand: template (jean-francois@famillegratton.net)
- Completed SEQUENCE update (jean-francois@famillegratton.net)
- Now SEQUENCE nextvals are updated after db import (jean-
  francois@famillegratton.net)
- new sample, doc update, specfile fix (jean-francois@famillegratton.net)
- Added an extra script (jean-francois@famillegratton.net)
- Forgot to bump version for deb package (jean-francois@famillegratton.net)

* Tue Sep 20 2022 builder <builder@famillegratton.net> 0.250-0
- Doc update, build scripts version bump (jean-francois@famillegratton.net)
- Doc updates, specfile cleanup (jean-francois@famillegratton.net)
- New table: templates (jean-francois@famillegratton.net)
- Synch across devices (jean-francois@famillegratton.net)

* Sat Sep 17 2022 builder <builder@famillegratton.net> 0.200-0
- Package build dry-run
