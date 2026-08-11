In this page we will discuss the predefined compound types, `pair`, `list`, `set`, `map`, and `tuple`.

There are also generic versions of all of these except `tuple`, as will be discussed [here](https://github.com/tim-hardcastle/Pipefish/wiki/Validation-and-parameterized-types#generics). But for many purposes you can get along fine without them. 

## `pair`

The `pair` type is formed using the `::` infix, e.g. `42::"foo"`, or `"walrus"::true`.

Unlike in some other functional languages, you're not meant to construct list-like structures out of the `pair` type. Rather it exists to represent things that naturally come in pairs: in particular key-value pairs in maps, field-value pairs in structs, and lower and upper limits of numeric ranges.

Besides equality/inequality, and the `::`, the only operation defined on a pair is indexing, the left element having index `0` and the right element having index `1`: e.g. `("walrus"::true)[0]` evaluates to `"walrus"`.

## `list`

The `list` type consists of an ordered, indexable, nestable collection of values. List literals consist of the elements of the list separated by commas and enclosed in square brackets, e.g. `["a", "b", "c", "d"]` or `[42, true, "bananas"]`, or, since lists can contain lists, `[1, [2, 3], 4]`. An empty list is denoted by `[]`.

A list, like a string, can be indexed and sliced using the same notation, and as with strings the `len` function will return its length. Also, the `in` operator will evaluate whether an element is in a list.

<img width="710" height="207" alt="image" src="https://github.com/user-attachments/assets/b3007854-7851-4667-9cf1-e054b0220298" />

Notice that when we slice a list or a string, what we are doing is indexing it by something of type `pair`.

Elements can be prepended or appended to a list using the `&` operator, paralleling the syntax used to prepend and append runes to strings.

## `set`

The `set` type is an unordered, unindexed collection of values: ` set (1, 2, 3)` or `set (42, "walrus")`. The empty set is denoted by `set ()`.

Since a set is unordered, it cannot be indexed or sliced. `len` and `in` work as for lists, as do `+` and `-`, which behave as you would expect: `set(1, 2, 3) - set(2, 4, 6)` returns `set(1, 3)`. `/\` is used as an intersection operator.

## `map`

A map is a collection of key-value pairs which can be indexed by the keys. Its literals are formed by `map` followed by a tuple of key-value pairs, e.g: `map "foo"::"bar", 42::"walrus"`.

Indexing is done as usual with square brackets: `(map "foo"::"bar", 42::"walrus")[42]` will return `walrus`.

An error will be returned if the key is not in the map.

The `len` function returns the number of keys in the map.

The `keys` function returns the keys as a list.

## `tuple`

The `tuple` type is an ordered, indexable, unnestable collection of values. It is represented by a comma-separated sequence of values, e.g. `1, 2, 3`, which may optionally be enclosed in parentheses: `(1, 2, 3)`. The empty tuple is represented by `()`; for clarity you may wish to write `tuple ()`.

By "unnestable", we mean that for example `((1, (2, 3)), ((4)))` is the same as `1, 2, 3, 4`. 

The `+` operator is not implemented for the `tuple` type, since the fact that tuples are unnestable means that they can just be joined together with commas.

The index and slice operators work for tuples just as for lists, as does `len`.

As with all languages that implement tuples, this type is the Cinderella of the type system. It is meant to silently, invisibly pass values to and from functions. If you find yourself doing anything else with it, you should ask yourself why and then use a list instead.