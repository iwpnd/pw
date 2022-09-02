# pw

pw is a minimal password generator for me lazy fk

## Installation
### cli

```
go install github.com/iwpnd/pw/cmd/pw@latest
```

```bash
➜ pw create --help
NAME:
   pw create - create a new password

USAGE:
   pw create [command options] [arguments...]

OPTIONS:
   --length value, -l value  length of password (default: 50)
   --no-lower, --nl          password contains NO lowercase characters (default: false)
   --no-numbers, --nn        password contains NO number characters (default: false)
   --no-special, --ns        password contains NO special characters (default: false)
   --no-upper, --nu          password contains NO uppercase characters (default: false)
```

### package
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

        // optional
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

inspired by [@briandowns](https://github.com/briandowns) pass that bubbled up in my timeline and I used as foundation to play with strings and go a little.

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/pw](https://github.com/iwpnd/pw)
