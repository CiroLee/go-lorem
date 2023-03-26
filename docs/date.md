## Timestamp     
return a random Unix timestamp, support `sec` for second, `milli` for millisecond and `nano` for nanosecond      
signature:    
```go
func Timestamp(from, to time.Time, format string) int
```
example:    
```go
from := time.Date(2023, time.Month(1), 1,0,0,0,0, time.Now().Location())
to := time.Date(2023, time.Month(12), 23,59,59,0,0, time.Now().Location()) 

lorem.Timestamp(from, to, "sec") // 1673732208
lorem.Timestamp(from, to, "milli") // 1697955943535868
lorem.Timestamp(from, to, "nano") // 1685705856559827527
```

## Date      
return a random UTC date     
signature:   
```go
func Date(from, to time.Time) time.Time
```
example:    
```go
from := time.Date(2023, time.Month(1), 1,0,0,0,0, time.Now().Location())
to := time.Date(2023, time.Month(12), 23,59,59,0,0, time.Now().Location()) 

lorem.Date(from, to) // 2023-05-07 17:00:02.878863734 +0000 UTC
```

## FormatDate      
return a format date via layout       
signature:    
```go
func FormatDate(from, to time.Time, layout string) string
```
example:     
```go
import (
  "github.com/CiroLee/go-lorem/consts"
)
from := time.Date(2023, time.Month(1), 1,0,0,0,0, time.Now().Location())
to := time.Date(2023, time.Month(12), 23,59,59,0,0, time.Now().Location()) 
// use built-in layout
lorem.FormatDate(from, to, consts.ZhLayout) // 2023年3月12日 15时04分05秒
```
## Week    
return a random week day support Chinese, English, and English supports abbreviation if set `abbr` true    
signature:   
```go
func Week(lang string, abbr bool) string
```
example:     
```go
lorem.Week("en", false) // Wednesday
lorem.Week("en", true) // Mon.
lorem.Week("zh", false) // 六
```
## Month    
return a random month, support Chinese, English, and English supports abbreviation if set `abbr` true      
signature:     
```go
func Month(lang string, abbr bool) string
```
example:    
```go
lorem.Month("en", false) // August
lorem.Month("en", true) // Apr.
lorem.Month("zh", false) // 十一月
```