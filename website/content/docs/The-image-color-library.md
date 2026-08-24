## Overview 

The `color` package implements types for representing colors in the usual color models.  

## Types 

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
## Functions 

### `swatch(c Color)`

Returns a string consisting of "█" with control codes so that it'll display the given color.  

