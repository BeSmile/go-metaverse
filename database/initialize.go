package database

func Setup() {
	var db = new(Mysql) // 初始化mysql
	db.Setup()
}
