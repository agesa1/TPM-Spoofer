# TPM-Spoofer

A Windows tool for spoofing the TPM 2.0 Endorsement Key (EK) by modifying persistent handles.  
Built using the `go-tpm (legacy)` library.

This tool performs operations similar to the Linux `tpm2_evictcontrol` command.

> ⚠️ WARNING:
> - This tool performs low-level TPM operations.
> - A **TPM CLEAR is REQUIRED before spoofing**.

---

## Features

- Manage TPM 2.0 persistent handles
- Modify Endorsement Key (EK) handle
- Lightweight and fast (written in Go)
- Windows compatible

---

## Requirements

- Go 1.19 or higher
- TPM 2.0 enabled in BIOS:
  - Intel: PTT
  - AMD: fTPM
- Administrator privileges

---

## ⚠️ REQUIRED: TPM CLEAR (BEFORE SPOOFING)

You **MUST clear the TPM before using this tool**.

### Step 1: Windows Security

1. Open **PowerShell**
2. Type **Clear-TPM**
3. Run as admin tpm-spoofer.exe
4. Check your TPM serials

https://www.youtube.com/watch?v=AUCJlpRXFKA

---

## Installation & Build

```cmd
go mod init tpm-spoof

go get github.com/google/go-tpm/legacy/tpm2

go get github.com/google/go-tpm/tpmutil

go mod tidy

go build -o tpm-spoof.exe
