## Overview 

Package `bmp` supplies a .bmp encoder and decoder. The .bmp format is specified at http://www.digicamsoft.com/bmp/bmp.html. The `image` library supplies an `Image` type and some basic functions to handle it. The `color` package implements types for representing colors in the usual color models.  

## Modules 

### ` ("golang.org/x/image/bmp")`

### ` ("image/color")`

### ` ("image")`

### ` ("os")`

### ` ("fmt")`
## Types 

### `Image = clone list`

Contains a bitmap, a clone of `list` in which the elements are the columns of the bitmap, i.e. lists with elements of type `Color`.  

### `BmpFile = struct (name string)`

Contains the name of the file to be read or written to.  

### `Config = struct (colorModel type, w int, h int)`

The metadata for an `Image`.  

### `Cmyk = struct (c int, m int, y int, k int)`

Represents a fully opaque color in the CMYK color model, with 8 bit for each dimension.  

### `Gray = struct (g int)`

Represents a shade of gray in 8-bit gray scale.  

### `Rgba = struct (r int, g int, b int, a int)`

Represents a color in the RGBA color model, with 8 bits for each dimension.  

### `Rgb = struct (r int, g int, b int)`

Like `Rgba`, but every color is fully opaque.  

### `Rgba_64 = struct (r int, g int, b int, a int)`

Represents a color in the RGBA color model, with 16 bits for each dimension.  

### `Rgb_64 = struct (r int, g int, b int)`

Like `Rgba_64`, but every color is fully opaque.  

### `YCbCr = struct (y int, cb int, cr int)`

Representation of a fully opaque 24-bit Y'CbCr color, having 8 bits each for one luma and two chroma components.  

### `Color = abstract Cmyk/Gray/Rgb/Rgba/Rgb_64/Rgba_64/YCbCr`

The abstract type containing the various color types, `Rgb`, `Rgba`, `Cmyk`, etc.  
## Commands 

### `get(x ref) from (b BmpFile)`

Populates the reference variable with an `Image` loaded from the given file.  

### `post(I Image) to (b BmpFile)`

Saves the `Image` to the given file as a .bmp file.  
## Functions 

### `newImage(w int, h int, fillColor Color)`

Constructs a new image of the given dimensions and filled with the given color.  

### `width(I Image) -> int`

Returns the width of the image.  

### `height(I Image) -> int`

Returns the height of the image.  

### `(I Image) with (args ... pair) -> Image`

Overloads the standard `with` to perform validation.  

### `model(I Image) -> type`

Returns the color model of the image, i.e. a type in the abstract type `Color`.  

### `string(I Image)`

Returns the metadata of the image as a string.  

### `swatch(c Color)`

Returns a string consisting of "█" with control codes so that it'll display the given color.  

