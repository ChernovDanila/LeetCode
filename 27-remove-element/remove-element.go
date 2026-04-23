func removeElement(nums []int, val int) int {
    
    slow := len(nums)-1
    fast := 0
    counter := 0
    for fast <= slow{
        if nums[fast]==val {
            counter++
            nums[fast] = nums[slow]
            slow--
        }else {
            fast++
        }
    }
    return len(nums) - counter;
}