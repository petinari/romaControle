package seguranca

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
	"time"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uuid.UUID, tenantId uuid.UUID) (string, error) {
	exp, _ := strconv.Atoi(os.Getenv("TEMPO_EXP"))
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * time.Duration(exp)).Unix()
	permissoes["usuarioId"] = usuarioID
	permissoes["tenantId"] = tenantId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(os.Getenv("SECRETKEY")))
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(g *gin.Context) error {
	tokenString := ExtrairToken(g)
	token, e := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})
	if e != nil {
		return e
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// ExtrairUsuarioID retorna o usuarioId que está salvo no token

func ExtrairDadosUsuarioToken(g *gin.Context) (usuarioID *uuid.UUID, tenantID *uuid.UUID, erro error) {
	tokenString := ExtrairToken(g)
	token, e := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if e != nil {
		return nil, nil, e
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := uuid.Parse(fmt.Sprintf("%v", permissoes["usuarioId"]))
		if erro != nil {
			return nil, nil, erro
		}
		tenantID, erro := uuid.Parse(fmt.Sprintf("%v", permissoes["tenantId"]))
		if erro != nil {
			return nil, nil, erro
		}

		return &usuarioID, &tenantID, nil
	}
	return nil, nil, errors.New("erro ao extrair dados do token do usuário")

}

func ExtrairDadosUsuarioTokenInvalido(g *gin.Context) (usuarioID *uuid.UUID, tenantID *uuid.UUID, erro error) {
	tokenString := ExtrairToken(g)
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {

		return nil, nil, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok {
		usuarioID, erro := uuid.Parse(fmt.Sprintf("%v", permissoes["usuarioId"]))
		if erro != nil {
			return nil, nil, erro
		}
		tenantID, erro := uuid.Parse(fmt.Sprintf("%v", permissoes["tenantId"]))
		if erro != nil {
			return nil, nil, erro
		}

		return &usuarioID, &tenantID, nil
	}
	return nil, nil, errors.New("erro ao extrair dados do token do usuário")

}

func ExtrairToken(g *gin.Context) string {
	token := g.Request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

/*
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return os.Getenv("SECRETKEY"), nil
}
*/
