## Amount       
return a random amount with special prefix        
```go
lorem.Amount(min float64, max float64, decimal uint, prefix string) string    
```
example:      
```go
lorem.Amount(1000, 99999, 4, "$") // $49,605.6515 
```

## BankCardNumber      
return a random bank card number. support Visa, Amex, Jcb, MasterCard and UnionPay card issuer        
signature:        
```go
func BankCardNumber(issuer string) string
```
example:      
```go
lorem.BankCardNumber(lorem.Visa) // 40299297434197151
```