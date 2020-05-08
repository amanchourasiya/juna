package heap

import (
	"fmt"
)

type Comparable interface {
	GetKey() float32
	Compare(Comparable) int
	ToString() string
}

var (
	arr = make([]Comparable, 0)
	n   = 0
)

func swim(k int) {
	for {
		if k == 0 || less((k-1)/2, k) {
			break
		}
		swap(k, (k-1)/2)
		k = (k - 1) / 2
	}

}

func sink(k int) {
	j := 2*k + 1
	for {
		if j >= n {
			break
		}
		if j < n && less(j+1, j) {
			j++
		}
		if less(k, j) {
			break
		}
		swap(k, j)
		k = j
	}

}

func swap(i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func less(i, j int) bool {
	//log.Printf("i: %d j: %d\n", i, j)
	if arr[i].Compare(arr[j]) == -1 {
		return true
	}
	return false
}

func Insert(element Comparable) {
	arr = append(arr, element)
	//log.Printf("Insert calling swim %d\n", n)
	TraverseHeap()
	//log.Printf("Len of arr %d\n", len(arr))
	swim(n)
	n++
}

func Delete() Comparable {
	res := arr[0]
	swap(0, n-1)
	n--
	//arr = append(arr[0 : n-1])
	//arr[n] = nil
	sink(0)
	return res
}

func GetAllElements() []Comparable {
	return arr
}

func TraverseHeap() {
	for _, v := range arr {
		fmt.Printf("%s", v.ToString())
	}
}

func Sort(elements []Comparable) {

}
