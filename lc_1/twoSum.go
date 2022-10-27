package lc_1

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

// // 双重循环，判断i和j相加是否等于target
// func Answer(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				s := []int{i, j}
// 				return s
// 			}
// 		}
// 	}
// 	return nil
// }

// 单循环 判断是否有nums[]中是否存在target-nums[i]
// go中没有indexOf之类的属性 不好实现

// map实现
// 将i, nums[i]存入key中，得到myMap

// myMap[nums[i]] = i;
// 判断 myMap[target - nums[i]]是否存在

// myMap[i] = nums[i]
// 判断 myMap中是否有target-num[i]
// 但这样就失去了map的意义，和数组一样

// func Answer(nums []int, target int) []int {
// 	myMap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		myMap[nums[i]] = i
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if v, s := myMap[target-nums[i]]; s && (v != i) {
// 			return []int{i, v}
// 		}
// 	}
// 	return nil
// }

// Map 优化
// 优化点：
// 1. 不用先把所有值都存入map，一边查找一边存储即可
// 这样做好处：减少循环、每次查找完再存可以保证不会重复
// 2. 用range循环？
// 3. 不用make创建map?
// 返回的数据注意先p后k，因为查找的是k之前的，即nums[0] ~ nums[k-1]之间，p<k
func Answer(nums []int, target int) []int {
	myMap := map[int]int{}
	for k, v := range nums {
		if p, ok := myMap[target-v]; ok {
			return []int{p, k}
		}
		myMap[v] = k
	}
	return nil
}
