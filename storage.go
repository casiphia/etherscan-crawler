package main

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Storage 后端存储接口
type Storage struct {
	Database string
	URI      string
	Client   *mongo.Client
	DB       *mongo.Database
	visited  *mongo.Collection
	cookies  *mongo.Collection
}

func NewStorage() *Storage {
	return &Storage{
		Database: Database,
		URI:      fmt.Sprintf("mongodb://%s:%d", Ip, Port),
	}
}

// Init 初始化mongoDB数据库
func (s *Storage) Init() error {

	var err error
	if s.Client, err = mongo.NewClient(options.Client().ApplyURI(s.URI)); err != nil {
		return err
	}
	if err = s.Client.Connect(context.Background()); err != nil {
		return err
	}

	s.DB = s.Client.Database(s.Database)
	s.visited = s.DB.Collection("opensea_visited")
	s.cookies = s.DB.Collection("opensea_cookies")

	return nil

}

// Visited 请求记录
func (s *Storage) Visited(requestID uint64) error {

	_, err := s.visited.InsertOne(context.Background(), bsonx.MDoc{
		"requestID": bsonx.String(strconv.FormatUint(requestID, 10)),
		"visited":   bsonx.Boolean(true),
	})
	return err
}

// IsVisited 是否请求记录
func (s *Storage) IsVisited(requestID uint64) (bool, error) {

	result := bsonx.MDoc{}
	err := s.visited.FindOne(nil, bsonx.MDoc{
		"requestID": bsonx.String(strconv.FormatUint(requestID, 10)),
	}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Cookies 请求缓存
func (s *Storage) Cookies(u *url.URL) string {

	result := bsonx.MDoc{}
	if err := s.cookies.FindOne(nil, bsonx.MDoc{
		"host": bsonx.String(u.Host),
	}).Decode(&result); err != nil {
		if err != mongo.ErrNoDocuments {
			logrus.Errorf("mongoDB ErrNoDocuments %v", err)
		}
		return ""
	}
	return result["cookies"].String()
}

// SetCookies 设置缓存
func (s *Storage) SetCookies(u *url.URL, cookies string) {

	if _, err := s.cookies.InsertOne(nil, bsonx.MDoc{
		"host":    bsonx.String(u.Host),
		"cookies": bsonx.String(cookies),
	}); err != nil {
	}
}
