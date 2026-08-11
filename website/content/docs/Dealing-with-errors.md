#### `hub why(errNo int)`

The initial report of an error is concise. `hub why` provides a more detailed explanation of what has gone wrong, and may suggest a help topic, or how it relates to other errors that are reported at the same time.

#### `hub where(errNo int)`

This shows the line where the error occurred with the exact place where the error occurred underlined in red, occasionally useful when the line is a long one.

#### `hub trace`

This gives the course a runtime error took as it passed up the stack. Unlike `why` and `where` it doesn't require an error number, because there can only ever be one runtime error thrown at a time, unlike syntax errors.

#### `hub values`

When a runtime error is the result of a type mismatch, the initial error report will give the expected and actual types, but won't give the actual value or values passed because this might be a firehose of information. This allows you to see the values should you wish.

When such values are available, the initial error report will contain a note to that effect, and `hub values` will show the values involved. As with `hub trace`, there is no need to specify an error number because only one runtime error can happen at a time.

#### `hub errors`

Repeats the list of the last set of errors thrown, to save you from having to scroll up looking for it.

#### `hub help (topic string)`

Gives helpful information on the topic. (Although at present the help topics are insufficient in number and idiosyncratically chosen.)

#### `hub help`

Lists the help topics.