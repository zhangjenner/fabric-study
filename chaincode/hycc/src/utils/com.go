package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	"math/rand"
	"time"
)

//=============================================================================
//公用常量
const PUBCHAN_NAME = ""    //公共通道名称
const PUBCC_NAME = "pubcc" //公共链码名称

//=============================================================================
//获取随机数
func GetRand(bits int, nonce int64) []byte {
	src := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	dst := make([]byte, bits, bits)
	r := rand.New(rand.NewSource(nonce))
	srcLen := len(src)
	for i := 0; i < bits; i++ {
		dst[i] = src[r.Intn(srcLen)]
	}
	return dst
}

//获取提案提案者信息
func GetCreator(stub shim.ChaincodeStubInterface) (mspid string, cert string) {
	creator, err := stub.GetCreator()
	if err != nil {
		panic(Error(err))
	}
	sid := msp.SerializedIdentity{}
	err = proto.Unmarshal(creator, &sid)
	if err != nil {
		panic(Error(err))
	}
	return sid.GetMspid(), string(sid.GetIdBytes())
}

//从证书中获取公钥(PEM格式)
func GetPubKeyFromCert(pemCert string) (pubkey string) {
	block, _ := pem.Decode([]byte(pemCert))
	if block == nil {
		panic(Errorf("Invalid pemCert:%s", pemCert))
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(Error(err))
	}
	PubKeyBytes, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		panic(Error(err))
	}
	nblock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: PubKeyBytes,
	}
	return string(pem.EncodeToMemory(nblock))
}

//效率测试
func GetTime() time.Time {
	return time.Now()
}

func SubTime(msg string, st time.Time) time.Time {
	et := time.Now()
	fmt.Printf("%s:%v\n", msg, et.Sub(st))
	return et
}
