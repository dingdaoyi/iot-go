package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iot-go/internal/config"
	"iot-go/internal/model"
	"iot-go/pkg/model/result"
	"log"
	"net/http"
	"strconv"
)

// RegisterDepartmentRoutes 注册 department 相关的路由
func RegisterDepartmentRoutes(router *gin.RouterGroup) {
	router.POST("/create", CreateDepartment)
	router.GET("/:id", GetById)
	router.GET("/tree", Tree)
	router.DELETE("/:id", DeleteById)
}

// DeleteById 删除
func DeleteById(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("未提供 ID"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("无效的 ID"))
		return
	}
	// 判断是否可以删
	validError := validateDelete(id)
	if validError != nil {
		c.JSON(http.StatusBadRequest, result.Error(validError))
		return
	}
	tx := config.Db.Delete(&model.Department{}, id)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("删除错误:"+tx.Error.Error()))
		return
	}
	if tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("删除结果为0"))
		return
	}
	c.JSON(http.StatusOK, result.Data("删除成功"))
}

func validateDelete(id int) error {
	var exists bool
	config.Db.Model(&model.Department{}).Where("parent_id = ?", id).Select("count(*) > 0").Scan(&exists)
	if exists {
		return fmt.Errorf("存在子级部门, 无法删除")
	}
	return nil
}

func CreateDepartment(c *gin.Context) {
	db := config.Db
	d := &model.Department{}
	err := c.BindJSON(d)
	if err != nil {
		log.Printf("保存错误:%s", err)
		c.JSON(http.StatusBadRequest, result.Error(err))
		return
	}
	var count int64
	cr := db.Model(&model.Department{}).Where("name = ?", d.Name).Count(&count)
	if cr.Error != nil {
		log.Printf("查询出错: %v", cr.Error)
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("查询出错"))
		return
	}
	if count > 0 {
		// 记录存在
		log.Printf("记录存在，数量: %d", count)
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("部门已经存在"))
		return
	}
	r := db.Create(d)
	if r.Error != nil || r.RowsAffected == 0 {
		log.Printf("保存失败:%s", err)
		c.JSON(http.StatusBadRequest, result.Error(err))
		return
	}
	c.JSON(http.StatusOK, result.Ok())
}

// GetById 	根据id获取
func GetById(c *gin.Context) {
	d := &model.Department{}
	id := c.Param("id")
	r := config.Db.First(d, id)
	if r.Error != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("不存在"))
		return
	}
	c.JSON(http.StatusOK, result.Data(d))
}
func Tree(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("参数ID不存在"))
		return
	}
	parseUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("参数ID为int类型"))
		return
	}
	children, erMsg := getDepartmentWithChildren(config.Db, uint(parseUint))
	if erMsg != nil {
		c.JSON(http.StatusNotFound, result.ErrorWithMessage("notFount"))
		return
	}

	c.JSON(http.StatusOK, result.Data(children))
}
func getDepartmentWithChildren(db *gorm.DB, id uint) (*model.Department, error) {
	var department model.Department
	if err := db.Preload("Children.Children").First(&department, id).Error; err != nil {
		return nil, err
	}
	return &department, nil
}
