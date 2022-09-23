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
- Some libvirtd-related packages: if you installed the RPM/DEB/APK package, it should have either installed the dependencies, or notified you about

## Once installed, what to do
1. If you've met all of the pre-requisites, the first thing is to run `vmman db bootstrap`
2. **TODO TODO TODO TODO** next steps (do I really need db commands to populate the DB ?)

## Issues, FIXME, ETC

### Issues
There's a website for that :p

### FixMe:
None, yet

### TODO:
Read **TODO.md** : this is my roadmap, my upcoming bugfixes, etc.