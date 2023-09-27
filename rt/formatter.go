package rt

type RTFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatRT(rt RT) RTFormatter {
	formatter := RTFormatter{
		ID:   rt.ID,
		Code: rt.Code,
		Name: rt.Name,
	}

	return formatter
}

func FormatRTs(rts []RT) []RTFormatter {
	var rtsFormatter []RTFormatter

	for _, rt := range rts {
		rtFormatter := FormatRT(rt)
		rtsFormatter = append(rtsFormatter, rtFormatter)
	}

	return rtsFormatter
}
