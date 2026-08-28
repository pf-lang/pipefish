## Enums

A Pipefish enum is a very simple thing. The script `examples/docs/enums.pf` shows how to declare enums:

```
newtype

Suit = enum CLUBS, HEARTS, SPADES, DIAMONDS

Value = enum ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, 
          .. EIGHT, NINE, TEN, JACK, QUEEN, KING
```

These have only the most basic built-in functions, plus each enum type has a constructor which will make an element of the type out of an integer. We demonstrate this as usual in the TUI:

```tui
→ hub run "examples/docs/enums.pf"
Starting script "enums.pf" as service "enums". 
enums → HEARTS == DIAMONDS
false
enums → CLUBS in Suit
true
enums → type JACK
Value
enums → Suit(2)
SPADES
enums →    
```

## General note on type declaration

So far we have seen how to declare structs and enums. Let's look at what they have in common.

* They both appear under the `newtype` headword.
* The declarations are parallel in form: struct declarations begin with `TypeName = struct`; enum declarations begin with `TypeName = enum`.
* The types are conventionally named in `PascalCase`.
* Each struct or enum type declared automatically has a constructor with the same name as the type. This is an ordinary function and can be overloaded like other functions.

This pattern will be repeated with other type declarations.