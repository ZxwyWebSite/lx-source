package mg

import "lx-source/src/sources"

// const (
// 	q_128k = `PQ`
// 	q_320k = `HQ`
// 	q_flac = `SQ`
// 	q_fl24 = `ZQ`
// )

var (
	qualityMap = map[string]string{
		sources.Q_128k: `PQ`,
		sources.Q_320k: `HQ`,
		sources.Q_flac: `SQ`,
		sources.Q_fl24: `ZQ`,
	}
	// qualityMapReverse = map[string]string{
	// 	q_128k: sources.Q_128k,
	// 	q_320k: sources.Q_320k,
	// 	q_flac: sources.Q_flac,
	// 	q_fl24: sources.Q_fl24,
	// }
	qualitys = map[string]string{
		sources.Q_128k: `1`,
		sources.Q_320k: `2`,
		sources.Q_flac: `3`,
		sources.Q_fl24: `4`,

		sources.Q_master: `5`,
	}
	// qualitysReverse = map[string]string {
	// 	`000009`: sources.Q_128k,
	// 	`020010`: sources.Q_320k,
	// 	`011002`: sources.Q_flac,
	// 	`011005`: sources.Q_fl24,
	// }
	mgheader = map[string]string{
		`Origin`:  `https://m.music.migu.cn`,
		`Referer`: `https://m.music.migu.cn/v4/`,
		`By`:      ``,
		`channel`: ``,
		`Cookie`:  ``,
	}
)
