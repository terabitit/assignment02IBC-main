package assignment02IBC

import (
	"crypto/sha256"
	"fmt"
)

const miningReward = 100
const rootUser = "Satoshi"

type IndividualBalance struct {
	UserName string
  UserAmount int
}



type BlockData struct {
	Title    string
	Sender   string
	Receiver string
	Amount   int
}
type Block struct {
	Data        []BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {
	tempX := chainHead
	totalBalance := 0
  for tempX != nil{
		for i, _ := range tempX.Data{
			if tempX.Data[i].Sender == userName{
				totalBalance -= tempX.Data[i].Amount
			}else if tempX.Data[i].Receiver == userName{
				totalBalance += tempX.Data[i].Amount
			}
		}

    tempX = tempX.PrevPointer
  }
  return totalBalance
}
func CalculateHash(inputBlock *Block) string {
	Transaction := fmt.Sprintf("%v", inputBlock.Data)
	Result := fmt.Sprintf("%x\n", sha256.Sum256([]byte(Transaction)))
	return Result
}
func VerifyTransaction(transaction *BlockData, chainHead *Block) bool {
	X := CalculateBalance(transaction.Sender,chainHead)
	if (X < transaction.Amount){
		// failed
		return false
	}else{
		// success
		return true
	}
}
func InsertBlock(blockData []BlockData, chainHead *Block) *Block {
/*  for i, _ := range blockData{
        CalculateBalance(blockData[i].Sender, chainHead)
        if VerifyTransaction(&blockData[i], chainHead) == false {
            fmt.Printf("%s Failed Transaction \n", blockData[i].Title)
            return chainHead
        }
    }*/
		var individual_balance []IndividualBalance
    for i, _ := range blockData {
			if VerifyTransaction(&blockData[i], chainHead) == false {
					fmt.Printf("%s Failed Transaction \n", blockData[i].Title)
					return chainHead
			} else {
            search := false
            for j, _ := range individual_balance{
							if individual_balance[j].UserName != blockData[i].Sender {
								search = false
							}else if individual_balance[j].UserName == blockData[i].Sender {
								current_balance := blockData[i].Amount
								current_balance = individual_balance[j].UserAmount - current_balance
								if current_balance > 0{
									search = true
								}else if current_balance < 0 {
                        fmt.Printf("%s - Transaction Failed\n", blockData[i].Title)
                        return chainHead
                    }
                    search = true
                    break
                }
            }
            if search == false {
                c_balance := CalculateBalance(blockData[i].Sender, chainHead) - blockData[i].Amount
                person := IndividualBalance{UserName: blockData[i].Sender, UserAmount: c_balance}
                individual_balance = append(individual_balance, person)
            }

        }
    }





  var tempBlock = new(Block)
  coin_base_transaction := BlockData{Title: "Coinbase", Sender: "System", Receiver: rootUser, Amount: miningReward}
  blockData = append(blockData, coin_base_transaction)
  tempBlock.Data = blockData
  tempBlockHash := CalculateHash(tempBlock)
  tempBlock.CurrentHash = tempBlockHash
  if (chainHead != nil) {
		tempBlock.PrevPointer = chainHead
    tempBlock.PrevHash = chainHead.CurrentHash
  } else if (chainHead == nil){
		tempBlock.PrevPointer = nil
    tempBlock.PrevHash = ""
    chainHead = tempBlock
  }
  return tempBlock
}
func ListBlocks(chainHead *Block) {

	fmt.Printf("\n\n....... Blockchain ......\n")
  tempX := chainHead
  for i := 1; tempX != nil; i++ {
    fmt.Printf("\n\nBlock - %d\n", i)
    for j := 0; j < len(tempX.Data); j++ {
      fmt.Printf(" \nTransaction :- %d \n", j+1)
      fmt.Printf("%s -> %s = %d",tempX.Data[j].Sender,tempX.Data[j].Receiver,tempX.Data[j].Amount)
  }

  tempX = tempX.PrevPointer
    }

}
func VerifyChain(chainHead *Block) {

	tempX := chainHead
	for tempX != nil{
		if tempX.CurrentHash == CalculateHash(tempX){
		}else{
			fmt.Printf("Blockchain is compromised. Not verified")
			return
		}
		tempX = tempX.PrevPointer
	}
	fmt.Printf("Blockchain is not compromised. Blockchain verified")


}
func PremineChain(chainHead *Block, numBlocks int) *Block {
        for i := 0; i < numBlocks; i++ {
            var transaction []BlockData
            chainHead = InsertBlock(transaction, chainHead)
        }
    return chainHead
}
