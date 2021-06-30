package antibots

import "errors"

type ReCaptcha struct {
}

func (r *ReCaptcha) GetCookie(data ...interface{}) (string, error) {
	panic(errors.New("unimplemented"))
}
