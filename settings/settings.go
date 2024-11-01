package settings

import (
	"os"
	"log"
)


const BaseDir = "~/.local/share/TemplateFilesManager/"

// путь до файлов установленной утилиты
var Filepath string = BaseDir + "files"

// путь до SQLite3 БД
// var PathDB string = BaseDir + "sqlite3.db"
var PathDB string = "sqlite3.db"

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		fatalLog := log.New(os.Stderr, "[FATAL]\t", log.Ldate|log.Ltime|log.Lshortfile)
		fatalLog.Panic(err)
	}
}
