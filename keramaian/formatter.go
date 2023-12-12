package keramaian

type KeramaianFormatter struct {
	Keramaian
}

func FormatKeramaian(keramaian Keramaian) KeramaianFormatter {
	formatter := KeramaianFormatter{
		Keramaian: keramaian,
	}

	return formatter
}

func FormatKeramaians(sktms []Keramaian) []KeramaianFormatter {
	var sktmsFormatter []KeramaianFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatKeramaian(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
