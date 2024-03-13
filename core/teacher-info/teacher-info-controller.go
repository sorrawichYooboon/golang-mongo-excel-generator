package teacherinfo

import "github.com/gin-gonic/gin"

type ITeacherInfoController interface {
	PingTeacherInfo(c *gin.Context)
	GenerateTeacherInfoExcel(c *gin.Context)
}

type TeacherInfoController struct {
}

func NewTeacherInfoController() ITeacherInfoController {
	return &TeacherInfoController{}
}

func (controller TeacherInfoController) PingTeacherInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from teacher info!",
	})
}

func (controller TeacherInfoController) GenerateTeacherInfoExcel(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from teacher info!",
	})
}
