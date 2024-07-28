Viedo Date: 2023/12/16
Made By:	BIRKHOFF
Date:	2024-07-04

��18-1-MongoDB�ͻ��ˡ�
��18-2-���������
��18-3-��ѯ������
��18-4-������ѯ����ҳ������
��18-5-���º�ɾ��������
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
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

// ������ѯ
func FindOne() {
	//����
	// filter := bson.D{{"name", "HEYE"}} //Slice
	filter := bson.D{{"name", bson.D{{"$eq", "tom"}}}}
	// filter := bson.M{"name": "tom"} //MAP
	// filter := bson.M{"name": bson.M{"$ne": "jerry"}}
	// filter := bson.D{} //û������ȫ�������� bson.M
	var u User
	err := users.FindOne(context.TODO(), filter).Decode(&u) //bin -> object
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//˵��û���κ�ƥ���ĵ�
			log.Println("û���κ�ƥ���ĵ�")
			return
		}
		log.Fatal(err)
	}
	fmt.Println(u)
}

// ������ѯ
func findMany1() {
	filter := bson.M{} //��������ȫ������
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //�ر��α�

	var results []*User
	for cursor.Next(context.TODO()) {
		var u User
		err = cursor.Decode(&u) //bin -> object
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &u) //װ������
	}
	fmt.Println(results)
}

// ��ѯ����������װ������
func findMany2() {
	filter := bson.D{} //��������ȫ������
	var results []*User

	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //�ر��α�

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	for i, r := range results {
		fmt.Println(i, r)
	}
}

// ��ѯ����
// ���������findMany2 ����������ʹ���������в�ͬfilter
func findByFilter(filter interface{}) {
	var results []*User
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //�ر��α�

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

// ���캯��findByFilterΪfindAll
func findAll(filter interface{}, opt *options.FindOptions) {
	var results []*User
	cursor, err := users.Find(context.TODO(), filter, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //�ر��α�

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

// ͶӰ
func TOUYING() {
	filter := bson.M{"age": bson.M{"$gt": 18}}
	opt := options.Find()
	opt.SetProjection(bson.M{"name": false, "age": false}) //name,age�ֶβ�ͶӰ������ʾ��ֵ
	findAll(filter, opt)
}

func init() {
	var err error
	dsn := "mongodb://localhost:27017/" //�ͻ��˿��Զ�ȡȫ�ֻ����������������ļ�
	opts := options.Client()
	opts.ApplyURI(dsn).SetConnectTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), opts) //localhost:27017 ���÷����ָ���û���������

	//ָ���� use�� //Ŀǰû�б�
	db := client.Database("GoGo1")
	users = db.Collection("users") //ָ�����Ͼ��Ǳ�

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("client: ", client)
}

// ����
func UpdateOne() {
	// filter := bson.M{"age": bson.M{"$exist": true}} //����age�ֶε��ĵ�
	filter := bson.M{"name": "tom"}                         //����age�ֶε��ĵ�
	update := bson.M{"$set": bson.M{"name": "BIRKHOFFXIA"}} //age�ֶμ���5
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

// ɾ��
func deleteOne() {
	filter := bson.M{} //û��������ƥ�������ĵ�
	dr, err := users.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(dr.DeletedCount)
}
func deleteMany() {
	filter := bson.M{} //û��������ƥ�������ĵ�
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

	//����
	// insertOne()
	// insertMany()

	// ���ҵ��С�����
	// FindOne()
	// findMany1()

	//���� Filter
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

	// //����
	// filter := bson.M{}
	// opt := options.Find()
	// opt.SetSort(bson.M{"age": -1}) //1 (for ascending) or -1 (for descending)
	// //��ҳ
	// opt.SetSkip(3) //offset
	// findAll(filter, opt)

	//ͶӰ
	// TOUYING()

	// ����
	// UpdateOne()

	//ɾ��
	// deleteOne()

	deleteMany()
}


��18-6-������Context��
=====================================================================

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)
EG��
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

	//���ؽӿ�
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
