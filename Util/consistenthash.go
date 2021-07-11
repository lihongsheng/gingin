package Util

import (
	"hash/crc32"
)

/**
一致性hash函数
 */

const multiple = 3
const hashLen = 1<<32 - 1

type HashFun func(str string) int

type ConsistendHash interface {
	Add(string2 string)
	Get(string2 string)
}

func HashCrc(str string) int  {
	return int(crc32.ChecksumIEEE([]byte(str)))
}

type hashMap struct {
	Hash HashFun
	HashValue []int
	Mapping map[int]string
}

func NewHashMap(f HashFun) *hashMap {
	h := &hashMap{
		Hash: f,
		HashValue: make([]int,0,100),
		Mapping: map[int]string{},
	}
	return h
}

//添加节点
func (h *hashMap) Add(str string)  {
	key := h.Hash(str)
	h.addSlice(key)
	h.Mapping[key] = str
}
//加入和排序节点的可以
func (h *hashMap) addSlice(key int)  {
	//l := len(h.HashValue)
	if len(h.HashValue) == 0 {
		h.HashValue = append(h.HashValue,key)
	} else {
		tmp := key
		for i, v := range h.HashValue{
			if tmp < v {
				t := h.HashValue[i]
				h.HashValue[i] = tmp
				tmp = t
			}
		}
		h.HashValue = append(h.HashValue,tmp)
	}
}
//获取可以所属的节点
func (h *hashMap) Get(str string) string  {
	k :=h.Hash(str)
	key := h.HashValue[0]
	for i, v := range h.HashValue {
		if (k < v && (i-1) > 0) {
			key = h.HashValue[i-1]
			break
		}
	}
	return h.Mapping[key]
}

func (h *hashMap) GetMap() map[int]string {
	return h.Mapping
}