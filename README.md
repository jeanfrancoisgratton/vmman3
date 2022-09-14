# vmman3
A Go-based libvirtd client, using the Cobra CLI manager.

From this tool you can:
- create/delete VMs
- start/stop VMs
- manage VM snapshots
- manage VM clusters
- edit resources (vpus, vmem, storage) of a given VM

The tool is hypervisor-aware, meaning it can manage VMs across a farm of hypervisors

### Issues:
None known, so far :p

### FixMe:
None, yet

### TODO:
- Some packages (db !) are messy
- Refactoring is needed : move more functions in the helpers package
- virt-sysprep (or its equivalent) is still not libvirtd-native....
- ssh handling ? not there yet, I sure hope it works better than in Python (vmman1) !
- vm_management is not yet database-aware
- inventory: same