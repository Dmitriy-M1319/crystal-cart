package db

import (
	"context"

	"slices"

	"github.com/Dmitriy-M1319/crystal-cart/internal/crystal-cart/models"
	"github.com/redis/go-redis/v9"
)

type RedisCartStorage struct {
	client *redis.Client
}

func NewRedisCartStorage(c *redis.Client) *RedisCartStorage {
	return &RedisCartStorage{client: c}
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

func (s *RedisCartStorage) AddProductToCart(email string, productId int64, count int64) error {
	cart, err := s.GetCartProducts(email)
	if err != nil {
		return nil
	}

	cart = append(cart, models.CartProductInfo{ProductID: productId, Count: count})
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
	for i := range cart {
		if cart[i].ProductID == productId {
			newCart := slices.Delete(cart, i, i+1)
			err = s.client.Set(ctx, email, newCart, 0).Err()
			return err
		}
	}
	return nil
}

func (s *RedisCartStorage) GetCartProducts(email string) (models.Products, error) {
	ctx := context.Background()
	var cart models.Products
	err := s.client.Get(ctx, email).Scan(&cart)
	return cart, err
}
