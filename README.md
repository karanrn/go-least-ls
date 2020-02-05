# go-least-ls
A command line utility to list (ls) least recently used files.<br>
A simple command line tool to find and clear up old and unused files/directory in your system.

Usage:
```
./go-least-ls -count=10 -older=60 -filetype .txt [-all]
```
Flags:
```
-count    : Number of files to show/list, oldest file will be at top. 
            Default value is 5. It should be an integer.
-older    : How older a file should be shown? Value is number of days. 
            Default value is 30. It should be an integer.
-filetype : File type, extension of files to be searched. It should be a string. 
            Ex: -filetype .txt
-all      : By default, utility does not list hidden files. Mention this flag to list hidden files.
```
**NOTE**: Currently, Hidden file check is only for linux/unix OS type.

Examples:<br>
Default usage:<br>
![alt text](screenshots/usage.png "Usage with default flag values")

Using `-count` and `-older` flags to list files based on the values passed:<br>
![alt text](screenshots/countFlag.png "Usage of count flag to override the default value.")<br>
![alt text](screenshots/olderFlag.png "Usage of older flag to override the default value.") 

Using `-all` to list hidden files:<br>
![alt text](screenshots/hiddenfiles.png "Usage of all flag to list hidden files.") 

Using `-filetype` to list files based on extension mentioned:<br>
![alt text](screenshots/filetypeFlag.png "Usage of filetype flag to list files based on extension type.") 