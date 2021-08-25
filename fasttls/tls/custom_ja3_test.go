package tls

import (
	"testing"
)

func TestStringToSpec(t *testing.T) {
	chromeStr := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0"
	chromeSpec, err := StringToSpec(chromeStr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	str := SpecToString(&chromeSpec)
	if str != chromeStr {
		t.Fatalf("Creating custom spec from chrome fail, got %s\nwant %s", chromeStr, str)
	}
}
