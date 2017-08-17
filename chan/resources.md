# Resources
Some links to related informations.

## [blog.golang](https://blog.golang.org/) - Articles
- [Pipelines](https://blog.golang.org/pipelines)
- [Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)
- [Concurrency is not parallelism](http://blog.golang.org/concurrency-is-not-parallelism)
- [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/go-concurrency-patterns-timing-out-and)
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)

## [YouTube](http://www.youtube.com/) - Videos
- [Go Concurrency Patterns](http://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Concurrency Patterns](http://www.youtube.com/watch?v=QDDwwePbDtw)

## other blogs
- [Dave Cheney: ](https://dave.cheney.net/resources-for-new-go-programmers)

- [Golang channels tutorial](http://guzalexander.com/2013/12/06/golang-channels-tutorial.html) by [alexander Guz](http://guzalexander.com/)

## books
- [Communicating Sequential Processes (CSP)](http://www.usingcsp.com/cspbook.pdf) *The* CSPbook
- **The Go Programming Language (Addison-Wesley Professional Computing Series)**
    * Author: Alan A.A. Donovan and Brian Kernighan
    * Publication Date: November, 2015
    * ISBN: 978-0134190440
    * Reference: http://www.gopl.io/

## further readings
- [Golang Internals Part 2: Nice benefits of named return values](https://blog.minio.io/golang-internals-part-2-nice-benefits-of-named-return-values-1e95305c8687)

- [Go by Example: Channels](https://gobyexample.com/channels)
  "*Channels* are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine."
  "Channels are typed by the values they convey."

---
- [Simple Data Processing Pipeline with Golang](https://www.hugopicado.com/2016/09/26/simple-data-processing-pipeline-with-golang.html) by [Hugo Picado](https://www.hugopicado.com/) Sep. 26, 2016
- [Sources](https://github.com/picadoh/gostreamer)

"In this example we are building a simple processing pipeline that consumes a text line from a socket and sends it through a series of processes to extract independent words, filter the ones starting with # and printing the result to the console. For this, a set of structures and functions were created so we can try around and build other kind of pipelines at will."

Has a `Collector.Execute` as `Fan-In(cap=1)` and a `Processor.Execute`,and a `ChannelDemux.Execute` for non-random FanOut. TODO: ReView for FanOut.

Uses
```go 
	type ProcessFunction func(name string, input Message, out chan Message) 
```

Code [flavour](chan/flavour.md) is `ssss`

---
- [Fan-out-Fan-in/package](https://go.hotlibs.com/github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/package)
  [repo](github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in)

*Warning*: These are just misnamed copies:
- [Fan-out example](https://gist.github.com/mchirico/df9fad3e7a5ea0c4527a)
  [same](https://www.snip2code.com/Snippet/1043022/Go-(Golang)-Fan-out-example/)
  `merge` - a Fan-In, not -Out! (with sync.WaitGroup for closer, and 'done' as context)

- [Fan-out-Fan-in/package](https://go.hotlibs.com/github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/package)
  [repo](github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in)

---
[Channels Are Not Enough](https://gist.github.com/kachayev/21e7fe149bc5ae0bd878) ... or Why Pipelining Is Not That Easy -
by [@kachayev](https://twitter.com/kachayev)
- Unicorn Cartoon :-)
- Fan-In sample from above, "(shamelessly stolen from [here](http://blog.golang.org/pipelines))" 
- Delves into "channel is a functor" and "Futures & Promises", but does not distinguish supply and demand (but uses it)

---
[Buffered Channels In Go — What Are They Good For?](https://medium.com/capital-one-developers/buffered-channels-in-go-what-are-they-good-for-43703871828) by [Jon Bodner](https://medium.com/@jon_43067)

Buffered channels are useful when you know how many goroutines you have launched, want to limit the number of goroutines you will launch, or want to limit the amount of work that is queued up.

- Parallel Processing
- Creating a Pool

---
[Abundant concurrency in Go](https://hunterloftis.github.io/2017/07/14/abundant-concurrency/)

Contrasts his 'JavaScripter’s mindset' with an 'abundance mindset'. 

---
### [Experience Reports - Generics](https://github.com/golang/go/wiki/ExperienceReports#generics)

---
[Using code generation to survive without generics in Go](https://dev.to/joncalhoun/using-code-generation-to-survive-without-generics-in-go)
Uses simple {{.Name}} & {{.Type}} templates (no multi-types). Leaves 'import' to `goimport`.

---
[gen](http://clipperhouse.github.io/gen/) uses typewriters ...

[genny](https://github.com/cheekybits/genny) uses generic.Type in master sample

[gengen](https://github.com/joeshaw/gengen)

[StaticTemplate](https://github.com/bouk/statictemplate)
is a code generator for Go's text/template and html/template packages.
- It works by reading in the template files, and generating the needed functions based on the combination of requested function names and type signatures.
- Please read [my blogpost](http://bouk.co/blog/code-generating-code/) about this project for some background.
- TODO: Would allow to create go code from the chan templates 
---
[Closures are the Generics for Go](https://medium.com/capital-one-developers/closures-are-the-generics-for-go-cb32021fb5b5) by [Jon Bodner](https://medium.com/@jon_43067)

"We could use `interface{}` as a way to pass untyped input and output parameters around but that misses the point. It creates ugly code that requires casts and subverts the type system that helps us write correct code."

Shows clearly, that 'generics' can be attacked and solved using
- generated code - for generic algorithms & data structures
- closures - as shown here
  - see below
  - `type sorter struct { ... }` - see below - aka `dancing` in dlx

[use a closure](https://play.golang.org/p/dNmhg_6x9T)

	package main

	import (
		"fmt"
	)

	func outer2(name string) int {
		var total int
		myClosure := func(x int) {
			total = len(name) * x
		}
		helper(myClosure)
		return total
	}

	func helper(f func(int)) {
		f(4)
	}

	func main() {
		fmt.Println(outer2("hello"))
	}

Result: 20 
- = 5 * 4
- = _`(len("hello"))`_ * 4 )
- = _`(len("hello") * x )`_(4)
- = _`(len(name)`("hello")` * x )`_(4)

"Look at `outer2`. Its local variable `total` was modified when `myClosure` was passed to `helper` and called from there. There’s no reference to `total` in `helper`, but using a closure allowed it to be modified. Just like structs in Go, closures have state. This state provides the solution to our problem."

---
[Closure: sorter](https://play.golang.org/p/hYMcQ81AvN)

	package main

	import (
		"fmt"
		"sort"
	)

	type sorter struct {
		len  int
		swap func(i, j int)
		less func(i, j int) bool
	}
	
	func (x sorter) Len() int           { return x.len }
	func (x sorter) Swap(i, j int)      { x.swap(i, j) }
	func (x sorter) Less(i, j int) bool { return x.less(i, j) }
	
	func Sort(n int, swap func(i, j int), less func(i, j int) bool) {
		sort.Sort(sorter{len: n, swap: swap, less: less})
	}

	func main() {
		a := []int{5, 4, 3, 2, 1}
		Sort(
			len(a),
			func(i, j int) {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			},
			func(i, j int) bool {
				return a[i] < a[j]
			})
		fmt.Println(a)

		b := []string{"bear", "cow", "ant", "chicken", "dog"}
		Sort(
			len(b),
			func(i, j int) {
				temp := b[i]
				b[i] = b[j]
				b[j] = temp
			},
			func(i, j int) bool {
				return b[i] < b[j]
			})
		fmt.Println(b)
	}

---
TODO: Add a comment about [dotgo](https://githib.com/GoLangsam/dotgo) & [AnyType](https://githib.com/GoLangsam/AnyType)

---
## Pitfalls

### `case <-done` - The Ordering Trap
[Notifications on the channels in Golang](http://blog.atte.ro/2017/07/09/golang-channel-notification-select.html) - The Ordering Trap
- A Pitfall in `select` re order between `case <-inp` and `case <-done`

NoGood:
	func (f *Foo) Number() (int, error) {
	    select {
	    case <-f.doneChan:
	        return 0, errors.New("canceled")
	    case <-f.numberChan:
	        return f.number, nil // Deliver result
	    }
	}

Better: Deliver also upon done, if available
	func (f *Foo) Number() (int, error) {
	    select {
	    case <-f.doneChan:
	        select {
	        case <-f.numberChan:
	            return f.number, nil // Deliver result
	        default:
	            return 0, errors.New("canceled")
	        }
	    case <-f.numberChan:
	        return f.number, nil // Deliver result
	    }
	}

---
[](https://www.cockroachlabs.com/blog/squashing-a-schroedinbug-with-strong-typing/)

### Lesson learned

Imagine you have a toy, and the object of the toy is to fill it by passing its contents through appropriate slots. Ideally, you’d want the toy to tell you that you’re doing it wrong by using different shapes for the different objects and holes. Instead, with CockroachDB, all the objects and holes were shaped the same and the rule was instead to “pay attention to the color.” What’s worse is that Go didn’t help us realize we were also color blind. A schrodinbug in CockroachDB

**Strong(er) typing really helped us put things back in shape.**

We’ll do it more from now on.

=> container/fs - documentation

---
[Buffered Channels](https://medium.com/capital-one-developers/buffered-channels-in-go-what-are-they-good-for-43703871828)

See do_experiments/pool
TODO: Make this pool more generic
TODO: Use 'pool' for concurrent output into different directories

---
[](https://deedlefake.com/2017/07/the-problem-with-interfaces/)

"Generics, specifically type parameters, are the standard solution to these issues. Though it may not be the best solution, examples of the need for a solution are so rampant that they can even be found throughout the standard library, and particularly in the various `container` subpackages, and most recently in the form of `sync.Map`.
Just the fact that there are so few packages under `container` could be considered an example of these problems."

---
[Fuzzing](https://en.wikipedia.org/wiki/Fuzzing)
TODO: apply to `dotgo`, and `dotpath` especially.

---
[Text based user interfaces](https://appliedgo.net/tui/)

## Licenses
- [Choose a license](https://choosealicense.com/)
- [Open Source Guide](https://opensource.guide/legal/#which-open-source-license-is-appropriate-for-my-project)

## People

### The Project Team
started the project in late 2007

- "Rob 'Commander' Pike" <r@google.com> - [@robpike](https://github.com/robpike) -
[@ Google+](https://plus.google.com/101960720994009339267) -
[@BlogSpot](https://robpike.blogspot.com/)
- "Ken Thompson" <ken@google.com>
- "Robert Griesemer" <gri@golang.org> `gri`

---
joined in 2008
- Ian Lance Taylor
  - e.g. `gccgo`
- "Russ Cox" <rsc@swtch.com> - [@rsc](https://github.com/rsc) - ![rsc](rsc.png)
  - [swtch.com](http://research.swtch.com/)

---
- "Brian Mills" <bcmills@google.com> - [@bcmills](https://github.com/bcmills)
  - e.g. `syncmap`

---
### Contributors
