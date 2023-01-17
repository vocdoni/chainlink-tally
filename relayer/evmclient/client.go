package evmclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	results "github.com/jordipainan/vocdoni-onchain-results/relayer/evmclient/contracts"
)

// Client is a wrapper around the ethclient.Client
type Client struct {
	url string
	ethclient.Client
	contracts map[string]*EVMContract
}

// NewClient creates a new Client
func NewClient(ctx context.Context, url string) *Client {
	return &Client{url: url, contracts: make(map[string]*EVMContract)}
}

// newEVMContract initializes the EVM contract
func (c *Client) newEVMContract(ctx context.Context,
	name, _abi, ENSDomain string,
	bytecode []byte,
	address common.Address) error {
	c.contracts[name] = &EVMContract{
		ENSDomain: ENSDomain,
		Bytecode:  bytecode,
		Address:   address,
	}
	if err := c.contracts[name].SetABI(_abi); err != nil {
		return fmt.Errorf("couldn't set contract %s ABI: %w", name, err)
	}
	switch name {
	case "results":
		results, err := results.NewContractResults(c.contracts[name].Address, &c.Client)
		if err != nil {
			return fmt.Errorf("cannot instanciate contract: %w", err)
		}
		c.contracts[name].Instance = results
	}
	return nil
}

// NewResultsContract initializes the results contract
func (c *Client) NewResultsContract(ctx context.Context, contractAddress string) error {
	if err := c.newEVMContract(ctx,
		"results",
		results.ContractResultsABI,
		"",
		[]byte{},
		common.HexToAddress(contractAddress),
	); err != nil {
		return fmt.Errorf("cannot initialize results contract: %w", err)
	}
	return nil
}

// GetContract returns a contract given its name
func (c *Client) GetContract(name string) (*EVMContract, error) {
	if c.contracts[name] == nil {
		return nil, fmt.Errorf("contract %s not found", name)
	}
	return c.contracts[name], nil
}

// Connect connects to the EVM endpoint
func (c *Client) Connect(ctx context.Context) error {
	client, err := ethclient.DialContext(ctx, c.url)
	if err != nil {
		return fmt.Errorf("cannot connect to EVM endpoint: %w", err)
	}
	if _, err := c.NetworkID(ctx); err != nil {
		return fmt.Errorf("cannot connect to EVM endpoint: %w", err)
	}
	c.Client = *client
	return nil
}

// CheckConnection checks if the client is connected to the EVM endpoint
func (c *Client) CheckConnection(ctx context.Context) error {
	if _, err := c.NetworkID(ctx); err != nil {
		return fmt.Errorf("cannot connect to EVM endpoint: %w", err)
	}
	return nil
}

// CloseConnection closes the connection to the EVM endpoint
func (c *Client) CloseConnection() {
	c.Close()
}

// EVMContract wraps an EVM smart contract
type EVMContract struct {
	// ABI is the contract ABI
	ABI abi.ABI
	// Bytecode is the contract bytecode
	Bytecode []byte
	// Domain is the ENS domain of the contract
	ENSDomain string
	// Address is the contract address
	Address common.Address
	// Instance is the contract instance
	Instance interface{}
}

// SetABI sets an EVM contract ABI
func (ec *EVMContract) SetABI(_abi string) error {
	var err error
	if ec.ABI, err = abi.JSON(strings.NewReader(_abi)); err != nil {
		return fmt.Errorf("cannot read results contract ABI: %w", err)
	}
	return nil
}
