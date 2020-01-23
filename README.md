# Diceware 

[![Go Report Card](https://goreportcard.com/badge/github.com/nouney/diceware)](https://goreportcard.com/report/github.com/nouney/diceware)
[![GoDoc](https://godoc.org/github.com/nouney/diceware?status.svg)](https://godoc.org/github.com/nouney/diceware)

Diceware is a small Golang package implementing the [Diceware method](https://en.wikipedia.org/wiki/Diceware).
It uses [word lists from the EFF](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases)

The code is based on this awesome [StackOverflow answer](https://stackoverflow.com/a/31832326/2432477) by [icza](https://stackoverflow.com/users/1705598/icza).

## Installation

To install `diceware`, run the following command:

 `$ go get github.com/nouney/diceware`

## Example

```golang
package main

import (
    "fmt"
    "log"

    "github.com/nouney/diceware"
)

func main() {
    // Pick a word from the large list
    word, err := diceware.Pick(diceware.LargeList)
    if err != nil {
        panic(err)
    }

    // Generate a passphrase
    passphrase, err := diceware.Passphrase(4, diceware.LargeList)
    if err != nil {
        panic(err)
    }
}

```
