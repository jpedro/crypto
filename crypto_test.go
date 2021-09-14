// Swine leberkas venison
//
// Burgdoggen sirloin biltong chuck drumstick shank capicola porchetta. Turkey pork loin
// chuck fatback jowl. T-bone short ribs turducken cupim, brisket cow pork belly leberkas.
// Landjaeger ham hock fatback pig corned beef bresaola beef ribs. Pork pork chop boudin
// strip steak landjaeger, pork belly kevin pork loin capicola ham. Pastrami spare ribs
// porchetta, drumstick leberkas t-bone short loin doner filet mignon hamburger corned
// beef. Venison short loin flank, cupim fatback spare ribs pork loin buffalo turducken
// tail.package main
package crypto

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	text := "test"
	pass := "test"

	encrypted, err := Encrypt(text, pass)
	if err != nil {
		panic(err)
	}

	decrypted, err := Decrypt(encrypted, pass)
	if err != nil {
		panic(err)
	}

	if decrypted != text {
		t.Error("Expected", text, "got", decrypted)
	}
}
