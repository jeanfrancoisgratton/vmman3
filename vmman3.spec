%define debug_package   %{nil}
%define _build_id_links none
%define _name   vmman3
%define _prefix /opt
%define _version 0.250
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
Obsoletes: vmman1

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

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{name}


%changelog
* Sat Sep 17 2022 builder <builder@famillegratton.net> 0.200-0
- Package build dry-run
