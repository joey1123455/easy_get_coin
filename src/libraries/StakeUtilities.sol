// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.18;

library StakeUtilities {

    // @notice Represent the data used to generate a stake ID
    // @creator - the address of the user stacking the tokens
    // @amount - the amount staked 
    // @stakedAt - the time stacked at
        struct StakeKeyData {
        address creator;
        uint256 amount;
        uint256 stackedAt;
    }

    /// @notice Calculate the key for a staking program
    /// @param key The components used to compute the incentive identifier
    /// @return stakeID The identifier for the incentive
    function compute(StakeKeyData memory key) internal pure returns (bytes32 stakeID) {
        return keccak256(abi.encode(key));
    }

    function calculate(uint256 amount, uint256 bps) public pure returns (uint256) {
        return amount * bps / 10_000;
    }

    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }
}