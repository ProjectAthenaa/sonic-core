package antibots

type Antibot interface {
	GetCookie(data ...interface{}) (string, error)
}
