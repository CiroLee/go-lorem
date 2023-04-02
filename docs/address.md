## Province        
return a random Chinese Mainland province name and code         
signature:       
```go
func Province() AddressBase

type AddressBase struct {
	Code string
	Name string
}
```
example:          
```go
lorem.Province() // {210000 辽宁省}
```

## City       
return a random Chinese Mainland city name and code          
signature:              
```go
func City() AddressBase
```
example:      
```go
lorem.City() // {321000 扬州市}
```

## County           
return a random Chinese Mainland county name and code           
signature:      
```go
func County() AddressBase
```
example:      
```go
lorem.County() // {370911 岱岳区}
```

## Address        
return a random Chinese Mainland address string that include province, city and county, separated by space          
signature:      
```go
func Address() string
```
example:       
```go
lorem.Address() // 浙江省 台州市 路桥区
```

## ZipCode         
return a random Chinese Mainland zip code          
signature:       
```go
func ZipCode() int
```
example:       
```go
lorem.ZipCode() // 460986
```

## LongAngLat        
return a random longitude and latitude array.support deg format and dms format                   
signature:        
```go
func LongAngLat(format string) [2]string
```
example:      
```go
lorem.LongAndLat("deg") // [-45.9817° -67.0966°]
lorem.LongAndLat("dms") // [21°26′47″ 4°50′30″]
```