



//Approach 1:
/*
Time: O(n+m)
Space: O(1)
*/

import(
    "sort"
)
func merge(nums1 []int, m int, nums2 []int, n int)  {
   
   for i,j := 0,m ; i < n ; i++{
        nums1[j] = nums2[i]
        j++
    }
    
    sort.Ints(nums1)

}

//Three Pointer Approach.
/*
func merge(nums1 []int, m int, nums2 []int, n int)  {
  ptr1, ptr2, ptr3 := m-1, n-1, m+n-1

    for ; ptr1 >= 0 && ptr2 >= 0 ; ptr3--{
        if nums2[ptr2] >= nums1[ptr1]{
            nums1[ptr3] = nums2[ptr2]
            ptr2--
        }else{
            nums1[ptr3] = nums1[ptr1]
            ptr1--
        }
    }
    if ptr2 >= 0{
        copy(nums1[:ptr3+1],nums2[:ptr2+1])
    }
}
*/

