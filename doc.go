/*
Package skia provides generated bindings for some of [skia]'s API.

A simple example of it's use can be found in [demo].

# naming

Typically skia types that start with "Sk" have the Sk stripped from the
corresponding Go type. For example the [Paint] type represents the skia [SkPaint] class.

# constructors and destructors

Most skia classes have one or more constructors.
They have names of the form `NewSomeclassname[SomeSuffix]`.
For example [NewPaint] and [NewPaintCopy] are constructors for the [Paint] type.
Not all constructors are configured to be generated.

Destructors are available for classes that have public destructors.
Destructors are expose as Delete methods, such as [Paint.Delete].

Classes that are derived from SkRefCnt have an Unref method,
such as [SVGDOM.Unref].

# memory management

Class instances are not automatically deleted by garbage collection when their
Go types are no longer reachable.
In other words there is no use of [runtime.SetFinalizer] to call a destructor.
An application using this package is responsible for calling Delete and Unref
methods when appropriate.

# documentation

Most of the documentation comes directly from documentation comments in skia C++ header files.
So some of it may not be entirely correct for the generated Go code.

[skia]: https://skia.org/
[demo]: https://github.com/pekim/go-skia/blob/main/demo/main.go
[SkPaint]: https://api.skia.org/classSkPaint.html
*/
package skia
