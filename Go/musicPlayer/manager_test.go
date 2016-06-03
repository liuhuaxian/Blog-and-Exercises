package library

import {
    "testing"
}

func TestOps(t *testing.T) {
    mm := NewMusicManager()

    if mm == nil {
        t.Error("NewMusicManager failed.")
    }

    if mm.len() != 0 {
        t.Error("NewMusicManager failed, not empty.")
    }

    m0 := &MusicEntry {
        "1", "My Heart Will Go On", "Celion Dion", "Pop",  "http://qbox.me", "MP3"
    }
    mm.Add(m0)

    if mm.len() != 1 {
        e.Error("MusicManager.Add() failed.")
    }

    m := mm.Find(m1.name)

    if m == nil {
        t.Error("MusicManager.Find() failed.")
    }

    if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name ||
    m.Genre !=  m0.Genre || m.Source != m0.Source || m.Type != m0.Type {
        t.Error("MusicManager.Find() failed. Found item mismatch.")
    }

    m, err : = mm.Get(0)

    if m == nil {
        t.Error("MusicManager.Get() failed.", err)
    }

    m = mm.Remove(0)
    if m == nil || mm.len() != 0 {
        t.Error("MusicManger.Remove() failed.", err)
    }
}
