package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o Mysql
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "murilo:03mu03wk05@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
