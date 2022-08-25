%define debug_package   %{nil}
%define _name   vmman3
%define _prefix /opt
%define _homedir %{_prefix}/%{_name}
%define _version 1.000
%define _rel 0
%define _arch amd64

Name:       vmman3
Version:    %{_version}
Release:    %{_rel}
Summary:    vmman

Group:      virtualMachines/orchestration
License:    GPL2.0
URL:        http://git.famillegratton.net:3000/devops/vmman1

Source0:    %{name}-%{version}.tar.gz
BuildRoot:  %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)
BuildArchitectures: x86_64
#BuildRequires: gcc
#Requires: python3,python3-devel,libvirt-devel,python3-libvirt,python3-pip,libvirt-daemon-driver-qemu,virt-clone,sudo

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

%{__mkdir_p} "$RPM_BUILD_ROOT%{_homedir}"
install -Dpm 0755 %{_name}.exe "$RPM_BUILD_ROOT%{_homedir}/%{_name}"
%{__mkdir_p} "$RPM_BUILD_ROOT%{_prefix}/etc/%{_name}"
%{__install} -D -m 644 %{_name}_%{_version}-%{_rel}_%{_arch}%{_prefix}/etc/%{_name}/* "$RPM_BUILD_ROOT%{_prefix}/etc/%{_name}/"

%post
chmod 755 /opt/vmman3/vmman3 /opt/vmman3/create-disk-pool.sh
chown 0:devops /opt/etc/vmman3 && chmod 770 /opt/etc/vmman3
touch /var/log/vmman3.log && chown 0:kvm /var/log/vmman3.log && chmod 770 /var/log/vmman3.log

%preun


%postun

%files
%defattr(-,root,root,-)
%{_homedir}
%{_prefix}/etc/%{_name}


%changelog

