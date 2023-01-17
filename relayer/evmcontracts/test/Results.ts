import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("Results", function () {
  
  const ErrCallerIsNotTheOwner = "Ownable: caller is not the owner";

  const ChainlinkTokenAddress = "0x326C977E6efc84E512bB9C30f76E30c160eD06FB";
  const ChainlinkOracleAddress = "0xCC79157eb46F5624204f47AB42b3906cAA40eaB7";
  const ChainlinkJobId = "0x3764613237303266333766643438653562316239613537313565333530396236";
  const ChainlinkFee = 100000000000000000n; // 0.1 LINK
  const endpoint = "https://dev.api.vocdoni.net/elections/";
  
  const electionId = "0x3764613237303266333766643438653562316239613537313565333530396236";
  const organizationId = "0x326C977E6efc84E512bB9C30f76E30c160eD06FB";
  const censusRoot = "0x3764613237303266333766643438653562316239613537313565333530396236";
  const sourceNetworkAddress = "0xCC79157eb46F5624204f47AB42b3906cAA40eaB7";
  const scrutiny = [[1, 2, 3],[4, 5, 6]]

  
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployResultsFixture() {
    // Contracts are deployed using the first signer/account by default
    const [owner, otherAccount] = await ethers.getSigners();

    const Results = await ethers.getContractFactory("Results");
    const results = await Results.deploy(
      ChainlinkTokenAddress,
      ChainlinkOracleAddress,
      ChainlinkJobId,
      10 // 0.1 LINK
    );

    return { results, owner, otherAccount };
  }

  describe("Results contract", function () {
    it("Should deploy with the right constructor parameters", async function () {
      const { results, owner } = await loadFixture(deployResultsFixture);
      
      expect(await results.getChainLinkToken()).to.be.equal(ChainlinkTokenAddress);
      expect(await results.getChainLinkOracle()).to.be.equal(ChainlinkOracleAddress);
      expect(await results.jobId()).to.be.equal(ChainlinkJobId);
      expect(await results.fee()).to.be.equal(ChainlinkFee);
      expect(await results.owner()).to.be.equal(owner.address);
    });

    it("Should set the endpoint only if owner", async function () {
      const { results, owner, otherAccount } = await loadFixture(deployResultsFixture);
      await results.setEndpoint("dev", endpoint, {from: owner.address});
      expect (await results.endpoints("dev")).to.be.equal(endpoint);
      await expect(results.connect(otherAccount).setEndpoint("fake", endpoint)).to.be.revertedWith(ErrCallerIsNotTheOwner);
    });
    
    it("Should set the fee only if owner", async function () {
      const { results, owner, otherAccount } = await loadFixture(deployResultsFixture);
      await results.setFee(100, {from: owner.address});
      expect (await results.fee()).to.be.equal(10000000000000000n);
      await expect(results.connect(otherAccount).setFee(1000)).to.be.revertedWith(ErrCallerIsNotTheOwner);
    });

    it("Should set the link token only if owner", async function () {
      const { results, owner, otherAccount } = await loadFixture(deployResultsFixture);
      await results._setChainlinkToken(ChainlinkOracleAddress, {from: owner.address});
      expect (await results.getChainLinkToken()).to.be.equal(ChainlinkOracleAddress);
      await expect(results.connect(otherAccount)._setChainlinkToken(ChainlinkTokenAddress)).to.be.revertedWith(ErrCallerIsNotTheOwner);
    });

    it("Should set the link oracle only if owner", async function () {
      const { results, owner, otherAccount } = await loadFixture(deployResultsFixture);
      await results._setChainlinkOracle(ChainlinkTokenAddress, {from: owner.address});
      expect (await results.getChainLinkOracle()).to.be.equal(ChainlinkTokenAddress);
      await expect(results.connect(otherAccount)._setChainlinkOracle(ChainlinkOracleAddress)).to.be.revertedWith(ErrCallerIsNotTheOwner);
    });

    it("Should transfer ownership only if owner", async function () {
      const { results, owner, otherAccount } = await loadFixture(deployResultsFixture);
      await results.transferOwnership(otherAccount.address, {from: owner.address});
      expect(await results.owner()).to.be.equal(otherAccount.address);
      await expect(results.connect(owner).transferOwnership(owner.address)).to.be.revertedWith(ErrCallerIsNotTheOwner);
    });

    it("Should encode a result", async function () {
      const { results } = await loadFixture(deployResultsFixture);
      const encodedResult = await results.encodeResult(electionId, organizationId, censusRoot, sourceNetworkAddress, scrutiny);
      const abiEncodedResult = ethers.utils.defaultAbiCoder.encode(
        ["bytes32", "address", "bytes32", "address", "uint256[][]"],
        [electionId, organizationId, censusRoot, sourceNetworkAddress, scrutiny]
      );
      expect(encodedResult).to.be.equal(abiEncodedResult);
    });

    it("Should decode a result", async function () {
      const { results } = await loadFixture(deployResultsFixture);
      const abiEncodedResult = ethers.utils.defaultAbiCoder.encode(
        ["bytes32", "address", "bytes32", "address", "uint256[][]"],
        [electionId, organizationId, censusRoot, sourceNetworkAddress, scrutiny]
      );
      const decodedResult = await results.decodeResult(abiEncodedResult);
      expect(decodedResult.electionId).to.be.equal(electionId);
      expect(decodedResult.organizationId).to.be.equal(organizationId);
      expect(decodedResult.censusRoot).to.be.equal(censusRoot);
      expect(decodedResult.sourceNetworkAddress).to.be.equal(sourceNetworkAddress);
      for (let i = 0; i < decodedResult.scrutiny.length; i++) {
        for (let j = 0; j < decodedResult.scrutiny[i].length; j++) {
          expect(decodedResult.scrutiny[i][j]).to.be.equal(scrutiny[i][j]);
        }
      }
    });
  });
});
