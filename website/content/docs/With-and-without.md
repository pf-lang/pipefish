## `with`

We've already seen the use of the `with` operator to construct new struct values, e.g. `Person with name::"Joseph", age::22`, where `Person` is the name of a type and `name::"Joseph", age::22` is a tuple of pairs.

We use the same `with` operator and the same syntax to copy and modify existing values.

This is important because in Pipefish, all values are immutable. You cannot change the nth item of a list or struct, or add a key to a map, or change the value associated with an existing key. Instead, you can ask Pipefish for *a copy of* the list/map/struct with an element or elements changed. Let's demonstrate this in the TUI, first running a little script to declare some data for us to copy-and-modify. Here is the script, in the file `examples/docs/with.pf`.

```
newtype

Person = struct(name string, age int)

var

john = Person("John", 22)
myList = ["a", "b", "c", "d"]
myMap = map("a"::[0, 1], "b"::[2, 3])
```

Let's run it in the TUI.

```tui
→ hub run "examples/docs/with.pf"
Starting script "with.pf" as service "with". 
with → john with name::"Susan"
Person with (name::Susan, age::22)
with → // The value of John is unchanged.
with → john
Person with (name::John, age::22)
with → // We can change several elements at a time.
with → myList with 1::"w", 3::"x"
[a, w, c, x]
with → // We can change an element of an element.
with → myMap with ["b", 0]::99
map(a::[0, 1], b::[99, 3])
with → // And we can add a key-value pair to a map.
with → myMap with "c"::[4, 5]
map(a::[0, 1], b::[2, 3], c::[4, 5])
with →
```

## `without`

The `without` operator for maps returns a copy of the map with the given key or keys removed. 

```
with → myMap without "a"
map(b::[2, 3])
with → myMap without "a", "b"
map()
with → 
```

No error is returned if the key is not in the map: Pipefish consistently treats deletion as an idempotent operation.