package core

import "github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"

var siteNeedsAccount = map[product.Site]bool{
	product.SiteFinishLine:     false,
	product.SiteJD_Sports:      false,
	product.SiteYeezySupply:    false,
	product.SiteSupreme:        false,
	product.SiteEastbay_US:     false,
	product.SiteChamps_US:      false,
	product.SiteFootaction_US:  false,
	product.SiteFootlocker_US:  false,
	product.SiteBestbuy:        false,
	product.SitePokemon_Center: false,
	product.SitePanini_US:      false,
	product.SiteTopss:          false,
	product.SiteNordstorm:      false,
	product.SiteEnd:            false,
	product.SiteTarget:         true,
	product.SiteAmazon:         false,
	product.SiteSolebox:        false,
	product.SiteOnygo:          false,
	product.SiteSnipes:         false,
	product.SiteSsense:         false,
	product.SiteWalmart:        false,
	product.SiteHibbet:         false,
	product.SiteNewBalance:     false,
}