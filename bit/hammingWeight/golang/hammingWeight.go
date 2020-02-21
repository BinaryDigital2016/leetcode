/*
统计1的个数
*/
// func hammingWeight(num uint32) int {
//     count := 0
//     for num > 0{
//         if num & 1 == 1{
//             count++
//         }
//         num >>= 1
//     }
//     return count
// }

// n&(n-1)将最低位1置0
func hammingWeight(num uint32) int {
    count := 0
    for num > 0{
        count++
        num = num & (num-1)
    }
    return count
}
