package Util

import "sync"

/**
滑动窗口
 */
type record struct {
	sucess int
	fail int
	timeout int
	requestNum int
}

type Sliding struct {
	link []record
	maxLength int
	length int
	head int
	end int
	lock sync.Mutex
}

func New(maxLength int) *Sliding  {
	return &Sliding{
		maxLength: maxLength,
		link: make([]record,maxLength),
		length: 0,
	}
}

func (s *Sliding) addSuccess()  {

}
func (s *Sliding) addFail()  {

}
func (s *Sliding) addTimeout()  {

}