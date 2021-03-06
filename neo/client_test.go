package neo_test

import (
	"testing"

	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	nodeURI := selectTestNode()

	t.Run("NewClient()", func(t *testing.T) {
		client := neo.NewClient(nodeURI)
		assert.Equal(t, nodeURI, client.NodeURI)
		assert.IsType(t, neo.Client{}, client)
	})

	t.Run(".GetBestBlockHash()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			blockHash, err := client.GetBestBlockHash()
			assert.NoError(t, err)
			assert.NotEmpty(t, blockHash)
		})
	})

	t.Run(".GetBlockByHash()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testBlock := range testBlocks {
				t.Run(testBlock.id, func(t *testing.T) {
					block, err := client.GetBlockByHash(testBlock.hash)
					assert.NoError(t, err)

					assert.Equal(t, testBlock.hash, block.Hash)
					assert.Equal(t, testBlock.index, block.Index)
					assert.Equal(t, testBlock.merkleRoot, block.Merkleroot)
					assert.Len(t, block.Transactions, testBlock.numberOfTransactions)
				})
			}
		})
	})

	t.Run(".GetBlockByIndex()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testBlock := range testBlocks {
				t.Run(testBlock.id, func(t *testing.T) {
					block, err := client.GetBlockByIndex(testBlock.index)
					assert.NoError(t, err)

					assert.Equal(t, testBlock.index, block.Index)
					assert.Equal(t, testBlock.hash, block.Hash)
					assert.Equal(t, testBlock.merkleRoot, block.Merkleroot)
					assert.Len(t, block.Transactions, testBlock.numberOfTransactions)
				})
			}
		})
	})

	t.Run(".GetBlockCount()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			blockCount, err := client.GetBlockCount()
			assert.NoError(t, err)
			assert.NotEmpty(t, blockCount)
		})
	})

	t.Run(".GetBlockHash()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)
			blockIndex := int64(316675)
			expectedBlockHash := "3f0b498c0d57f73c674a1e28045f5e9a0991f9dac214076fadb5e6bafd546170"

			blockHash, err := client.GetBlockHash(blockIndex)
			assert.NoError(t, err)
			assert.NotEmpty(t, blockHash)
			assert.Equal(t, expectedBlockHash, blockHash)
		})
	})

	t.Run(".GetConnectionCount()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			blockCount, err := client.GetConnectionCount()
			assert.NoError(t, err)
			assert.NotEmpty(t, blockCount)
		})
	})

	t.Run(".GetTransaction()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testTransaction := range testTransactions {
				t.Run(testTransaction.id, func(t *testing.T) {
					transaction, err := client.GetTransaction(testTransaction.hash)
					assert.NoError(t, err)

					assert.Equal(t, testTransaction.hash, transaction.ID)
					assert.Equal(t, testTransaction.size, transaction.Size)
				})
			}
		})
	})

	t.Run(".GetTransactionOutput()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			for _, testTransactionOutput := range testTransactionOutputs {
				t.Run(testTransactionOutput.id, func(t *testing.T) {
					transactionOutput, err := client.GetTransactionOutput(
						testTransactionOutput.hash, testTransactionOutput.index,
					)
					assert.NoError(t, err)

					assert.Equal(t, testTransactionOutput.asset, transactionOutput.Asset)
					assert.Equal(t, testTransactionOutput.value, transactionOutput.Value)
				})
			}
		})
	})

	t.Run(".GetUnconfirmedTransactions()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			_, err := client.GetUnconfirmedTransactions()
			assert.NoError(t, err)
		})
	})

	t.Run(".Ping()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			client := neo.NewClient(nodeURI)

			ok := client.Ping()
			assert.True(t, ok)
		})

		t.Run("SadCase", func(t *testing.T) {
			testCases := []struct {
				description string
				uri         string
			}{
				{
					description: "InvalidURI",
					uri:         ")£*&%(£*&Q",
				},
				{
					description: "OfflineURI",
					uri:         "/foo",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					client := neo.NewClient(testCase.uri)

					ok := client.Ping()
					assert.False(t, ok)
				})
			}
		})
	})
}
