// Chain-block representation of a HashTable or Set.

package chain_hash_table

import (
    "fmt"
)

// If a number is already at the given hash location, instead
// of trying to find a new place to put it, we just chain it
// on to the chain already there.

type chain struct {
    val  float64
    next *chain
}

// Store an array of 
type ChainHashTable []*chain;

func (h ChainHashTable) String() string {
    // this is basically the toString method

    s := "";
    s += "{"
    first := false
    for i, v := range(h) {
        if v != nil {
            if i != len(h) && first {
                s += "\n" // This could also be s += ", " to keep it on one line.
            }
            s += fmt.Sprintf("%.3f", (*v).val)
            if !first { first = true }
            if (*v).next != nil {
                data := (*v).next
                for ; data != nil; data = (*data).next {
                    // Use chains of arrows to denote data chains.
                    s += fmt.Sprintf(" -> %.3f", (*data).val)
                }
            }
        }
    }
    s += "}"
    return s
}

func (h *ChainHashTable) Add(val float64) {
    // Adds a value to the HashTable

    // Perhaps if a lot of duplicate calls would be made, it would make
    // sense to first check this. This could be commented out if most
    // values would be unique.
    if h.Contains(val) { return }

    location := h.HashFunc(val)
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

func (h *ChainHashTable) Remove(val float64) error {
    // Removes a value from the set

    if (!h.Contains(val)) {
        // If it isn't in there, raise an error
        err := fmt.Errorf("Set: Value %f not in set!", val)
        return err
    }

    // First check to see if it is in the first position.
    location := h.HashFunc(val)
    data := (*h)[location]
    if (*data).val == val {
        // If it is, change the actual HashTable slice index.
        (*h)[location] = (*h)[location].next
        return nil
    }

    // If it isn't, go through the chain at that location
    var last *chain
    for (*data).val != val {
        last = data
        data = (*data).next
    }

    // Once it is found, change the pointer that was going to val to
    // the pointer that val is pointing to.
    // a -> b -> c becomes a -> c (assuming we are looking for b)

    last.next = data.next
    return nil
}

func (h *ChainHashTable) Contains(val float64) bool {
    // checks the HashTable to see if a value is present.

    location := h.HashFunc(val)
    data := (*h)[location]
    for data != nil {
        if (*data).val == val {
            return true
        }
        data = (*data).next

    }
    return false
}

func (h *ChainHashTable) HashFunc(val float64) int {
    // This can be modified and as long as the values are fit to the length of
    // the HashTable array. This implementation sums up every two digits.

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
