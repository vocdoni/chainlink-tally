package main

import (
	"context"

	"github.com/jordipainan/vocdoni-onchain-results/relayer"
	flag "github.com/spf13/pflag"
	"go.vocdoni.io/dvote/log"
)

func main() {

	// config
	relayerCfg := &relayer.RelayerCfg{}
	relayerCfg.LogLevel = *flag.String("logLevel", "info", "logging level")
	relayerCfg.EVMEndpoint = *flag.String("evmEndpoint", "127.0.0.1:8545", "evm endpoint to connect")
	relayerCfg.SignerPrivateKey = *flag.String("signerPrivateKey", "", "signer private key")
	relayerCfg.ResultsContractAddress = *flag.String("resultsContractAddress", "", "results contract address")
	relayerCfg.ElectionID = *flag.String("electionID", "", "election identifier i.e '0xaa0000000000000000000000000000000000000000000000000000000000000ff'")
	relayerCfg.VocdoniChain = *flag.String("vocdoniChain", "azeno", "vocdoni chain identifier i.e 'azeno'")
	flag.Parse()

	// logging
	log.Init(relayerCfg.LogLevel, "stdout")

	// start relayer
	log.Info("Starting relayer ...")
	log.Info("For uploading results to the results contract, remember that you need to have ETH on the signer for paying the execution")
	relayer, err := relayer.NewRelayer(context.Background(), relayerCfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := relayer.ConnectClient(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := relayer.NewResultsContract(context.Background(), relayerCfg.ResultsContractAddress); err != nil {
		log.Fatal(err)
	}

	// relay result
	if err := relayer.RelayResult(context.Background(), relayerCfg.ElectionID, relayerCfg.VocdoniChain); err != nil {
		log.Fatal(err)
	}
}
