package status

type StatusFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatStatus(status Status) StatusFormatter {
	formatter := StatusFormatter{
		ID:   status.ID,
		Code: status.Code,
		Name: status.Name,
	}

	return formatter
}

func FormatStatuss(statuss []Status) []StatusFormatter {
	var statussFormatter []StatusFormatter

	for _, status := range statuss {
		statusFormatter := FormatStatus(status)
		statussFormatter = append(statussFormatter, statusFormatter)
	}

	return statussFormatter
}
