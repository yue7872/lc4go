package lc_167

// 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列 ，请你从数组中找出满足相加之和等于目标数 target 的两个数。
// 如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
// 你所设计的解决方案必须只使用常量级的额外空间。

// 和两数之和类似
// 顺带回顾一下

// // shamefully暴力解
// func Answer(numbers []int, target int) []int {
// 	for i := 0; i < len(numbers); i++ {
// 		for j := i + 1; j < len(numbers); j++ {
// 			if numbers[i]+numbers[j] == target {
// 				return []int{i + 1, j + 1}
// 			}
// 		}
// 	}
// 	return nil
// }

// map
// 遍历数组，存入myMap[]，令 myMap[nums[i]] = i;
// 判断myMap是否存在target - nums[i]
// 判断方法：
// if _,ok:= myMap[targe-nums[i]]; ok {
// }

// func Answer(numbers []int, target int) []int {
// 	myMap := map[int]int{}
// 	for idx, v := range numbers {
// 		if _, ok := myMap[target-v]; ok {
// 			return []int{myMap[target-v] + 1, idx + 1}
// 		}
// 		myMap[v] = idx
// 	}
// 	return nil
// }

// 注意到这道题 非递减顺序排列，可以考虑用双指针/二分
// numbers = [2,7,11,15], target = 9
// i := 0; j:= len(nums)

// [-2, 2,3, 4,5,7,9] 8
// if i+j < target { i< 0 必成立 否则没有满足题意的值}
// then i++ => i+j > target    j--
// if i+j < target { i++ }

// 输出：[1,2]

// []int{2, 7, 11, 15} 9

// func Answer(numbers []int, target int) []int {
// 	for i := 0; i < len(numbers); i++ {
// 		for j := len(numbers) - 1; j > i; j-- {
// 			if numbers[i]+numbers[j] == target {
// 				return []int{i + 1, j + 1}
// 			}
// 			if numbers[i]+numbers[j] < target {
// 				continue
// 			}
// 		}
// 	}
// 	return nil
// }

// 改进

func Answer(numbers []int, target int) []int {
	// 定义双指针
	n := len(numbers)
	// j定义在外面的好处是不用在循环里写判断条件 即不用每次j--都判断是否相等/而是j--到 和<target才跳出循环判断（可以与上一种方法对比）
	j := n - 1
	for i := 0; i < n; i++ {
		// for循环 循环条件是 i<j 且 nums[i] + nums[j] > target 此时j-- 否则跳出循环 判断相等 不等则i++
		// 这里相当于上一种写法的简化
		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

// TODO：严格的二分查找
