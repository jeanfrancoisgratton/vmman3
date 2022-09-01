%define debug_package   %{nil}
%define _name   vmman3
%define _prefix /opt
%define _version 1.000
%define _rel 0
%define _arch amd64

Name:       vmman3
Version:    %{_version}
Release:    %{_rel}
Summary:    vmman

Group:      virtualMachines/orchestration
License:    GPL2.0
URL:        http://git.famillegratton.net:3000/devops/vmman3

Source0:    %{name}-%{version}.tar.gz
BuildRoot:  %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)
BuildArchitectures: x86_64
BuildRequires: libvirt-devel
Requires: libvirt-devel,libvirt,virt-clone,sudo

%description
GoLang-based libvirt client

%prep
%setup -q

%build
go build -v -o %{_name}.exe

%clean
rm -rf $RPM_BUILD_ROOT

%pre
/usr/sbin/groupadd kvm 2> /dev/null
/usr/sbin/groupadd -g 2500 devops 2> /dev/null
exit 0

%install

%{__mkdir_p} "$RPM_BUILD_ROOT"
install -Dpm 0755 "$RPM_BUILD_ROOT/%{_name}.exe" "$RPM_BUILD_ROOT%{_prefix}/%{_name}"

%post

%preun


%postun

%files
%defattr(-,root,root,-)
%{_prefix}/bin/%{_name}


%changelog
* Thu Sep 01 2022 builder <builder@famillegratton.net> 1.000-0
- new package built with tito

* Thu Sep 01 2022 builder <builder@famillegratton.net> 1.000-0
- new package built with tito


