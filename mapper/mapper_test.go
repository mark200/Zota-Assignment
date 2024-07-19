package mapper

import (
	"testing"

	"example/zota_assignment/model"
)

func TestGenerateSignature(t *testing.T) {
	order := model.Order{
		OrderAmount: 0,
	}

	expectedSignature := "007fcb36d445dcbc22f29223477d7f3f85876830ee27b50d89b72e0c515cbaa1"

	signature := generateSignature(order)

	if signature != expectedSignature {
		t.Errorf("Expected signature %s, but got %s", expectedSignature, signature)
	}
}
