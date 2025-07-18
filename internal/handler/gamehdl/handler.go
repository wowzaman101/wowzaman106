package gamehdl

import "github.com/gofiber/fiber/v3"

type handler struct {
}

type Handler interface {
	Play(ctx fiber.Ctx) error
}

func New() Handler {
	return &handler{}
}

type Request struct {
	PlayHands []Hand `json:"playHands"`
	GameType  int    `json:"gameType"` // 1, 2
}

type Response []Play

func (h *handler) Play(ctx fiber.Ctx) error {
	var req Request
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response := make(Response, len(req.PlayHands))
	for i, hand := range req.PlayHands {
		response[i] = PlayCard(hand)
	}

	// Example handler logic
	// This could be replaced with actual game logic
	return ctx.JSON(response)
}

func PlayCard(hands Hand) Play {
	if hands.Sum() < 6 {
		return PlayHit
	}

	return PlayStand
}
