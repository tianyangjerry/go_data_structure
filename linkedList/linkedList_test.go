package linkedList_test

import (
	"math/rand"
	"testing"
	"time"
)

const testTime = 100

func TestLinkedList(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < testTime; i++ {
		targetArr := make([]int, 100)
		arr := make([]int, 100)

		for j := 0; j < 100; j++ {
			targetArr[j] = rand.Intn(1000)
			arr[j] = targetArr[j]
		}

		for j := 0; j < 100; j++ {
			opt := rand.Intn(4)

			if opt == 0 {
				arrAdd(rand.Intn(1000), targetArr)

			} else if opt == 1 {
				delNum := rand.Intn(1000)
				arrDelete(targetArr, delNum)
			} else if opt == 2 {
				chaIndex := rand.Intn(100)
				chaValue := rand.Intn(1000)
				arrChange(targetArr, chaIndex, chaValue)
			} else if opt == 3 {
				findNum := rand.Intn(1000)
				_, _ = arrFind(targetArr, findNum)
			}
		}

	}
}

func arrAdd(value int, arr []int) {
	arr = append(arr, value)
}

func arrSize(arr []int) int {
	return len(arr)
}

func arrDelete(arr []int, value int) bool {
	flag := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			arr = append(arr[:i], arr[i+1:]...)
			flag = true
		}
	}
	return flag
}

func arrFind(arr []int, value int) (int, bool) {
	flag := false
	for index, element := range arr {
		if element == value {
			flag = true
			return index, flag
		}
	}
	return -1, flag
}

func arrChange(arr []int, index, value int) bool {
	if arrSize(arr) <= index {
		return false
	}

	arr[index] = value
	return true
}
