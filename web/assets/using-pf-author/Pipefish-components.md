**Pipefish components** are components which contain Pipefish code and/or a Pipefish service.

As JavaScript components with attributes and methods they are [described in the main Pipefish documentation](****). Here they are described only in their role as components in the `pf-book` framework, where most of their JavaScript API is hidden from you.

## pf-service

This is a non-visual component which can be used to power widgets in the page when you don't need to be able to see either the code or the tui.

## pf-tui

A tui backed by and giving access to a Pipefish service.

## pf-text

Displays Pipefish code with the appropriate highlighting.

## pf-reader

Displays files in tabs, using `pf-text` to display highlighted Pipefish code, and also capable of displaying raw or rendered `.md` files, displaying `.bmp`, `.jpg`, and `.png` files.

## pf-editor

This adds to the features of `pf-reader` by letting you create, edit, and delete files in the [virtual file system](*********).

## pf-coder

This contains an editor and a service. It is not particularly useful on its own, but exists so that we can make the next component.

## pf-ide

This consists of a pf-coder, and a pf-tui which has been told that the pf-coder is its service. Hence you can livecode in the coder and evaluate expressions in the tui in the context of whatever sevice is defined in the coder.

