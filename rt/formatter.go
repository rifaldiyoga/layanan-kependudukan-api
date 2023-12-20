package rt

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/rw"
)

type RTFormatter struct {
	ID        int            `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	RwID      int            `json:"rw_id"`
	Rw        rw.RWFormatter `json:"rw" `
}

func FormatRT(rt RT) RTFormatter {
	formatter := RTFormatter{
		ID:        rt.ID,
		Code:      rt.Code,
		Name:      rt.Name,
		RwID:      rt.RwID,
		CreatedAt: helper.FormatDateToString(rt.CreatedAt),
		Rw:        rw.FormatRW(rt.RW),
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
