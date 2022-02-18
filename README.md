# 🗂️ fog

Files Organizer in Go

A little tool to help organize files easier.

## Installation

### Install Go

Get Go from [https://go.dev/doc/install](https://go.dev/doc/install)

Or via package manager: [chocolatey](https://chocolatey.org/):

```shell
choco install go
```

Please making sure your go version is up to date

```shell
# Current latest version 1.77.7
go version 
```

### Install/Update this tool

Install via go command:

```shell
go install github.com/kwchang0831/fog
```

## Note

Not fully tested yet. Only ran on Windows x64.  
Not recommended to use in any production environment.  
Please use it at your own risk.  

Every committing changes will generate a log for you to revert back.  
Might not support network drive.  

## Usages

### folderout

Move files in the folder in current directory out of their folders.

```shell
fog folderout "." -w
```

`"."` : Target directory. Defaults to `"."` if omitted.

`-w` : Commit changes. Without the flag wil be dry run.

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── SomeVideo_S01E01
    ├── SomeVideo_S01E01.mp4
    ├── SomeVideo_S01E01.jpg 
├── SomeVideo_S01E02
    ├── SomeVideo_S01E02.avi
            </code></pre></td>
            <td><pre><code>.
├── SomeVideo_S01E01.mp4
├── SomeVideo_S01E01.jpg 
├── SomeVideo_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

### folderin

Move SomeVideo_SXXEXX into their individual folder

Dry Run on current directory  

```shell
fog folderin "." -w
```

`"."` : Target directory. Defaults to `"."` if omitted.

`-w` : Commit changes. Without the flag wil be dry run.

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── SomeVideo_S01E01.mp4
├── SomeVideo_S01E01.jpg
├── SomeVideo_S01E02.avi
            </code></pre></td>
            <td><pre><code>.
├── SomeVideo_S01E01
    ├── SomeVideo_S01E01.mp4
    ├── SomeVideo_S01E01.jpg
├── SomeVideo_S01E02
    ├── SomeVideo_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

### revert

Undo previous actions

```shell
fo revert -w [pathToTheLog]
```

`-w` : Commit changes. Without the flag wil be dry run.

### replacename

#### Remove [Bad] in file name only

```shell
fog replacename "\[Bad\]" "" -d "." -w
```

`-d`: Set directory  

`-w` : Commit changes. Without the flag wil be dry run.

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [Bad]SomeVideo_Folder[Bad] 
├── [Bad]SomeVideo_S01E01[Bad].mp4
├── [Bad]SomeVideo_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── [Bad]SomeVideo_Folder[Bad]  
├── SomeVideo_S01E01.mp4
├── SomeVideo_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

#### Replace [Bad] in folder name only

```shell
fog replacename "\[Bad\]" "" -d "." -m1 -w
```

`-d`: Set directory  

`-w` : Commit changes. Without the flag wil be dry run.

`-m1`: Mode 1: Folder name only  

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [Bad]SomeVideo_Folder[Bad]   
├── [Bad]SomeVideo_S01E01[Bad].mp4
├── [Bad]SomeVideo_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── SomeVideo_Folder
├── [Bad]SomeVideo_S01E01[Bad].mp4
├── [Bad]SomeVideo_S01E02[Bad].avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

#### Remove [Bad] in both folder name and file name

```shell
fog replacename "\[Bad\]" "" -d "." -m2 -w
```

`-d`: Set directory  

`-w` : Commit changes. Without the flag wil be dry run.

`-m2`: Mode 2: Folder and File name  

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [Bad]SomeVideo_Folder[Bad]   
├── [Bad]SomeVideo_S01E01[Bad].mp4
├── [Bad]SomeVideo_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── SomeVideo_Folder
├── SomeVideo_S01E01.mp4
├── SomeVideo_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

#### Replace filename with regex

```shell
fog replacename "(.*)(SomeVideo)(.*)(S[0-9]+E[0-9]+)(.*)(\.(mp4|avi))" "$2-$4$6" -d "." -w
```

`-d`: Set directory  

`-w` : Commit changes. Without the flag wil be dry run.

Learn more and try regex, see [regex101](https://regex101.com/).  

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [20240202]SomeVideo_S01E01[Bad].mp4
├── [20240207]SomeVideo_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── SomeVideo-S01E01.mp4
├── SomeVideo-S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

### Show Help

For more information, please check help command

```shell
fog
```

Output

```shell
File Organizer in Go provides commands to help you batch edit filenames and organize files.

Usage:
  fog [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  folderin          move files into their own folders.
  folderout         Move files out of folders
  help              Help about any command
  move              Move matching files/folders into target directory.
  renameafterfolder Rename files inside matching folders with the folder name.
  replacename       Replace name using the search pattern and replace pattern.
  revert            revert commands issued.
  rmemptydir        remove empty folders.

Flags:
  -h, --help   help for fog

Use "fog [command] --help" for more information about a command.
```

## Development

Get dependencies

```shell
go mod tidy
```

Build

```shell
go build
```
