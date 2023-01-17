import { ethers } from "hardhat";

async function main() {
  
  const Results = await ethers.getContractFactory("Results");
  // THIS MUST BE CHANGED AND ADAPTED TO YOUR OWN DEPLOYMENT
  const results = await Results.deploy(
    "0x326C977E6efc84E512bB9C30f76E30c160eD06FB",
    "0xCC79157eb46F5624204f47AB42b3906cAA40eaB7",
    "0x3764613237303266333766643438653562316239613537313565333530396236",
    10 // 0.1 LINK
  );

  await results.deployed();

  console.log(`Results deployed to ${results.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
