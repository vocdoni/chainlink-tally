package relayer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jordipainan/vocdoni-onchain-results/relayer/evmclient"
	results "github.com/jordipainan/vocdoni-onchain-results/relayer/evmclient/contracts"
	vocdoniSigner "go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/log"
)

// Relayer is the main struct of the relayer
type Relayer struct {
	Signer    *vocdoniSigner.SignKeys
	EVMClient *evmclient.Client
}

// RelayerCfg is the configuration for the relayer
type RelayerCfg struct {
	LogLevel               string
	EVMEndpoint            string
	SignerPrivateKey       string
	ResultsContractAddress string
	ElectionID             string
	VocdoniChain           string
}

// NewRelayer creates a new Relayer
func NewRelayer(ctx context.Context, config *RelayerCfg) (*Relayer, error) {
	if config.EVMEndpoint == "" {
		log.Fatal("no evm endpoint provided")
	}
	if config.SignerPrivateKey == "" {
		log.Fatal("no signer private key provided")
	}
	signer := vocdoniSigner.NewSignKeys()
	if err := signer.AddHexKey(config.SignerPrivateKey); err != nil {
		return nil, fmt.Errorf("cannot create signer: %w", err)
	}
	return &Relayer{
		EVMClient: evmclient.NewClient(ctx, config.EVMEndpoint),
		Signer:    signer,
	}, nil
}

// ConnectClient connects to the EVM endpoint
func (r *Relayer) ConnectClient(ctx context.Context) error {
	if err := r.EVMClient.Connect(ctx); err != nil {
		return fmt.Errorf("cannot connect to EVM endpoint: %w", err)
	}
	return nil
}

// NewResultsContract initializes the Results contract
func (r *Relayer) NewResultsContract(ctx context.Context, address string) error {
	if err := r.EVMClient.NewResultsContract(ctx, address); err != nil {
		return fmt.Errorf("cannot initialize results contract: %w", err)
	}
	return nil
}

// RelayResult calls the results smartcontract SetResults(electionId) function
func (r *Relayer) RelayResult(ctx context.Context, electionID, chainName string) error {
	if r.EVMClient == nil {
		return fmt.Errorf("no client available")
	}
	if r.Signer == nil {
		return fmt.Errorf("no signer available")
	}
	contract, err := r.EVMClient.GetContract("results")
	if err != nil {
		return fmt.Errorf("cannot get results contract: %w", err)
	}
	var contractTransactor results.ContractResultsTransactor
	if contract.Instance == nil {
		contractTransactor = contract.Instance.(results.ContractResults).ContractResultsTransactor
	} else {
		return fmt.Errorf("results contract not initialized")
	}
	chainID, err := r.EVMClient.Client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("cannot get chainID: %w", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(&r.Signer.Private, chainID)
	if err != nil {
		return fmt.Errorf("cannot create transaction options: %w", err)
	}
	tx, err := contractTransactor.SetResult(opts, common.HexToHash(electionID), chainName)
	if err != nil {
		return fmt.Errorf("cannot call SetResults: %w", err)
	}
	log.Infof("setResults(%s) transaction with hash %s sent", electionID, tx.Hash().Hex())
	return nil
}
