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

Source0:    %{name}-%{_version}.tar.gz
BuildRoot:  %{_tmppath}/%{name}_%{version}-%{_rel}-root-%(%{__id_u} -n)
BuildArchitectures: x86_64
BuildRequires: libvirt-devel,wget,gcc
Requires: libvirt-devel,libvirt,virt-clone,sudo

%description
GoLang-based libvirt client

%prep
%setup -q

%build
#cd "$RPM_BUILD_ROOT/source"
cd %{_sourcedir}/%{_name}-%{_version}/source
PATH=$PATH:/opt/go/bin go build -o "$RPM_BUILD_ROOT/%{_name}.exe" .

%clean
rm -rf $RPM_BUILD_ROOT
microdnf remove -y libvirt-devel gcc

%pre
/usr/sbin/groupadd kvm 2> /dev/null
/usr/sbin/groupadd -g 2500 devops 2> /dev/null
exit 0

%install
%{__mkdir_p} "$RPM_BUILD_ROOT%{_prefix}/bin"
#install -Dpm 0755 "$RPM_BUILD_ROOT/%{_name}.exe" "$RPM_BUILD_ROOT%{_prefix}/bin/%{_name}"

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_prefix}/bin/%{_name}


%changelog
