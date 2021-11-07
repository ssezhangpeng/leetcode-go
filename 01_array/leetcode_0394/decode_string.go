package leetcode0394

func decodeString(s string) string {
	var t string
	var s_num []int
	var s_str []string

	cnt := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			cnt = cnt*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			s_num = append(s_num, cnt)
			s_str = append(s_str, t)
			cnt = 0
			t = ""
		} else if s[i] == ']' {
			k := s_num[len(s_num)-1]
			s_num = s_num[:len(s_num)-1]
			for j := 0; j < k; j++ {
				s_str[len(s_str)-1] += t
			}
			t = s_str[len(s_str)-1]
			s_str = s_str[:len(s_str)-1]
		} else {
			t += string(s[i])
		}
	}

	return t
}
