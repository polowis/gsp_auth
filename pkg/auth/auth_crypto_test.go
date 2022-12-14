package auth

import (
	"testing"
)

func TestSimpleEncryptCoordinate(t *testing.T) {
	key := "123456"
	plaintext := "12_34;50_23;"
	ciphertext := EncryptCoordinates(plaintext, key)
	expectedCipher := "MDAwMDVmMDcwMTNiMDQwMjVmMDYwNjNi"
	if ciphertext != expectedCipher {
		t.Fatalf("base64 not equal!")
	}
}

func TestSimpleDecryptCoordinate(t *testing.T) {
	key := "123456"
	plaintext := "12_34;50_23;"
	ciphertext := EncryptCoordinates(plaintext, key)
	expectedCipher := "MDAwMDVmMDcwMTNiMDQwMjVmMDYwNjNi"
	if ciphertext != expectedCipher {
		t.Fatalf("base64 not equal!")
	}
	decryptedText := DecryptCoordinates(ciphertext, key)
	if decryptedText != plaintext {
		t.Fatalf("decrypted text not equal!")
	}
}

func TestSimpleDecryptLongCoordinate(t *testing.T) {
	key := "123456"
	plaintext := "12_34;50_23;13_90;125_212;120_222;124_212;124_212;195_212;125_218;125_232;30_21"
	ciphertext := EncryptCoordinates(plaintext, key)
	expectedCipher := "MDAwMDVmMDcwMTNiMDQwMjVmMDYwNjNiMDAwMTVmMGQwNTNiMDAwMDA2NWYwNzA3MDMzYjAyMDYwNTVmMDMwMDAxM2IwNDA0MDU1ZjAxMDUwNzNiMDAwMDA3NWYwNzA3MDMzYjAyMGQwMDVmMDMwMzAxM2IwNDA0MDQ1ZjAxMDUwZDNiMDAwMDA2NWYwNzA1MDMzYjAwMDQ1ZjA0MDA="
	if ciphertext != expectedCipher {
		t.Fatalf("base64 not equal!")
	}
	decryptedText := DecryptCoordinates(ciphertext, key)
	if decryptedText != plaintext {
		t.Fatalf("decrypted text not equal!")
	}
}

func TestDecryptionLongCoordinateWithWrongKey(t *testing.T) {
	key := "123456"
	plaintext := "12_34;50_23;13_90;125_212;120_222;124_212;124_212;195_212;125_218;125_232;30_21"
	ciphertext := EncryptCoordinates(plaintext, key)
	expectedCipher := "MDAwMDVmMDcwMTNiMDQwMjVmMDYwNjNiMDAwMTVmMGQwNTNiMDAwMDA2NWYwNzA3MDMzYjAyMDYwNTVmMDMwMDAxM2IwNDA0MDU1ZjAxMDUwNzNiMDAwMDA3NWYwNzA3MDMzYjAyMGQwMDVmMDMwMzAxM2IwNDA0MDQ1ZjAxMDUwZDNiMDAwMDA2NWYwNzA1MDMzYjAwMDQ1ZjA0MDA="
	if ciphertext != expectedCipher {
		t.Fatalf("base64 not equal!")
	}
	wrongKey := "678902"
	expectedWrong := "67_621;25_636;66_45;6762_755;58635_579;463_9607;6763_755;5840_549;462_96061;6762_775;861_66"
	decryptedText := DecryptCoordinates(ciphertext, wrongKey)
	if decryptedText != expectedWrong {
		t.Fatalf("wrong hash")
	}
}
