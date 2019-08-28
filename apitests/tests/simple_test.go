package tests

import (
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/insolar/insolar/apitests/apiclientdeprecated"
	// - Big numbers to store signatures.
)

var Wg sync.WaitGroup
var Logger *log.Logger

// TODO move these tests and remove file
// The identifier is to be incremented in every request and each response will contain a corresponding one.
// The transfer request sends an amount of funds to member identified by a reference:
func TestCreateTransferGetBalance(t *testing.T) {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	rootMember := apiclientdeprecated.GetRootMember()
	member1 := apiclientdeprecated.GetNewMember(rootMember)
	member2 := apiclientdeprecated.GetNewMember(rootMember)

	count := 5
	Wg.Add(count * 2)

	for i := 0; i < count; i++ {
		go transfer(member1, member2)
		go transfer(member2, member1)
	}
	Wg.Wait()
	Logger.Println("finished")
}

func transfer(memberFrom apiclientdeprecated.MemberObject, memberTo apiclientdeprecated.MemberObject) {
	defer Wg.Done()
	result := memberFrom.TransferMoney(memberTo, "100")
	Logger.Printf("Transfer from: %v to %v",
		memberFrom.MemberResponse.Result.CallResult.Reference,
		memberTo.MemberResponse.Result.CallResult.Reference)
	Logger.Println("result: " + string(result))
	time.Sleep(100 + time.Millisecond)
}
