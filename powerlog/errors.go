package powerlog

type InvalidUsagePooledEncoderError string

func (err InvalidUsagePooledEncoderError) Error() string {
	return string(err)
}
