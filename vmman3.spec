%define debug_package   %{nil}
%define _build_id_links none
%define _name   vmman3
%define _prefix /opt
%define _version 0.400
%define _rel 0
%define _arch x86_64

Name:       vmman3
Version:    %{_version}
Release:    %{_rel}
Summary:    vmman

Group:      virtualMachines/orchestration
License:    GPL2.0
URL:        http://git.famillegratton.net:3000/devops/vmman3

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: libvirt-devel,wget,gcc
Requires: libvirt-devel,libvirt,virt-clone,sudo
Obsoletes: vmman1 > 1.140

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
* Sun Oct 16 2022 builder <builder@famillegratton.net> 0.400-0
- Needed to compile 0.400 (jean-francois@famillegratton.net)
- Automatic commit of package [vmman3] release [0.400-0].
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

* Sun Oct 16 2022 builder <builder@famillegratton.net>
- Needed to compile 0.400 (jean-francois@famillegratton.net)
- Automatic commit of package [vmman3] release [0.400-0].
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
- new sample, dock update, specfile fix (jean-francois@famillegratton.net)
- Added an extra script (jean-francois@famillegratton.net)
- Forgot to bump version for deb package (jean-francois@famillegratton.net)

* Tue Sep 20 2022 builder <builder@famillegratton.net> 0.250-0
- Doc update, build scripts version bump (jean-francois@famillegratton.net)
- Doc updates, specfile cleanup (jean-francois@famillegratton.net)
- New table: templates (jean-francois@famillegratton.net)
- Synch across devices (jean-francois@famillegratton.net)

* Sat Sep 17 2022 builder <builder@famillegratton.net> 0.200-0
- Package build dry-run
