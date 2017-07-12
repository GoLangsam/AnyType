# Gain concurrency - use go-channels for piping (and more)

## Piping
*Piping* is a popular concept, method, facility, mechanism ... whatever You call it.

And(!) piping is a convenient way to design a process - a process with **concurrent** parts. Batteries included :-)

Each such part (and any composite hereof(!)) may be seen as a _pipeable-unit_.

This project [chan](https://github.com/GoLangsam/AnyType/chan/) provides facilites and mechanisms to build, invoke and manage such pipeable-unit.

They come in a variety of [flavours](flavours.md), have different [sizes](sizes.md), obey to strict and consistent [namings](namings.md).

Several such pipeable-unit become connected into a **networked** ensemble.

This is useful to desing, build, create and operate a process.
And: Such process is concurrent by design!

Thus: it's components may execute as parallel(!)
(as much as Your environemt permits/supports).

Really(!) parallel (on multi-cores), that is.
Not only quasi-parallel (via some pre-emptive multitasking).

Note: Mind You: Piping is a mechanism as much as a channel is not (just) a type, but should better be seen as a mechanism, a means to some end.)

## pipeable-unit

Any _pipeable-unit_ has (zero or more) input(s) and/or output(s).

Note: It could also be named 'transitor' (not transi`s`tor), as it facilitates stuff to transit thru it. Or 'pipe-station', or 'pipe-tube'). [Some](https://blog.golang.org/pipelines) call it 'stage'.

As all available literature about piping and concurrency (so far we've read *both* good books available to us) is inconsistent and/or incomplete in terms of nomenclatura, we take liberty to introduce some freshly invented [vocabulary](Vocabulary.md), complimenting the [namings](namings.md).

Hint: If You like it: smile. If You don't: smile anyway, and You'll be (a little more) happy.
