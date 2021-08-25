package tls

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrInvalidExtension string

func (e ErrInvalidExtension) Error() string {
	return fmt.Sprintf("utls: extension does not exist%s\n", string(e))
}

func StringToSpec(ja3 string) (ClientHelloSpec, error) {
	// extMap maps extension values to the TLSExtension object associated with the
	// number. Some values are not put in here because they must be applied in a
	// special way. For example, "10" is the SupportedCurves extension which is also
	// used to calculate the JA3 signature. These JA3-dependent values are applied
	// after the instantiation of the map.
	extMap := map[string]TLSExtension{
		"0": &SNIExtension{},
		"5": &StatusRequestExtension{},
		// These are applied later
		// "10": &SupportedCurvesExtension{...}
		// "11": &SupportedPointsExtension{...}
		"13": &SignatureAlgorithmsExtension{
			SupportedSignatureAlgorithms: []SignatureScheme{
				ECDSAWithP256AndSHA256,
				ECDSAWithP384AndSHA384,
				ECDSAWithP521AndSHA512,
				PSSWithSHA256,
				PSSWithSHA384,
				PSSWithSHA512,
				PKCS1WithSHA256,
				PKCS1WithSHA384,
				PKCS1WithSHA512,
				ECDSAWithSHA1,
				PKCS1WithSHA1,
			},
		},
		"16": &ALPNExtension{
			AlpnProtocols: []string{"h2", "http/1.1"},
		},
		"18": &SCTExtension{},
		"21": &UtlsPaddingExtension{GetPaddingLen: BoringPaddingStyle},
		"23": &UtlsExtendedMasterSecretExtension{},
		"27": &CertCompressionAlgsExtension{
			[]CertCompressionAlgo{
				CertCompressionBrotli,
			},
		},
		"28": &FakeRecordSizeLimitExtension{},
		"35": &SessionTicketExtension{},
		"43": &SupportedVersionsExtension{Versions: []uint16{
			GREASE_PLACEHOLDER,
			VersionTLS13,
			VersionTLS12,
			VersionTLS11,
			VersionTLS10}},
		"44": &CookieExtension{},
		"45": &PSKKeyExchangeModesExtension{
			Modes: []uint8{
				PskModeDHE,
			}},
		"51":    &KeyShareExtension{KeyShares: []KeyShare{{Group: X25519}}},
		"13172": &NPNExtension{},
		"65281": &RenegotiationInfoExtension{
			Renegotiation: RenegotiateOnceAsClient,
		},
	}

	tokens := strings.Split(ja3, ",")
	// version := tokens[0]
	ciphers := strings.Split(tokens[1], "-")
	extensions := strings.Split(tokens[2], "-")
	curves := strings.Split(tokens[3], "-")
	if len(curves) == 1 && curves[0] == "" {
		curves = []string{}
	}
	pointFormats := strings.Split(tokens[4], "-")
	if len(pointFormats) == 1 && pointFormats[0] == "" {
		pointFormats = []string{}
	}

	// parse curves
	var targetCurves []CurveID
	for _, c := range curves {
		cid, err := strconv.ParseUint(c, 10, 16)
		if err != nil {
			return ClientHelloSpec{}, err
		}
		targetCurves = append(targetCurves, CurveID(cid))
	}
	extMap["10"] = &SupportedCurvesExtension{Curves: targetCurves}

	// parse point formats
	var targetPointFormats []byte
	for _, p := range pointFormats {
		pid, err := strconv.ParseUint(p, 10, 8)
		if err != nil {
			return ClientHelloSpec{}, err
		}
		targetPointFormats = append(targetPointFormats, byte(pid))
	}
	extMap["11"] = &SupportedPointsExtension{SupportedPoints: targetPointFormats}

	// build extenions list
	var exts []TLSExtension
	for _, e := range extensions {
		te, ok := extMap[e]
		if !ok {
			return ClientHelloSpec{}, ErrInvalidExtension(e)
		}
		exts = append(exts, te)
	}
	// build SSLVersion
	//vid64, err := strconv.ParseUint(version, 10, 16)
	//if err != nil {
	//	return ClientHelloSpec{}, err
	//}
	// vid := uint16(vid64)

	// build CipherSuites
	var suites []uint16
	for _, c := range ciphers {
		cid, err := strconv.ParseUint(c, 10, 16)
		if err != nil {
			return ClientHelloSpec{}, err
		}
		suites = append(suites, uint16(cid))
	}

	return ClientHelloSpec{
		//TLSVersMin:         vid,
		//TLSVersMax:         vid,
		CipherSuites:       suites,
		CompressionMethods: []byte{compressionNone},
		Extensions:         exts,
	}, nil
}

// SpecToString returns the ja3 string of a client hello spec. If this string
// is hashed with MD5, you can get the ja3 hash of a spec
func SpecToString(spec *ClientHelloSpec) string {
	version := spec.TLSVersMax
	ciphers := strings.Trim(strings.Replace(fmt.Sprint(spec.CipherSuites), " ", "-", -1), "[]")

	var curves []CurveID
	var curveFormats []uint8
	var extensions []uint16
	for _, e := range spec.Extensions {
		switch v := e.(type) {
		case *SNIExtension:
			extensions = append(extensions, 0)
		case *StatusRequestExtension:
			extensions = append(extensions, 5)
		case *SupportedCurvesExtension:
			curves = v.Curves
			extensions = append(extensions, 10)
		case *SupportedPointsExtension:
			curveFormats = v.SupportedPoints
			extensions = append(extensions, 11)
		case *SignatureAlgorithmsExtension:
			extensions = append(extensions, 13)
		case *ALPNExtension:
			extensions = append(extensions, 16)
		case *SCTExtension:
			extensions = append(extensions, 18)
		case *UtlsPaddingExtension:
			extensions = append(extensions, 21)
		case *UtlsExtendedMasterSecretExtension:
			extensions = append(extensions, 23)
		case *CertCompressionAlgsExtension:
			extensions = append(extensions, 27)
		case *FakeRecordSizeLimitExtension:
			extensions = append(extensions, 28)
		case *SessionTicketExtension:
			extensions = append(extensions, 35)
		case *SupportedVersionsExtension:
			extensions = append(extensions, 43)
		case *CookieExtension:
			extensions = append(extensions, 44)
		case *PSKKeyExchangeModesExtension:
			extensions = append(extensions, 45)
		case *KeyShareExtension:
			extensions = append(extensions, 51)
		case *NPNExtension:
			extensions = append(extensions, 13172)
		case *RenegotiationInfoExtension:
			extensions = append(extensions, 65281)
		}
	}
	extStr := strings.Trim(strings.ReplaceAll(fmt.Sprint(extensions), " ", "-"), "[]")
	curvesStr := strings.Trim(strings.ReplaceAll(fmt.Sprint(curves), " ", "-"), "[]")
	curveFmtStr := strings.Trim(strings.ReplaceAll(fmt.Sprint(curveFormats), " ", "-"), "[]")
	ja3 := fmt.Sprint(version) + "," + ciphers + "," + extStr + "," + curvesStr + "," + curveFmtStr
	return ja3
}
