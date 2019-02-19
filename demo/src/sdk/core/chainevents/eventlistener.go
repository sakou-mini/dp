package chainevents

import (
    rlog "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"../ethereum/events"
	"time"
)

type ContractInfo struct {
    Address string
    Abi string
    EventNames []string
}

func ListenEvent(conn *ethclient.Client, contracts []ContractInfo,
            fromBlock uint64, interval time.Duration,
            dataChannel chan events.Event, errorChannel chan error) bool {
	rv := true
	rlog.Info("start listening events...")

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("Failed to listen event. error:", err)
			rv = false
		}
	}()

    if len(contracts) == 0 {
        rlog.Error("invalid contracts parameter")
        return false
    }

    builder := events.NewScanBuilder()
    for _, v := range contracts {
        builder.SetContract(common.HexToAddress(v.Address), v.Abi, v.EventNames...)
    }

	recp, err := builder.SetClient(conn).
		SetFrom(fromBlock).
		SetTo(0).
		SetGracefullExit(true).
		SetDataChan(dataChannel, errorChannel).
		SetInterval(interval).
		BuildAndRun()
	if err != nil {
		rlog.Error("failed to listen to events.", err)
		return false
	}

	recp.WaitChan()

	return rv
}
