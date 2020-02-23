/*
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:

输入: s = "egg", t = "add"
输出: true

示例 2:

输入: s = "foo", t = "bar"
输出: false

示例 3:

输入: s = "paper", t = "title"
输出: true

说明:
你可以假设 s 和 t 具有相同的长度。
*/
// //使用map保存映射
// func isIsomorphic(s string, t string) bool {
//     return isIsomorphicCore(s,t)&&isIsomorphicCore(t,s) //两边都要检查
// }

// func isIsomorphicCore(s string, t string) bool{
//     m := make(map[uint8]uint8)
//     n := len(s)
//     for i:=0;i<n;i++{
//         if e,ok:=m[s[i]];ok{
//             if e != t[i]{
//                 return false
//             }
//         } else {
//             m[s[i]] = t[i]
//         }
//     }
//     return true
// }

//映射成第三种模式，略
