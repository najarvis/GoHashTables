package hash_table

import (
    "fmt"
)

type HashTable []*float64;

/*
func main() {
    size := 157 // Prime number
    h := make(HashTable, size);
    h.add(25)
    h.add(29)
    fmt.Println(h.contains(30))
    fmt.Println(h.contains(25))
    h.remove(25)
    fmt.Println(h.contains(25))
    fmt.Println(h)
}*/

func (h HashTable) String() string {
    // ToString method
    s := "";
    s += "{"
    first := false
    for _, v := range(h) {
        if v != nil {
            if first {
                s += ", "
            }
            s += fmt.Sprintf("%.3f", *v)
            if !first { first = true }
        }
    }
    s += "}"
    return s
}

func (h *HashTable) Add(val float64) {
    location_offset := 0
    location := h.HashFunc(val)
    for (*h)[location] != nil {
        if *((*h)[location]) == val {
            return
        }
        // Quadratic probing
        location_offset += 1
        location = location + location_offset * location_offset

    }
    (*h)[location] = &val
}

func (h *HashTable) Remove(val float64) error {
    if !h.Contains(val) {
        return fmt.Errorf("Value %.3f not in HashTable", val)
    }

    location_offset := 0
    location := h.HashFunc(val)

    for (*h)[location] != nil {
        if *((*h)[location]) == val {
            (*h)[location] = nil
        }
        location_offset += 1
        location = location + location_offset * location_offset
    }
    return nil
}

func (h *HashTable) Contains(val float64) bool {
    location_offset := 0
    location := h.HashFunc(val)
    for (*h)[location] != nil {
        if *((*h)[location]) == val {
            return true
        }
        // Quadratic probing
        location_offset += 1
        location = location + location_offset * location_offset

    }
    return false
}

func (h *HashTable) HashFunc(val float64) int {
    var sum float64
    curr := int(val)
    for curr >= 1 {
        sum += float64(curr)
        curr /= 100
    }
    s := 1.0
    if val < 0 { s = -1.0 }
    return int(((sum * s) + 1) * 96821) % cap(*h)
}
