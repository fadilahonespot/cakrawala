package cached

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func SetupCache() RedisClient {
	fmt.Println("Connect Redis Client.....")

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),     // Redis server address
		Password: os.Getenv("REDIS_PASSWORD"), // No password for your local setup
		Username: os.Getenv("REDIS_USERNAME"), // Username
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return &cache{client: client}
}

func (w *cache) SetEmailVerification(ctx context.Context,  email, code string) (err error) {
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", email, code)
	err = w.client.Set(ctx, email, code, time.Duration(24)*time.Hour).Err()
	return
}

func (w *cache) GetEmailVerification(ctx context.Context, email string) (code string, err error) {
	fmt.Printf("[CACHED GET] key: %v \n", email)
	err = w.client.Get(ctx, email).Scan(&code)
	return
}

func (w *cache) DeleteEmailVerification(ctx context.Context, email string) (err error) {
	fmt.Printf("[CACHED DEL] key: %v \n", email)
	err = w.client.Del(ctx, email).Err()
	return
}

func (w *cache) GetCity(ctx context.Context, provinceId string) (data string) {
	key := fmt.Sprintf("%v-%v", KeyGetCity, provinceId)
	fmt.Printf("[CACHED GET] key: %v \n", key)
	w.client.Get(ctx, key).Scan(&data)
	return
}

func (w *cache) SetCity(ctx context.Context, provinceId string, value string) (err error) {
	key := fmt.Sprintf("%v-%v", KeyGetCity, provinceId)
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", key, value)
	err = w.client.Set(ctx, key, value, time.Duration(48)*time.Hour).Err()
	return
}

func (w *cache) GetProvince(ctx context.Context) (data string) {
	key := fmt.Sprintf("%v", KeyGetProvince)
	fmt.Printf("[CACHED GET] key: %v \n", key)
	w.client.Get(ctx, key).Scan(&data)
	return
}

func (w *cache) SetProvince(ctx context.Context, value string) (err error) {
	key := fmt.Sprintf("%v", KeyGetProvince)
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", key, value)
	err = w.client.Set(ctx, key, value, time.Duration(48)*time.Hour).Err()
	return
}

func (w *cache) GetCityProvince(ctx context.Context, provinceId, cityId string) (data string) {
	key := fmt.Sprintf("%v-%v-%v", KeyGetCityAndProvince, provinceId, cityId)
	fmt.Printf("[CACHED GET] key: %v \n", key)
	w.client.Get(ctx, key).Scan(&data)
	return
}

func (w *cache) SetCityProvince(ctx context.Context, provinceId, cityId, value string) (err error) {
	key := fmt.Sprintf("%v-%v-%v", KeyGetCityAndProvince, provinceId, cityId)
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", key, value)
	err = w.client.Set(ctx, key, value, time.Duration(48)*time.Hour).Err()
	return
}

func (w *cache) GetEmailNotif(ctx context.Context, externalId string) (email string) {
	key := fmt.Sprintf("%v_%v", KeyNotifEmail, externalId)
	fmt.Printf("[CACHED GET] key: %v \n", key)
	w.client.Get(ctx, key).Scan(&email)
	return
}

func (w *cache) SetEmailNotif(ctx context.Context, externalId, value string) (err error) {
	key := fmt.Sprintf("%v_%v", KeyNotifEmail, externalId)
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", key, value)
	err = w.client.Set(ctx, key, value, time.Duration(24)*time.Hour).Err()
	return
}

func (w *cache) GetGenerateText(ctx context.Context, productId string) (resp GenerateText) {
	key := fmt.Sprintf("%v_%v", KeyGenerateText, productId)
	fmt.Printf("[CACHED GET] key: %v \n", key)
	value := ""
	w.client.Get(ctx, key).Scan(&value)
	if value != "" {
		json.Unmarshal([]byte(value), &resp)
	}

	return
}

func (w *cache) SetGenerateText(ctx context.Context, productId int, value GenerateText) (err error) {
	key := fmt.Sprintf("%v_%v", KeyGenerateText, productId)
	valueByte, _ := json.Marshal(value)
	fmt.Printf("[CACHED SET] key: %v, value: %v \n", key, string(valueByte))
	err = w.client.Set(ctx, key, string(valueByte), time.Duration(24)*time.Hour).Err()
	return
}

