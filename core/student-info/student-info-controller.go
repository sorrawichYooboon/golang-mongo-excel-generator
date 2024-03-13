package studentinfo

import (
	"fmt"
	"gomongexcelgenerator/constants"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type IStudentInfoController interface {
	PingStudentInfo(c *gin.Context)
	GetAllStudentInfo(c *gin.Context)
	GenerateStudentInfoExcelSaveNew(c *gin.Context)
	GenerateStudentInfoExcelMemoryNew(c *gin.Context)
	GenerateStudentInfoExcelStreamNew(c *gin.Context)
	GenerateStudentInfoExcelStreamRandom(c *gin.Context)
}

type StudentInfoController struct {
	StudentInfoModel IStudentInfoModel
}

func NewStudentInfoController(studentInfoModel IStudentInfoModel) IStudentInfoController {
	return &StudentInfoController{
		StudentInfoModel: studentInfoModel,
	}
}

func (controller StudentInfoController) PingStudentInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from student info!",
	})
}

func (controller StudentInfoController) GetAllStudentInfo(c *gin.Context) {
	students := controller.StudentInfoModel.GetStudentInfo()
	c.JSON(200, gin.H{
		"students": students,
	})
}

func (controller StudentInfoController) GenerateStudentInfoExcelSaveNew(c *gin.Context) {
	students := controller.StudentInfoModel.GetStudentInfo()

	// Create a new Excel file
	f := excelize.NewFile()

	// Add a new worksheet named "Students"
	// f.NewSheet("Students")

	// Define headers
	headers := []string{"Grade", "Room", "Gender", "Name", "StudentID"}

	// Write headers to the first row (index starts from 1)
	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 65+i, 1)
		f.SetCellValue(constants.SheetOne, cell, header)
	}

	// Write student data starting from the second row (index 2)
	row := 2
	for _, student := range students {
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("A%d", row), student.Grade)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("B%d", row), student.Room)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("C%d", row), student.Gender)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("D%d", row), student.Name)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("E%d", row), student.StudentID)
		row++
	}

	// Save the Excel file on local storage
	err := f.SaveAs("temp/students-info-save-new.xlsx")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Send the Excel file as a response
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=students-info-save-new.xlsx")
	c.File("temp/students-info-save-new.xlsx")
}

func (controller StudentInfoController) GenerateStudentInfoExcelMemoryNew(c *gin.Context) {
	students := controller.StudentInfoModel.GetStudentInfo()

	// Create a new Excel file
	f := excelize.NewFile()

	// Add a new worksheet named "Students"
	// f.NewSheet("Students")

	// Define headers
	headers := []string{"Grade", "Room", "Gender", "Name", "StudentID"}

	// Write headers to the first row (index starts from 1)
	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 65+i, 1)
		f.SetCellValue(constants.SheetOne, cell, header)
	}

	// Write student data starting from the second row (index 2)
	row := 2
	for _, student := range students {
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("A%d", row), student.Grade)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("B%d", row), student.Room)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("C%d", row), student.Gender)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("D%d", row), student.Name)
		f.SetCellValue(constants.SheetOne, fmt.Sprintf("E%d", row), student.StudentID)
		row++
	}

	// Save the Excel file in memory
	excelData, err := f.WriteToBuffer()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Send the Excel file as a response and set filename
	fileName := "students-info-memory-new.xlsx" // Adjust as needed
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", excelData.Bytes())
}

func (controller StudentInfoController) GenerateStudentInfoExcelStreamNew(c *gin.Context) {
	f := excelize.NewFile()

	sw, err := f.NewStreamWriter(constants.SheetOne)
	if err != nil {
		fmt.Println(err)
		return
	}

	students := controller.StudentInfoModel.GetStudentInfo()

	headers := []string{"Grade", "Room", "Gender", "Name", "StudentID"}
	headersInterface := make([]interface{}, len(headers))
	for i, header := range headers {
		headersInterface[i] = header
	}

	if err := sw.SetRow("A1", headersInterface); err != nil {
		fmt.Println(err)
		return
	}

	for rowID, student := range students {
		row := []interface{}{student.Grade, student.Room, student.Grade, student.Name, student.StudentID}
		cell, err := excelize.CoordinatesToCellName(1, rowID+2)
		if err != nil {
			fmt.Println(err)
			break
		}
		if err := sw.SetRow(cell, row); err != nil {
			fmt.Println(err)
			break
		}
	}
	if err := sw.Flush(); err != nil {
		fmt.Println(err)
		return
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=students-info-stream-new.xlsx")
	f.Write(c.Writer)
	c.Writer.Flush()
}

func (controller StudentInfoController) GenerateStudentInfoExcelStreamRandom(c *gin.Context) {
	f := excelize.NewFile()

	sw, err := f.NewStreamWriter(constants.SheetOne)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rowID := 2; rowID <= 102400; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, err := excelize.CoordinatesToCellName(1, rowID)
		if err != nil {
			fmt.Println(err)
			break
		}
		if err := sw.SetRow(cell, row); err != nil {
			fmt.Println(err)
			break
		}
	}
	if err := sw.Flush(); err != nil {
		fmt.Println(err)
		return
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=students-info-stream-new.xlsx")
	f.Write(c.Writer)
	c.Writer.Flush()
}
