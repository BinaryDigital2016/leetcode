func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    result := ""
	s0 := strs[0]
	for k, i := range s0 {
		si := string(i)
		for _, s := range strs {
            // 判断长度，否则会panic，若s0使用长度最短str则不用判断
			if len(s) <=k || s0[k] != s[k]{
				return result
			}
		}
		result = fmt.Sprintf("%s%s", result, si)
	}

	return result
}
