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
```
### pointers
- `struct pointers` are `automatically dereferenced` 
```go
// This is valid
func (w *Wallet) Balance() int {
	return (*w).balance
}
// But it equals to 
func (w *Wallet) Balance() int {
	return w.balance
}
```
### type alias
- domain specific types 
```go
// custom type
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}
``` 
### Stringer
```go
// Defined in fmt package
// Defines how type is printer when using %s prints
type Stringer interface {
	String() string
}
// from: "got 0 want 10"
// to:  "got 0 BTC want 10 BTC"
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```
## errcheck
```bash
go install github.com/kisielk/errcheck@latest

errcheck . 
```

## pointers
- when there is a need to mutate state, use `pointers`
- example usage
    - referencing large data structures
    - db connection pools
## nils
- pointers can be nil
- when using function that returns pointers, check of nil
    - else might raise runtime panic
- useful when want to describe a value that is missing
## errors
- signify failure when calling function or method
- check for errors, but also handle them gracefully
