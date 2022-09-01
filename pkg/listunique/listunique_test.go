package listunique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenDuplicateItemsList_When_ThenItemsReturnedAreUnique(t *testing.T) {
	list1 := []int{42, 42, 5, 5, 5, 1, 26, 26}

	newList := GetListOfUniqueItems(list1)

	expectedList := []int{42, 5, 1, 26}
	assert.Equal(t, expectedList, newList)
}

func Test_GivenNoDuplicateItemsList_When_ThenItemsReturnedAreUnique(t *testing.T) {
	list1 := []int{42, 5, 1, 26}

	newList := GetListOfUniqueItems(list1)

	expectedList := []int{42, 5, 1, 26}
	assert.Equal(t, expectedList, newList)
}

func Test_GivenNilList_When_ThenItemsReturnedAreEmpty(t *testing.T) {
	var list1 []int

	newList := GetListOfUniqueItems(list1)

	var expectedList []int
	assert.Equal(t, expectedList, newList)
}

func Test_GivenEmptyList_When_ThenItemsReturnedAreEmpty(t *testing.T) {
	list1 := make([]int, 0)

	newList := GetListOfUniqueItems(list1)

	expectedList := make([]int, 0)
	assert.Equal(t, expectedList, newList)
}

func Test_GivenNewUniqueList_WhenAddOne_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.Add(1)
	uniqueList.Add(5)
	uniqueList.Add(1)
	uniqueList.Add(2)

	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
}

func Test_GivenNewUniqueList_WhenAddAll_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
}

func Test_GivenNewUniqueList_WhenAddAllList_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll([]int{1, 5, 1, 2}...)

	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
}

func Test_GivenNewUniqueList_WhenDelete_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	wasDeleted := uniqueList.Delete(5)
	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 2}, list1)
	assert.True(t, wasDeleted)
}

func Test_GivenNewUniqueListNoItemFound_WhenDelete_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	wasDeleted := uniqueList.Delete(42)
	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
	assert.False(t, wasDeleted)
}

func Test_GivenNewUniqueList_WhenUpdating_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	wasUpdated := uniqueList.Update(2, 42)
	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 42}, list1)
	assert.True(t, wasUpdated)
}

func Test_GivenNewUniqueListUpdateDuplicate_WhenUpdating_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	wasUpdated := uniqueList.Update(1, 2)
	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
	assert.False(t, wasUpdated)
}

func Test_GivenNewUniqueListOutOfBoundsPositive_WhenUpdating_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	wasUpdated := uniqueList.Update(uniqueList.GetSize(), 42)
	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
	assert.False(t, wasUpdated)
}

func Test_GivenNewUniqueListOutOfBoundsNegative_WhenUpdating_ThenMakeSureItemsAreUnique(t *testing.T) {
	uniqueList := NewUniqueList[int]()
	uniqueList.AddAll(1, 5, 1, 2)

	wasUpdated := uniqueList.Update(-1, 42)
	list1 := uniqueList.GetList()

	assert.Equal(t, []int{1, 5, 2}, list1)
	assert.False(t, wasUpdated)
}
