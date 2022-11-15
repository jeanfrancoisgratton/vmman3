__legend__<br>
`x` means enough done for a release, more to come in subsequent releases<br>
`xx` means that this item is feature-completed<br>
`-` means that this is slated for a subsequent release<br>
`..` means that this was finally rejected
` ` means it hasn't been started.
<br><br><br>
__ROADMAP__

[ xx ] - QEMU connection string defaults to root@some-host.
This needs to change with both a -u param, and/or a default value in some JSON file, and/or a DB value.

[    ] - Allows for some locking mechanism:
Say that you have a vm create running against a specific template, on a given hypervisor.
If you wanted to create a VM with the same template on another hypervisor, we'll need a mechanism to notify the
Process that it's OK (or not) to proceed.

[ .. ] - Everything under helpers/ should be under its proper package.
That's a relic from the Helper class in Python3's version, vmman1.

[-] - Learn GoLang Generics:
This might simplify (A LOT) the code, especially in the db/ package.
That package is just ugly.

[ xx ] - ALTER SEQUENCES : START value resets at one after db import.
This needs to be reset (drop sequence, recreate with higher value).

[ - ] - More db subcommands to edit/insert/delete data, list tables, etc.

[ - ] - Template subcommands.

[ xx ] - Completion of the snapshot package.

[ x ] - Completion of the vm_management package: Stop[All]/Start[All]/Reset[All].

[ xx ] - New package: pool_storage.

[ - ] - New package: resources (? could be folded into `vm_management` ?).

[ .. ] - New package : cluster (? might not need this one anymore, unsure)

[ - ] - Minor: show uptime when VM is up, otherwise show Last status change

[ xx ] - Hypervisor management

[ - ] - Better help handling than the defaults provided with `cobra`

Nice to have:<br>
`func (d *Domain) GetGuestInfo(types DomainGuestInfoTypes, flags uint32) (*DomainGuestInfo, error)`<br>
https://gitlab.com/libvirt/libvirt-go-module/-/blob/v1.8007.0/domain.go#L5527


Copy templated volume to new vol:
virsh help vol-create-from