package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit btc into wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBitcoin(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}

		_ = wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(0)

		assertBitcoin(t, got, want)
	})

	t.Run("Withdraw, insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(5)}

		err := wallet.Withdraw(Bitcoin(10))

		if err == nil {
			t.Error("expected error, but didn't get one")
		}
	})
}

func assertBitcoin(t *testing.T, got, want Bitcoin) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
