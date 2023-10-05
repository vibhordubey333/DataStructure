/*
Monotonic means either the array is in increasing order or in decreasing order
*/
func isMonotonic(nums []int) bool {
    increase := true
    decrease := true
    for i := 0 ; i < len(nums) - 1; i++{
        if nums[i] > nums[i+1]{
            increase = false
        }
    }


    for i := 0 ; i < len(nums) -1  ; i++{
        if nums[i] < nums[i+1]{
            decrease = false
        }
    }

    
    return  increase || decrease 
}
