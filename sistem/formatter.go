package sistem

type SistemFormatter struct {
	Sistem
}

func FormatSistem(sistem Sistem) SistemFormatter {
	formatter := SistemFormatter{
		Sistem: sistem,
	}

	return formatter
}

func FormatSistems(sistems []Sistem) []SistemFormatter {
	var sistemsFormatter []SistemFormatter

	for _, sistem := range sistems {
		sistemFormatter := FormatSistem(sistem)
		sistemsFormatter = append(sistemsFormatter, sistemFormatter)
	}

	return sistemsFormatter
}
