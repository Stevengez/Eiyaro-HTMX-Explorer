package models

type Translation map[string]interface{}

// ** Access a nested map of the translation map
func (d Translation) M(k string) Translation {
	return d[k].(map[string]interface{})
}

// ** Access a string value of the translation map
func (d Translation) S(k string) string {
	return d[k].(string)
}
