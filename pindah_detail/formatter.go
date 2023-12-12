package pindah_detail

type PindahDetailFormatter struct {
	PindahDetail
}

func FormatPindahDetail(berpergian PindahDetail) PindahDetailFormatter {
	formatter := PindahDetailFormatter{
		PindahDetail: berpergian,
	}

	return formatter
}

func FormatPindahDetails(sktms []PindahDetail) []PindahDetailFormatter {
	var sktmsFormatter []PindahDetailFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatPindahDetail(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
