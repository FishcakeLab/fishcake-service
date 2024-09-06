package main

import (
	"fmt"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
)

func main() {
	merchantAbi, _ := abi.FishcakeEventManagerMetaData.GetAbi()
	nftTokenAbi, _ := abi.NftManagerMetaData.GetAbi()
	ActivityAdd := merchantAbi.Events["ActivityAdd"].ID.String()
	fmt.Println("ActivityAdd = ", ActivityAdd)

	ActivityFinish := merchantAbi.Events["ActivityFinish"].ID.String()
	fmt.Println("ActivityFinish = ", ActivityFinish)

	Drop := merchantAbi.Events["Drop"].ID.String()
	fmt.Println("Drop = ", Drop)

	CreateNFT := nftTokenAbi.Events["CreateNFT"].ID.String()
	fmt.Println("CreateNFT = ", CreateNFT)

}
