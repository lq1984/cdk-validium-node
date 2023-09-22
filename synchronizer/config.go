package synchronizer

import (
	"github.com/0xPolygon/cdk-validium-node/config/types"
)

// Config represents the configuration of the synchronizer
type Config struct {
	// SyncInterval is the delay interval between reading new rollup information
	// NOTE: 同步定时器
	SyncInterval types.Duration `mapstructure:"SyncInterval"`

	// NOTE: 批量扫描区块日志
	// SyncChunkSize is the number of blocks to sync on each chunk
	SyncChunkSize uint64 `mapstructure:"SyncChunkSize"`
	// TrustedSequencerURL is the rpc url to connect and sync the trusted state

	// TODO: 得看下用了sequencer 什么方法, 查询了下批次号信息
	TrustedSequencerURL string `mapstructure:"TrustedSequencerURL"`
}
