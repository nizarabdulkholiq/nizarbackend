package nizarbackend

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertTagihanRegistrasi(t *testing.T) {
	dbname := "tagihan"
	tagihanregis := TagihanRegistrasi{
		ID:      primitive.NewObjectID(),
		NamaMahasiswa: "Nizar",
 		NIM:           "nizar@gmail.com",
 		Semester:      "dua",
 		BiayaRegistrasi:      "10000000",
 	}
 	insertedID := InsertTagihanRegistrasi(dbname, tagihanregis)
 	if insertedID == nil {
 		t.Error("Failed to insert tagihan")
 	}
 }
 func TestInsertTagihanSPP(t *testing.T) {
 	dbname := "tagihan"
 	tagihanspp:= TagihanSPP{
		ID:      primitive.NewObjectID(),
 		NamaMahasiswa: "Nizar",
 		NIM:           "nizar@gmail.com",
 		Semester:      "dua",
 		BiayaSPP:      "10000000",
 	}
 	insertedID := InsertTagihanSPP(dbname, tagihanspp)
 	if insertedID == nil {
 		t.Error("Failed to insert tagihan")
 	}
 }

//  func InsertTagihanSPP(db string, tagihanspp TagihanSPP) (insertedID interface{}) {
//  	insertResult, err := MongoConnect(db).Collection("tagihanspp").InsertOne(context.TODO(), tagihanspp)
//  	if err != nil {
//  		fmt.Printf("InsertTagihanSPP: %v\n", err)
//  	}
//  	return insertResult.InsertedID
//  }
//  func InsertTagihanRegistrasi(db string, tagihanregis TagihanRegistrasi) (insertedID interface{}) {
//  	insertResult, err := MongoConnect(db).Collection("tagihanregis").InsertOne(context.TODO(), tagihanregis)
//  	if err != nil {
//  		fmt.Printf("InsertTagihanRegistrasi: %v\n", err)
//  	}
//  	return insertResult.InsertedID
//  }

 //func GetTagihanRegistrasi(stats string) (data []TagihanRegistrasi) {
 //	user := MongoConnect("tagihan").Collection("tagihanregist")
// 	filter := bson.M{"semester": stats}
// 	cursor, err := user.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetTagihanRegistrasi :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }
// func GetTagihanSPP(stats string) (data []TagihanSPP) {
// 	user := MongoConnect("tagihan").Collection("tagihanspp")
// 	filter := bson.M{"semester": stats}
// 	cursor, err := user.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetTagihanSPP :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }
// func InsertMahaTag(db *mongo.Database, collect string, NamaMahasiswa string, NIM string) (InsertedID interface{}) {
// 	var srt Mahasiswa
// 	srt.NamaMahasiswa = NamaMahasiswa
// 	srt.NIM = NIM
// 	return InsertOneDoc(db, collect, srt)
// }

func TestGetTagihanRegistrasi(t *testing.T) {
	stats := "biaya_registrasi"
	data := GetTagihanRegistrasi(stats)
	fmt.Println(data)
}
func TestGetTagihanSPP(t *testing.T) {
	stats := "biaya_spp"
	data := GetTagihanSPP(stats)
	fmt.Println(data)
}


