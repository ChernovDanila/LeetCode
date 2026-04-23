func firstStableIndex(nums []int, k int) int {

    n := len(nums)
    minRight := make([]int, n)
    minRight[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i--{
        minRight[i] = min(nums[i], minRight[i+1])
    }

    maxLeft := nums[0]
    if maxLeft - minRight[0] <= k{
        return 0
    }

    for i:=1; i < n; i ++{
        maxLeft = max(maxLeft, nums[i])
        if maxLeft - minRight[i] <= k{
            return i
        }
    }    
    return -1
}