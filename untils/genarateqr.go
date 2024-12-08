package untils

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

type QRCode struct {
	id int
}

func GenerateQRCode(data int) ([]byte, error) {
	dataStr := fmt.Sprintf("%d", data)
	qr, err := qrcode.Encode(dataStr, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qr, nil
}
