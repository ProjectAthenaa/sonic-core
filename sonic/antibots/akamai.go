package antibots

import "errors"

type Akamai struct {
}

func (a *Akamai) GetCookie(data ...interface{}) (string, error) {
	panic(errors.New("unimplemented"))
}
