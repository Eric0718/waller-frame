package utils

// If 模拟三元运算
func TernaryOperations(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
