package controllers

import "net/http"

// CriarUsuario insere um usuário no BD
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

// BuscarUsuarios busca todos os usuários no BD
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

// BuscarUsuario busca um usuário no BD
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

// AtualizarUsuario altera as informações de um usuário no BD
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário!"))
}

// DeletarUsuario exclui as informações de um usuário no BD
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
