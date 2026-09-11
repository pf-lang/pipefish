## Introduction

As the `pf-service` component is meant to work with Pipefish-agnostic components, there is no need to produce a set of Pipefish-compatible buttons and widgets: they are all Pipefish-compatible.

We do however need components which can work with Pipefish code as text, correctly highlighting it and dealing with indents and outdents, and so on.

The components fit together in a nestable structure represented by the following diagram, where the `pf-` prefixes of the components are omitted throughout.

```box
                                          |-->service
                               |-->coder--|
                               |     ^    |-->editor-->reader-->highlighter
                      |-->ide--|     |
                      |        |     |
author-->book-->page--|        |--->tui--highlighter       
                      |
                      |-->image            
```

- The `pf-service` you have already met. Let's take the others more or less from the inside out.

- The `pf-highlighter` is a non-visual component with one method, `highlight`, which takes Pipefish code and returns it as highlighted HTML.

- The `pf-reader` is a box with one method, `display`, which takes Pipefish code and displays it highlighted in the box using the highlighter.

- The `pf-editor` allows you to edit the box, and has a `display` method wrapping around the `pf.reader`'s display method.

- The `pf-coder` has a `pf-editor` and a `pf-service`. It has `compile` and `do` methods which wrap around the methods of the `pf-service`, and a `display` method wrapping around the `pf-editor`'s `display` method. This is for when you want to be able to edit the code and compile and execute it but you don't need a TUI for anything.

- The `pf-tui` has the same `pf-highlighter` as the `pf-reader`. It also has a `pf-service` or something which, like the `pf-coder`, looks like one by having the same methods. The point of this is that if you just hook it up to a `pf-service`, then you get the `pf-tui` backed by whatever code the `pf-service` is running; but if you hook it up to the coder then it's backed by the `pf-coder` and the `pf-service` inside it.

- The `pf-ide` contains a `pf-coder` and a `pf-tui`. It has `compile`, `do`, and `display` methods which wrap around the `pf-coder`'s equivalent methods.

- A `pf-image` is a simple (indeed, Pipefish-agnostic) component which displays an image loaded from a resource file and updates itself whenever the file is changed. It has a single method, It has one method, `create`, giving the location of the image file. This is the exception to the rule that we don't need to supply widgets: the `pf-page` needs something to compile markdown images to.

- A `pf-page` is one page rendered from a `.md` file, with fenced code compiled to Pipefish. The `pf-page` will contain `pf-ide`s, or `pf-reader`s, or whatever, depending on information in the fence tags of the fenced code, and `pf-images` if the markdown has images. It has one method, `create`, giving the location of the `.md` file.

- The `pf-book` generates `pf-pages` from a flat folder of `.md` files.

- The `pf-author` component allows you to edit the markdown file of a `pf-book` in your browser and see the HTML update in real time and use the code in the `pf-ide`s, etc.





```box
                |-->service
     |-->coder--|
     |     ^    |-->editor-->reader-->highlighter
ide--|     |
     |     |
     |--->tui--highlighter                 
```

- The `pf-service` you have already met. Let's take the others more or less from the inside out.

- The `pf-highlighter` is a non-visual component with one method, `highlight`, which takes Pipefish code and returns it as highlighted HTML.

- The `pf-reader` is a box with one method, `display`, which takes Pipefish code and displays it highlighted in the box using the highlighter.

- The `pf-editor` allows you to edit the box, and has a `display` method wrapping around the `pf.reader`'s display method.

- The `pf-coder` has a `pf-editor` and a `pf-service`. It has `compile` and `do` methods which wrap around the methods of the `pf-service`, and a `display` method wrapping around the `pf-editor`'s `display` method. This is for when you want to be able to edit the code and compile and execute it but you don't need a TUI for anything.

- The `pf-tui` has the same `pf-highlighter` as the `pf-reader`. It also has a `pf-service` or something which, like the `pf-coder`, looks like one by having the same methods. The point of this is that if you just hook it up to a `pf-service`, then you get the `pf-tui` backed by whatever code the `pf-service` is running; but if you hook it up to the coder then it's backed by the `pf-coder` and the `pf-service` inside it.

- The `pf-ide` contains a `pf-coder` and a `pf-tui`. It has `compile`, `do`, and `display` methods which wrap around the `pf-coders` equivalent methods.