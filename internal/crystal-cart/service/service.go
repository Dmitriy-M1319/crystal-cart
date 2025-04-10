package service

import (
	"github.com/Dmitriy-M1319/crystal-cart/internal/crystal-cart/models"
	"github.com/rs/zerolog/log"
)

type CartRepository interface {
	CreateUserCart(email string) error
	DeleteUserCart(email string) error
	AddProductToCart(email string, productId int64, count int64) error
	RemoveProductFromCart(email string, productId int64) error
	GetCartProducts(email string) (models.Products, error)
}

type CartService struct {
	repository CartRepository
}

func NewCartService(repo CartRepository) *CartService {
	return &CartService{repository: repo}
}

func (s *CartService) CreateCart(email string) error {
	err := s.repository.CreateUserCart(email)
	if err != nil {
		// прокинуть трассу
		log.Err(err).Msgf("Failed to create cart for user %s", email)
	}
	return err
}

func (s *CartService) DeleteCart(email string) error {
	err := s.repository.DeleteUserCart(email)
	if err != nil {
		// прокинуть трассу
		log.Err(err).Msgf("Failed to delete cart for user %s", email)
	}
	return err
}

func (s *CartService) AddProduct(email string, productId int64, count int64) error {
	err := s.repository.AddProductToCart(email, productId, count)
	if err != nil {
		// прокинуть трассу
		log.Err(err).Msgf("Failed to add products for user %s", email)
	}
	return err
}

func (s *CartService) RemoveProduct(email string, productId int64) error {
	err := s.repository.RemoveProductFromCart(email, productId)
	if err != nil {
		// прокинуть трассу
		log.Err(err).Msgf("Failed to remove products for user %s", email)
	}
	return err
}
func (s *CartService) GetProducts(email string) (models.Products, error) {
	products, err := s.repository.GetCartProducts(email)
	if err != nil {
		// прокинуть трассу
		log.Err(err).Msgf("Failed to get products from cart for user %s", email)
	}
	return products, err
}
