# 🔐 Epstein Project

A simple file encryption CLI written in Go.

Epstein encrypts and decrypts **any file type** using modern authenticated encryption. It stores encrypted data inside a custom binary container together with the metadata required for decryption.

> **⚠️ Educational project**
>
> This project was created for learning purposes to explore cryptography, binary file formats, and CLI application architecture in Go.

---

## Features

* 🔒 Encrypt any file
* 🔓 Decrypt encrypted files
* 🛡️ XChaCha20-Poly1305 authenticated encryption
* 🔑 Argon2id password-based key derivation
* 📦 Custom binary container format
* 💻 Cross-platform CLI

---

## Cryptography

| Component      | Algorithm          |
| -------------- | ------------------ |
| Encryption     | XChaCha20-Poly1305 |
| Key Derivation | Argon2id           |
| Randomness     | `crypto/rand`      |

Passwords are **never stored** inside encrypted files.

Every encrypted file contains:

* file format version
* random salt
* random nonce
* original filename
* encrypted payload

---

## Installation

```bash
git clone https://github.com/kyotomin/epstein.git
cd epstein

go build -o epst
```

---

## Usage

### Encrypt

```bash
epst encrypt photo.png
```

Output:

```text
photo.png.epst
```

---

### Decrypt

```bash
epst decrypt photo.png.epst
```

The original filename is restored automatically.

---

## Container Format

```text
+-----------+
| Magic     |
+-----------+
| Version   |
+-----------+
| Salt      |
+-----------+
| Nonce     |
+-----------+
| Name Len  |
+-----------+
| Filename  |
+-----------+
| Ciphertext|
+-----------+
```

---

## Project Structure

```text
.
├── cmd/
├── internal/
│   ├── container/
│   ├── crypto/
│   ├── prompt/
│   ├── service/
│   └── ui/
├── main.go
└── README.md
```

---

## Roadmap

* [x] File encryption
* [x] File decryption
* [x] Custom binary container
* [x] Password-based encryption
* [ ] Streaming encryption
* [ ] Progress bar
* [ ] Configurable Argon2 parameters
* [ ] Unit tests

---

## License

MIT
