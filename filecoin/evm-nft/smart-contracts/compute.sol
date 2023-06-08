// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TaubyteNFTCompute {
    address public owner;
    uint256 public fee;
    uint256 public currentJobId;  // New variable to keep track of the current jobId
    mapping(uint256 => bool) public jobs;

    event ComputeSubmitted(uint256 jobId, string gender, uint256 count, address sender);

    constructor(uint256 _fee) {
        owner = msg.sender;
        fee = _fee;
        currentJobId = 1;  // Initialize currentJobId to 1
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can call this function");
        _;
    }

    function submitCompute(string memory gender, uint256 count) external payable {
        uint256 totalFee = fee * count;
        require(msg.value == totalFee, "Incorrect fee amount");

        jobs[currentJobId] = true;

        emit ComputeSubmitted(currentJobId, gender, count, msg.sender);

        currentJobId++;  // Increment the currentJobId for the next execution
    }

    function withdrawBalance() external onlyOwner {
        payable(owner).transfer(address(this).balance);
    }

    function updateFee(uint256 newFee) external onlyOwner {
        fee = newFee;
    }
}