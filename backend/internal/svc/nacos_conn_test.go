package svc

import (
	// "fmt"
	// "context"
	// "fmt"
	"context"
	"fmt"
	"rank_list/internal/config"
	"testing"
)

func TestNacosConnection(t *testing.T) {
	svc := NewServiceContext(config.Config{})
	
	res, err := svc.Sql.FindOne(context.Background(), 1)
	if err != nil {
		fmt.Printf("Error fetching data: %v", err)
	}
	fmt.Println(res)
}
