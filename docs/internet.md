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

## Url      
return a random url via the option struct            
signature:    
```go
type UrlOption struct {
	SubDir       uint      // layers of sub-directory,default is between [1,4]
	Protocol     string    // protocol of url,default is random
	Suffix       string    // suffix of the url, such as .com, .org,default is random
	DomainLevel  uint      // level of sub domain, default is between [1,3]
 }

 func Url(option UrlOption) string
```
example:      
```go
lorem.Url(lorem.UrlOption{
  SubDir:      2,
  Suffix:      ".com.cn",
  DomainLevel: 3,
})
// ftp://kgqje.c7r62fxqb.fol.com.cn/u/ksfnh
```

## SimpleUrl      
return a random url with no params. It will generate common url with https protocol       
signature:     
```go
func SimpleUrl() string
```
example:     
```go
lorem.SimpleUrl() // https://bzqy--3.t05wq8hf92.com/r/jmw
```

## Email     
return a random email      
signature:     
```go
func Email() string
```
example:     
```go
lorem.Email() // g5mVoEWH3v/@foxmail.com
```