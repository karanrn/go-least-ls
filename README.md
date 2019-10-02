# go-least-ls
A command line utility to list (ls) least recently used files.<br>
A simple command line tool to check old files in the system.

Usage:
```
./go-least-ls -count=10 -older=60
```
Utility lists all the least recently files in the current directory older <br>
than `-older=<value>`.

![alt text](./Usage.png "Usage.")


`count` and `older` flags has to be integer.

![alt text](./IntegerCheck.png "Invalid flag values.")