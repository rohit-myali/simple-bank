package db

import (
	"context"
	"testing"
	"time"

	"github.com/rohit-myali/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, createdAccount1, createdAccount2 Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: createdAccount1.ID,
		ToAccountID:   createdAccount2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createdAccount1 := createRandomAccount(t)
	createdAccount2 := createRandomAccount(t)
	createRandomTransfer(t, createdAccount1, createdAccount2)
}

func TestReadTransfer(t *testing.T) {
	createdAccount1 := createRandomAccount(t)
	createdAccount2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, createdAccount1, createdAccount2)

	readTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, readTransfer)

	require.Equal(t, readTransfer.Amount, transfer.Amount)
	require.Equal(t, readTransfer.ID, transfer.ID)
	require.Equal(t, readTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, readTransfer.ToAccountID, transfer.ToAccountID)
	require.WithinDuration(t, readTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	createdAccount1 := createRandomAccount(t)
	createdAccount2 := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, createdAccount1, createdAccount2)
		createRandomTransfer(t, createdAccount2, createdAccount1)
	}

	arg := ListTransfersParams{
		FromAccountID: createdAccount1.ID,
		ToAccountID:   createdAccount1.ID,
		Limit:         5,
		Offset:        5,
	}
	listTransfer, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listTransfer, 5)

	for _, transfer := range listTransfer {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == createdAccount1.ID || transfer.ToAccountID == createdAccount1.ID)
	}
}
