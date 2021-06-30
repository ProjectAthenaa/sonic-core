package antibots

import "errors"

type PerimeterX struct {

}

func (p *PerimeterX) GetCookie(data ...interface{}) (string, error){
	panic(errors.New("unimplemented"))
}
