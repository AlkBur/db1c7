package db1c7

type User struct {
	//Основные
	Name string
	FullName string
	Path string  //Путь к настройкам пользователя
	Type string //???????

	//Права
	password string
	Rule string
	Interface string

	//служебные
	file string //Путь для чтения дополнительных данных
}