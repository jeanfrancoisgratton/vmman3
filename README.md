# vmman3
A Go-based libvirtd client, using Cobra to manage CLI params.

From this tool you can:
- create/delete VMs
- start/stop VMs
- manage VM snapshots
- manage VM clusters.
- manage templated VMs (the templated VMs are the "source" VMs that get copied into your target VM)
- edit resources (vpus, vmem, storage) of a given VM
- manage all back-end related data

The tool is hypervisor-aware, meaning it can manage VMs across a farm of hypervisors.<br>
This actually is the default behaviour.

## Pre-requisites
- A working PostgreSQL server where you have an admin account (an account that can create ROLEs and DATABASEs).
- Some libvirtd-related packages: if you installed the RPM/DEB/APK package, it should have either installed the dependencies, or have notified you about the missing deps
- A templated VM to be used as a source VM (read more about templates, below)
## Once installed, what to do
1. If you've met all of the pre-requisites, the first thing is to run `vmman db bootstrap`
2. You will need an _environment file_ . See the section **ENVIRONMENT FILE** below for what this is all about
3. Until I've automated the initial DB populate subcommand, you will need to run the following commands against your PGSQL database:
```postgresql
INSERT INTO hypervisors (hname, haddress, hconnectinguser) VALUES ('AAA', 'BBB', 'CCC');
INSERT INTO storagepools(spname, sppath, spowner) VALUES ('DDD', 'EEE', 'AAA');
INSERT INTO vmstates(vmname, vmip, vmoperatingsystem, vmhypervisor, vmstoragepool) VALUES ('FFF','GGG','HHH','AAA','DDD');
INSERT INTO disks(dname, dpool, dvm, dhypervisor) VALUES ('III','DDD','FFF','AAA')

Replace the VALUES (AAA, BBB, etc) with these values:
AAA: the name your first (maybe sole) hypervisor
BBB: its IP address, or hostname if resolvable
CCC: the user connecting to it
DDD: the storage pool name
EEE: the storage path (if needed)
FFF: the name of the VM
GGG: its ip address
HHH: its operating system
III: the first (maybe sole) disk used by that VM
```
3. **TODO TODO TODO TODO** next steps (do I really need db commands to populate the DB ?)

## Issues, FIXME, ETC

### Issues
There's a website for that :p

### FixMe:
None, yet

### TODO:
Read **TODO.md** : this is my roadmap, my upcoming bugfixes, etc.

# The templates
Templates are the basis of how vmman works. What vmman does is to actually copy a templated VM as a new VM, with the hostname and network config that you've provided.

This means that for the tool to properly work, you'd need an already-working VM.
It does not matter what you've installed in that VM (of course, it should be as lean as possible, being a template; the target VM is where you'd install stuff)
You do need a running ssh server that allows root access with key (the key should already be in /root/.ssh/authorized_keys)