# Gain concurrency by using go-channels for piping (and more)

Piping is a popular concept, method, facility, mechanism ... whatever You call it.

And(!) piping is a convenient way to design a process - a process with concurrent parts. Batteries included :-)

Each such part (and any composite hereof(!)) may be seen as a _pipeable-unit_ (or pipe-station, or pipe-tube).

This project _chan_ provides facilites and mechanisms to build, invoke and manage some pipeable-unit.

Several such pipeable-unit become connected into a networked ensemble.

This is useful to desing, build, create and operate a process.
And: Such process is concurrent by design!

Thus: it's components may execute as parallel(!)
(as much as Your environemt permits/supports).

Really(!) parallel (on multi-cores), that is.
Not only quasi-parallel (via some pre-emptive multitasking).

Note: Mind You: Piping is a mechanism as much as a channel is not (just) a type, but should better be seen as a mechanism, a means to some end.)

## pipeable-unit

Any pipeable-unit has (zero or more) input(s) and/or (zero or more) output(s).

Note: It could also be named transitor, as it facilitates stuff to transit thru it.

As all available literature about piping and concurrency (so far we've read *both* good books available to us) is inconsistent and/or incomplete in terms of nomenclatura, we take liberty to introduce some freshly invented vocabulary below.

Hint: If You like it: smile. If You don't: smile anyway, and You'll be (a little more) happy.

### shape

The shape/form of a pipeable-unit can vary:

#### one-sided - unary
* pumpfeed (or pumpfill):
	* has only output, no input
* sinkhole
	* has only input, no output

#### two-sided - binary
* extender
	* one input & one output
	* both of same type
* converter
	* one input & one output
	* of different type
* adapter (or caster)
	* one input & one output
	* of different type
	* connected by nothing but a type-conversion/extraction/injection

#### fan-shaped - unary/n-ary
* fan-in (= n-ary/unary)
	* several inputs & one output
	* all of same type
* fan-out (= unary/n-ary)
	* one input & several outputs
	* all of same type

Warning: some stupid & distracting question/bad joke/pun is ahead - please feel free to consider to skip next line!
	* A fan is also called a gopher, is it not? Okey - when it's a fan of go, at least ... , is it not? Yawn ... )

#### double-fan - gopher-shaped
* hourglass
	* several input & several output
	* all of same type

* bottleneck
	* several input & several output
	* each side of different type

* gopher-task :-)
	* several input & several output
	* all of different type

#### afterthoughts
Note:
* Being short of phantasy, we do not (yet) provide any hourglass, bottleneck or (atomic) gopher-task. ;-)
* Note: Please feel free to pass enlightening suggestions/remarks/comments on to us.

Note:
* Intentionally we do not supply any 'real' sinkhole as such would give more pain than gain.
* ( and should better be named 'wormhole', as it's use would open a can of bugs - I mean: worms. )
* Instead, we recommend to use the "done-idiom" and supply respective DoneXyz-functions, which return a single-event-channel and signal completion hereon.
* Note: Well, if You don't care to listen - it's Your choice: sink & stink & abort prematurely. RIP.

### plugability
The plugability of a pipeable-unit can vary (as much as the arguments passed to and the results returned from functions (in go)):

* only out (see Generators)
* ...
* same-same (the easy one)
* in-out (limitless combinations - beyond imagination)
* ...
* two-out
* in-two
* ...
* two-two

Not to mention go-idioms such as "comma-ok" or "comma-error" ...

## implementation

### functions

Any pipeable-unit is modeled by a suitable function.

Being lazy, we pack and provide a library (and intentionally(!) *not* a framework).
A package with a library of functions => a library of pipable-unit. Batteries included :-)

Being *very* lazy, we (have to) generate go-code for any type.

### _chan_ - go-channels
We use (mostly) nonblocking go-channels to model the inputs and outputs of any pipeable-unit.

More specifically, 

They may have various types.
The strict type-system of go assures for any joined connection to be a proper match
and conveniently signals eventual errors at compile time already..

#### send-only channel of SomeType
Pass type:

	chan<- SomeType

Allows sends but not receives.

#### receive-only channel of SomeType
Pass type:

	<-chan SomeType

Allows receives but not sends.

Hint:
* The position of the <- arrow relative to the chan keyword is a mnemonic.
* Note: Violations of this discipline are detected at compile time.

Note:
* In order to keep things easy (and safe), we focus on the receive-only kind!
* The supplied functions accept and/or supply/return receive-only channels only.
* The sending into the returned channel is launched as part of the function.
* Batteries included :-)

## flavours

Any such pipeable-unit comes in different flavours (or "Sizes").
( Mind You: "one size fits all" ain't real - it's just plain nonsense (popularised by lazy/eager dress dealers) :-) ).

Such as:
* *SSS*	super small & simple
* *SS*	super simple
* *S* 	supply-driven channel types (think: easy, push, provide, eager)
* *L* 	demand-driven channel types (think: lazy, poll, request, demand, sleuth)
* *XS*	eXtended (with Context and functionality) easy supply-driven channel types 
* *XL*	eXtended (with Context and functionality) lazy demand-driven channel types

Note:
* _S_ and _L_ are just symmetric opposites
* some people believe, s is easier to understand and to reason about.
* Thus: s might be more appropiate for smaller brains ... ;-)

In order to bridge/convert between _S_ and _L_, we have *SL* (like the famous Mercedes sport-car) for those, who need to mix supply-driven and demand-driven channels.

## Naming conventions
Let *Pack* be the ID of a package such as strings, ioutil or os

Let *pack* be the import-path for _Pack_ such as "strings", "io/ioutil" or "os"

Let *Type* be the Name of a type implemented/supported by _Pack_, such as "String", "Writer", "Header", "File" ... (or ""; if it's id is obvious for the given _Type_)

Let *type* be the actual type underlying _Type_ such as string or *io.Writer

Let *Func* be the ID of an exported function to be wrapped, such as ToUpper, Open, ...
or some newly invented Name to characterise the lauched functionality

### Prefixes
* ChanXyz ([args] as-needed) <-chan _type_
	* Inputs: anything but a read-only chan to be consumed
	* Launch: a producer
	* Return: a read-only chan

* PipeXyz (inp <-chan _type1_ [, args]) ( <-chan _type2_ [, more] )
	* Launch: a producer for _type2_ which consumes _type1_ and closes inp

* SinkXyz (inp <-chan _type_ ) // or NullXyz ?
	* Launch: a simple drainer - intentionally *not* provided!

* DoneXyz (inp <-chan _type_ ) <- chan struct{}
	* Launch: a simple drainer, and give ONE signal on the returned channel when done


### Postfixes
#### Postfix "_" - ignore secondary results such as ok's/errors.
Many functions return multiple values - especially the "comma-error" idiom is very popular.

In order to deal conveniently therewith, we use _ as postfix for chan-wrapper-implementations which ignore errors (or other secondary data such as bytes written...) and return a chan with the 'main' data only.
Any such quiet/skipping/ignoring implementation shall be complimented with it's fully returning companion. Thus, implementors are free to choose.

#### Postfix "s" - Plural via variadic arguments

Note: In templated code, only used by generators (Prefix "Chan").

#### Postfix "S" - Plural via Slice 

Note: In templated code, only used by generators (Prefix "Chan").

### Infixes
#### Infix "_" to omit optional arguments
"_" is used as Infix when secondary arguments are ignored/not to be supplied,
and reasonable and documented(!) defaults are subsituted.

Note: A twin without "_" should behave exactly similar, if called with nil argument(s).

Note: Yes, this is syntactic sugar.
Yes, this is useful and practical (and thus almost mandatory)
as long as go does not allow optional arguments.

### Generic functions
#### Creators
* Make _Type_ ()  <-chan _type_

#### Generators
Hint: A [Generator](https://en.wikipedia.org/wiki/Generator_(computer_programming)) yields a sequence of values ... one at a time.
* Chan _Type_ s (inp... _type_ ) <-chan _type_
* Chan _Type_ S (inp [] _type_ ) <-chan _type_

For _Func_ in _pack_ with signature func ( [args] ) ( _type_ [, error]/[, more] )

* Chan _Type_ _Func_ 
* Chan _Type_ _Func_ _ (if [, error] )

#### Pipetubes
* Pipe _Type_ Func  (inp [] _type_ , act func( _type_ ) _type_ ) <-chan _type_
	Note: it 'should' be Pipe _Type_ Map for functional people,
	alas: map has a different meaning in go ...

* Pipe _Type_ Filter  (inp [] _type_ , pass func( _type_ ) bool ) <-chan _type_
* Pipe _Type_ Skiper  (inp [] _type_ , skip func( _type_ ) bool ) <-chan _type_

* Done _Type_ (inp <-chan _type_ [, args]) <-chan struct{}
	an "informative sinkhole" with a simple event channel aka "<-done"-idiom.

For _Func_ in _pack_ with signature func ( _type_ ) ( _type_ [, error]/[, more] )

* Pipe _Type_ _Func_ 
* Pipe _Type_ _Func_ _ (if [, error] )

For _Func_ in _pack_ with signature func ( _type_ ) ( _type-out_ [, error]/[, more] )

* Pipe _Type-out_ _Func_ _Type_ 
* Pipe _Type-out_ _Func_ _Type_ _ (if [, error] )

#### Fanning
* FanOut
* FanIns
* FanInS
* FanOutUpTo (with a workLimit semaphore)


### Source file names
Note: currently, there is exactly one file named "chan.go" (and some related chan*_test.go)
in "runtime": src/runtime/chan.go. The implementation of channels.

Thus: We feel free to choose and use "chan" as prefix for all our source files.
Even if we 'inject' such into existing packages in order to enhance them in place.

More specifically:
* "chan.go"
  for manual contributions

* "chan\_generic.go"
  for generated source files ( if _Type_ == "" )
* "chan\_ _Type_ \_generic.go"
  for generated source files ( if _Type_ != "" )


### misc notes:
Hint: SeeAlso crypto/md5/gen.go
// +build ignore
//go:generate echo Success
//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p arm64
	must be 