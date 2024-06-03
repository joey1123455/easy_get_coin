// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.18;

library ProgramID {

    // @notice Represent the data used to generate a program ID
    // @stakeDuration - the amount of time the token while be stacked for
    // @creator - the address of the user stacking the tokens
    // @rewardPercentage - the percentage to calculate rewards with 
    // @tokenTicker - the ticker of the token related to the program
    struct ProgramKeyData {
        uint stakeDuration;
        address creator;
        uint rewardPercentage;
        string tokenTicker;
    }


    /// @notice Calculate the key for a staking program
    /// @param key The components used to compute the incentive identifier
    /// @return programID The identifier for the incentive
    function compute(ProgramKeyData memory key) internal pure returns (bytes32 programID) {
        return keccak256(abi.encode(key));
    }
}