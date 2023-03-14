## Ipv4    
return a random ipv4 address     
signature:    
```go
func Ipv4() string
```
example:    
```go
lorem.Ipv4() // 198.24.24.125
```

## Ipv6     
return a random ipv6 address       
signature:      
```go
func Ipv6() string
```
example:     
```go
lorem.Ipv6() // 1f2c:3ff1:4016:7734:e014:fa37:becd:e726
```

## Tld     
return a random top-level-domain. `tldType` has two types: iTLD and ccTLD, means international Top-Level Domain name and country code Top-Level Domain              
signature:      
```go
func Tld() (tld, tldType string)
```
example:    
```go
lorem.Tld() // .gov iTLD
```