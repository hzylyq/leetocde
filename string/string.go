package string

// 1576. 替换所有的问号
func ModifyString(s string) string {
	temp := make([]byte, len(s))
	copyStr := []byte(s)

	for idx, char := range s {
		if char != '?' {
			temp[idx] = s[idx]
			continue
		}

		for j := byte('a'); j <= 'z'; j++ {
			if (idx == 0 || j != copyStr[idx-1]) && (idx == len(s)-1 || j != copyStr[idx+1]) {
				temp[idx] = j
				copyStr[idx] = j
				break
			}
		}
	}

	return string(temp)
}

// 1678. 设计 Goal 解析器
func Interpret(command string) string {
	var result string
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			result += "G"
		} else if command[i] == '(' {
			if command[i+1] == 'a' {
				result += "al"
				i += 3
			} else {
				result += "o"
				i += 1
			}
		}
	}

	return result
}
