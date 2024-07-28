Viedo Date: 2023/12/16
Made By:	BIRKHOFF
Date:	2024-07-04

【18-1-MongoDB客户端】
【18-2-插入操作】
【18-3-查询操作】
【18-4-条件查询、分页、排序】
【18-5-更新和删除操作】
package main

import (
	//标准库提供的统一编程接口
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users *mongo.Collection
var client *mongo.Client

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("<%s: %s,%d>", u.ID, u.Name, u.Age)
	// return fmt.Sprintf("<%s: %s,%d>", u.Name, u.Age)
}

func insertOne() {
	tom := User{Name: "xks", Age: 30}
	insertResult, err := users.InsertOne(context.TODO(), tom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
}

func insertMany() {
	u1 := User{Name: "Ben", Age: 20}
	u2 := User{Name: "HEYE", Age: 16}
	insertManyResult, err := users.InsertMany(context.TODO(), []interface{}{&u1, &u2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------")
	fmt.Println(insertManyResult.InsertedIDs...)
}

// 单条查询
func FindOne() {
	//条件
	// filter := bson.D{{"name", "HEYE"}} //Slice
	filter := bson.D{{"name", bson.D{{"$eq", "tom"}}}}
	// filter := bson.M{"name": "tom"} //MAP
	// filter := bson.M{"name": bson.M{"$ne": "jerry"}}
	// filter := bson.D{} //没有条件全部都符合 bson.M
	var u User
	err := users.FindOne(context.TODO(), filter).Decode(&u) //bin -> object
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//说明没有任何匹配文档
			log.Println("没有任何匹配文档")
			return
		}
		log.Fatal(err)
	}
	fmt.Println(u)
}

// 多条查询
func findMany1() {
	filter := bson.M{} //无条件，全部符合
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //关闭游标

	var results []*User
	for cursor.Next(context.TODO()) {
		var u User
		err = cursor.Decode(&u) //bin -> object
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &u) //装入容器
	}
	fmt.Println(results)
}

// 查询多条，成批装入容器
func findMany2() {
	filter := bson.D{} //无条件，全部符合
	var results []*User

	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //关闭游标

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	for i, r := range results {
		fmt.Println(i, r)
	}
}

// 查询条件
// 改造上面的findMany2 函数，可以使用下面表格中不同filter
func findByFilter(filter interface{}) {
	var results []*User
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //关闭游标

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

// 改造函数findByFilter为findAll
func findAll(filter interface{}, opt *options.FindOptions) {
	var results []*User
	cursor, err := users.Find(context.TODO(), filter, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //关闭游标

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

// 投影
func TOUYING() {
	filter := bson.M{"age": bson.M{"$gt": 18}}
	opt := options.Find()
	opt.SetProjection(bson.M{"name": false, "age": false}) //name,age字段不投影，都显示零值
	findAll(filter, opt)
}

func init() {
	var err error
	dsn := "mongodb://localhost:27017/" //客户端可以读取全局环境变量或者配置文件
	opts := options.Client()
	opts.ApplyURI(dsn).SetConnectTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), opts) //localhost:27017 配置服务端指定用户名和密码

	//指定库 use库 //目前没有表
	db := client.Database("GoGo1")
	users = db.Collection("users") //指定集合就是表

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("client: ", client)
}

// 更新
func UpdateOne() {
	// filter := bson.M{"age": bson.M{"$exist": true}} //所有age字段的文档
	filter := bson.M{"name": "tom"}                         //所有age字段的文档
	update := bson.M{"$set": bson.M{"name": "BIRKHOFFXIA"}} //age字段减少5
	ur, err := users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(ur.MatchedCount, ur.ModifiedCount)
}

func disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

// 删除
func deleteOne() {
	filter := bson.M{} //没有条件，匹配所有文档
	dr, err := users.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(dr.DeletedCount)
}
func deleteMany() {
	filter := bson.M{} //没有条件，匹配所有文档
	dr, err := users.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(dr.DeletedCount)
}
func main() {
	// defer disconnect()

	fmt.Println("users: ", users)

	u1 := User{}
	fmt.Printf("%#v\n", u1)

	//插入
	// insertOne()
	// insertMany()

	// 查找单行、多行
	// FindOne()
	// findMany1()

	//过滤 Filter
	// filter := bson.M{"age": bson.M{"$gte": 20}}
	// filter := bson.M{"name": bson.M{"$in": []string{"Tom", "Ben"}}}
	// filter := bson.D{{"name", bson.D{{"$nin", []string{"Tom"}}}}}
	// filter := bson.M{"name": "tom", "age": bson.M{"$eq": 20}} //and
	// filter := bson.M{"$and": []bson.M{{
	// 	"name": bson.M{"$eq": "tom"},
	// }, {
	// 	"age": bson.M{"$gt": 15},
	// }}}
	// filter := bson.M{"age": bson.M{"$not": bson.M{"$gte": 40}}}
	// filter := bson.M{"age": bson.M{"$type": 16}}
	// filter := bson.M{"gender": bson.M{"$exists": false}}
	// findByFilter(filter)

	// filter := bson.M{}
	// findAll(filter, options.Find().SetLimit(2))

	// //排序
	// filter := bson.M{}
	// opt := options.Find()
	// opt.SetSort(bson.M{"age": -1}) //1 (for ascending) or -1 (for descending)
	// //分页
	// opt.SetSkip(3) //offset
	// findAll(filter, opt)

	//投影
	// TOUYING()

	// 更新
	// UpdateOne()

	//删除
	// deleteOne()

	deleteMany()
}


【18-6-上下文Context】
=====================================================================

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)
EG：
package main

import (
	"context"
	"fmt"
)

func main() {
	c1 := context.WithValue(context.Background(), "k1", "v1")
	c2 := context.WithValue(context.Background(), "k2", "v2")
	c3 := context.WithValue(c1, "k3", "v3")
	c4 := context.WithValue(c3, "k4", "v4")

	//返回接口
	//	Deadline() (deadline time.Time, ok bool)
	//	Done() <-chan struct{}
	//	Err() error
	//	Value(key any) any

	names := []string{"k1", "k2", "k3", "k4"}
	for _, name := range names {
		fmt.Printf("%s:%v\t%v\t%v\t%v\n", name, c1.Value(name), c2.Value(name), c3.Value(name), c4.Value(name))
	}
	// k1:v1	 <nil>	 v1	     v1
	// k2:<nil>	 v2	     <nil>	 <nil>
	// k3:<nil>	 <nil>	 v3	     v3
	// k4:<nil>	 <nil>	 <nil>	 v4
}
