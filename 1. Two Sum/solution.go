// Time 73.89%
// Memory 27.25%

func twoSum(nums []int, target int) []int {
    m := map[int]int{} // map for storing sub of element and target and element index
    for i, el := range nums {
		// if we already met element with value of target - element, then it is match
        if _, ok := m[el]; ok {
            return []int{m[el], i}
		// else add element to map
        } else {
            m[target - el] = i
        }
    }
    return []int{0, 0}
}