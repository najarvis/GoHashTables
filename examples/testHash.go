package main

import (
    "fmt"
    "hash_table/hash_table"
    "hash_table/chain_hash_table"
)

func main() {
    ht := make(hash_table.HashTable, 100)
    cht := make(chain_hash_table.ChainHashTable, 10)
    for i := 0.0; i < 100; i++ {
        ht.Add(i)
        cht.Add(i)
    }
    fmt.Println("Regular Hash-Table")
    fmt.Println(ht)
    fmt.Println("\nNow for the Chain Hash-Table")
    fmt.Println(cht)
}
