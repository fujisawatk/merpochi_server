package repository

// UserRepository userPersistenceの抽象依存
type AuthRepository interface {
	SignIn(string, string) (string, error)
}
