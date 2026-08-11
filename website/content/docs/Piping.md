This page explains the use of the `->`, `>>`, and `?>` operators.

## The piping operator `->`

The `->` operator takes a value, on the left, and insets it into a function, on the right. So for example `5 -> triangularNumber` is just the same as `triangularNumber 5`.

This is useful because it allows us to talk about operations on data from left to right. An expression like `"hElLo wOrlD" -> toLowerCase -> reverse -> capitalizeInitials -> reverse` is easier to read, and for you to evaluate as `"hellO worlD"`, than if it was written as `reverse capitalizeInitials reverse toLowerCase "hElLo wOrlD"`.

Instead of a function, we can write an expression which uses the word `that`: for example `3 -> that + 5 -> that * that` will evaluate to 64.

## The mapping operator `>>`

This takes a list on the left and a function on the right, and returns the result of applying the function to each element of the list. So `["fee" "fie", "fo", "fum"] >> len` would evaluate to `[3, 3, 2, 3]`.

Again, we can use the `that` keyword to write mapping expressions: `[1, 2, 3, 4, 5] >> that + 5 >> that * that` would evaluate to `[36, 49, 64, 81, 100]`.

## The filter operator `?>`

This takes a list on the left and a function returning a boolean value on the right, and returns those elements for which the application of the function to the element returns `true`. So for example `[1, 2, 3, 4, 5] ?> that % 2 == 1` will select the odd elements and return `[1, 3, 5]`.