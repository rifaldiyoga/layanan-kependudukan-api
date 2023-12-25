package position

type PositionFormatter struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Jabatan string `json:"jabatan"`
}

func FormatPosition(position Position) PositionFormatter {
	formatter := PositionFormatter{
		ID:      position.ID,
		Code:    position.Code,
		Jabatan: position.Jabatan,
	}

	return formatter
}

func FormatPositions(positions []Position) []PositionFormatter {
	var positionsFormatter []PositionFormatter

	for _, position := range positions {
		positionFormatter := FormatPosition(position)
		positionsFormatter = append(positionsFormatter, positionFormatter)
	}

	return positionsFormatter
}
