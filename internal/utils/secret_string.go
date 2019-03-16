package utils

// SecretString скрывает значение данных при конвертации в строку
type SecretString string

// String возвращает обфусцированное значение
func (p SecretString) String() string {
	return "******"
}

// Value возвращает не обфусцированное значение
func (p SecretString) Value() string {
	return string(p)
}
