# 🗂️ fog

Files Organizer in Go

A little tool to help organize files easier.

## Note

This project is a small prototype from my early days of learning Golang, approximately 10 years ago.  
While it may still be functional, it has undergone limited testing.  
**Caution is advised, particularly in production environments.**  
Every commit generates a log, enabling easy reversion if necessary.  
However, please be mindful of potential risks before proceeding.

## Installation

### Install Go

Get Go from [https://go.dev/doc/install](https://go.dev/doc/install)

Or via package manager: [chocolatey](https://chocolatey.org/):

```shell
choco install go
```

Check go version

```shell
go version
```

### Install/Update this tool

Install via go command:

```shell
go install github.com/kwchang0831/fog@latest
```

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

Move SomeVideo_SXXEXX into their individual folder.

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

Undo committed actions from target log file.

```shell
fo revert -w [pathToTheLog]
```

`-w` : Commit changes. Without the flag wil be dry run.

### replacename

#### Remove [Bad] in file name only

```shell
fog replacename "\[Bad\]" "" -d "." -w
```

`-d`: Set directory.

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

`-d`: Set directory.

`-w` : Commit changes. Without the flag wil be dry run.

`-m1`: Mode 1: Folder name only.

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

`-d`: Set directory.

`-w` : Commit changes. Without the flag wil be dry run.

`-m2`: Mode 2: Folder and File name.

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

`-d`: Set directory.

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

For more information, please check help command.

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

Upgrade dependencies

```go
go get -u
```

Get dependencies

```shell
go mod tidy
```

Build

```shell
go build
```
