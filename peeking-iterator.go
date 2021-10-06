/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
package main

// 284 https://leetcode-cn.com/problems/peeking-iterator/
// 迭代器
type PeekingIterator struct {
	iter  *Iterator
	cache []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	if len(this.cache) > 0 {
		return true
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if len(this.cache) > 0 {
		ans := this.cache[0]
		this.cache = this.cache[1:]
		return ans
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if len(this.cache) > 0 {
		return this.cache[0]
	}

	this.cache = append(this.cache, this.iter.next())
	return this.cache[0]
}
