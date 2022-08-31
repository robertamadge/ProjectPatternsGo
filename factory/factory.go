package factory

import "fmt"

//Nesse padrão criamos um objeto sem expor sua lógica ao cliente.
//No exemplo abaixo a interface única pode ser chamada por 2 bancos de dados diferentes e o cliente não saberá.
type (
	mongoDB struct {
		db map[string]string
	}
	sqlite struct {
		db map[string]string
	}

	//Database implements two functions
	Database interface {
		GetData(string) string
		PutData(string, string)
	}
)

//GetData e PutData são os dois métodos que a interface utilizará
func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.db[query]; !ok {
		return ""
	}

	fmt.Println("MongoDB")
	return mdb.db[query]
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.db[query]; !ok {
		return ""
	}

	fmt.Println("Sqlite")
	return sql.db[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.db[query] = data
}

func (sql sqlite) PutData(query string, data string) {
	sql.db[query] = data
}

//DatabaseFactory escolhe qual dos dois bancos será utilizado pelo cliente
//dependendo do ambiente.
func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongoDB{
			db: make(map[string]string),
		}
	case "development":
		return sqlite{
			db: make(map[string]string),
		}
	default:
		return nil
	}
}
