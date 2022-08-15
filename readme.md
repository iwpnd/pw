# pw

pw is a minimal password generator for me lazy fk

## Installation

```
go get -u github.com/iwpnd/pw
```

## Usage

```go
package main

import (
  "fmt"

  "github.com/iwpnd/pw"
  )


func main() {
    // use all characters
    p := pw.NewPassword(50)
    fmt.Println("default password: ", p)

    cp := pw.NewPassword(
        50,
        // use numbers in password
        WithNumbers(),
        // use uppercase characters in password
        WithUpperCase(),
        // use lowercase characters in password
        WithLowerCase(),
        // use special characters in password
        WithSpecial(),
    )
    fmt.Println("custom password: ", p)
}
```

## License

MIT

## Acknowledgement

inspired by @briandowns pass that bubbled up in my timeline and I used as foundation to play with strings and go a little.

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/pw](https://github.com/iwpnd/pw)
