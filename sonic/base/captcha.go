package base

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/godtoy/autosolve"
)

type CaptchaType int

const (
	ReCaptchaV2Checkbox   CaptchaType = 0
	ReCaptchaV2Invisible  CaptchaType = 1
	ReCaptchaV3           CaptchaType = 2
	HCaptchaCheckbox      CaptchaType = 3
	HCaptchaInvisible     CaptchaType = 4
	GeeTest               CaptchaType = 5
	ReCaptchaV3Enterprise CaptchaType = 6
)

type Opts struct {
	ReCaptchaAction   string
	ReCaptchaMinScore float32
}

func (tk *BTask) StatusListener(status autosolve.Status) {
	tk.SetStatus(module.STATUS_SOLVING_CAPTCHA, string("Autosolve "+status))
}

func (tk *BTask) ErrorListener(err error) {
	tk.SetStatus(module.STATUS_ERROR, "Autosolve "+err.Error())
}

func (tk *BTask) CaptchaTokenCancelResponseListener(response autosolve.CaptchaTokenCancelResponse) {
}

func (tk *BTask) CaptchaTokenResponseListener(response autosolve.CaptchaTokenResponse) {
	callbackChannel, loaded := tk.autosolveChannels.LoadAndDelete(response.TaskId)
	if loaded {
		defer close(callbackChannel.(chan autosolve.CaptchaTokenResponse))
		callbackChannel.(chan autosolve.CaptchaTokenResponse) <- response
	}
}
