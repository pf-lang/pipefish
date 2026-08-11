Type aliases are constructed in the `newtype` section, as you would expect, and take the form `<typename> = alias <type to be aliased>, e.g:

```
newtype

ℤ = alias int
Strings = alias list{string}
```

These will alias both the types and their constructors, e.g. `ℤ("42")` means the same as `int("42")`; or `Strings["fee", "fie", "fo", "fum"]` means the same as `list{string}["fee", "fie", "fo", "fum"]`.