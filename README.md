pangu.go
========

Paranoid text spacing for good readability, to automatically insert whitespace between CJK (Chinese, Japanese, Korean) and half-width characters (alphabetical letters, numerical digits and symbols).

* Go version: [pangu.go](https://github.com/vinta/pangu)
* JavaScript version: [pangu.js](https://github.com/vinta/paranoid-auto-spacing)
* Node.js version: [pangu.node](https://github.com/huei90/pangu.node)
* Python version: [pangu.py](https://github.com/vinta/pangu.py)
* Java version: [pangu.java](https://github.com/vinta/pangu.java)

## Installation

To install the package for using in your Go programs:

``` bash
$ go get -u github.com/vinta/pangu
```

To install the command-line tool, `pangu-axe`:

``` bash
$ go get -u github.com/vinta/pangu/pangu-axe
```

## Usage

### Package

TODO

### Command-line Interface

``` bash
$ pangu-axe text "所以,請問Jackey的鼻子有幾個?3.14個"
所以, 請問 Jackey 的鼻子有幾個? 3.14 個

$ pangu-axe file your_file.txt
your_file.pangu.txt
```

## Documentation

TODO
