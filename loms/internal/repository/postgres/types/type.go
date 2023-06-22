package types

import "fmt"

var (
	ErrItemNotFound       = fmt.Errorf("item not found")
	ErrFailedBuildQuery   = fmt.Errorf("failed to build query")
	ErrFailedExecuteQuery = fmt.Errorf("failed to execute query")
	ErrFailedToScanRow    = fmt.Errorf("failed to scan row")
)
