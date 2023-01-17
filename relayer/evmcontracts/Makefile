.PHONY: clean

# Variables
OUTPUT_FILE=makefile.log

# Target: clean
# Description: Remove the node_modules folder.
clean:
	@echo "Cleaning node_modules folder..."
	@rm -rf node_modules || (echo "Error: Cleaning node_modules folder failed" >> $(OUTPUT_FILE) && exit 1)

# Target: install
# Description: Install all npm packages.
install:
	@echo "Installing npm packages..."
	@yarn install || (echo "Error: Installing npm packages failed" >> $(OUTPUT_FILE) && exit 1)

# Target: compile
# Description: Compile the contracts.
compile:
	@echo "Compiling contracts..."
	@npx hardhat compile || (echo "Error: Compiling contracts failed" >> $(OUTPUT_FILE) && exit 1)

# Target: test
# Description: Run the tests.
test:
	@echo "Running tests..."
	@npx hardhat test || (echo "Error: Running tests failed" >> $(OUTPUT_FILE) && exit 1)

# Target: copy-abi
# Description: Copy the ABI key data of the JSON of the Results contract to the artifacts folder.
copy-abi:
	@echo "Copying ABI data to artifacts folder..."
	@jq -r '.abi' ./artifacts/contracts/Results.sol/Results.json > ./artifacts/contracts/Results.sol/Results.abi || (echo "Error: Copying ABI data failed" >> $(OUTPUT_FILE) && exit 1)

# Target: generate-go-bindings
# Description: Generate go bindings of the Results contract ABI in a file named results.go and storing them in the go-bindings folder.
generate-go-bindings:
	@echo "Generating Go bindings..."
	@mkdir -p ./artifacts/contracts/Results.sol/go-bindings || (echo "Error: Creating go-bindings dir" >> $(OUTPUT_FILE) && exit 1)
	@abigen --abi ./artifacts/contracts/Results.sol/Results.abi --pkg contractResults --out ../../relayer/evmclient/contracts/results.go || (echo "Error: Generating Go bindings failed" >> $(OUTPUT_FILE) && exit 1)

# Target: all
# Description: Run all the targets.
all: clean install compile test copy-abi generate-go-bindings
# Target: bindings
# Description: Run the copy-abi and generate-go-bindings targets to generate golang bindings.
bindings: copy-abi generate-go-bindings
