## `with`

We've already seen the use of the `with` operator to construct new struct values, e.g. `Person with name::"Joseph", age::22`, where `Person` is the name of a type and `name::"Joseph", age::22` is a tuple of pairs.

We use the same `with` operator and the same syntax to copy and modify existing values.

This is important because in Pipefish, all values are immutable. You cannot change the nth item of a list or struct, or add a key to a map, or change the value associated with an existing key. Instead, you can ask Pipefish for *a copy* of the list/map/struct with an element or elements changed. Let's demonstrate this in the TUI, first running a little script to declare some data for us to copy-and-modify.

<img width="710" height="557" alt="image" src="https://github.com/user-attachments/assets/0bd597f7-1fac-4403-9721-2b52665a1104" />

## `without`

The `without` operator for maps returns a copy of the map with the given key or keys removed. 

<img width="710" height="89" alt="image" src="https://github.com/user-attachments/assets/2305c22d-934a-40bb-8cf7-f58422f2b314" />

No error is returned if the key is not in the map: Pipefish consistently treats deletion as an idempotent operation.