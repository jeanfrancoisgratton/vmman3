[ ] - QEMU connection string defaults to root@some-host.
    This needs to change with both a -u param, and a default value in some JSON file.

[ ] - Allows for some locking mechanism:
    Say that you have a vm create running against a specific template, on a given hypervisor.
    If you wanted to create a VM with the same template on another hypervisor, we'll need a mechanism to notify the
    Process that it's OK (or not) to proceed.

    Either this, or having multiple copies of the same templated VM.

[ ] - Everything under helpers/ should be under its proper package.
    That's a relic from the Helper class in Python3's version, vmman1.

[ ] - Learn GoLang Generics:
    This might simplify (A LOT) the code, especially in the db/ package.
    That package is just ugly.

[x] - ALTER SEQUENCES : START value resets at one after db import.
    This needs to be reset (drop sequence, recreate with higher value).

[ ] - More db subcommands to edit/insert/delete data, list tables, etc

[ ] - template subcommands

[ ] - 
