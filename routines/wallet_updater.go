package routines

import (
	"fmt"
	"myapp/models"
	"time"
)

func UpdateWalletsDaily() {
	userIDs, err := models.GetAllUserIDs()
	if err != nil {
		fmt.Println("Error retrieving user IDs:", err)
		return
	}

	for _, userID := range userIDs {
		err := updateWalletForUser(userID)
		if err != nil {
			fmt.Printf("Error updating wallet for user %s: %v\n", userID, err)
			continue
		}
		fmt.Printf("Wallet updated successfully for user %s\n", userID)
	}
}

func updateWalletForUser(userID string) error {
	wallet, err := models.GetWallet(userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve wallet for user %s: %v", userID, err)
	}

	updatedWallet := matchInvestmentsWithExpectedGain(wallet)
	if updatedWallet == nil {
		return fmt.Errorf("failed to match investments with expected gain for user %s", userID)
	}

	err = models.UpdateWallet(userID, updatedWallet.Tickers, updatedWallet.AmountInvested, updatedWallet.ExpectedGain)
	if err != nil {
		return fmt.Errorf("failed to update wallet for user %s: %v", userID, err)
	}

	return nil
}

func matchInvestmentsWithExpectedGain(wallet *models.Wallet) *models.Wallet {
	//TODO: Implement logic to match investments with expected gain
	return wallet
}

func StartDailyUpdateRoutine() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go UpdateWalletsDaily()
		}
	}
}
