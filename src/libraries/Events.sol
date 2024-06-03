// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.18;

library EgcEvents {
       // @notice Emits contract owner set
    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    // @notice Emits token staked
    event TokenStaked(address indexed _owner, bytes32 indexed _programID, bytes32 indexed _stakeID, uint256 _amount);

    // @notice Emits token unstaked
    event TokenUnstaked(address indexed _claimingAccount, bytes32 indexed _programID, bytes32 indexed _stakeID, uint256 _amount);

    // @notice EMits program created
    event ProgramCreated(address indexed _creator, bytes32 indexed _programID, uint duration);
}