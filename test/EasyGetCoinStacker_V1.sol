// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.18;

import "forge-std/Test.sol";
import "../src/EasyGetCoinStacker_V1.sol";

contract TokenStackerTest is Test {
    TokenStacker tokenStacker;

    function setUp() public {
        tokenStacker = new TokenStacker();
    }

    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }

    function testListProgramIDs() public {
        // Create a program first to have something to list
        bytes32 programId = tokenStacker.createStakeProgram(30, 100, "TEST");

        // Retrieve and verify the list of program IDs
        bytes32[] memory programIds = tokenStacker.listProgramIDs();
        assertTrue(findIndex(programIds, programId), "Created program ID should be in the list");
    }

    function testListProgramsID() public {
        // Create a program first to have something to retrieve by ticker
        tokenStacker.createStakeProgram(30, 100, "TEST");
        tokenStacker.createStakeProgram(50, 60, "NOT_TEST");
        tokenStacker.createStakeProgram(10, 1200, "TEST");

        bytes32[] memory programs = tokenStacker.listProgramIDs();
        assertTrue(programs.length == 3, "There should be three programs for the ticker");
    }

    function testGetProgramByID() public {
        // Create a program first to have something to retrieve
        bytes32 programId = tokenStacker.createStakeProgram(30, 100, "TEST");

        // Retrieve the program by ID and verify its properties
        TokenStacker.Program memory program = tokenStacker.getProgramByID(programId);
        bool strMatch = compareStrings(program.tokenTicker, "TEST");
        
        assertTrue(program.unclaimedTokens == 0, "Unclaimed tokens should initially be 0");
        assertTrue(program.stakeDuration == 30, "Stake duration should match");
        assertTrue(program.rewardPercentage == 100, "Reward percentage should match");
        assertTrue(strMatch == true, "Token ticker should match");
    }

    function testGetProgramsByTicker() public {
        // Create a program first to have something to retrieve by ticker
        tokenStacker.createStakeProgram(30, 100, "TEST");
        tokenStacker.createStakeProgram(50, 60, "NOT_TEST");
        tokenStacker.createStakeProgram(10, 1200, "TEST");

        // Retrieve the programs by ticker and verify the count
        TokenStacker.Program[] memory testPrograms = tokenStacker.getProgramsByTicker("TEST");
        TokenStacker.Program[] memory notTestPrograms = tokenStacker.getProgramsByTicker("NOT_TEST");
        assertTrue(testPrograms.length == 2, "There should be two programs for the ticker");
        assertTrue(notTestPrograms.length == 1, "There should be a single program for the ticker");
    }

    // Helper function to find an element in an array
    function findIndex(bytes32[] memory arr, bytes32 val) private pure returns (bool) {
        for (uint i = 0; i < arr.length; i++) {
            if (arr[i] == val) {
                return true;
            }
        }
        return false;
    }
}
