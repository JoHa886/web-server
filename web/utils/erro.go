package utils

const (
	RECODE_OK      = "0"
	RECODE_ERR     = "1"
	RECODE_UNKNOWN = "-1"
)

var recodeText = map[string]string{
	RECODE_OK:      "成功",
	RECODE_ERR:     "失败",
	RECODE_UNKNOWN: "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWN]
}
