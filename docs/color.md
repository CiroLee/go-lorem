## Hex        
return a random hex color. set alpha true,will return a hex with alpha channel              
signature:    
```go
func Hex(alpha bool) string
```
example:     
```go
lorem.Hex(true)   // #e82e86a9
lorem.Hex(false)  // #c40ad9
```

## RgbArray       
return a random rgb color array      
signature:     
```go
func RgbArray() [3]int
```
example:     
```go
lorem.RgbArray() // [80 238 187]
```

## RgbaArray      
return a random rgba color array      
signature:     
```go
func RgbaArray() [4]float32
```
example:     
```go
lorem.RgbaArray() // [118 2 23 0.73]
```

## RgbStr      
return a random rgb color string. support legacy and modern style     
signature:    
```go
func RgbStr(style string) string
```
example:    
```go
lorem.RgbStr("legacy") // rgb(31, 88, 91)
lorem.RgbStr("modern") // rgb(6 183 203)
```

## RgbaStr      
return a random rgba color string. support legacy and modern style     
signature:    
```go
func RgbaStr(style string) string
```
example:      
```go
lorem.RgbaStr("legacy") // rgba(232, 29, 74, 0.46)
lorem.RgbaStr("modern") // rgb(210 190 112 / 20%)
```

## HslArray     
return a random hsl color array. saturation and lightness will be decimal between 0 and 1(present percent)     
signature:      
```go
func HslArray() [3]float32
```
example:     
```go
lorem.HslArray() // [196 0.92 0.71]
```

## HslaArray      
return a random hsla color array. saturation and lightness will be decimal between 0 and 1(present percent)         
signature:    
```go
func HslaArray() [4]float32
```
example:     
```go
lorem.HslaArray() // [194 0.35 0.96 0.67]
```

## HslStr      
return a random hsl color string. support legacy and modern style       
signature:     
```go
func HslStr(style string) string
```
example:      
```go
lorem.HslStr("legacy") // hsl(51, 90%, 6%)
lorem.HslStr("modern") // hsl(338 55% 87%) 
```

## HslaStr     
return a random hsla color string. support legacy and modern style       
signature:     
```go
func HslaStr(style string) string
```
example:     
```go
lorem.HslaStr("legacy") // hsla(124, 72%, 49%, 1%)
lorem.HslaStr("modern") // hsl(300 45% 21% / 6%)
```