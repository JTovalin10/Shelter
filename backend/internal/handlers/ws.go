package handlers

// Docs: https://pkg.go.dev/github.com/gorilla/websocket
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
	"shelter/backend/internal/monitor"
)