// Create controller

package productController

import (
	"fmt"
	"net/http"

	"github.com/bagaskaramadhan/go-project-toko/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(res *gin.Context) {
	var products []models.Product
	models.DB.Find((&products))

	res.JSON(http.StatusOK, gin.H{"products": products})

}

func GetById(res *gin.Context) {
	var product models.Product
	id := res.Param("id")

	err := models.DB.First(&product, id).Error
	fmt.Println("ERROR @GetById ==> ", err)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return

		default:
			res.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	res.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(res *gin.Context) {
	var product models.Product

	err := res.ShouldBindJSON(&product)
	if err != nil {
		res.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	res.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(res *gin.Context) {
	var product models.Product
	id := res.Param("id")

	err := res.ShouldBindJSON(&product)
	fmt.Println("ERROR @Update ==> ", err)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return

		default:
			res.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		res.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res.JSON(http.StatusOK, gin.H{"message": "Success updated"})
}

func Delete(res *gin.Context) {
	var product models.Product
	id := res.Param("id")

	err := models.DB.First(&product, id).Error
	fmt.Println("ERROR @GetById ==> ", err)
	if err != nil {
		res.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
		return
	}
	models.DB.Delete(&product, id)
	res.JSON(http.StatusOK, gin.H{"message": "Data has been deleted"})
}
