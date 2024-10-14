# learning go with tests
## rules of thumb
- the test speaks to us more clearly, as if it were an assertion of truth, `not a sequence of operations` 

## arrays and slices
- ref: https://go.dev/blog/slices-intro
### arrays
- size is encoded in its type
- do NOT need to be initialized
- Go's arrays are values; not a pointer to the first array element like in C 
- AKA array is a soft of struct but with indexed rather named fields: fixed-size composite value 
```go
a[2] == 0 // ready to use
myArray := [4]int{1,2,3} // prints to [1,2,3,0]
// literal 
b := [2]string{"apple", "orange"}
// compiler count for us 
b := [...]string{"apple", "orange"}

```
### slices
- do NOT encode size of collection and instead has dynamic size
```go
// literal
mySlice := []int{1,2,3}
// via make function
mySlice := make([int{}])


```

### structs
```go
// %#V prints struct and values 
// %g - max default precision
t.Errorf("%#v - got %g want %g", tt.shape, got, tt.want)
```