package plugin

type Authentification interface {
	Authentificate(username, password string) (bool, error)
}
