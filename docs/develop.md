## GitCommitSha        
return a random git commit sha        
signature:     
```go
func GitCommitSha() string
```
example:        
```go
lorem.GitCommitSha() // 2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6
```

## GitCommitShortSha        
return a random git commit short sha         
signature:      
```go
func GitCommitShortSha() string
```
example:      
```go
lorem.GitCommitShortSha() // e0BE3cc
```

## GitBranch          
return a random git branch        
signature:       
```go
func GitBranch() string
```
example:      
```go
lorem.GitBranch() // hotfix/critical-bug-fix
```

## Version       
return a random software version         
signature:      
```go
func Version() string
```
example:     
```go
lorem.Version() // 6.34.12
```

## Password      
return a random password,support specified length and strength(low, medium, high)                     
signature:       
```go
func Password(length uint, strength string) string
```

example:        
```go
lorem.Password(12, "low") // pvljwcssbdbt
lorem.Pasword(12, "medium") // zl22LdITIdcI
lorem.Password(12, "high") // yCH7LsmWQsA!
```

## MD5       
return a random md5 string           
signature:        
```go
func MD5() string
```

example:       
```go
lorem.MD5() // 7755d659cbc55125faa2384e90dc7ff7
```