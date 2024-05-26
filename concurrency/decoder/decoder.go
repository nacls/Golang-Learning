package decoder

func StartDecipher(senderChan chan string, decipherer func(encrypted string) string) chan string {
	outputChan := make(chan string, 5)

	go func() {
		for {
			encryptedData := <-senderChan
			decryptedData := decipherer(encryptedData)
			outputChan <- decryptedData
		}
	}()

	return outputChan
}
