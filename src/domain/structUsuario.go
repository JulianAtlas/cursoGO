package domain

//Mail Mail del usuario
var Mail string

//Username User de nuestros usuarios
var Username string

//Usuario mi usuario tiene mail y nombre
type Usuario struct {
	id       int
	mail     string
	username string
}

//GetID getter id
func (us *Usuario) GetID() int {
	return us.id
}

//GetMail getter mail
func (us *Usuario) GetMail() string {
	return us.mail
}

//GetUsername getter username
func (us *Usuario) GetUsername() string {
	return us.username
}

//SetID setter id
func (us *Usuario) SetID(id int) {
	us.id = id
}

//SetMail setter mail
func (us *Usuario) SetMail(mail string) {
	us.mail = mail
}

//SetUsername setter username
func (us *Usuario) SetUsername(username string) {
	us.username = username
}
