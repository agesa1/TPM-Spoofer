package main

import (
	"os"

	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

func main() {
	rwc, err := tpmutil.OpenTPM()
	if err != nil {
		os.Exit(1)
	}
	defer rwc.Close()

	tmpl := tpm2.Public{
		Type:       tpm2.AlgRSA,
		NameAlg:    tpm2.AlgSHA256,
		Attributes: tpm2.FlagFixedTPM | tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth | tpm2.FlagRestricted | tpm2.FlagDecrypt,
		RSAParameters: &tpm2.RSAParams{
			Symmetric: &tpm2.SymScheme{
				Alg:     tpm2.AlgAES,
				KeyBits: 128,
				Mode:    tpm2.AlgCFB,
			},
			KeyBits: 2048,
		},
	}

	handle, _, err := tpm2.CreatePrimary(rwc, tpm2.HandleEndorsement, tpm2.PCRSelection{}, "", "", tmpl)
	if err != nil {
		os.Exit(1)
	}
	defer tpm2.FlushContext(rwc, handle)

	err = tpm2.EvictControl(rwc, "", tpm2.HandleOwner, handle, 0x81010001)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}