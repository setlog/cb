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

Copy stdin contents to clipboard, truncating 1 trailing line break if it exists:
```sh
echo foo | cb -n
```

## Installation
```sh
go get -u github.com/setlog/cb
```
