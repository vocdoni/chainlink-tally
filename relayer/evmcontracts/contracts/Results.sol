// SPDX-License-Identifier: AGPL3.0
pragma solidity ^0.8.9;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Results is ChainlinkClient, Ownable {
    
    using Chainlink for Chainlink.Request;

    // string constant for the api response key
    string constant JSONRESPONSEPATH = "abiEncoded";

    /** @notice Event emitted when a new result is requested and the request is fulfilled */
    event RequestFulfilled(
        bytes32 indexed requestId,
        bytes32 indexed electionId,
        bytes indexed rawResults
    );

    /** @notice The job ID to use for the Chainlink request */
    bytes32 public jobId;
    /** @notice The fee to pay the Chainlink oracle */
    uint256 public fee;

    /** @notice The endpoint to use for the Chainlink request (chainName => endpoint) */
    mapping(string => string) public endpoints;
    
    /**
     * @notice Result struct
     * @param electionId id of the election
     * @param organizationId id of the organization
     * @param censusRoot merkle root of the merkle tree containing all elegible voters
     * @param sourceNetworkAddress address of the evm contract linked to this election, usually contains the census
     * i.e An ERC20 token contract contains the census if the election is a token based voting
     * @param scrutiny uint256[][] aggregated results of the election https://blog.aragon.org/vocdoni-ballot-protocol/
     */
    struct Result {
        address organizationId;
        bytes32 censusRoot;
        address sourceNetworkAddress;
        uint256[][] scrutiny;
    }

    /** @notice electionId => Result */
    mapping(bytes32 => Result) public results;

    /**
     * @notice Constructor
     * @param _chainLinkToken address of the chainlink token
     * @param _chainlinkOracle address of the chainlink oracle
     * @param _jobId oracle job id
     * @param _fee uint256 (i.e _fee = 10 -> 0,1 LINK * 10**18) (Varies by network and job)
     * Goerli Testnet details:
     * Link Token: 0x326C977E6efc84E512bB9C30f76E30c160eD06FB
     * Oracle: 0xCC79157eb46F5624204f47AB42b3906cAA40eaB7 (Chainlink DevRel)
     * JobId: 0x3764613237303266333766643438653562316239613537313565333530396236
     * Fee = (1 * LINK_DIVISIBILITY) / 10; // 0,1 LINK
     */
    constructor(
        address _chainLinkToken,
        address _chainlinkOracle,
        bytes32 _jobId,
        uint256 _fee) Ownable() {
        jobId = _jobId;
        fee = (1 * LINK_DIVISIBILITY) / _fee;
        setChainlinkToken(_chainLinkToken);
        setChainlinkOracle(_chainlinkOracle);
    }

    /**
     * @notice Get result for a given election
     * @param _electionId id of the election
     * @return organizationId id of the organization
     * @return censusRoot merkle root of the merkle tree containing all elegible voters
     * @return sourceNetworkAddress address of the evm contract linked to this election, usually contains the census
     * i.e An ERC20 token contract contains the census if the election is a token based voting
     * @return scrutiny uint256[][] aggregated results of the election https://blog.aragon.org/vocdoni-ballot-protocol/
     */
    function getResult(bytes32 _electionId) public view returns (
        address organizationId,
        bytes32 censusRoot,
        address sourceNetworkAddress,
        uint256[][] memory scrutiny)
    {
        organizationId = results[_electionId].organizationId;
        censusRoot = results[_electionId].censusRoot;
        sourceNetworkAddress = results[_electionId].sourceNetworkAddress;
        scrutiny = results[_electionId].scrutiny;
    }

    /**
     * @notice Get the chainlink token address
     * @return address of the chainlink token
     */
    function getChainLinkToken() public view returns (address) {
        return chainlinkTokenAddress();
    }

    /**
     * @notice Get the chainlink oracle address
     * @return address of the chainlink oracle
     */
    function getChainLinkOracle() public view returns (address) {
        return chainlinkOracleAddress();
    }

    /**
     * @notice Set endpoint to get data from
     * @param _name name identifying the endpoint
     * @param _endpoint the endpoint
     */
    function setEndpoint(string memory _name, string memory _endpoint) public onlyOwner {
        endpoints[_name] = _endpoint;
    }

    /**
     * @notice Set fee in LINK (Granularity: 10**18)
     * @param _fee uint256 (i.e _fee = 10 -> 0,1 LINK * 10**18) (Varies by network and job)
     */ 
    function setFee(uint256 _fee) public onlyOwner {
        fee = (1 * LINK_DIVISIBILITY) / _fee;
    }

    /**
     * @notice Set chainlinkToken
     * @param _chainlinkToken address of the chainlink token
     */
    function _setChainlinkToken(address _chainlinkToken) public onlyOwner {
        setChainlinkToken(_chainlinkToken);
    }

    /**
     * @notice Set chainlinkOracle
     * @param _chainlinkOracle address of the chainlink oracle
     */
    function _setChainlinkOracle(address _chainlinkOracle) public onlyOwner {
        setChainlinkOracle(_chainlinkOracle);
    }
    
    /**
     * @notice Set the result for a given election
     * @param _electionId id of the election
     * @param _chainName name of the endpoint to get data from
     * @return requestId id of the request
     */
    function setResult(bytes32 _electionId, string memory _chainName) public returns (bytes32 requestId) {
        // check result not already set for this election by checking if the organizationId is 0x0
        require(results[_electionId].organizationId == address(0), "Result already set for this election");

        Chainlink.Request memory req = buildChainlinkRequest(
            jobId,
            address(this),
            this.fulfillRequest.selector
        );
        
        string memory requestUrl = string(abi.encodePacked(endpoints[_chainName], _electionId, "/scrutiny"));
        req.add(
            "get",
            requestUrl
        );
        // parse JSON response matching key "abiEncoded"
        req.add("path", JSONRESPONSEPATH);

        // example response:
        // {
        //  "abiEncoded": "0x3764613237303266333766643438653562316239613537313565333530396236
        //                 000000000000000000000000cc79157eb46f5624204f47ab42b3906caa40eab737
        //                 646132373032663337666434386535623162396135373135653335303962360000
        //                 00000000000000000000cc79157eb46f5624204f47ab42b3906caa40eab7000000
        //                 00000000000000000000000000000000000000000000000000000000a000000000
        //                 000000000000000000000000000000000000000000000000000000020000000000
        //                 000000000000000000000000000000000000000000000000000040000000000000
        //                 00000000000000000000000000000000000000000000000000c000000000000000
        //                 000000000000000000000000000000000000000000000000030000000000000000
        //                 000000000000000000000000000000000000000000000001000000000000000000
        //                 000000000000000000000000000000000000000000000200000000000000000000
        //                 000000000000000000000000000000000000000000030000000000000000000000
        //                 000000000000000000000000000000000000000003000000000000000000000000
        //                 000000000000000000000000000000000000000400000000000000000000000000
        //                 000000000000000000000000000000000000050000000000000000000000000000
        //                 000000000000000000000000000000000006"
        // }

        // send the request
        return sendChainlinkRequest(req, fee);
    }

    /**
     * @notice Get the result from a given request and stores them in the results mapping
     * @param _requestId bytes32
     * @param _rawResult bytes
     */
    function fulfillRequest(
        bytes32 _requestId,
        bytes memory _rawResult
    ) public recordChainlinkFulfillment(_requestId) {
        Result memory r;
        uint256[][] memory scrutiny;
        bytes32 electionId;
        (electionId, r.organizationId, r.censusRoot, r.sourceNetworkAddress, scrutiny) = decodeResult(_rawResult);
        r.scrutiny = scrutiny;
        results[electionId] = r;
        emit RequestFulfilled(_requestId, electionId, _rawResult);

    }

    /**
     * @notice Encode a result into bytes
     * @param _electionId bytes32
     * @param _organizationId address
     * @param _censusRoot bytes32
     * @param _sourceNetworkAddress address
     * @param _scrutiny uint256[][]
     * @return bytes
     */
    function encodeResult(
        bytes32 _electionId,
        address _organizationId,
        bytes32 _censusRoot,
        address _sourceNetworkAddress,
        uint256[][] memory _scrutiny
    ) public pure returns(bytes memory) {
        return abi.encode(_electionId, _organizationId, _censusRoot, _sourceNetworkAddress, _scrutiny);
    }

    /**
     * @notice Decode an abi encoded result
     * @param _result bytes
     * @return electionId bytes32
     * @return organizationId address
     * @return censusRoot bytes32
     * @return sourceNetworkAddress address
     * @return scrutiny uint256[][]
     */
    function decodeResult(
        bytes memory _result
    ) public pure returns(
        bytes32 electionId,
        address organizationId,
        bytes32 censusRoot,
        address sourceNetworkAddress,
        uint256[][] memory scrutiny
    ) {
        (electionId, organizationId, censusRoot, sourceNetworkAddress, scrutiny) = abi.decode(_result, (bytes32, address, bytes32, address, uint256[][]));
    }

    /**
     * @notice Allow withdraw of Link tokens from the contract
     */
    function withdrawLink() public onlyOwner {
        LinkTokenInterface link = LinkTokenInterface(chainlinkTokenAddress());
        require(
            link.transfer(msg.sender, link.balanceOf(address(this))),
            "Unable to transfer"
        );
    }
}
