package studentinfo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentInfo struct {
	Grade     string `json:"Grade" bson:"grade"`
	Room      string `json:"Room" bson:"room"`
	Gender    string `json:"Gender" bson:"gender"`
	Name      string `json:"Name" bson:"name"`
	StudentID string `json:"StudentId" bson:"studentid"`
}

type IStudentInfoModel interface {
	GetStudentInfo() []StudentInfo
}

type StudentInfoModel struct {
	Collection *mongo.Collection
}

func NewStudentInfoModel(mongodb *mongo.Database) IStudentInfoModel {
	return &StudentInfoModel{
		Collection: mongodb.Collection("students"),
	}
}

func (repo StudentInfoModel) GetStudentInfo() []StudentInfo {
	var students []StudentInfo
	cursor, err := repo.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return students
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var student StudentInfo
		cursor.Decode(&student)
		students = append(students, student)
	}

	return students
}
