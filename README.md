# go-lorem
> A golang library for generating mock data

**note**: Not stable,still in rapid development
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
