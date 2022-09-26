%define debug_package   %{nil}
%define _build_id_links none
%define _name   vmman3
%define _prefix /opt
%define _version 0.300
%define _rel 1
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
