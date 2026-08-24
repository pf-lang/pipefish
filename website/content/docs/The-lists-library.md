## Overview 

The `lists` library contains some standard functions taking the `list` type, and usually its clones, as arguments.  

## Functions 

### `binarySearch(L clones{list}, x any?) -> int?`

Performs a binary search on a list `L` which is presumed sorted, for an element `x`, returning its index in `L` if it's found or `NULL` if it isn't.  

### `compact(L list) -> list`

Returns a compacted version of the list, i.e. one with consecutive identical elements removed.  

### `count(L clones{list}, x any?) -> int`

Returns the count of the number of elements in the list equal to `x`.  

### `count(L clones{list}, F func) -> int`

Returns the count of the number of elements in the list for which `F` returns `true`.  

### `find(L clones{list}, x any?) -> int?`

Performs a search on a list `L` for an element `x`, returning its index in `L` if it's found or `NULL` if it isn't.  

### `find(L clones{list}, F func) -> int?`

Performs a search on a list `L` for an element for which `F` returns `true`, returning its index in `L` if it's found or `NULL` if it isn't.  

### `findAll(L clones{list}, x any?) -> list`

Performs a search on a list `L` for all the elements in it equal to `x`, and returns a list of their indices in `L`.  

### `findAll(L clones{list}, F func) -> list`

Performs a search on a list `L` for all the elements in it for which `F` is true, and returns a list of their indices in `L`.  

### `fold(L clones{list}, F func)`

Folds the list by the function `F`.  

### `insert(L list, x any?) -> list`

Inserts an element `x` into sorted position in a list `L` presumed to be already sorted.  

### `isSorted(L clones{list}) -> bool`

Tests whether a given list is sorted in ascending order.  

### `max(L clones{list})`

Returns the maximum value in a list.  

### `max(L clones{list}, F func)`

Returns the element of the list for which `F` attains its maximum.  

### `merge(J list, K list) -> list`

Merges two lists assumed to be sorted into a sorted list.  

### `min(L clones{list})`

Returns the maximum value in a list.  

### `min(L clones{list}, F func)`

Returns the element of the list for which `F` attains its maximum.  

### `product(L clones{list})`

Returns the list folded by the `*` operator.  

### `repeat(n int, x ... any?) -> list`

Returns `n` copies of `x`.  

### `reverse(L list) -> list`

Returns a reversed copy of `L`.  

### `sort(L list) -> list`

Returns the list sorted by the `<` operator.  

### `sum(L clones{list})`

Returns the list folded by the `+` operator.  

### `transpose(ls ... list) -> list`

Returns a list in which element `i` is a list of the `i`th element of each list. Hence transposing two lists would be equivalent to a `zip` function.  

