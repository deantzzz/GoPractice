package geometry

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	fmt.Printf("------ x == %d-------\n", x)
	// 每一个字都有64个二进制位，所以为了定位x的bit位，我们用了x/64的商作为字的下标
	// 确定x属于哪个bucket: 1/64 == 0 表示在第0个桶中
	word := x / 64
	// 用x%64得到的值作为这个字内的bit的所在位置index
	// 确定x在bucket的第几位: 1%64 == 1 表示在第1位 [...0000010]
	// word和index一起确定了x的最终位置：第0个桶的第2个位置，bucket[0]:byte[...000010]
	index := uint(x % 64)
	bool1 := word < len(s.words)
	bt := 1 << index
	fmt.Println("bt ", bt)
	// s.words[word] == ...00000010 ，比较: 以1为例, s.words[word]，word是bucket，表示取出第0个桶，桶里目前只有一个数1
	// 1<<bit		 ==			 10	，把1左移index个位置,其它位都是0
	// ’&‘ 与运算	 == ...00000010  == 1 ，确定桶0的位置1上的值为1，表示1这个数在集合中
	andResult := s.words[word] & 1 << index
	bool2 := andResult != 0
	fmt.Printf("has : bucket: %d, index: %b\n", word, andResult)
	return bool1 && bool2
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	fmt.Printf("%s\n", s.words[word])
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			//bit位的“或”逻辑操作符号|来一次完成64个元素的或计算
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
			if word&(1<<uint(j)) != 0 {
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
	return len(s.words)
}

func (s *IntSet) Remove(x int) {

	s.words[x/64] = 0
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	n := IntSet{}
	n.words = make([]uint64, 0, len(s.words)+1)
	for i := range s.words {
		if s.words[i] != 0 {
			n.words = append(n.words, s.words[i])
		}
	}
	return &n
}

func (s *IntSet) AddAll(vals ...int) {
	for i := range vals {
		s.Add(vals[i])
	}
	fmt.Println(len(s.words))
}
