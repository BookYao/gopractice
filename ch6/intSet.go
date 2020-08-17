/**
 * @Author: BookYao
 * @Description: Add IntSet operation method
	练习6.1: 为bit数组实现下面这些方法
 * @File:  intSet
 * @Version: 1.0.0
 * @Date: 2020/8/17 11:18
 */

package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && (s.words[word] & (1<<bit)) != 0
}

func (s *IntSet) UniconWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word & (1 << uint64(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Clear() {
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint64(j)) != 0 {
				s.words[i] ^= (1<<uint(j))
			}
		}
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= s.words[word] ^ (1 << bit)
}

func (s *IntSet) Copy() *IntSet {
	var result IntSet
	for _, word := range s.words {
		result.words = append(result.words, word)
	}
	return &result
}

func (s *IntSet) AddAll(element...int) {
	for _, val := range element {
		s.Add(val)
	}
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(4)
	fmt.Println("x.string:", x.String())

	y.Add(10)
	y.Add(122)
	fmt.Println("y.string:", y.String())

	fmt.Println("2 is exist ?", x.Has(2))
	x.UniconWith(&y)
	fmt.Println("x unicon y:", x.String())

	fmt.Println("x len:", x.Len())

	x.Remove(2)
	fmt.Println("2 is exist ?", x.Has(2))

	j := x.Copy()
	fmt.Println("Copy X:", j)
	x.Clear()
	fmt.Println("Clear X:", x.String())

	x.AddAll(100, 200, 300)
	fmt.Println("AddAll x:", x.String())
}

  