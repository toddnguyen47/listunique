package listunique

import mapset "github.com/deckarep/golang-set/v2"

func GetListOfUniqueItems[T comparable](inputList []T) []T {
	if inputList == nil || len(inputList) <= 0 {
		return inputList
	}

	set1 := mapset.NewSet[T]()
	uniqueList1 := make([]T, 0)
	for _, elem := range inputList {
		if set1.Contains(elem) {
			continue
		}
		set1.Add(elem)
		uniqueList1 = append(uniqueList1, elem)
	}
	return uniqueList1
}

type UniqueList[T comparable] struct {
	innerList []T
	innerSet  mapset.Set[T]
}

func NewUniqueList[T comparable]() UniqueList[T] {
	return UniqueList[T]{
		innerList: make([]T, 0),
		innerSet:  mapset.NewSet[T](),
	}
}

func (u2 *UniqueList[T]) Add(elem T) {
	if u2.innerSet.Contains(elem) {
		return
	}
	u2.innerSet.Add(elem)
	u2.innerList = append(u2.innerList, elem)
}

func (u2 *UniqueList[T]) AddAll(elems ...T) {
	for _, elem := range elems {
		u2.Add(elem)
	}
}

func (u2 *UniqueList[T]) GetList() []T {
	return u2.innerList
}

func (u2 *UniqueList[T]) GetSize() int {
	return len(u2.innerList)
}

// Update - will return false if the new element is already in the list or if index is out of bounds
func (u2 *UniqueList[T]) Update(index int, elem T) bool {
	if !isWithinBounds(index, u2.innerList) || u2.innerSet.Contains(elem) {
		return false
	}

	// Get old element and remove it from Set
	oldElem := u2.innerList[index]
	u2.innerSet.Remove(oldElem)

	// Add new element now
	u2.innerList[index] = elem
	u2.innerSet.Add(elem)

	return true
}

// Delete - unfortunately, `Delete` will be a linear operation and very slow. Use sparingly!
// Will return a bool to denote if element was found and removed or not
func (u2 *UniqueList[T]) Delete(elem T) bool {
	desiredIndex := -1
	for i, listElem := range u2.innerList {
		if listElem == elem {
			desiredIndex = i
			break
		}
	}
	if desiredIndex < 0 {
		return false
	}

	u2.innerList = append(u2.innerList[:desiredIndex], u2.innerList[desiredIndex+1:]...)
	u2.innerSet.Remove(elem)
	return true
}

func isWithinBounds[T comparable](index int, list1 []T) bool {
	return index >= 0 && index < len(list1)
}
