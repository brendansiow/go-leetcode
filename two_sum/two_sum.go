package goleetcode

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		subject := target - num
		if j, exist := numMap[subject]; exist {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}
