package jobs

type Option func(j *job)

func WithTrackError() Option {
	return func(j *job) {
		j.trackError = true
	}
}
