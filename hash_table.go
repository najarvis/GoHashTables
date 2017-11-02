package main

import (
    "fmt"
)

type HashTable []*int;

func main() {
    size := 157 // Prime number
    h := make(HashTable, size);
    h.add(25)
    h.add(29)
    fmt.Println(h.contains(30))
    fmt.Println(h.contains(25))
    fmt.Println(h)
}

func (h HashTable) String() string {
    // ToString method
    s := "";
    s += "{"
    first := false
    for i, v := range(h) {
        if v != nil {
            if i != len(h) && first {
                s += ", "
            }
            s += fmt.Sprintf("%d", *v)
            if !first { first = true }
        }
    }
    s += "}"
    return s
}

func (h *HashTable) add(val int) {
    location_offset := 0
    location := h.hashFunc(val)
    data := (*h)[location]
    for data != nil {
        if *data == val {
            return
        }
        // Quadratic probing
        location_offset++
        location := location + location_offset * location_offset
        data = (*h)[location]

    }
    (*h)[location] = &val
}

func (h *HashTable) contains(val int) bool {
    location_offset := 0
    location := h.hashFunc(val)
    data := (*h)[location]
    for data != nil {
        if *data == val {
            return true
        }
        // Quadratic probing
        location_offset++
        location := location + location_offset * location_offset
        data = (*h)[location]

    }
    return false
}

func (h *HashTable) hashFunc(val int) int {
    var sum int
    curr := val
    for curr > 1 {
        sum += curr % 100
        curr /= 100
    }
    return sum % cap(*h)
}
