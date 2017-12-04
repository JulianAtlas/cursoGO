package domain

type Tweet interface{
	GetText() string 
	GetID() int 
	GetUser() *Usuario
	SetID(id int)
	SetText(nuevoTexto string)
}