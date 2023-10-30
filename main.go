package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// https://www.heise.de/hintergrund/Historische-Kryptografie-Vigenere-Chiffre-in-Python-programmiert-9339405.html?seite=2
// var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var alphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {

	encryptFlag := flag.String("encrypt", "", "zu verschlüsselnder Text.")
	keyFlag := flag.String("key", "", "zeichen zum ver- und entschlüsseln.")
	decryptFlag := flag.String("decrypt", "", "zu entschlüsselnder Text")

	var key *string
	var encrypted *string
	var decrypted *string

	flag.Parse()

	if len(*keyFlag) != 0 {
		l := len(*keyFlag)
		key = removeIllegalChars(keyFlag)

		if len(*key) == 0 || l != len(*key) {
			fmt.Printf("Key ist ungültig!\nGültige Zeichen: %s\n", string(alphabet))
			flag.PrintDefaults()
			os.Exit(1)
		}
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(*encryptFlag) != 0 {
		encrypted = removeIllegalChars(encryptFlag)
		encrypted := encrypt(encrypted, key)
		fmt.Println(*encrypted)
		os.Exit(0)
	}
	if len(*decryptFlag) != 0 {
		decrypted = removeIllegalChars(decryptFlag)
		decrypted = decrypt(decrypted, key)
		fmt.Println(*decrypted)
		os.Exit(0)
	}
	flag.PrintDefaults()
	os.Exit(1)

}

func removeIllegalChars(text *string) *string {
	internAlphabet := string(alphabet)
	aText := strings.ToUpper(*text)

	newString := strings.Builder{}
	for _, c := range aText {
		if strings.ContainsRune(internAlphabet, c) {
			newString.WriteRune(c)
		}
	}
	s := newString.String()
	return &s
}

func encrypt(toEncrypt *string, key *string) *string {
	toEncryptRunes := []rune(*toEncrypt)
	keyRunes := []rune(strings.ToUpper(*key))
	keyLen := len(keyRunes)
	result := strings.Builder{}

	for i, c := range toEncryptRunes {

		// Rune an der Position i in key
		kv := keyRunes[i%keyLen]
		// wert der position keyvalue in alphabet
		keyValue := strings.IndexRune(string(alphabet), kv)
		// position der Rune c im string alphabet
		textValue := strings.IndexRune(string(alphabet), c)
		v := (textValue + keyValue) % len(alphabet)
		//fmt.Printf("keyValue: %d, textValue: %d, zusammen: %d", keyValue, textValue, v)
		result.WriteRune(rune(alphabet[v]))
	}
	s := result.String()
	return &s
}

func decrypt(toDecrypt *string, key *string) *string {
	internAlphabet := string(alphabet)
	toDecryptRunes := []rune(*toDecrypt)
	keyRunes := []rune(strings.ToUpper(*key))
	keyLen := len(keyRunes)
	result := strings.Builder{}

	for i, c := range toDecryptRunes {

		// Rune an der Position i in key
		kv := keyRunes[i%keyLen]
		// wert der position kv in alphabet
		keyValue := strings.IndexRune(internAlphabet, kv)
		// position der Rune c im string alphabet
		textValue := strings.IndexRune(internAlphabet, c)
		v := (textValue - keyValue) % len(internAlphabet)
		if v < 0 {
			v = v * -1
		}
		//fmt.Printf("keyValue: %d, textValue: %d, zusammen: %d", keyValue, textValue, v)
		result.WriteRune(rune(alphabet[v]))
	}
	s := result.String()
	fmt.Println("decrypted: ", s)
	return &s
}
