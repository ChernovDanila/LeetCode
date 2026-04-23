func searchInsert(nums []int, target int) int {

    first := 0 
    last := len(nums)-1
    for first <= last {
        mide := (last + first)  / 2
        if target < nums[mide]{
            last = mide- 1
        }else if target > nums[mide]{
            first = mide + 1
        }else{
            return mide
        }               
    }

   return first 
   
}