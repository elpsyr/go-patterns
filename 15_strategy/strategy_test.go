package strategy

import "testing"

func TestPayment_Pay(t *testing.T) {
	payment := NewPayment("Ada", "id-123", 100, &Cash{})
	payment.Pay()

	payment2 := NewPayment("Ada", "id-123", 100, &Bank{})
	payment2.Pay()
}
