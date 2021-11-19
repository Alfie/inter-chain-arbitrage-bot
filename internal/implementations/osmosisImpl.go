package osmosisImpl

import(
 "github.com/cosmos/cosmos-sdk/simapp"
 "github.com/cosmos/cosmos-sdk/testutil/testdata"
 banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
 types "github.com/cosmos/cosmos-sdk/types"
 cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
 "github.com/cosmos/cosmos-sdk/types/tx/signing"
 xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
 tx "github.com/cosmos/cosmos-sdk/client/tx"
)

//TODO: this should return pools
func getPools() {}

//TODO: these two functions likely have to sign and broadcast an ibc-transfer over rpc
func deposit(){}

func Withdraw() (string,error) {
//test data    
priv1, _, addr1 := testdata.KeyTestPubAddr()
priv2, _, addr2 := testdata.KeyTestPubAddr()
priv3, _, addr3 := testdata.KeyTestPubAddr()
    // Choose your codec: Amino or Protobuf. Here, we use Protobuf, given by the
    // following function.
    encCfg := simapp.MakeTestEncodingConfig()

    // Create a new TxBuilder.
    txBuilder := encCfg.TxConfig.NewTxBuilder()

// Define two x/bank MsgSend messages:
    // - from addr1 to addr3,
    // - from addr2 to addr3.
    // This means that the transactions needs two signers: addr1 and addr2.
    msg1 := banktypes.NewMsgSend(addr1, addr3, types.NewCoins(types.NewInt64Coin("atom", 12)))
    msg2 := banktypes.NewMsgSend(addr2, addr3, types.NewCoins(types.NewInt64Coin("atom", 34)))

    err := txBuilder.SetMsgs(msg1, msg2)
    if err != nil {
        return "", err
    }

    //TODO: set these values properly
    txBuilder.SetGasLimit(3)
    txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("atom", 12)))
    txBuilder.SetMemo("test")
    //get current block height and add some value
    txBuilder.SetTimeoutHeight(3)
    
    //TODO: use real accounts and account sequence #s, these are test values
    privs := []cryptotypes.PrivKey{priv1, priv2, priv3}
    accNums:= []uint64{7,7,7} // The accounts' account numbers
    accSeqs:= []uint64{7,7,7} // The accounts' sequence numbers

    // First round: we gather all the signer infos. We use the "set empty
    // signature" hack to do that.
    var sigsV2 []signing.SignatureV2
    for i, priv := range privs {
        sigV2 := signing.SignatureV2{
            PubKey: priv.PubKey(),
            Data: &signing.SingleSignatureData{
                SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
                Signature: nil,
            },
            Sequence: accSeqs[i],
        }

        sigsV2 = append(sigsV2, sigV2)
    }
    err = txBuilder.SetSignatures(sigsV2...)
    if err != nil {
        return "", err
    }

    // Second round: all signer infos are set, so each signer can sign.
    sigsV2 = []signing.SignatureV2{}
    for i, priv := range privs {
        signerData := xauthsigning.SignerData{
            ChainID:       "chain-o0PZgl",
            AccountNumber: accNums[i],
            Sequence:      accSeqs[i],
        }
        sigV2, err := tx.SignWithPrivKey(
            encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
            txBuilder, priv, encCfg.TxConfig, accSeqs[i])
        if err != nil {
            return "", err
        }

        sigsV2 = append(sigsV2, sigV2)
    }
    err = txBuilder.SetSignatures(sigsV2...)
    if err != nil {
        return "", err
    }

   //print something
/*       txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
    if err != nil {
        return err
    } */

    // Generate a JSON string.
    txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
    if err != nil {
        return "", err
    }
    txJSON := string(txJSONBytes)

   return txJSON, nil
}

//TODO: input types most likely should not be ints
func swap(in int, out int, amount int, pools int){ }

