package application

type IUserApplication interface {
 	GetUserByID(id int)
}