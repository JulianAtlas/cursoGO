package domain

//Mail Mail del usuario
var Mail string

//Username User de nuestros usuarios
var Username string

//Usuario mi usuario tiene mail y nombre
type Usuario struct {
	id         int
	mail       string
	Username   string
	seguidos   []*Usuario
	seguidores []*Usuario
	favoritos  []*Tweet
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
	return us.Username
}

//GetSeguidos usuario que sigo
func (us *Usuario) GetSeguidos() []*Usuario {
	return us.seguidos
}

//GetSeguidores usuarios que me siguen
func (us *Usuario) GetSeguidores() []*Usuario {
	return us.seguidores
}

//GetFavoritos getter de favoritos
func (us *Usuario) GetFavoritos() []*Tweet {
	return us.favoritos
}

//AddSeguidos agrego un usuario que comienzo a seguir
func (us *Usuario) AddSeguidos(user *Usuario) {
	us.seguidos = append(us.seguidos, user)
}

//AddSeguidor agrego al usuario a mis seguidores
func (us *Usuario) AddSeguidor(user *Usuario) {
	us.seguidores = append(us.seguidores, user)
}

//AddFavoritos agrego un tweet a los favoritos del user
func (us *Usuario) AddFavoritos(tw Tweet) {
	us.favoritos = append(us.favoritos, &tw)
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
	us.Username = username
}
