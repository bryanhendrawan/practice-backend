package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetArticles(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetArticles(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateArticle(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateArticle(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
