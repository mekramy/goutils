# Go Utils Library

This library provides a collection of utility functions for Go, covering various domains such as pointers, numbers, strings, web, and file operations.

## Installation

To install the library, use the following command:

```sh
go get github.com/mekramy/goutils
```

## Usage

### Pointer Utilities

#### `PointerOf`

Returns the pointer of a given value.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    value := 42
    ptr := goutils.PointerOf(value)
    fmt.Println(*ptr) // Output: 42
}
```

#### `SafeValue`

Returns the value of a pointer or an empty value if the pointer is nil.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    var ptr *int
    value := goutils.SafeValue(ptr)
    fmt.Println(value) // Output: 0
}
```

#### `ValueOf`

Returns the value of a pointer or a fallback value if the pointer is nil.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    var ptr *int
    value := goutils.ValueOf(ptr, 10)
    fmt.Println(value) // Output: 10
}
```

#### `Alter`

Returns the value of a pointer or a fallback value if the pointer is nil or zero.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    var ptr int
    value := goutils.Alter(ptr, 10)
    fmt.Println(value) // Output: 10
}
```

#### `NullableOf`

Returns nil if the value is zero.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    value := 0
    ptr := goutils.NullableOf(&value)
    fmt.Println(ptr) // Output: <nil>
}
```

#### `IsEmpty`

Checks if a pointer is nil or zero.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    var ptr *int
    isEmpty := goutils.IsEmpty(ptr)
    fmt.Println(isEmpty) // Output: true
}
```

#### `IsSame`

Checks if two pointer values are equal.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    a := 42
    b := 42
    isSame := goutils.IsSame(&a, &b)
    fmt.Println(isSame) // Output: true
}
```

### Number Utilities

#### `Abs`

Returns the absolute value of a number.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    value := -42
    absValue := goutils.Abs(value)
    fmt.Println(absValue) // Output: 42
}
```

#### `RoundUp`

Returns the nearest larger integer (ceil).

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    value := 42.3
    roundedValue := goutils.RoundUp[int](value)
    fmt.Println(roundedValue) // Output: 43
}
```

#### `Round`

Returns the nearest integer, rounding half away from zero.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    value := 42.5
    roundedValue := goutils.Round[int](value)
    fmt.Println(roundedValue) // Output: 43
}
```

#### `RoundDown`

Returns the nearest smaller integer.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    value := 42.9
    roundedValue := goutils.RoundDown[int](value)
    fmt.Println(roundedValue) // Output: 42
}
```

#### `Min`

Returns the smallest value among the given numbers or zero.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    values := []int{42, 10, 56}
    minValue := goutils.Min(values...)
    fmt.Println(minValue) // Output: 10
}
```

#### `Max`

Returns the largest value among the given numbers or zero.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    values := []int{42, 10, 56}
    maxValue := goutils.Max(values...)
    fmt.Println(maxValue) // Output: 56
}
```

### String Utilities

#### `ExtractNumbers`

Extracts numbers from a string.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    str := "abc123def"
    numbers := goutils.ExtractNumbers(str)
    fmt.Println(numbers) // Output: 123
}
```

#### `ExtractAlphaNum`

Extracts alphanumeric characters from a string.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    str := "abc123!@#"
    alphaNum := goutils.ExtractAlphaNum(str)
    fmt.Println(alphaNum) // Output: abc123
}
```

#### `RandomString`

Returns a random string from a character set.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    randomStr := goutils.RandomString(10, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    fmt.Println(randomStr) // Output: Random 10 character string
}
```

#### `RandomNumeric`

Returns a random numeric string.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    randomNum := goutils.RandomNumeric(10)
    fmt.Println(randomNum) // Output: Random 10 digit number
}
```

#### `RandomAlphaNum`

Returns a random alphanumeric string.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    randomAlphaNum := goutils.RandomAlphaNum(10)
    fmt.Println(randomAlphaNum) // Output: Random 10 character alphanumeric string
}
```

#### `Slugify`

Makes a URL-friendly slug from strings.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    slug := goutils.Slugify("Hello World!")
    fmt.Println(slug) // Output: hello-world
}
```

#### `Concat`

Returns concatenated non-empty strings with a separator.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    concatenated := goutils.Concat("-", "Hello", "", "World")
    fmt.Println(concatenated) // Output: Hello-World
}
```

#### `FormatNumber`

Formats a number with a comma separator.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    formatted := goutils.FormatNumber("%d Dollars", 100000)
    fmt.Println(formatted) // Output: 100,000 Dollars
}
```

#### `FormatRx`

Formats a string using a regex pattern.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    formatted, _ := goutils.FormatRx("123456", `(\d{3})(\d{2})(\d{1})`, "($1) $2-$3")
    fmt.Println(formatted) // Output: (123) 45-6
}
```

### Web Utilities

#### `RelativeURL`

Returns the relative URL path of a file with respect to the root directory.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    relativeURL := goutils.RelativeURL("/root", "/root/path/to/file")
    fmt.Println(relativeURL) // Output: path/to/file
}
```

#### `AbsoluteURL`

Returns the absolute URL path of a file with respect to the root directory.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    absoluteURL := goutils.AbsoluteURL("/root", "/root/path/to/file")
    fmt.Println(absoluteURL) // Output: /path/to/file
}
```

#### `SanitizeRaw`

Sanitizes input to raw text.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    sanitized := goutils.SanitizeRaw("<script>alert('xss')</script> and more", true)
    fmt.Println(sanitized) // Output: and more
}
```

#### `SanitizeCommon`

Sanitizes input to HTML with common allowed tags.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    sanitized := goutils.SanitizeCommon("<b>bold</b><script>alert('xss')</script>", true)
    fmt.Println(sanitized) // Output: <b>bold</b>
}
```

### File Utilities

#### `NormalizePath`

Joins and normalizes a file path.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    path := goutils.NormalizePath("/root", "path//to\\some", "file")
    fmt.Println(path) // Output: /root/path/to/some/file
}
```

#### `CreateDirectory`

Creates a nested directory.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    err := goutils.CreateDirectory("/path/to/dir")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Directory created")
    }
}
```

#### `IsDirectory`

Checks if a path is a directory.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    isDir, err := goutils.IsDirectory("/path/to/dir")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(isDir) // Output: true or false
    }
}
```

#### `GetSubDirectory`

Returns a list of subdirectories.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    subDirs, err := goutils.GetSubDirectory("/path/to/dir")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(subDirs) // Output: [subdir1 subdir2 ...]
    }
}
```

#### `ClearDirectory`

Deletes all files and subdirectories in a directory.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    err := goutils.ClearDirectory("/path/to/dir")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Directory cleared")
    }
}
```

#### `FileExists`

Checks if a file exists.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    exists, err := goutils.FileExists("/path/to/file")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(exists) // Output: true or false
    }
}
```

#### `FindFile`

Searches a directory for a file with a pattern and returns the first file.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    file := goutils.FindFile("/path/to/dir", `admin[\d]*\.txt`)
    if file != nil {
        fmt.Println(*file) // Output: /path/to/dir/admin32.txt
    } else {
        fmt.Println("File not found")
    }
}
```

#### `FindFiles`

Searches a directory for files with a pattern.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    files := goutils.FindFiles("/path/to/dir", `user.*\.txt`)
    fmt.Println(files) // Output: [/path/to/dir/user23.txt /path/to/dir/userjohn.txt ...]
}
```

#### `GetMime`

Returns file MIME info from content.

```go
package main

import (
    "fmt"
    "goutils"
    "os"
)

func main() {
    data, _ := os.ReadFile("/path/to/file.txt")
    mime := goutils.GetMime(data)
    fmt.Println(mime.String()) // Output: text/plain; charset=utf-8
}
```

#### `GetExtension`

Returns the file extension.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    ext := goutils.GetExtension("/path/to/file.txt")
    fmt.Println(ext) // Output: txt
}
```

#### `GetFilename`

Returns the file name without extension.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    filename := goutils.GetFilename("/path/to/file.txt")
    fmt.Println(filename) // Output: file
}
```

#### `TimestampedFile`

Returns a file name with a timestamp prefix.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    timestamped := goutils.TimestampedFile("file.txt")
    fmt.Println(timestamped) // Output: file-1633024800000.txt
}
```

#### `NumberedFile`

Generates a unique numbered file name.

```go
package main

import (
    "fmt"
    "goutils"
)

func main() {
    numbered, err := goutils.NumberedFile("/path/to/dir", "file.txt")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(numbered) // Output: file-1.txt
    }
}
```

## License

This project is licensed under the MIT License.
