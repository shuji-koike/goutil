package goutil

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSliceMap(t *testing.T) {
	slice := []int{1, 2, 3}
	var dist []string
	dist = SliceMap(slice, func(e int) string { return fmt.Sprintf("int:%d", e) }).([]string)
	if len(dist) != 3 {
		t.Fail()
	}
	if dist[0] != "int:1" {
		t.Fail()
	}
	if dist[1] != "int:2" {
		t.Fail()
	}
	if dist[2] != "int:3" {
		t.Fail()
	}
}

func TestSliceReduce(t *testing.T) {
	slice := []int{1, 2, 3}
	acc := SliceReduce(slice, 0, func(acc, e int) int { return acc + e }).(int)
	if acc != 6 {
		t.Fail()
	}
}

func TestSliceReduceToString(t *testing.T) {
	slice := []int{1, 2, 3}
	acc := SliceReduce(slice, "acc", func(acc string, e int) string {
		return fmt.Sprintf("%s,%d", acc, e)
	}).(string)
	if acc != "acc,1,2,3" {
		t.Fail()
	}
}

func BenchmarkSliceMap(b *testing.B) {
	slice := []int{1, 2, 3}
	for n := 0; n < b.N; n++ {
		dist1 := SliceMap(slice, func(e int) string { return fmt.Sprintf("%d", e) }).([]string)
		dist2 := SliceMap(dist1, func(e string) int { v, _ := strconv.Atoi(e); return v }).([]int)
		for i := range slice {
			if slice[i] != dist2[i] {
				b.Fail()
				return
			}
		}
	}
}

func BenchmarkFor(b *testing.B) {
	slice := []int{1, 2, 3}
	for n := 0; n < b.N; n++ {
		var (
			dist1 = make([]string, len(slice))
			dist2 = make([]int, len(slice))
		)
		for i := range slice {
			dist1[i] = fmt.Sprintf("%d", slice[i])
		}
		for i := range dist2 {
			dist2[i], _ = strconv.Atoi(dist1[i])
		}
		for i := range slice {
			if slice[i] != dist2[i] {
				b.Fail()
				return
			}
		}
	}
}
