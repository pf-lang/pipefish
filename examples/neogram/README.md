## Introduction

This is an experimental piece of software used to model sound changes in the evolution of languages. It is a work in progress, which doesn't yet support all the planned phonological features (e.g. aspirated and labialized consonants, a distinction between stressed and unstressed syllables, etc).

## Neogram by example

The best way to explain neogram is just to show you what it does. Let's take it for a spin!

### Phonology

The `phonemes` function returns a set of phonemes, e.g.

```
neogram → phonemes -- dental fricative
set(ð, θ)
neogram → phonemes -- front vowel
set(a, e, e̞, i, y, æ, ø, ø̞, œ, ɛ, ɪ, ɶ, ʏ)
neogram → phonemes -- open front vowel
set(a, ɶ)
neogram → phonemes -- rounded open front vowel
set(ɶ)
neogram →
```

And so on. Each word `open`, `close`, `dental`, etc names a set, and the whole phrase means their intersection. A list of all the built-in phoneme classes can be found at the end of this document.

We can use `/` or `or` to specify more complicated groups. Note the difference:

```
neogram → phonemes -- close back/central
set(u, ɨ, ɯ, ʉ)
neogram → phonemes -- close back or central
set(u, ä, ɐ, ɘ, ə, ɜ, ɞ, ɨ, ɯ, ɵ, ʉ)
neogram → 
```

The first of these is parsed as `close (back / central)` while the second is `(close back) / central`. You can also use parentheses to clarify your intent: as the old saying goes, they're free.

There is also an `&` operator, but you usually won't want it since e.g. `dental & fricative` means just the same as `dental fricative`.

You can also just name phonemes and get the set containing them:

```
neogram → phonemes -- s
set(s)
neogram → 
```

This on its own is not very useful, but you can combine it with other sets using the operators:

```
neogram → phonemes -- s or dental fricative
set(s, ð, θ)
neogram → 
```

Conversely, we can ask neogram to describe the features of a phoneme:

```
neogram → describe -- l
alveolar approximant lateral voiced consonant
neogram → describe -- ts
affricate alveolar sibilant unvoiced consonant
neogram → describe -- a
front open unrounded vowel
neogram → 
```

## Adding vocabulary

Now we need some words for neogram to work on. We can add vocabulary:

```
neogram → add --
  prila - walrus
  ŋɑli - insecticide
  zɑga - pineapple
          
OK
neogram → 
```

We can also merge vocabulary into the language from a file:

```
neogram → merge "morewords.txt"
OK
neogram →
```

And, of course, we can inspect it.

```
neogram → vocab
▪ beni - goatherd
▪ gluze - unicycle
▪ krupə - contretemps
▪ moge - hypotenuse
▪ prila - walrus
▪ saŋgə - moon
▪ zibi - acupuncture
▪ zɑdə - pancreas
▪ zɑga - pineapple
▪ ŋɑli - insecticide
▪ χozɑ - parasol

neogram → 
```

### Sound-changes

Now let's do our first, very simple, sound change.

```
neogram → shift -- initial bilabial plosive > fricative
OK
neogram → vocab
▪ gluze - unicycle
▪ krupə - contretemps
▪ moge - hypotenuse
▪ saŋgə - moon
▪ zibi - acupuncture
▪ zɑdə - pancreas
▪ zɑga - pineapple
▪ ŋɑli - insecticide
▪ ɸrila - walrus
▪ βeni - goatherd
▪ χozɑ - parasol

neogram →   
```

Notice what's happened. Neogram has *opinions* about phonology. Since we told it to turn bilabial plosives into fricatives, it knows which fricatives would make most sense for which bilablial plosives, based on their phonological characteristics.

We can use the `diff` command to see only the words that have just changed, and we'll usually do this from now on to save space.

```
neogram → diff
▪ ɸrila - walrus
▪ βeni - goatherd

neogram →  
```

That was a one-off sound change, a shift. The `fix` command imposes a permanent morphological rule that will be applied to words borrowed into the language or produced by other sound-changes. E.g:

```
neogram → fix -- sibilant, front V > ʃ, V
OK
neogram → diff
▪ gluʃe - unicycle
▪ ʃaŋgə - moon
▪ ʃibi - acupuncture

neogram →  
```

So now if we add a little vowel-harmony, we can see that doing so will not only cange `zɑga` to `zaga`, but to `ʃaga`, because the language no longer allows the sequence `za`.

```
neogram → shift -- V, C, front V > front V, C, V
OK
neogram → diff
▪ glyʃe - unicycle
▪ møge - hypotenuse
▪ ŋali - insecticide
▪ ʃaga - pineapple

neogram →  
```

We can refer to the phonemes by position, useful if we want to insert or delete phonemes. Let's made initial consonant clusters illegal by insisting that the consonants need a schwa betwen them:

```
neogram → fix -- initial C, C > C, ə, C.2
OK
neogram → diff
▪ gəlyʃe - unicycle
▪ kərupə - contretemps
▪ ɸərila - walrus

neogram → 
```

We didn't have to say that the first consonant on the right-hand side of the sound-change was C.1, though we could have, because that's where neogram expects the first consonant to go.

We can modify phonemes while addressing them by number.

```
neogram → shift -- V, plosive C, V > V, nasal C.2, C.2, V.3
OK
neogram → diff
▪ kərum̥pə - contretemps
▪ møŋge - hypotenuse
▪ zɑndə - pancreas
▪ ʃaŋga - pineapple
▪ ʃimbi - acupuncture

neogram →  
```

Note that the numbering is according to the position of the phnemes as phnemes, rather than numbering the vowels and consonants in separate series: the vowel-plosive-vowel pattern we're looking for is considered to consist of `V.1, C.2, V.3`, not `V.1, C.1, V.2`. 

Using `drop` will get rid of a phoneme:

```
neogram → shift -- voiced C, final a/ə > C, drop
OK
```

And now if we look back at our whole evolved word-list ...

```
neogram → vocab
▪ gəlyʃe - unicycle
▪ kərum̥pə - contretemps
▪ møŋge - hypotenuse
▪ zɑnd - pancreas
▪ ŋali - insecticide
▪ ɸəril - walrus
▪ ʃaŋg - pineapple, moon
▪ ʃimbi - acupuncture
▪ βeni - goatherd
▪ χozɑ - parasol

neogram →    
```

... we see among other things that the words for "pineapple" and "moon" have assimilated. If we want to look at the historical process by which all this happened, it was recorded:

```
neogram → history
0.   shift -- initial bilabial plosive > fricative
1.   fix -- sibilant, front V > ʃ, V
2.   shift -- V, C, front V > front V, C, V
3.   fix -- initial C, C > C, ə, C.2
4.   shift -- V, plosive C, V > V, nasal C.2, C.2, V.3
5.   shift -- voiced C, final a/ə > C, drop

neogram → 
```

## Retcon and borrow

Because we keep track of the history, we can retcon words into the language --- at any stage in its history, but by default into the protolanguage, before any rules were applied.

```
neogram → retcon --
            boti - wheelbarrow
            nuda - anteater
            zuli - telescope
          
OK
neogram → diff
▪ nynd - anteater
▪ ʃyli - telescope
▪ βøn̥ti - wheelbarrow

neogram → 
```

We can also combine the `merge` command with `retcon` to retcon a file of vocabulary into the language, etc. A full list of commands is found in the last section of this document below, and in the [README file of the repo]().

The `inventory` command gets the set of all the phonemes represented in the vocabulary.

```
neogram → inventory
set(a, b, d, e, g, i, k, l, m, m̥, n, n̥, o, p, r, t, u, y, z, ø, ŋ, ɑ, ə, ɸ, ʃ, β, χ)
neogram → 
```

The `borrow` command will take a list of words and their definitions and try to assimilate them into the language by (a) only using phonemes in the inventory, and (b) applying all the fixed rules. So if we try to borrow the words "zip" and "crook" from English ...

neogram → borrow --
            zɪp - zip
            krʊk - crook
          
OK
neogram → diff
▪ kərɑk - crook
▪ ʃip - zip

neogram →   

And of course there are various commands to save and open projects, delete words from the vocab, and other mundane housekeeping tasks. The full list of public commands and functions follows.

## Commands

### `add(sn snippet)`

Adds a list of words to the vocab list.  

### `borrow(sn snippet)`

Borrows words into the language, conforming them to the phoneme inventory.  

### `delete(sn snippet)`

Deletes words from the vocab list.  

### `diff`

Shows which entries in the word list have been changed by the most recent command.  

### `fix(sn snippet)`

Creates a permanent law of the phonology.  

### `history`

Shows the history of the sound changes.  

### `inventory`

Lists the phonemes currrently in the language.  

### `merge(filename string)`

Merges a file containing a list of words into the vocab.  

### `note(sn snippet)`

Adds a note to the history.  

### `nuke`

Deletes the project in-memory, though not on disk if you've saved it.  

### `open(filename string)`

Opens a saved project.  

### `retcon(sn snippet)`

Retcons a list of words into the vocab list at the start.  

### `retcon(i int, sn snippet)`

Retcons a list of words into the vocab list at a point just prior to event `i`.  

### `retcon merge (filename string)`

Retcons a file containing a list of words into the vocab at the start.  

### `retcon merge (i int, filename string)`

Retcons a file containing a list of words into the vocab at the start.  

### `save as (filename string)`

Saves the project to the given file.  

### `save`

### `shift(sn snippet)`

Applies a sound change.  

### `vocab`

Displays the vocab list.

## Functions 

### `describe(sn snippet) -> string`

Describes the phonological properties of a phoneme.  

### `phonemes(sn snippet)`

Returns the phonemes matching a given pattern. 

