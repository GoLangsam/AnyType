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
