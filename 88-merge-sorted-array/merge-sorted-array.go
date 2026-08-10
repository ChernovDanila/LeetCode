 func merge(nums1 []int, m int, nums2 []int, n int)  {

    i := m - 1
    j := n - 1

    lenResult := m + n - 1
    
    for j >= 0 {
        
        if i<0 || nums1[i] < nums2[j] {
            nums1[lenResult] = nums2[j]
            j--
        }else{
            nums1[lenResult] = nums1[i]
            i--    
        }
        lenResult --

    }
}