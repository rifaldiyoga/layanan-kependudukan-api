package job

import "layanan-kependudukan-api/helper"

type JobFormatter struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func FormatJob(job Job) JobFormatter {
	formatter := JobFormatter{
		ID:        job.ID,
		Code:      job.Code,
		Name:      job.Name,
		CreatedAt: helper.FormatDateToString(job.CreatedAt),
	}

	return formatter
}

func FormatJobs(jobs []Job) []JobFormatter {
	var jobsFormatter []JobFormatter

	for _, job := range jobs {
		jobFormatter := FormatJob(job)
		jobsFormatter = append(jobsFormatter, jobFormatter)
	}

	return jobsFormatter
}
