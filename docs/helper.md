## Elements     
return random elements of the array at special num         
signature:     
```go
func Elements[T any](arr []T, num uint) []T
```
example:     
```go
lorem.Elements([]int{1,2,3,4,5,6,7,8}, 2) // []int{2,7}
```