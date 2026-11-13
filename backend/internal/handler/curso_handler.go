package handler

import (
	"net/http"
	"strconv"

	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CursoHandler struct {
	service *service.CursoService
}

func NewCursoHandler(service *service.CursoService) *CursoHandler {
	return &CursoHandler{service: service}
}

type cursoResponse struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

func (h *CursoHandler) Listar(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page debe ser un entero mayor que cero",
		})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "size debe ser un entero mayor que cero",
		})
		return
	}

	cursos, total, err := h.service.Listar(page, size, c.Query("q"))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "no se pudieron obtener los cursos",
		})
		return
	}

	respuesta := make([]cursoResponse, 0, len(cursos))
	for _, curso := range cursos {
		respuesta = append(respuesta, cursoResponse{
			ID:     curso.ID,
			Nombre: curso.Titulo,
		})
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	c.JSON(http.StatusOK, respuesta)
}
