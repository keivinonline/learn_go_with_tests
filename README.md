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
## mocking
- important skill to slice up requirements as small as you can
### what mocking can do but tests can't
- e.g. we have a countdown app that has 3 seconds of sleep
```go
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}
```
- slow tests ruin dev productivity
- it has a dependency on `Sleep`ing 
- we could *mock* `time.Sleep` instead of *real* `time.Sleep` and we can spy on the calls
- if mocking becomes too hard,
    - break the module apart
    - dependencies are too fine-grained
    - tests are too concerned with implementation details
        - test the behavior instead
- focus on testing *useful behavior* unless certain implementation details are critical
### test guidance
- `refactoring`
    - code changes but behavior remains the same 
    - "am I testing the behavior I want, or the implementation details ?"
    - "if I were to refactor this code, would I have to make lots of changes to tests ?"
- focus on testing `public functions` as `private functions` are part of the implementation detail to support the former
- if a test is working with more than 3 mocks - it is a red flag. time to rethink the design
- `use spies with caution`
    - only use it if you really need to spy on the inner workings
    - has tight coupling
### what can mocking do ?
- calling a service that *can fail*
- testing `states` of a system
- mocking dbs
- mocking webservices
### test double
- mocks are also known as `test double`
## concurrency 
- `having more than 1 thing in progress` 
- e.g. waiting the kettle to boil while making toast
### CheckWebsites
- instead of waiting for a website to respond before sending a request to next site, tell the program to make the next request white it is waiting - AKA do not site idle 
- *blocking* - normally, when we call `DoSomething()` func, computer will wait for it to finish
- *goroutine* - when an operation does not `block` in Go and run in a separate process
- Analogy
    - Process - reading down the page of Go code from top to bottom
    - Separate process - another *reader* begins reading inside each function while letting the original reader carry on ready down a page. AKA like spawning a sub process ? 
- use `go doSomething()` to start a goroutine
```go

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// // Blocking operation
		// results[url] = wc(url)

		// // Goroutine via anonymous functions
        // // each iteration of the loop starts a new goroutine
		go func() {
			results[url] = wc(url)
		}()

	}
	return results
}

```
### Anonymous functions
- can be executed at the same time they are declared via the `()` 
- maintain access to all scope defined up till the point of anon func
### Handling concurrency
- hard to predict when not handled properly. Hence tests are required
- the following failed the test as it gives an empty `map[]` as 
    - none of the goroutines spawned had enough time to `add their results to the results map` which ended up being empty
- `fatal error: concurent map writes` 
    - when 2 goroutines write to the results map at the same time AKA *race condition*
- use the race detector 
```go
go test -race
```
```bash
==================
WARNING: DATA RACE
Write at 0x00c000120420 by goroutine 7:
  runtime.mapaccess2_faststr()
...

Previous read at 0x00c000120420 by goroutine 6:
  reflect.mapiterelem()
      /Users/keivinc/.gvm/gos
```
### channels
- data races can be solved by coordinating goroutines using `channels`
- channels can do both  
    - `receive data`
    - `send data`
- 
### benchmarking
- before
```bash
❯ gtb
goos: darwin
goarch: arm64
pkg: github.com/keivinonline/elastic-go-examples/learn_go_with_tests/concurrency
BenchmarkCheckWebsites-8               1        2328199000 ns/op
PASS
ok      github.com/keivinonline/elastic-go-examples/learn_go_with_tes
```
- after
```bash
goos: darwin
goarch: arm64
pkg: github.com/keivinonline/elastic-go-examples/learn_go_with_tests/concurrency
BenchmarkCheckWebsites-8              56          21107448 ns/op
PASS
ok      github.com/keivinonline/elastic-go-examples/learn_go_with_tests/concurrency     1.432s
```
### summary
- `goroutines` are the basic units of concurrency in Go
- `anonymous` functions are used to start each of the concurrent processes
- `channels` - help to organize and control communication between diff processes and avoid race condition
- `race detector` helps to debug problems with concurrent code