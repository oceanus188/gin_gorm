package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"protect_gin_gorm/utils"
	"reflect"
)

type List struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(20);not null;primary key;auto increment" binding:"required"`
	State   string `json:"state" gorm:"type:varchar(20);not null" binding:"required"`
	Phone   string `json:"phone" gorm:"type:varchar(20);not null" binding:"required"`
	Email   string `json:"email" gorm:"type:varchar(50);not null" binding:"required"`
	Address string `json:"address" gorm:"type:varchar(20);not null" binding:"required"`
}

func (l *List) BeforeDelete(Db *gorm.DB) error {

	fmt.Println("list id =", l.ID)
	fmt.Println("before delete func")
	return nil
}

//存入数据
func _post(c *gin.Context) {

	var data []List
	c.ShouldBindJSON(&data)
	fmt.Println("data= ", data)

	res := utils.Db.Create(&data)
	fmt.Println("affect rows", res.RowsAffected)
	c.JSON(200, gin.H{
		"msg":  "添加成功",
		"data": data,
		"code": 200,
	})
}



//删除数据
func _delete(c *gin.Context) {
	//根据主键删除
	//接收参数
	id := c.Param("id")
	res := utils.Db.Delete(&List{}, id)
	//in
	//utils.Db.Delete(&List{}, []int{3, 4})
	//where条件
	result := utils.Db.Where("phone=?", "13838383840").Delete(&List{})

	fmt.Println("delete res error=", res.Error, "affect rows", result.RowsAffected)
}

//修改数据
func _exit(c *gin.Context) {
	var list List
	utils.Db.First(&list)

	list.Address = "2353sfsf@yeah.net"
	list.Phone = "18932245643"
	res := utils.Db.Save(&list)
	fmt.Println("affect rows count=", res.RowsAffected)

}

func _test(c *gin.Context) {

	var mapDemo []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["id"] = "1"
	m1["name"] = "jack"
	m1["title"] = "test1"
	mapDemo = append(mapDemo, m1)

	m2 := make(map[string]interface{})
	m2["id"] = "2"
	m2["name"] = "peter"
	m2["title"] = "test2"
	mapDemo = append(mapDemo, m2)
	fmt.Println("mapDemo=", mapDemo)

	jsonMapDemo, _ := json.Marshal(mapDemo)
	fmt.Println("jsonDemo=", string(jsonMapDemo))
	fmt.Println("jsonDemo type=", reflect.TypeOf(jsonMapDemo))

	var mapDemoUnmarshal []map[string]interface{}
	json.Unmarshal(jsonMapDemo, &mapDemoUnmarshal)
	fmt.Println("unmarshal = ", mapDemoUnmarshal)

	strJson := `[{"id":"100","name":"test1","title":"title_test1"},{"id":"200","name":"test2","title":"title_test2}]`
	fmt.Println("strJson type=", reflect.TypeOf([]byte(strJson)))

	var mapUnmarshal []map[string]interface{}
	err := json.Unmarshal([]byte(strJson), &mapUnmarshal)
	fmt.Println("unmarshal err=", err)

	fmt.Println(mapUnmarshal)
}

//查询数据
func _list(c *gin.Context) {
	//接收查询参数
	name := c.Param("name")
	var listSlice []List
	result := utils.Db.Where("name like ?", "%"+name+"%").Find(&listSlice)

	//utils.Db.First(&listSlice)
	for _, list := range listSlice {
		fmt.Println(list.Email)
		fmt.Println("list=", list)
	}

	fmt.Println("rows count", result.RowsAffected)
	//查询至map中
	mapList := map[string]interface{}{}
	utils.Db.Model(&List{}).First(&mapList)
	fmt.Println("map is ", mapList["name"])

	res := utils.Db.Table("list").Take(&mapList)
	fmt.Println("map2 is ", mapList)
	fmt.Println("rows affect is ", res.RowsAffected)

	//
}

func main() {

	//自动生成数据库
	//utils.Db.AutoMigrate(&List{})

	router := gin.Default()

	//测试连接

	router.POST("/user/add", _post)
	router.DELETE("/user/delete/:id", _delete)
	router.GET("/user/get/:name", _list)
	router.GET("/user/update", _exit)
	router.GET("/test", _test)

	router.Run(":3001")
}
