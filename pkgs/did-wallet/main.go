package main

import (
	"fmt"
	"os"

	"github.com/ewangplay/did-wallet"
)

func main() {
	w, err := wallet.NewFileSystemWallet("./wallet")
	if err != nil {
		fmt.Printf("New filesytem wallet failed: %v", err)
		os.Exit(1)
	}

	r, err := createRawIdentity()
	if err != nil {
		fmt.Printf("Create Raw identity failed: %v", err)
		os.Exit(1)
	}

	err = w.Put("User1", r)
	if err != nil {
		fmt.Printf("Put User1 identity to wallet failed: %v", err)
		os.Exit(1)
	}

	x, err := createX509Identity()
	if err != nil {
		fmt.Printf("Create X509 identity failed: %v", err)
		os.Exit(1)
	}

	err = w.Put("User2", x)
	if err != nil {
		fmt.Printf("Put User2 identity to wallet failed: %v", err)
		os.Exit(1)
	}

	l, err := w.List()
	if err != nil {
		fmt.Printf("Get wallet list failed: %v", err)
		os.Exit(1)
	}

	for _, label := range l {
		id, err := w.Get(label)
		if err != nil || id == nil {
			continue
		}

		fmt.Printf("%s:\n", label)
		fmt.Printf("%v\n", id.Did())
		fmt.Printf("%v\n", id.Type())

		switch id.Type() {
		case wallet.X509IdentityType:
			x, ok := id.(*wallet.X509Identity)
			if ok {
				fmt.Printf("%s\n", x.Certificate())
				fmt.Printf("%s\n\n", x.Key())
			}
		case wallet.RawIdentityType:
			r, ok := id.(*wallet.RawIdentity)
			if ok {
				fmt.Printf("%s\n", r.KeyType())
				fmt.Printf("PublicKeyHex: %s\n", r.PublicKeyHex())
				fmt.Printf("PrivatekeyHex: %s\n\n", r.PrivateKeyHex())
			}
		}
	}
}

func createRawIdentity() (wallet.Identity, error) {
	id := "did:example:3dda540891d14a1baec2c7485c273c00"
	keyID := "keys-1"
	privateKeyHex := "a889f4da49ff8dd6b03d4334723fe3e5ff55ae6a2483de1627bec873b0b73e1e86eabd6abce2f96553251de61def0265784688ff712ce583621a5b181ef21639"
	publicKeyHex := "86eabd6abce2f96553251de61def0265784688ff712ce583621a5b181ef21639"
	key := &wallet.Key{
		ID:            keyID,
		Type:          wallet.Ed25519KeyType,
		PrivateKeyHex: privateKeyHex,
		PublicKeyHex:  publicKeyHex,
	}
	return wallet.NewRawIdentity(id, key)
}

func createX509Identity() (wallet.Identity, error) {
	id := "did:example:4ddajfjsfj3kjf2099uhej485c273c89"
	cert := "-----BEGIN CERTIFICATE-----\nMIIBWjCB3qADAgECAgYBbvXSw4QwDQYJKoZIhvcNAQELBQAwEzERMA8GA1UEAwwI\nSm9obiBEb2UwHhcNMTkxMjEwMTYzNzQwWhcNMTkxMTI2MDA1NTUxWjATMREwDwYD\nVQQDDAhKb2huIERvZTB2MBAGByqGSM49AgEGBSuBBAAiA2IABBIV2OGF/VkRcQTf\n5NjLpQMIW+kc6VmBdpd7+YJ4CrpxtCISiMcDf4LxQ2QdVhkM0FSiYCFLxnDOg8u6\nTm+uKVzlH0HEKkPycoDk784dcvyXiUuWuo6ZHXaCQJfEHNldPzANBgkqhkiG9w0B\nAQsFAANoADBlAjEAoNys0S+/R9/w3bUMwohRN7NuIh2JYmxy3oEafunF4LaNaRd8\ndG9gLBn/7LQZGUu7AjBLQQMV0GPZCNl6JN4TZyxcARxDCmpiuIAzwZuFRYpaAVTO\npJgR6ICTZ0Ko3rz4cT4=\n-----END CERTIFICATE-----\n"
	key := "-----BEGIN PRIVATE KEY-----\nMIG/AgEAMBAGByqGSM49AgEGBSuBBAAiBIGnMIGkAgEBBDAAL3tEAlZDEPZiOxZp\njjGncTzZtLBbtO30tqT+WdTbRqwF9OpGLBAgsbzzo9nhqBagBwYFK4EEACKhZANi\nAAQSFdjhhf1ZEXEE3+TYy6UDCFvpHOlZgXaXe/mCeAq6cbQiEojHA3+C8UNkHVYZ\nDNBUomAhS8ZwzoPLuk5vrilc5R9BxCpD8nKA5O/OHXL8l4lLlrqOmR12gkCXxBzZ\nXT8=\n-----END PRIVATE KEY-----\n"
	return wallet.NewX509Identity(id, cert, key)
}
