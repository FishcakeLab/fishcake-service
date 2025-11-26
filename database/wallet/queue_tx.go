package wallet

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"
)

type QueueTx struct {
	ID              string `json:"id" gorm:"column:id;primaryKey;default:replace((uuid_generate_v4())::text, '-'::text, ''::text)"`
	RawTx           string `json:"raw_tx" gorm:"raw_tx"`
	Result          string `json:"result" gorm:"result"`
	TransactionHash string `gorm:"transaction_hash" json:"transactionHash"`
	Status          int8   `json:"status" gorm:"column:status;default:0"`
	Timestamp       uint64 `json:"timestamp" gorm:"timestamp"`
}

type QueueTxDB interface {
	ExistQueueTx(rawTx string) bool
	StoreRawTx(string, string) error
	QueryRawTxInfoByStatus(int8) ([]QueueTx, error)
	MarkedTxToSentOrSuccess([]QueueTx) error
	QueryTxInfoByHash(txHash string) (*QueueTx, error)
}

type queueTxDB struct {
	db *gorm.DB
}

func NewQueueTxDB(db *gorm.DB) QueueTxDB {
	return &queueTxDB{db: db}
}

func (w queueTxDB) StoreRawTx(rawTx string, txHash string) error {
	queueTx := QueueTx{
		RawTx:           rawTx,
		Result:          "receive transaction from frontend",
		TransactionHash: txHash,
		Status:          0,
		Timestamp:       0,
	}
	if err := w.db.Table("queue_tx").Create(&queueTx).Error; err != nil {
		return err
	}
	return nil
}

func (w queueTxDB) QueryRawTxInfoByStatus(status int8) ([]QueueTx, error) {
	var queueTxList []QueueTx
	err := w.db.Table("queue_tx").Where("status = ?", status).Find(&queueTxList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return queueTxList, nil
		}
		log.Error("get unhandled tx fail", "err", err)
		return nil, err
	}
	return queueTxList, nil
}

func (w *queueTxDB) MarkedTxToSentOrSuccess(queueTxList []QueueTx) error {
	for i := 0; i < len(queueTxList); i++ {
		var queueTxItem = QueueTx{}
		result := w.db.Table("queue_tx").Where(&QueueTx{ID: queueTxList[i].ID}).Take(&queueTxItem)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			log.Error("Marked tx status fail", "err", result.Error)
			return result.Error
		}

		log.Info("MarkedTxToSentOrSuccess queue tx status", "Status", queueTxList[i].Status)

		queueTxItem.Status = queueTxList[i].Status
		queueTxItem.Result = queueTxList[i].Result
		queueTxItem.TransactionHash = queueTxList[i].TransactionHash
		err := w.db.Table("queue_tx").Save(queueTxItem).Error
		if err != nil {
			log.Error("Update queue tx status Fail", "err", result.Error)
			return err
		}
	}
	return nil
}

func (w queueTxDB) ExistQueueTx(rawTx string) bool {
	var queueTx QueueTx
	err := w.db.Table("queue_tx").Where("raw_tx = ?", rawTx).Take(&queueTx).Error

	if err == nil {
		return true
	}
	// 查不到时，并不保证返回 ErrRecordNotFound
	// 有些 GORM 配置会 fallback 拿第一条记录
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return false
}

func (w queueTxDB) QueryTxInfoByHash(txHash string) (*QueueTx, error) {
	var queueTx QueueTx
	result := w.db.Table("queue_tx").Where("transaction_hash = ?", txHash).Take(&queueTx)
	if result.Error != nil {
		log.Info("query fail", "err", result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &queueTx, nil
}
