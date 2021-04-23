package two_sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.
//
//
//
//Example 1:
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Output: Because nums[0] + nums[1] == 9, we return [0, 1].
//Example 2:
//
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
//Example 3:
//
//Input: nums = [3,3], target = 6
//Output: [0,1]
//
//
//Constraints:
//
//2 <= nums.length <= 103
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//Only one valid answer exists.

//Approach 1: Brute Force
// loop through each element x and find if there is another value that equals to target - x
// Time complexity: O(n^2). For each element, we try to find its complement by looping through the rest of array which takes O(n) time. Therefore, the time complexity is O(n^2)
// Space complexity: O(1_
func TwoSum(nums []int, target int) []int {
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[j] == target-nums[i] {
	//			return []int{i, j}
	//		}
	//	}
	//}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target - nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil

}

// Approach 2: Two-pass Hash Table
// We can reduce the lookup time from O(n) to O(1) by trading space for speed.
// A hash table is for this. It supports fast look up in near constant time
// Time complexity: O(n). we traverse the list containing n elements exactly twice. Since the hash table reduces the look up time to O(1_, the time complexiity is O(n)
//Space complexity: O(n). The extran space required depends on the number of items stored in the hash table, which stores  exactly n elements.
func TwoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i, v := range nums {
		hashMap[v] = i
	}

	for i, v := range nums {
		complement := target - v
		if _, ok := hashMap[complement]; ok && complement != v {
			return []int{i, hashMap[complement]}
		}
	}

	return nil
}

func TwoSum2a(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i, v := range nums {
		hashMap[v] = i
	}

	for i, v := range nums {
		complement := target - v

		if _, ok := hashMap[complement]; ok && complement != v {
			return []int{i, hashMap[complement]}
		}
	}
	return nil
}

// Approach 3: One-pass Hash Table
// Complexity Analysis:
//
//Time complexity : O(n)O(n). We traverse the list containing nn elements only once. Each look up in the table costs only O(1)O(1) time.
//
//Space complexity : O(n)O(n). The extra space required depends on the number of items stored in the hash table, which stores at most nn elements.
func TwoSum3(nums []int, target int) []int {

	hashMap := make(map[int]int, len(nums))

	for i, v := range nums {
		complement := target - v

		if _, ok := hashMap[complement]; ok {
			return []int{hashMap[complement], i}
		}

		hashMap[v] = i
	}
	return nil
}

func TwoSum3a(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i, v := range nums {
		complement := target - v
		if _, ok := hashMap[complement]; ok {
			return []int{hashMap[complement], i}
		}
		hashMap[v] = i
	}
	return nil
}

func TwoSum3b(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i, v := range nums {
		complement := target - v
		if _, ok := hashMap[complement]; ok {
			return []int{hashMap[complement], i}
		}
		hashMap[v] = i
	}
	return nil
}
