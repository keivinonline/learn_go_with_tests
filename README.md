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

## maps 
- can modify maps without passing as an address to it e.g. `&myMap`
- `a map value is a pointer o a runtime.hmap structure`
- when passing a map to a function/method, it only copies the pointer part and not the underlying data structure that contains the data. 
- maps can be `nil`
- reading `nil` map returns empty map
- writing to `nil` map causes *runtime panic*
```go
// NEVER initialize nil map
var m map[string]string
// DO THIS INSTEAD
//// 1 - either this
var dict = map[string]string{} 
//// 2- or this 
var dict = make(map[string]string)
```

## dependency injection
- there are lots of common misunderstandings on this topic
### benefits
- don't need a framework
- does not overcomplicate design
- facilitates testing
- allows to write general-purpose functions
### details
- AKA `pass in` 
- function does NOT NEED TO CARE about
    - where the print happens
    - how the print happens 
- just need to accept an `interface` rather than `concrete` type 
- based on `fmt.Printf`,
    - From this we can infer that os.Stdout implements io.Writer; Printf passes os.Stdout to Fprintf which expects an io.Writer.
```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
### samples
```go
	// stdout - hard to capture using testing framework
	fmt.Printf("Hello, %v\n", name)
// Use io.Writer interface
//  - compatible with any other struct that implements io.Writer 
//  - e.g. os.Stdout, http.ResponseWriter
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, greetingPrefix+name)
}
// instead of bytes.Buffer 
//  - which does not implements io.Writer  interface
func Greet(writer *bytes.Buffer*, name string) {
	fmt.Fprintf(writer, greetingPrefix+name)
}

```
### summary
- Testing the code
    - DI motivates you to inject in a db dependency via interfaces which you can mock
- separate our concerns
    - decouple where the data goes and how to generate it 
    - e.g. of too many responsibilities 
        - method/func generating data AND writing to db 
        - method/func handling HTTP requests and doing domain level logic
- allow code to be reused AKA DRY
    
### io.Writer interface
- by using this interface, we can use `bytes.Buffer` in the tests
- can use other `Writer`s from std lib like `os.Stdout` to use the function the matches `io.Writer` 
