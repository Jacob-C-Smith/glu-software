// Package clause
package internal

// I found this little gem in the CS455-Resources repo. Thanks, Amit!
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
