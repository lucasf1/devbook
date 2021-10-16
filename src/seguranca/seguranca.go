package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e retorna um hash dela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha e um hash e retorna se são iguais
func VerificarSenha(senhaHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}
