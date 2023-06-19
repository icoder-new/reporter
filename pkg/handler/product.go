package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icoder-new/reporter/utils"
	"github.com/spf13/cast"
)

type productData struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var request productData

	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.CreateProduct(cat_id, request.Name, request.Description, request.Price)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    product,
	})
}

func (h *Handler) GetProducts(c *gin.Context) {
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	products, err := h.service.GetProducts(cat_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (h *Handler) GetProduct(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	product, err := h.service.GetProduct(id, cat_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    product,
	})
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	var request productData

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.UpdateProduct(id, cat_id, request.Name, request.Description, request.Price)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    product,
	})
}

func (h *Handler) UploadPictureProduct(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	image, err := c.FormFile("picture")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := utils.GenFilenameWithDir(image.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusUnsupportedMediaType, err.Error())
		return
	}

	product, err := h.service.UploadPictureProduct(id, cat_id, filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(image, fmt.Sprintf("./files/layouts/%s", filePath)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    product,
	})
}

func (h *Handler) ChangePictureProduct(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	image, err := c.FormFile("picture")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := utils.GenFilenameWithDir(image.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusUnsupportedMediaType, err.Error())
		return
	}

	product, err := h.service.ChangePictureProduct(id, cat_id, filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(image, fmt.Sprintf("./files/layouts/%s", filePath)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    product,
	})
}

// TODO
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	if err := h.service.DeleteProduct(id, cat_id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// TODO
func (h *Handler) RestoreProduct(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	_param := c.Query("category_id")
	if _param == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid query param for 'category_id'")
		return
	}
	cat_id := cast.ToInt(_param)

	if err := h.service.RestoreProduct(id, cat_id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
