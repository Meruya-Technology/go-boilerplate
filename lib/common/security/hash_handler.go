package security

import "golang.org/x/crypto/bcrypt"

type HashHandler struct {
}

func (j HashHandler) HashPassword(Password string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (j HashHandler) CheckPasswordHash(Password string, Hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Hash), []byte(Password))
	return err == nil
}
