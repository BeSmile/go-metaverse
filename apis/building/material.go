package building

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/helpers/app"
	"go-metaverse/models/building"
)

// GetMaterialLandAnimal 获取动物产出
// @Description  获取动物产出
// @ID         	 get-material-land-animal
// @Accept       json
// @Produce      json
// @Success      200      {object}  app.Response
// @Router		/api/v1/building/material/animals [get]
func GetMaterialLandAnimal(c *gin.Context) {
	data := make([]building.Animal, 0)
	cat := building.Animal{
		Name: "猫",
		Island: 1,
		Output: "甜草莓",
	}
	dog := building.Animal{
		Name: "狗",
		Island: 2,
		Output: "甜草莓",
	}
	cow := building.Animal{
		Name: "牛",
		Island: 3,
		Output: "甜草莓",
	}
	chicken := building.Animal{
		Name: "鸡",
		Island: 4,
		Output: "",
	}
	sheep := building.Animal{
		Name: "羊",
		Island: 5,
		Output: "",
	}
	data = append(data, dog, cat, cow, chicken, sheep)
	app.OK(c, data, "success")
}

// SaveMagic 保存魔物
// @Description 保存魔物
// @ID         	 building-material-save-magic
// @Accept       json
// @Produce      json
// @Success      200      {object}  app.Response
// @Router /api/v1/building/material/magic [post]
func SaveMagic(c *gin.Context) {
	var magic building.Magic
	magic.SaveMagic(c)
}