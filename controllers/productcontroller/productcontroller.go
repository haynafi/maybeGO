package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haynafi/maybeGO/models"
	"gorm.io/gorm"
)

// Index godoc
// @Summary      Show all product
// @Description  get all product
// @Tags         Get
// @Router       /products [get]
func Index(c *gin.Context) {

	var product []models.Product

	models.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{"products": product})
}

// Show godoc
// @Summary      Show an product
// @Description  get product by ID
// @Tags         Show
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Router       /product/{id} [get]
func Show(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

// Create godoc
// @Summary      Insert product
// @Description  insert product
// @Tags         Insert
// @Accept       json
// @Produce      json
// @Param        nama_product   path      string  true  "Nama Produk"
// @Param        deskripsi   path      string  true  "deskripsi"
// @Router       /product/{id} [post]
func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed Update Product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update OK"})

}

func Delete(c *gin.Context) {

	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
