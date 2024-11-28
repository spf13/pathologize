![pathologize](https://github.com/user-attachments/assets/629a2fd0-a939-42d3-884f-489942aa783f)


# Pathologize

Pathologize is a Go library designed to fix what ails the paths in your application.  It ensures that the paths you work with are safe to use on all modern operating and file systems.
It is intentionally restrictive and will not allow names that are too long, contain invalid characters, or are reserved names on any modern operating system, even if not the host OS.

## Why Use Pathologize?

Modern applications work across multiple platforms and file systems, and it is essential to ensure that filenames are safe to use across different operating systems.  While individual operating systems enforce their own rules, relying on the host OS to validate filenames can lead to issues when sharing files across different platforms. Pathologize provides a simple way to clean and validate filenames, ensuring they are safe to use on all modern operating systems. It does not impose restrictions from legacy systems (e.g. DOS 8.3 filenames) unless they are still in use on modern systems. 

Filenames can be problematic due to various restrictions imposed by different operating systems and file systems. These restrictions include:

- Invalid characters that are not allowed in filenames.
- Reserved names that cannot be used as filenames.
- Length restrictions on filenames.

Pathologize addresses these issues by providing a simple way to clean and validate filenames, ensuring they are safe to use across different platforms.

See [Wikipedia article on reserved filenames](https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words).

## Installation

To install Pathologize, use `go get`:

```sh
go get github.com/spf13/pathologize
```

## Usage

To use Pathologize, import the `pathologize` package and call the `Clean` function with the filename you want to clean. The function will return a cleaned version of the filename that is safe to use on all modern operating systems. 

```go
package main

import (
    "fmt"
    "github.com/spf13/pathologize"
)

func main()
    filename := "CON.."
    cleanFilename := pathologize.Clean(filename)
    fmt.Println(cleanFilename) // Output: CON_

```

In this example, the filename `CON..` is cleaned using Pathologize, and the output is `CON_`. The cleaned filename is safe to use on all modern operating systems.

`CleanPath` will clean a full path by running `Clean` on each directory segment and the file name. This is useful when you need to clean a full path, including the directory structure.    

```go
package main

import (
    "fmt"
    "github.com/spf13/pathologize"
)

func main()
    path := "C:/Users/dir:e*c?t<o>r|y/CON.."
    cleanPath := pathologize.CleanPath(path)
    fmt.Println(cleanPath) // Output: C:/Users/directory/CON_

```

Lastly there's a helper function `IsClean` that will return true if the filename is already clean, and false if it is not.

```go
package main

import (
    "fmt"
    "github.com/spf13/pathologize"
)

func main()
    filename := "CON.."
    isClean := pathologize.IsClean(filename)
    fmt.Println(isClean) // Output: false

```

## License
This project is licensed under the Apache 2.0 License. See the LICENSE file for details.
