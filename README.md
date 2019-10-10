# cb
## Usage
Copy file contents to clipboard:
```sh
cb file.txt
```

Copy stdin contents to clipboard:
```sh
echo foo | cb
```

Omit exactly one trailing line break if it exists:
```sh
echo foo | cb -n
# or
echo -n foo | cb
```

Manually write the clipboard:
```sh
cb
Hello World!^D # ^D = Ctrl+D (send EOF)
```

Write clipboard to stdout:
```sh
cb -o
```

Write clipboard to file:
```sh
cb -o file.txt
```

## Installation
```sh
go get -u github.com/setlog/cb
```
