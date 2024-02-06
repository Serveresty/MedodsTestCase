package jwtservice

import "math/rand"

// Генерация сида, который связывает Access и Refresh токены
func GenerateSeed() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	seed := make([]rune, 32)
	for i := range seed {
		seed[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(seed)
}
