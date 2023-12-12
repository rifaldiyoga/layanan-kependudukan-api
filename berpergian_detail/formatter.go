package berpergian_detail

type BerpergianDetailFormatter struct {
	BerpergianDetail
}

func FormatBerpergianDetail(berpergian BerpergianDetail) BerpergianDetailFormatter {
	formatter := BerpergianDetailFormatter{
		BerpergianDetail: berpergian,
	}

	return formatter
}

func FormatBerpergianDetails(sktms []BerpergianDetail) []BerpergianDetailFormatter {
	var sktmsFormatter []BerpergianDetailFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatBerpergianDetail(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
