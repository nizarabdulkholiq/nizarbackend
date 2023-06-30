package nizarbackend

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTagihanSPP(db string, tagihanspp TagihanSPP) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("tagihanspp").InsertOne(context.TODO(), tagihanspp)
	if err != nil {
		fmt.Printf("InsertTagihanSPP: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTagihanRegistrasi(db string, tagihanregis TagihanRegistrasi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("tagihanregis").InsertOne(context.TODO(), tagihanregis)
	if err != nil {
		fmt.Printf("InsertTagihanRegistrasi: %v\n", err)
	}
	return insertResult.InsertedID
}


func InsertMahasiswa(db *mongo.Database, collect string, NamaMahasiswa string, NIM string) (InsertedID interface{}) {
	var srt Mahasiswa
	srt.NamaMahasiswa = NamaMahasiswa
	srt.NIM = NIM
	return InsertOneDoc(db, collect, srt)
}
func GetTagihanRegistrasi(stats string) (data []TagihanRegistrasi) {
	user := MongoConnect("tagihan").Collection("tagihanregist")
	filter := bson.M{"semester": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetTagihanRegistrasi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetTagihanSPP(stats string) (data []TagihanSPP) {
	user := MongoConnect("tagihan").Collection("tagihanspp")
	filter := bson.M{"semester": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetTagihanSPP :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}