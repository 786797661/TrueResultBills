package biz

import "github.com/google/wire"

// ProviderSet is unit providers.
var ProviderSet = wire.NewSet(NewBillsResultBiz)
