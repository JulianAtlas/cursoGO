package domain

//Mail Mail del usuario
var Mail string

//Username User de nuestros usuarios
var Username string

//Usuario mi usuario tiene mail y nombre
type Usuario struct {
	ID       int
	Mail     string
	Username string
}
