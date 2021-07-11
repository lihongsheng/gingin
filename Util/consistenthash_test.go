package Util

import (
	"testing"
)

func TestNewHashMap(t *testing.T) {
	h := NewHashMap(HashCrc)
	ipList := []string{"192.168.0.1","192.168.0.2","192.168.0.3","192.168.0.4","192.168.0.1#1","192.168.0.2#2","192.168.0.3#3","192.168.0.4#4"}
	for _, v := range ipList{
		h.Add(v)
	}
	 k1 := h.Get("10.10.10.1")
	 k2 := h.Get("10.10.10.1")
	 if k1 != k2 {
	 	t.Error("fail hash map")
	 }
}
