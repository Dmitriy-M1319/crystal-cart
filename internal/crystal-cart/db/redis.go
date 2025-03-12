package db

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisCartStorage struct {
	client *redis.Client
}

func NewRedisCartStorage(c *redis.Client) *RedisCartStorage {
	return &RedisCartStorage{client: c}
}

type Products []int64

func (p Products) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p Products) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}

func (s *RedisCartStorage) CreateUserCart(email string) error {
	ctx := context.Background()
	err := s.client.Set(ctx, email, "", 0).Err()
	return err
}

func (s *RedisCartStorage) DeleteUserCart(email string) error {
	ctx := context.Background()
	err := s.client.Del(ctx, email).Err()
	return err
}

func (s *RedisCartStorage) AddProductToCart(email string, productId int64) error {
	cart, err := s.GetCartProducts(email)
	if err != nil {
		return nil
	}

	cart = append(cart, productId)
	ctx := context.Background()
	err = s.client.Set(ctx, email, cart, 0).Err()
	return err
}

func (s *RedisCartStorage) RemoveProductFromCart(email string, productId int64) error {
	cart, err := s.GetCartProducts(email)
	if err != nil {
		return nil
	}

	ctx := context.Background()
	for i := 0; i < len(cart); i++ {
		if cart[i] == productId {
			newCart := append(cart[:i], cart[i+1:]...)
			err = s.client.Set(ctx, email, newCart, 0).Err()
			return err
		}
	}
	return nil
}

func (s *RedisCartStorage) GetCartProducts(email string) ([]int64, error) {
	ctx := context.Background()
	var cart Products
	err := s.client.Get(ctx, email).Scan(&cart)
	return cart, err
}
