package main

import (
	"context"
	"fmt"
	studentinfo "gomongexcelgenerator/core/student-info"
	teacherinfo "gomongexcelgenerator/core/teacher-info"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Starting the server...")
	clientOptions := options.Client().ApplyURI("mongodb+srv://kopracticing:kopracticing@golang-mongo-excel-gene.xkmllym.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	clientDatabase := client.Database("school")

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	studentInfoModels := studentinfo.NewStudentInfoModel(clientDatabase)

	studentInfoController := studentinfo.NewStudentInfoController(studentInfoModels)
	teacherInfoController := teacherinfo.NewTeacherInfoController()

	router.GET("/ping/student-info", studentInfoController.PingStudentInfo)
	router.GET("/ping/teacher-info", teacherInfoController.PingTeacherInfo)

	router.GET("/get/student-info", studentInfoController.GetAllStudentInfo)

	router.GET("/generate/student-info/excel/save-new", studentInfoController.GenerateStudentInfoExcelSaveNew)
	router.GET("/generate/student-info/excel/memory-new", studentInfoController.GenerateStudentInfoExcelMemoryNew)
	router.GET("/generate/student-info/excel/stream-new", studentInfoController.GenerateStudentInfoExcelStreamNew)
	router.GET("/generate/student-info/excel/stream-random", studentInfoController.GenerateStudentInfoExcelStreamRandom)
	router.GET("/generate/teacher-info/excel", teacherInfoController.GenerateTeacherInfoExcel)

	router.Run(":3000")
}
