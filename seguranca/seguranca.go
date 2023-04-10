package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash recebe uma string e coloca um hash nela

func Hash(senha string) ([]byte, error) {
	senhaHash, e := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if e != nil {
		return nil, e
	}
	return senhaHash, nil
}

// VerificarSenha compara uma senha e um hash e retorna se elas são iguais

func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
