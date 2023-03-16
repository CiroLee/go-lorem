## Unit
return a random uint. note: min must be less than max, if not, will throw an error     
signature: 
```go
func Uint(min, max unit) (uint, error)
```
example:     
```go
n,_ := lorem.Unit(10, 100) // 14
```

## Int    
return a random int.       
signature:   
```go
func Int(min, max unit) int
```
example:     
```go
n,_ := lorem.Int(-10, 10) // -2
```

## IntBy      
return a random int via special digit. if set `positive` true, will return a positive int number, and set false, return a negative int number           
signature:     
```go
func IntBy(digit uint, positive bool) (int, error)
``` 
example:    
```go
lorem.IntBy(3, true) // 235
lorem.IntBy(3, false) // -374
```

## Float32     
return a random float32 number with `decimal` decimal places          
signature:    
```go
func Float32(min, max float32, decimal unit)(float32, error)
```
example:     
```go
n,_ := lorem.Float32(10, 100, 2) // 23.34
```

## Float64    
return a random float64 number with `decimal` decimal places     
signature:     
```go
func Float64(min, max float64, decimal unit) (float64, error)
```
example:     
```go
n,_ := lorem.Float64(10, 100, 6)  // 8.169075
```