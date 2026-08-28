You can wrap any Go type in a Pipefish type, declaring it like this, as in the `examples/docs/wrapper.pf` file.

<img width="710" height="97" alt="image" src="https://github.com/user-attachments/assets/094bddc8-47d3-4039-8b2a-57368987998b" />

This works just as well with imported types, as seen in the `math/big` standard library, which wraps `*big.Float`, `*big.Int` and `*big.Rat`.

However, merely declaring the wrapper type achieves little and leaves much of the heavy lifting to the implementer, so that this is something you should only do as the basis for a library exposing a nice sensible API for the type. The catch is that declaring the type doesn't tell Pipefish so much as how to construct it, or how to render it as a string or literal, or how to implement `==`, and so you have to do many or all of these things by hand, or your type won't actually do anything, including being instantiated. And, of course, you have to do them all in Go functions wrapped in Pipefish signatures, because that's where the basic operations on the type are defined.

Here's an example of this sort of boilerplate for declaring `Uint_32`. First we need a constructor and an overloading of the `string` function:

<img width="710" height="212" alt="image" src="https://github.com/user-attachments/assets/b4f5194e-ad80-4a85-bf2b-1f12914e29ac" />

But `==` and `literal` are not overloadable, because normally that would be a terrible idea. Instead, we implement functions `Equals` and `Literal` in Go *without* a Pipefish wrapper which in this special case are automagically used for the wrapper types. These functions should cater for every wrapper type in the source file. Since in this example we're only implementing one wrapper type, the switch statements aren't really necessary, but they indicate what you would do if there was more than one wrapper type in the file.

<img width="710" height="403" alt="image" src="https://github.com/user-attachments/assets/6bc2b9ea-5679-43c5-b0ab-a28370c5e2c7" />

For larger examples, see the `math/big` and `math/cmplx` standard libraries.