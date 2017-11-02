// Chain-block representation of a HashTable or Set.

package main

import (
    "fmt"
)

// If a number is already at the given hash location, instead
// of trying to find a new place to put it, we just chain it
// on to the chain already there.

type chain struct {
    val  int
    next *chain
}

// Store an array of 
type HashTable []*chain;

func main() {
    // A prime number is used as the length of the array for greater efficiency.
    // Increasing this number decreases collision frequency but increases
    // memory usage.
    size := 157
    h := make(HashTable, size);

    // Inserting a bunch of data.
    for i := 0; i < 1000; i++ {
        h.add(i)
    }

    //h.add(25)
    //h.add(29)
    fmt.Println(h.contains(1001)) // false
    fmt.Println(h.contains(25))   // true
    fmt.Println(h)
}

func (h HashTable) String() string {
    // this is basically the toString method

    s := "";
    s += "{"
    first := false
    for i, v := range(h) {
        if v != nil {
            if i != len(h) && first {
                s += "\n" // This could also be s += ", " to keep it on one line.
            }
            s += fmt.Sprintf("%d", (*v).val)
            if !first { first = true }
            if (*v).next != nil {
                data := (*v).next
                for ; data != nil; data = (*data).next {
                    // Use chains of arrows to denote data chains.
                    s += fmt.Sprintf(" -> %d", (*data).val)
                }
            }
        }
    }
    s += "}"
    return s
}

func (h *HashTable) add(val int) {
    // Adds a value to the HashTable

    // Perhaps if a lot of duplicate calls would be made, it would make
    // sense to first check this. This could be commented out if most
    // values would be unique.
    if h.contains(val) { return }

    location := h.hashFunc(val)
    data := (*h)[location]
    // If there isn't already a chain started at the hash location, start one.
    if data == nil {
        (*h)[location] = &chain{val, nil}
        return
    }

    // If there is one there, go to the end of the chain and place the value there.
    var last_chain *chain
    for data != nil {
        // Because this is a set, duplicate values are not stored.
        if (*data).val == val {
            return
        }
        last_chain = data
        data = data.next

    }
    (*last_chain).next = &chain{val, nil}
}


func (h *HashTable) contains(val int) bool {
    // checks the HashTable to see if a value is present.

    location := h.hashFunc(val)
    data := (*h)[location]
    for data != nil {
        if (*data).val == val {
            return true
        }
        data = (*data).next

    }
    return false
}

func (h *HashTable) hashFunc(val int) int {
    // This can be modified and as long as the values are fit to the length of
    // the HashTable array. This implementation sums up every two digits.

    var sum int
    curr := val
    for curr > 1 {
        sum += curr % 100
        curr /= 100
    }
    return sum % cap(*h)
}
