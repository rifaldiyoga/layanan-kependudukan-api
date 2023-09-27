package rw

type RWFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatRW(rw RW) RWFormatter {
	formatter := RWFormatter{
		ID:   rw.ID,
		Code: rw.Code,
		Name: rw.Name,
	}

	return formatter
}

func FormatRWs(rws []RW) []RWFormatter {
	var rwsFormatter []RWFormatter

	for _, rw := range rws {
		rwFormatter := FormatRW(rw)
		rwsFormatter = append(rwsFormatter, rwFormatter)
	}

	return rwsFormatter
}
