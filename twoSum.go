func twoSum(nums []int, target int) []int {
    var total []int
    for index1, number1 := range nums {
        for index2, number2 := range nums{
            if index1 == index2 {
                continue
            } else if number1 + number2 == target {
                total = []int{index1, index2}
            }
        }
    }
    return total
}
