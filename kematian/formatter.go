package kematian

type KematianFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatKematian(kematian Kematian) KematianFormatter {
	formatter := KematianFormatter{
		ID: kematian.ID,
	}

	return formatter
}

func FormatKematians(kematians []Kematian) []KematianFormatter {
	var kematiansFormatter []KematianFormatter

	for _, kematian := range kematians {
		kematianFormatter := FormatKematian(kematian)
		kematiansFormatter = append(kematiansFormatter, kematianFormatter)
	}

	return kematiansFormatter
}
