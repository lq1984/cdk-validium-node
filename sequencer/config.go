package sequencer

import (
	"github.com/0xPolygon/cdk-validium-node/config/types"
)

// Config represents the configuration of a sequencer
type Config struct {
	// WaitPeriodPoolIsEmpty is the time the sequencer waits until
	// trying to add new txs to the state
	WaitPeriodPoolIsEmpty types.Duration `mapstructure:"WaitPeriodPoolIsEmpty"`

	// BlocksAmountForTxsToBeDeleted is blocks amount after which txs will be deleted from the pool
	// NOTE: 交易超过多少个区块就会从池子中移除
	BlocksAmountForTxsToBeDeleted uint64 `mapstructure:"BlocksAmountForTxsToBeDeleted"`

	// FrequencyToCheckTxsForDelete is frequency with which txs will be checked for deleting
	// NOTE: 间隔时间去检测 交易是否超过指定数量的区块 ref: BlocksAmountForTxsToBeDeleted
	FrequencyToCheckTxsForDelete types.Duration `mapstructure:"FrequencyToCheckTxsForDelete"`

	// MaxTxsPerBatch is the maximum amount of transactions in the batch
	// NOTE: 每个批次最大的交易数量
	MaxTxsPerBatch uint64 `mapstructure:"MaxTxsPerBatch"`

	// MaxBatchBytesSize is the maximum batch size in bytes
	// (subtracted bits of all types.Sequence fields excluding BatchL2Data from MaxTxSizeForL1)
	// NOTE: 每一个批次最大的Bytes大小
	MaxBatchBytesSize uint64 `mapstructure:"MaxBatchBytesSize"`

	// MaxCumulativeGasUsed is max gas amount used by batch
	// NOTE: 每一个批次最大的
	MaxCumulativeGasUsed uint64 `mapstructure:"MaxCumulativeGasUsed"`

	// MaxKeccakHashes is max keccak hashes used by batch
	// NOTE: 每一个批次最多的keccak hash次数
	MaxKeccakHashes uint32 `mapstructure:"MaxKeccakHashes"`

	// MaxPoseidonHashes is max poseidon hashes batch can handle
	// NOTE: 每一个批次最大的 poseidon hash次数
	MaxPoseidonHashes uint32 `mapstructure:"MaxPoseidonHashes"`

	// MaxPoseidonPaddings is max poseidon paddings batch can handle
	// TODO: 波塞冬对齐？ 和 zkevm相关
	MaxPoseidonPaddings uint32 `mapstructure:"MaxPoseidonPaddings"`

	// MaxMemAligns is max mem aligns batch can handle
	// TODO: 内存对齐 和 zkevm相关
	MaxMemAligns uint32 `mapstructure:"MaxMemAligns"`

	// MaxArithmetics is max arithmetics batch can handle
	// TODO: 最大的运算次数？ 和 zkevm相关
	MaxArithmetics uint32 `mapstructure:"MaxArithmetics"`

	// MaxBinaries is max binaries batch can handle
	// TODO: 最大的二进制 ? 和 zkevm相关
	MaxBinaries uint32 `mapstructure:"MaxBinaries"`

	// MaxSteps is max steps batch can handle
	// TODO: 最大的步数？ 和 zkevm相关
	MaxSteps uint32 `mapstructure:"MaxSteps"`

	// TxLifetimeCheckTimeout is the time the sequencer waits to check txs lifetime
	// NOTE: 交易超时检测周期
	TxLifetimeCheckTimeout types.Duration `mapstructure:"TxLifetimeCheckTimeout"`

	// MaxTxLifetime is the time a tx can be in the sequencer/worker memory
	// NOTE: 交易超时时间
	MaxTxLifetime types.Duration `mapstructure:"MaxTxLifetime"`

	// Finalizer's specific config properties
	Finalizer FinalizerCfg `mapstructure:"Finalizer"`

	// DBManager's specific config properties
	DBManager DBManagerCfg `mapstructure:"DBManager"`

	// EffectiveGasPrice is the config for the gas price
	EffectiveGasPrice EffectiveGasPriceCfg `mapstructure:"EffectiveGasPrice"`
}

// FinalizerCfg contains the finalizer's configuration properties
type FinalizerCfg struct {
	// GERDeadlineTimeout is the time the finalizer waits after receiving closing signal to update Global Exit Root
	GERDeadlineTimeout types.Duration `mapstructure:"GERDeadlineTimeout"`

	// ForcedBatchDeadlineTimeout is the time the finalizer waits after receiving closing signal to process Forced Batches
	ForcedBatchDeadlineTimeout types.Duration `mapstructure:"ForcedBatchDeadlineTimeout"`

	// SleepDuration is the time the finalizer sleeps between each iteration, if there are no transactions to be processed
	SleepDuration types.Duration `mapstructure:"SleepDuration"`

	// ResourcePercentageToCloseBatch is the percentage window of the resource left out for the batch to be closed
	// NOTE: 当前批次占用资源百分比就会被close
	ResourcePercentageToCloseBatch uint32 `mapstructure:"ResourcePercentageToCloseBatch"`

	// GERFinalityNumberOfBlocks is number of blocks to consider GER final
	// NOTE: 全局状态敲定区块，也就是需要处理layer1的分叉
	GERFinalityNumberOfBlocks uint64 `mapstructure:"GERFinalityNumberOfBlocks"`

	// ClosingSignalsManagerWaitForCheckingL1Timeout is used by the closing signals manager to wait for its operation
	ClosingSignalsManagerWaitForCheckingL1Timeout types.Duration `mapstructure:"ClosingSignalsManagerWaitForCheckingL1Timeout"`

	// ClosingSignalsManagerWaitForCheckingGER is used by the closing signals manager to wait for its operation
	ClosingSignalsManagerWaitForCheckingGER types.Duration `mapstructure:"ClosingSignalsManagerWaitForCheckingGER"`

	// ClosingSignalsManagerWaitForCheckingL1Timeout is used by the closing signals manager to wait for its operation
	ClosingSignalsManagerWaitForCheckingForcedBatches types.Duration `mapstructure:"ClosingSignalsManagerWaitForCheckingForcedBatches"`

	// ForcedBatchesFinalityNumberOfBlocks is number of blocks to consider GER final
	// NOTE: forced batch交易敲定区块数量，也就是需要处理Layer1的分叉，因为这个forced tx 是发生在layer1上的， 所以需要处理分叉
	ForcedBatchesFinalityNumberOfBlocks uint64 `mapstructure:"ForcedBatchesFinalityNumberOfBlocks"`

	// TimestampResolution is the resolution of the timestamp used to close a batch
	// NOTE: 假设该配置是10s, 如果10秒内批次有交易，那么就立马打包
	TimestampResolution types.Duration `mapstructure:"TimestampResolution"`

	// StopSequencerOnBatchNum specifies the batch number where the Sequencer will stop to process more transactions and generate new batches. The Sequencer will halt after it closes the batch equal to this number
	// NOTE: 到指定batch number就停止，应该是测试用的
	StopSequencerOnBatchNum uint64 `mapstructure:"StopSequencerOnBatchNum"`

	// SequentialReprocessFullBatch indicates if the reprocess of a closed batch (sanity check) must be done in a
	// sequential way (instead than in parallel)
	SequentialReprocessFullBatch bool `mapstructure:"SequentialReprocessFullBatch"`
}

// DBManagerCfg contains the DBManager's configuration properties
type DBManagerCfg struct {
	PoolRetrievalInterval    types.Duration `mapstructure:"PoolRetrievalInterval"`
	L2ReorgRetrievalInterval types.Duration `mapstructure:"L2ReorgRetrievalInterval"`
}

// EffectiveGasPriceCfg contains the configuration properties for the effective gas price
type EffectiveGasPriceCfg struct {
	// MaxBreakEvenGasPriceDeviationPercentage is the max allowed deviation percentage BreakEvenGasPrice on re-calculation
	MaxBreakEvenGasPriceDeviationPercentage uint64 `mapstructure:"MaxBreakEvenGasPriceDeviationPercentage"`

	// L1GasPriceFactor is the percentage of the L1 gas price that will be used as the L2 min gas price
	L1GasPriceFactor float64 `mapstructure:"L1GasPriceFactor"`

	// ByteGasCost is the gas cost per byte
	ByteGasCost uint64 `mapstructure:"ByteGasCost"`

	// MarginFactor is the margin factor percentage to be added to the L2 min gas price
	MarginFactor float64 `mapstructure:"MarginFactor"`

	// Enabled is a flag to enable/disable the effective gas price
	Enabled bool `mapstructure:"Enabled"`

	// DefaultMinGasPriceAllowed is the default min gas price to suggest
	// This value is assigned from [Pool].DefaultMinGasPriceAllowed
	DefaultMinGasPriceAllowed uint64
}
