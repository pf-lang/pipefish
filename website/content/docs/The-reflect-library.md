## Overview 

The `reflect` library allows you to find the properties of types, and so all the functions of the `reflect` library take as their arguments things of the built-in type `type`.  

## Functions 

### `contains(t type, subtype type) -> bool`

Tests whether the second type is a subtype of the first.  

### `elements(t type) -> list`

Returns a list of the elements of an enum type; returns an error if the type is not concrete or not an enum.  

### `isAbstract(t type) -> bool`

Tests whether the type is abstract, i.e. if it is the union of more than one (or of zero) concrete types.  

### `isBuiltin(t type) -> bool`

Tests whether the given type is built-in.  

### `isClone(t type) -> bool`

Tests whether the given type is a clone.  

### `isConcrete(t type) -> bool`

Tests whether the given type is concrete.  

### `isEmpty(t type) -> bool`

Tests whether the given type is `empty`, the abstract type formed from the union of no concrete types.  

### `isEnum(t type) -> bool`

Tests whether the given type is an enum.  

### `isStruct(t type) -> bool`

Tests whether the given type is a struct.  

### `fieldLabels(t type) -> list`

If the type is a struct, this returns a list of the labels of its fields, and an error otherwise.  

### `fieldTypes(t type) -> list`

If the type is a struct, this returns a list of the labels of its fields, and an error otherwise.  

### `name(t type) -> string`

Returns the name of the type as a string.  

### `operator(t type) -> string`

Returns the operator of the type; that is, its name if it isn't parameterized and the part before the parameters if it is.  

### `parameterTypes(t type) -> list`

Returns the types of the type parameters (or an empty list if not parameterized) for clone or struct types, otherwise returns an error.  

### `parameterValues(t type) -> list`

Returns the values of the type parameters (or an empty list if not parameterized) for clone or struct types, otherwise returns an error.  

### `parent(t type) -> type`

Returns the parent of the type if it's a clone, and an error otherwise.  

### `subtypes(t type) -> set`

Returns a set of all the concrete types that are subtypes of the given type.  

