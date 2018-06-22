package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type Auth struct {
}

// Retrieve a token, saves the token, then returns the generated client.
func (a *Auth) 
