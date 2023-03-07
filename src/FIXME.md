- [x] vmman vm stopall fails on vm alpinedev or alpinestorage
- [x] vmman stopall -a will not work at all (silently returns)
- [x] vmman -c $hypervisor reset : double-stop, no start
- [x] having multiple users on a single hypervisor in DB will mess the -a option.... (hint: https://golangbyexample.com/get-current-username-golang/)

- [x] vmman1 legacy: I need to create my VM + pools under my own name.

- [x] vmman ls -a not working anymore ?
- [ ] vmman -c $HYPERVISOR vm xmldump not working on remote hypervisor
- [x] vmman ls -a crashes if one hypervisor is missing : needs better handling
- [x] vmman vm down $VM not working anymore
- [ ] vmman db drop -y is not working