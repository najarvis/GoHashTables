package chain_hash_table

import (
    "testing"
)

func TestSize(t *testing.T) {
    h := make(ChainHashTable, 223)
    if cap(h) != 223 {
        t.Error("Expected size of 223 in the HashTable")
    }
}


func TestContains(t *testing.T) {
    h := make(ChainHashTable, 223)
    h.Add(3)
    if !h.Contains(3) {
        t.Error("Expected 3 to be inside of h")
    }
}

func TestRemove(t *testing.T) {
    h := make(ChainHashTable, 223)
    for i := 0.0; i < 10; i++ {
        h.Add(i)
    }
    for i := 0.0; i < 10; i++ {
        if !h.Contains(i) {
            t.Errorf("Expected %.3f in h", i)
        }
        err := h.Remove(i)
        if err != nil {
            t.Error(err)
        }
        if h.Contains(i) {
            t.Errorf("Epected %.3f to NOT be in h", i)
        }
    }
}
