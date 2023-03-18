## Letter    
return a random letter, support Chinese(zh) and English(en)   
signature:    
```go
func Letter(lang string) string
```
example:    
```go
lorem.Letter("en") // a
lorem.Letter("zh") // 汉
```

## Word     
 return a random word consisting of `num` letters. At least two letters          
 signature:     
 ```go
 func Word(num uint, lang string) string 
 ```
 example:    
 ```go
lorem.Word(3, "en") // awe
lorem.Word(3, "zh") // 函还阿
```

## Sentence    
return a random sentence consisting of `num` of words. At least one word, and word is based on `baseNum = 10`      
signature:    
```go
func Sentence(num uint, lang string) string
```
example:     
```go
lorem.Sentence(2, "en") // Isakp krjnx
lorem.Sentence(2, "zh") // 泡缸弓，等兄
```

## Paragraph     
return a random paragraph consisting of `num` of sentences. At least two 1 sentence      
signature:    
```go
func Paragraph(num unit, lang string) string
```
example:      
```go
lorem.Paragraph(2, "en") 
// Vknxo ebsgai fiaxumn ynlnwy xpswu aegfkg dcy npcdyhb ajndxdtb.Mnw kikjfyfs jxw jyotivkh.
lorem.Paragraph(2, "zh")
// 舱幼移灰膝涉娱疲，早驴池代央。握时，墨留安诱，息孤省点轧删躺，婶孝阁董绩佛，飘攀辞柴秃腿尸颈，妻眼文应。
```

## Name    
return a random Chinese name or English name. And if want use capitalized English name, set second param `upper` to true. Only useful to English      
signature:    
```go
func Name(lang string, upper bool) string
```
example:     
```go
lorem.Name("en", false) // pearl
lorem.Name("en", true) // Vicki
lorem.Name("zh", false) // 胡东贵
```

## Str     
return a random string at special length. min length is 4             
signature:    
```go
func Str(length uint) string
```
example:     
```go
lorem.Str(6) // *%=}P^
```
## StrBy     
return a random string using special source        
signature:     
```go
func StrBy(length uint, source string) string
```
example:      
```go
d := lorem.StrBy(4, "abcdefghijklmnopqrstuvwxyz1234567890")  // htgp
```