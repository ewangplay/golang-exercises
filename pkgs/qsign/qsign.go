package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jerray/qsign"
)

const (
	ddo = `
	{
		"@context" : "https://www.w3.org/ns/did/v1",
		"id" : "did:example:fafdecaa29934fde9dcc5adaea8ea82b",
		"version" : 1,
		"publicKey" : [
		{
			"id" : "did:example:fafdecaa29934fde9dcc5adaea8ea82b#keys-1",
			"type" : "Ed25519",
			"publicKeyHex" : "4cd5d192e33f390d7f2a5cc948103a080c52f42fefe2d4d334f86e7ac78e0938"
		},
		{
			"id" : "did:example:fafdecaa29934fde9dcc5adaea8ea82b#keys-2",
			"type" : "Ed25519",
			"publicKeyHex" : "e1df3e6e58d51ba0217137224d6daef2a5d1a5790b6537d6f9830d59639e0826"
		}
		],
		"controller" : "did:example:fafdecaa29934fde9dcc5adaea8ea82b",
		"authentication" : [
		"did:example:fafdecaa29934fde9dcc5adaea8ea82b#keys-1"
		],
		"recovery" : [
		"did:example:fafdecaa29934fde9dcc5adaea8ea82b#keys-2"
		],
		"service" : null,
		"proof" : {
		"type" : "Ed25519",
		"creator" : "did:example:fafdecaa29934fde9dcc5adaea8ea82b#keys-1",
		"signatureValue" : "ak5hxLAke9UQ1ExSyYVu8r1GjAzlScP1DHHefQ9187c3MCBsoR8My5X5ex1IBHxexDYj9yBetIv24yu+Dh2sBg=="
		},
		"created" : "2022-07-06T13:52:26.793567+08:00",
		"updated" : "2022-07-06T13:52:26.793567+08:00"
	}
	`
)

// PublicKey represents publicKey in DID Document,
// used for digital signatures, encryption and other cryptographic operations
type PublicKey struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	PublicKeyHex string `json:"publicKeyHex"`
}

type PublicKeyList []PublicKey

func (kl PublicKeyList) Len() int           { return len(kl) }
func (kl PublicKeyList) Swap(i, j int)      { kl[i], kl[j] = kl[j], kl[i] }
func (kl PublicKeyList) Less(i, j int) bool { return strings.Compare(kl[i].ID, kl[j].ID) == -1 }

func (kl PublicKeyList) MarshalQsign() string {
	sort.Sort(kl)

	s := "["
	for i, k := range kl {
		if i != 0 {
			s += "|"
		}
		s += fmt.Sprintf("id=%s&publicKeyHex=%s&type=%s", k.ID, k.PublicKeyHex, k.Type)
	}
	s += "]"
	return s
}

type StringList []string

func (sl StringList) MarshalQsign() string {
	sort.Strings(sl)

	s := "["
	for i, a := range sl {
		if i != 0 {
			s += "|"
		}
		s += a
	}
	s += "]"
	return s
}

// ServiceType represents service type
type ServiceType string

// Service represent the services provided by DID Subject
type Service struct {
	ID              string      `json:"id"`
	Type            ServiceType `json:"type"`
	ServiceEndpoint string      `json:"serviceEndpoint"`
}

// Proof represents the signature signed by DID Controller
type Proof struct {
	Type           string `json:"type"`
	Creator        string `json:"creator"`
	SignatureValue string `json:"signatureValue"`
}

// DDO represents DID Document
type DDO struct {
	Context        string        `json:"@context" qsign:"-"`
	ID             string        `json:"id"`
	Version        int8          `json:"version"`
	PublicKey      PublicKeyList `json:"publicKey"`
	Controller     string        `json:"controller"`
	Authentication StringList    `json:"authentication"`
	Recovery       StringList    `json:"recovery"`
	Service        []Service     `json:"service" qsign:"-"`
	Proof          Proof         `json:"proof" qsign:"-"`
	Created        time.Time     `json:"created" qsign:"-"`
	Updated        time.Time     `json:"updated" qsign:"-"`
}

func main() {
	var err error
	var document DDO
	err = json.Unmarshal([]byte(ddo), &document)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	q := qsign.NewQsign(qsign.Options{
		SuffixGenerator: func() string {
			return "&key=192006250b4c09247ec02edce69f6a2d"
		},

		// The default filter
		// Filter: func(key, value string) {
		// 	return len(value) > 0
		// },

		// To use a hash.Hash other than md5
		Hasher: func() hash.Hash {
			return sha256.New()
		},

		// To use a encoding other than hex
		Encoder: func() qsign.Encoding {
			return base64.StdEncoding
		},
	})

	digest, err := q.Digest(document)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(digest))

	signature, err := q.Sign(document)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(signature))
}
