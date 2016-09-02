package mysql_escape

func MysqlEscapeString(source string) string {
	if len(source) == 0 {
		return ""
	}
	tempStr := source[:]

	var j int = 0
	desc := make([]byte, len(tempStr)*2)
	for i := 0; i < len(tempStr); i++ {
		var escape byte
		escape = 0
		switch tempStr[i] {
		case 0:
			escape = '0'
		case '\r':
			escape = 'r'
		case '\n':
			escape = 'n'
		case '\\':
			escape = '\\'
		case '\'':
			escape = '\''
		case '"':
			escape = '"'
		case '\032':
			escape = 'Z'
		}
		if escape != 0 {
			desc[j] = '\\'
			desc[j+1] = escape
			j += 2
		} else {
			desc[j] = tempStr[i]
			j += 1
		}
	}
	return string(desc[0:j])
}
