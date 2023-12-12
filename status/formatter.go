package status

import "layanan-kependudukan-api/helper"

type StatusFormatter struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func FormatStatus(status Status) StatusFormatter {
	formatter := StatusFormatter{
		ID:        status.ID,
		Code:      status.Code,
		Name:      status.Name,
		CreatedAt: helper.FormatDateToString(status.CreatedAt),
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
