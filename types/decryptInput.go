package types

// DecryptInput TODO
// 2019/08/16 19:53:47
type DecryptInput struct {
	Capsule     string `json:"capsule"`
	CipherText  string `json:"ciphertext"`
	FileName    string `json:"filename"`
	DecryptMode string `json:"decryptmode"`
	DecKey      string `json:"decryptkey"`
}
