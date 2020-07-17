package initialize

//mysql/mariaDB数据库连接
import (
	"fmt"
	"os"
	"perServer/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.GVA_CONFIG.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		fmt.Println("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GVA_DB.LogMode(admin.LogMode)
	}
}
