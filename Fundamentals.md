Go is a compiled language, it is static.
It is more faster and secure. 
It was created by Google to overcome some issues faced with python/C++.

### Code syntax 
package main: what this file belongs to

Every Go file must start with a package line. It says which package (group of code) the file is part of.

main is a special name. It tells Go: “this is a program you can run, not a library.” A main package must also contain a func main(), and that is where the program starts.

If you wrote package mathutils instead, Go would treat the file as reusable code that other programs import. It couldn’t be run on its own.

import: what other code this file uses

import brings in packages that someone else wrote, so you can use their functions.

