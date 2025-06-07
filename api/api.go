package api
import (
	"encoding/json"
	"net/http"
)

// Coin balance params
type CoinBalanceParams struct {
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string
}