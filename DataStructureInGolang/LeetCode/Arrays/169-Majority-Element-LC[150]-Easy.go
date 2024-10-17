
//HashMap
func majorityElement(nums []int) int {
   count := make(map[int]int)
   n := len(nums)

   for _,val := range nums{
    count[val]++
    if count[val] > n/2{
        return val
    }
   }
   return 0
}

//Sorting
// func majorityElement(nums []int) int {
//    sort.Ints(nums)
//    return nums[len(nums)/2]
// }
