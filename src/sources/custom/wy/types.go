package wy

import "lx-source/src/sources"

var (
	// brMap = map[string]string{
	// 	sources.Q_128k: `128000`,
	// 	sources.Q_320k: `320000`,
	// 	sources.Q_flac: `1000000`, //`743625`,`915752`
	// 	sources.Q_fl24: `2000000`, //`1453955`,`1683323`
	// }
	qualityMap = map[string]string{
		sources.Q_128k: `standard`,
		sources.Q_320k: `exhigh`,
		sources.Q_flac: `lossless`,
		sources.Q_fl24: `hires`,

		sources.Q_dolby:  `jyeffect`,
		sources.Q_sky:    sources.Q_sky,
		sources.Q_master: `jymaster`,
	}
	// 优化：返回音质与查询音质相同，完全可以直接比较，不用多一步Reverse
	// qualityMapReverse = map[string]string{
	// 	`standard`: sources.Q_128k,
	// 	`exhigh`:   sources.Q_320k,
	// 	`lossless`: sources.Q_flac,
	// 	`hires`:    sources.Q_fl24,
	// 	`jyeffect`: `dolby`,
	// 	`jysky`:    `sky`,
	// 	`jymaster`: `master`,
	// }
)
