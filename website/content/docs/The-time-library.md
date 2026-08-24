## Overview 

The `time` library implements a `Time` type, functions to handle it, and a getter to get the current time from the system clock.  

## Modules 

### ` ("time")`
## Types 

### `Weekday = enum SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY`

An enum containing the days of the week.  

### `Clock = struct `

A data structure to pass to `get ... from` to get the current time.  

### `Time = struct (year int, month int, day int, hour int, minute int, second int, nanosecond int, location string)`

A data structure representing a time.  
## Commands 

### `get(x ref) from (c Clock)`

Populates the reference variable with a `Time` representing the current time.  
## Functions 

### `add(t Time, duration int) -> Time`

`add(t, duration)` returns the time `t + duration`, with the duration given in nanoseconds.  

### `addDate(t Time, years int, months int, days int) -> Time`

`addDate(t, years, months, days)` returns the time corresponding to adding the given number of years, months, and days to `t`.  

### `after(t Time, u Time) -> bool`

`after(t, u)` reports whether the time instant `t` is after `u`.  

### `before(t Time, u Time) -> bool`

`before(t, u)` reports whether the time instant `t` is before `u`.  

### `compare(t Time, u Time) -> int`

`compare(t, u)` compares two times. The result is: 


- `-1` if `t` is before `u`
- `0` if `t` is equal to `u`
- `+1` if `t` is after `u`
 

### `equal(t Time, u Time) -> bool`

`equal(t, u)` reports whether `t` and `u` represent the same time instant.  

### `format(t Time, layout string) -> string`

`format(t, layout)` returns a textual representation of `t` formatted according to `layout`.  

### `isDst(t Time) -> bool`

`isDst(t)` reports whether `t` is in daylight saving time.  

### `ISOWeek(t Time) -> int, int`

`ISOWeek(t)` returns the ISO 8601 year and week number in which `t` occurs.  

### `isZero(t Time) -> bool`

`isZero(t)` reports whether `t` represents the zero time instant.  

### `local(t Time) -> Time`

`local(t)` returns `t` with its location set to the local time zone.  

### `parse(layout string, value string)`

`parse(layout, value)` parses a formatted time string and returns a Time value.  

### `parseDuration(s string)`

`parseDuration(s)` parses a duration string.  

### `round(t Time, duration int) -> Time`

`round(t, duration)` returns the result of rounding `t` to the nearest multiple of `duration` since the zero time.  

### `sub(t Time, u Time) -> int`

`sub(t, u)` returns the duration `t - u`.  

### `timeIn(t Time, location string)`

`timeIn(t, location)` returns `t` represented in the given IANA time zone location.  

### `timeToUnix(t Time) -> int`

`timeToUnix(t)` returns the Unix time, the number of seconds elapsed since January 1, 1970 UTC.  

### `timeToUnixMicro(t Time) -> int`

`timeToUnixMicro(t)` returns the Unix time in microseconds.  

### `timeToUnixMilli(t Time) -> int`

`timeToUnixMilli(t)` returns the Unix time in milliseconds.  

### `timeToUnixNano(t Time) -> int`

`timeToUnixNano(t)` returns the Unix time in nanoseconds.  

### `truncate(t Time, duration int) -> Time`

`truncate(t, duration)` returns the result of rounding `t` down to a multiple of `duration` since the zero time.  

### `unixToTime(sec int, nsec int) -> Time`

`unixToTime(sec, usec)` returns the local Time corresponding to the given Unix time.  

### `unixMicroToTime(usec int) -> Time`

`unixMicroToTime(usec)` returns the local Time corresponding to the given Unix time expressed in microseconds.  

### `unixMilliToTime(msec int) -> Time`

`unixMilliToTime(msec)` returns the local Time corresponding to the given Unix time expressed in milliseconds.  

### `UTC(t Time) -> Time`

`UTC(t)` returns `t` with its location set to UTC.  

### `weekday(t Time) -> Weekday`

`weekday(t)` returns the day of the week specified by `t`.  

### `goWeekday(t Time) -> int`

`goWeekday(t)` is an internal helper returning the weekday as an integer.  

### `yearDay(t Time) -> int`

`yearDay(t)` returns the day of the year specified by `t`, in the range 1 through 365 or 366 inclusive.  

