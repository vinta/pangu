# pangu.go

[![Go Reference](https://img.shields.io/badge/go.dev-reference-007d9c?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/vinta/pangu/v4)
[![Go Version](https://img.shields.io/github/go-mod/go-version/vinta/pangu?style=for-the-badge)](https://github.com/vinta/pangu/blob/main/go.mod)

Opinionated paranoid text spacing in Go: automatically inserts whitespace between CJK (Chinese, Japanese, Korean) and ANS (alphabetical letters, numerical digits and symbols).

- [pangu.js](https://github.com/vinta/pangu.js)
- [pangu.py](https://github.com/vinta/pangu.py)
- [pangu.go](https://github.com/vinta/pangu)
- [pangu.java](https://github.com/vinta/pangu.java)
- [pangu.space](https://github.com/vinta/pangu.space) (HTTP API)

## Installation

```bash
# as a library
$ go get github.com/vinta/pangu/v4

# as a CLI, installs both `pangu` and `pangu-go`
$ go install github.com/vinta/pangu/v4/cmd/...@latest
```

## Usage

### In Go

```go
package main

import (
	"fmt"

	"github.com/vinta/pangu/v4"
)

func main() {
	fmt.Println(pangu.SpaceText("你從什麼時候開始產生了我沒使用Monkey Patch的錯覺?"))
	// 你從什麼時候開始產生了我沒使用 Monkey Patch 的錯覺?

	fmt.Println(pangu.HasProperSpacing("聽說 Hadoop 工程師睡不著的時候都會 Map/Reduce 羊"))
	// true
}
```

### In CLI

```bash
$ pangu-go "為了讓公司的開發流程正常化，有人提議要導入DevOps，但是因為有部分工程師反對，主管決定讓大家投票表決，有三個選項1.導入2.不導入3.維持現狀"
為了讓公司的開發流程正常化，有人提議要導入 DevOps，但是因為有部分工程師反對，主管決定讓大家投票表決，有三個選項 1. 導入 2. 不導入 3. 維持現狀

$ pangu-go -t "為什麼小明有問題都不Google？因為他有Bing"
為什麼小明有問題都不 Google？因為他有 Bing

$ pangu-go -f path/to/file.txt
未來的某一天，Gmail 配備的 AI 可能會得出一個結論：想要消滅垃圾郵件最好的辦法就是消滅人類

$ pangu-go -c "心裡想的是Microservice，手裡做的是Distributed Monolith"; echo $?
Corrected: 心裡想的是 Microservice，手裡做的是 Distributed Monolith
1

$ echo "Workaround雖可恥但有用" | pangu-go
Workaround 雖可恥但有用

$ go run github.com/vinta/pangu/v4/cmd/pangu@latest "聽說桐島rm -rf /*了"
聽說桐島 rm -rf /* 了
```

`pangu` and `pangu-go` are the same command. Use `pangu-go` when pangu.js or pangu.py also installs a `pangu` on your `PATH`.

## License

Released under the [MIT License](https://opensource.org/licenses/MIT).

## Author

- GitHub: [@vinta](https://github.com/vinta)
- Twitter: [@vinta](https://twitter.com/vinta)
- Website: [vinta.ws](https://vinta.ws/code/)
