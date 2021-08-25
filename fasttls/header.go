package fasttls

type Headers map[string][]string

const (
	PseudoHeaderOrderKey = "PseudoHeaderOrderKey"
	HeaderOrderKey       = "HeaderOrderKey"
)

const (
	PseudoMethod    = ":method"
	PseudoAuthority = ":authority"
	PseudoScheme    = ":scheme"
	PseudoPath      = ":path"
)
