func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    m, n := len(nums1), len(nums2)
	
	if m > n{
        nums1, nums2, m, n = nums2, nums1, n, m 
    }

	low, high := 0, m
	
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m + n + 1) / 2 - partitionX
		maxX := getMax(nums1, partitionX)
		minX := getMin(nums1, partitionX)
		maxY := getMax(nums2, partitionY)
		minY := getMin(nums2, partitionY)	
		if maxX <= minY && maxY <= minX {
			if (m+n)%2 == 0 {
				return float64(max(maxX, maxY)+min(minX, minY)) / 2.0
			}
			return float64(max(maxX, maxY))
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}	
	}
	return 0.0

}

func getMax(nums []int, partition int) int {
	if partition == 0 {
		return -1 << 31
	}
	return nums[partition-1]
}
func getMin(nums []int, partition int) int {
	if partition == len(nums) {
		return 1<<31 - 1
	}
	return nums[partition]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
