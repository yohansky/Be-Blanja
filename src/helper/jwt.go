package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(secretKey, email string, role string) (string, error) {
	// Membuat token baru dengan metode penandatanganan HS256
	//Header(berisi jenis algoritma dan tipe tokennya)
	token := jwt.New(jwt.SigningMethodHS256)
	//payload = berisi informasi 

	// Mendapatkan klaim dari token sebagai MapClaims
	claims := token.Claims.(jwt.MapClaims)

	// Menambahkan klaim email ke token
	claims["email"] = email
	// claims["role"] // nanti ditambahkan dan di taruh disini
	//parsing lewat object	
	claims["role"] = role
	
	// Menambahkan waktu kedaluwarsa 1 jam
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Menandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// Mengembalikan token yang sudah ditandatangani
	return tokenString, err
}
//gimana carnaya login,cara loginnya harus ngedetect 