/**
 * @Author: BookYao
 * @Description: Add IntSet operation method
	update1:增加练习6.1: 为bit数组实现下面这些方法
 	update2:增加练习6.2. IntersectWith, DifferenceWith, SymmetricDifference, 该版本有一些问题，后续再优化
    update3:增加练习6.4: 实现一个Elems方法，返回集合中的所有元素
 * @File:  intSet
 * @Version: 1.0.0
 * @Date: 2020/8/17 11:18
 */

package main

import (
	"bytes"
	"fmt"
)

const bit = 32 << (^uint(0) >> 63)
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

func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	var z IntSet
	for i, tword := range t.words {
		if i < len(s.words) {
			z.words = append(z.words, s.words[i]&tword)
		}
	}
	return &z
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var z IntSet
	for i, tword := range t.words {
		if i < len(s.words) {
			tmp := s.words[i] & tword
			z.words = append(z.words, s.words[i] - tmp)
		}
	}
	return &z
}

// SymmetricDifference sets s to the difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet{
	var z IntSet
	if len(t.words) == 0 {
		return s
	}
	len1, len2 := len(s.words), len(t.words)
	n := len1
	if len1 < len2 {
		n = len2
	}

	for i:=0; i<n; i++ {
		if i < len1 && i < len2 {
			tmp := s.words[i] & t.words[i]
			z.words = append(z.words, (s.words[i] - tmp) | (t.words[i] - tmp))
		} else if i < len1 {
			z.words = append(z.words, s.words[i])
		} else {
			z.words = append(z.words, t.words[i])
		}
	}
	return &z
}

func (s *IntSet) Elems () []uint {
	var z []uint

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				z = append(z, uint(bit*i+j))
			}
		}
	}

	return z
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(4)
	fmt.Println("x.string:", x.String())

	y.Add(1)
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

	x.AddAll(100, 200, 300)
	fmt.Println("AddAll x:", x.String())
	z := x.IntersectWith(&y)
	fmt.Println("x IntSet:", x.String())
	fmt.Println("y IntSet:", y.String())
	fmt.Println("x,y Intersect IntSet:", z.String())

	z = x.DifferenceWith(&y)
	fmt.Println("x IntSet:", x.String())
	fmt.Println("y IntSet:", y.String())
	fmt.Println("x,y DifferenceWith IntSet:", z.String())

	z = x.SymmetricDifference(&y)
	fmt.Println("x IntSet:", x.String())
	fmt.Println("y IntSet:", y.String())
	fmt.Println("x,y SymmetricDifference IntSet:", z.String())

	fmt.Println("x IntSet:", x.String())
	t := x.Elems()
	fmt.Println("Element: ", t)

	x.Clear()
	fmt.Println("Clear X:", x.String())
}

  