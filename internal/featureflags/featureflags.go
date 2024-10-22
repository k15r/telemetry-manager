package featureflags

var f = &flags{
	v1beta1Enabled: false,
}

type flags struct {
	v1beta1Enabled bool
}

func SetV1beta1Enabled(enabled bool) {
	f.v1beta1Enabled = enabled
}

func IsV1beta1Enabled() bool {
	return f.v1beta1Enabled
}