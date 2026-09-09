## Introduction

This document explains the relationship of GitHub markdown to Pipefish markdown, and to the pages produced by the `pf-page` web component.

## Pipefish pages

In HTML generated from the markdown, the first-level heading `#` is reserved for the title of the article.

The second- and third-level headings, `##` and `###`, are used to generate the headings and subheadings in the page, and are automatically supplied with ancor links based on the text of the heading.

## Code blocks

The one advantage of a Pipefish book over a regualar wiki is that you can run Pipefish code in it wherever it appears within triple backticks. Any tag coming after the backticks will indicate something other than Pipefish. No other forms of highlighting are currently supported: at present this will just get you a plain code block. 

These will be added later, and for future compatibility, all tags just giving the name of the language should be in lower-case as though they were file extensions: `c`, `go`, `html`, `sql`, `ascii`. 

## Color tags

Pipefish markdown has tags for colored text in the form e.g. <R>red text</>. Yellow, green, cyan, blue, and purple are similarly supported. Color tags should *not* be used in the text of a Pipefish page, but only in error messages and docstrings in Pipefish code.

## Missing features

Apart from the color tags, Pipefish implements a strict *subset* of the GitHub markdown spec. In particular, it is confined to those features which can be rendered in the terminal, so features such as underlining and strikethrough are not supported. Besides that, there are some minor missing features which will be added in later versions of Pipefish.
