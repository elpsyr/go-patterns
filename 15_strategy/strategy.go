package strategy

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardId,
			Money:  money,
		},
		strategy: strategy}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)

}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type Cash struct {
}

func (c *Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct {
}

func (b *Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account", ctx.Money, ctx.Name)
}
