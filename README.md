# go-lorem
> A golang library for generating mock data           

[![codecov](https://codecov.io/gh/cirolee/go-lorem/branch/main/graph/badge.svg)](https://codecov.io/gh/cirolee/go-lorem/branch/main) ![GitHub](https://img.shields.io/github/license/CiroLee/go-lorem)

## install    
```shell
go get github.com/CiroLee/go-lorem
```
## usage     
```go
import 
  (
    "github.com/CiroLee/go-lorem"
  )

func main() {
  // generate a Chinese word with two characters
  t := lorem.Word(2, "zh")
  // 汉啊 
}
```

## Docs
detail apis see [docs](./docs/README.md)
