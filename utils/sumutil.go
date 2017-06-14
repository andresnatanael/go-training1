/*
sumutils

@Author Andres Natanael Soria <andresnatanael@gmail.com>
*/

package utils

//SumValues only sum the values of V
//If V is divisible by 3 or 5 and V>0
func SumValues(v int) int {
	sum := 0
	for v >= 0 {
		if v%5 == 0 || v%3 == 0 {
			sum = sum + v
		}
		v--
	}
	return sum
}
