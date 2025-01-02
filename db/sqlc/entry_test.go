package db

import (
	"context"
	"testing"
	"time"

	"github.com/rohit-myali/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, createdAccount Account) Entry {

	arg := CreateEntryParams{
		AccountID: createdAccount.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotEmpty(t, entry.CreatedAt)
	require.NotEmpty(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createdAccount := createRandomAccount(t)
	createRandomEntry(t, createdAccount)
}

func TestGetEntry(t *testing.T) {
	createdAccount := createRandomAccount(t)
	createdEntry := createRandomEntry(t, createdAccount)

	getEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntry)

	require.Equal(t, createdEntry.AccountID, getEntry.AccountID)
	require.Equal(t, createdEntry.Amount, getEntry.Amount)

	require.WithinDuration(t, createdEntry.CreatedAt, getEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	createAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, createAccount)
	}

	arg := ListEntriesParams{
		AccountID: createAccount.ID,
		Limit:     5,
		Offset:    5,
	}
	listEntries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listEntries, 5)

	for _, entry := range listEntries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
