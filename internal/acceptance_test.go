package internal



/*func TestAcceptance(t *testing.T) {
	sdk := GetSdk()

	t.Run("Basic test E2E", func(t *testing.T) {
		record := entity.FromString(randHex(64))
		records := []entity.RecordEntity{record}

		rr, err := sdk.SendRecords(records)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.RecordReceipt{}, rr[0])

		sdk.WaitAnchor(rr[0].Anchor, 5000)

		// Retrieving record proof
		proof, err := sdk.GetProof(records)
		assert.Nil(t, err)
		timestamp, err := sdk.VerifyProof(proof, entity2.BloockChain)
		assert.Greater(t, timestamp, 0)
	})

	t.Run("Test send records invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")
		records := []entity.RecordEntity{record}

		_, err := sdk.SendRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test send records invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.SendRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test send records invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.SendRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test send records empty record input", func(t *testing.T) {
		res, err := sdk.SendRecords([]entity.RecordEntity{})
		assert.Nil(t, err)
		assert.Equal(t, []entity.RecordReceipt{}, res)
	})

	t.Run("Test get records invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")

		records := []entity.RecordEntity{record}

		_, err := sdk.GetRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get records invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get records invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get anchor non existing anchor", func(t *testing.T) {
		_, err := sdk.GetAnchor(666666666666666666)
		assert.NotNil(t, err)
		assert.IsType(t, exception2.HttpRequestException{}, err)
		assert.Equal(t, "HttpClient was not successful: Anchor not found", err.Error())
	})

	t.Run("Test wait anchor non existing anchor", func(t *testing.T) {
		_, err := sdk.WaitAnchor(666666666666666666, 3000)
		assert.NotNil(t, err)
		assert.IsType(t, exception3.WaitAnchorTimeoutException{}, err)
		assert.Equal(t, "Timeout exceeded while waiting for anchor", err.Error())
	})

	t.Run("Test get proof invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")
		records := []entity.RecordEntity{record}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get proof invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get proof invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get proof none existing leaf", func(t *testing.T) {
		record := entity.FromHash("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")

		records := []entity.RecordEntity{record}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception2.HttpRequestException{}, err)
		assert.Equal(t, "HttpClient was not successful: Message '0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef' not found.", err.Error())
	})

	t.Run("Test verify records invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")
		records := []entity.RecordEntity{record}

		_, err := sdk.VerifyRecords(records, entity2.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test verify records invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.VerifyRecords(records, entity2.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test verify records invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.VerifyRecords(records, entity2.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test verify records none existing leaf", func(t *testing.T) {
		record := entity.FromHash("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")

		records := []entity.RecordEntity{record}

		_, err := sdk.VerifyRecords(records, entity2.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception2.HttpRequestException{}, err)
		assert.Equal(t, "HttpClient was not successful: Message '0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef' not found.", err.Error())
	})
}

func randHex(length int) string {
	maxlen := 8
	min := math.Pow(16, math.Min(float64(length), float64(maxlen)) - 1)
	max := math.Pow(16, math.Min(float64(length), float64(maxlen))) - 1
	n := int((rand.Float64() * (max - min + 1)) + min)
	r := strconv.Itoa(n)
	for len(r) < length {
		r += randHex(length - maxlen)
	}
	return r
}*/
