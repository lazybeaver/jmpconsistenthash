Description
===========
An implementation of the Jump Consistent Hash algorithm:
http://arxiv.org/ftp/arxiv/papers/1406/1406.2294.pdf

Example
=======

    package main

    import (
      "fmt"
      "github.com/lazybeaver/jmpconsistenthash"
    )

    func main() {
      var key uint64 = 0xdeadbeef
      var shards uint64 = 5

      hash := jmpconsistenthash.Hash(key, shards)
      fmt.Println(hash)
    }
