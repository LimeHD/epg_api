package channel

import "database/sql"

type PlaylistCountry struct {
	ShowAfterPurchase sql.NullInt32
}
