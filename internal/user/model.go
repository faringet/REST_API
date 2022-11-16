package user

/*
json - чтобы конвертнуть модель в json;
bson - для mongoDB, подсказываем самой mongo что это поле в базе называется как _id
что в свою очередь является системным полем - монго само генерит уникальный индетификатор;
omitempty - поле может быть пустым
почему может быть пустым ?
потому что может прилетать запрос на создание пользователя, а в нем нет id
*/

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
