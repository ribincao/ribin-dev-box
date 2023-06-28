package utils

import gonanoid "github.com/matoous/go-nanoid"

func GenRoomCode() string {
	return gonanoid.MustGenerate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 9)
}
