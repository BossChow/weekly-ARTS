package main

// FindRemovedIndex: 找出需要删除的元素（下标）
func FindRemovedIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	negativeCount := 0
	// 统计负值数量
	for _, i := range arr {
		if i < 0 {
			negativeCount++
		}
	}

	// 情况A: 负值数量为奇数
	// 删除负值中的最大值
	if negativeCount%2 != 0 {
		maxValIndex := 0
		for index, val := range arr {
			if val < 0 && arr[maxValIndex] > 0 {
				maxValIndex = index
			}
			if val < 0 && val > arr[maxValIndex] {
				maxValIndex = index
			}
		}
		return maxValIndex
	} else { // 情况B: 负值数量为偶数
		if negativeCount == len(arr) { // 情况B1:全部为负值，删除最小值
			minValIndex := 0
			for index, val := range arr {
				if val < arr[minValIndex] {
					minValIndex = index
				}
			}
			return minValIndex
		} else { // 情况B2:不全部为负值，删除正值中的最小值
			minValIndex := 0
			for index, val := range arr {
				if val > 0 && val < arr[minValIndex] {
					minValIndex = index
				}
			}
			return minValIndex
		}
	}
}
