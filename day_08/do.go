package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 設定 MongoDB 連線選項
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 建立 MongoDB 連線
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 檢查是否成功連線到 MongoDB 伺服器
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("無法連線到 MongoDB 伺服器:", err)
	}

	fmt.Println("成功連線到 MongoDB 伺服器")

	// 選擇要操作的資料庫和集合
	database := client.Database("shanetest")
	collection := database.Collection("try")

	// 準備要插入的資料
	data := map[string]interface{}{
		"測試":        "隨便算打看看會怎樣",
		"age":       18,
		"timestamp": time.Now(),
	}

	// 插入資料
	insertResult, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal("插入資料失敗:", err)
	}

	fmt.Println("插入成功，插入的 ID:", insertResult.InsertedID)
}
