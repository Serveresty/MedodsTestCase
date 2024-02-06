package jwtservice

import "math/rand"

func GenerateSeed() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	seed := make([]rune, 32)
	for i := range seed {
		seed[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(seed)
}
