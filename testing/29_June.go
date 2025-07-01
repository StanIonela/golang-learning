// TODO [Implement]
// - Function: Solve `TwoSum` using a map for O(n) lookup
// - Function: Implement `DuplicateFinder` using a map as a set

// TODO [Test]
// - Write and run test cases for both functions
// - Handle edge cases: empty input, single element, large input
// - Run with `go test` and verify correctness

// TODO [Analyze]
// - Add comments explaining time and space complexity
// - Benchmark map vs slice performance (optional)

// TODO [Practice]
// - Solve 1–2 related Exercism problems:
//   - https://exercism.org/tracks/go/exercises/isogram
//   - https://exercism.org/tracks/go/exercises/pangram

// TODO [Cleanup]
// - Save solutions: `twosum.go`, `duplicatefinder.go`
// - Document approach and gotchas in `notes/29_June_arrays_maps.md`
// - Commit progress to Git: `git add . && git commit -m "29 June: arrays and hashmaps practice"`
package main  
import (
  "fmt"
)

func TwoSum (s []int, t int) []int {
  store := make(map[int]int)
  for i, v := range s {
    complement := t - v
    if j, ok := store[complement]; ok {
      return []int{j, i}
    }
    store[v] = i
  }
  return nil
}

func DuplicateFinder(s []int) []int {
  appearance := make(map[int]int)
  result := []int{}
  for _, v := range s {
    appearance[v]++
    if appearance[v] > 1 {
      result = append(result, v)
    }
  }
  return result
}

func main()  {
  nums := []int{4, 8, 15, 25}
  target := 12
  result := TwoSum(nums, target)
  fmt.Println(result)

  nums1 := []int{1,2,3,4,5,6,5,4,4,1, 1}
  result1 := DuplicateFinder(nums1)
  fmt.Println(result1)
}
