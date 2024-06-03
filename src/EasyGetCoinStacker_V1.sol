// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.18;

import './libraries/ProgramID.sol';
import './libraries/StakeUtilities.sol';
import './libraries/Events.sol';
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TokenStacker {

    // @notice The wallet address deploying the contract
    address immutable owner;

    // @notice The address for egc tokens 
    address public egcTokenAdd;

    // @notice Represents a stacking reward program
    // @notice The ID of the token is gotten from the hash of the stake duration, base percentage and stake duration
    // @id - the stakes id
    // @unclaimedTokens - total of tokens in the staken pool
    // @numberOfStakes - total no of coins stacked
    // @stakeDuration - the amount of time the token while be stacked for
    // @creator - the address of the user stacking the tokens
    // @rewardPercentage - the percentage to calculate rewards with 
    // @tokenTicker - the ticker of the token related to the program
    // @tickerProgramIDX - the index of the program in the list of programs for a paticular ticker
    struct Program {
        bytes32 id;
        uint256 unclaimedTokens;
        uint96 stakesHistory;
        uint stakeDuration;
        address creator;
        uint256 rewardPercentage;
        string tokenTicker;
        uint tickerProgramIDX;
    }

    // @notice Represents a single stake
    // @notice The stake ID is gotten from the hash of the owner address, amount staked and time staked at
    // @id - the stakes id
    // @owner - the address that stacked the token
    // @stakedAt - block time staked
    // @claimAt - stake maturity date
    // @amountStaked - token sent
    // @rewardsToClaim - total to claim
    // @programID - the program being staked
    // @rewardWallet - wallet to recieve token
    // @tokenTicker - token being stacked
    // @cliamed - represents if a token has been claimed
    // @tickerStakeIDX - the index of the stake in the list of stakes for a paticular ticker
    // @addressStakeIDX - the index of the stake in the list of stakes for a paticular address
    struct Stake {
        bytes32 id; 
        address owner;
        uint256 stakedAt;
        uint256 claimAt;
        uint256 amountStacked;
        uint256 rewardsToClaim;
        bytes32 programID;
        address rewardWallet;
        string tokenTicker;
        bool claimed;
        uint tickerStakeIDX;
        uint addressStakeIDX;
    }

    // @notice Mapping of program ID to program
    mapping(bytes32 => Program) programs;

    // @notice Stores Exist status for program IDS
    mapping(bytes32 => bool) programExists;

    // @notice Stores Programs by token ticker
    mapping(string => Program[]) tickerPrograms;

    // @notice List of program IDS
    bytes32[] programsID;

    // @notice Stores stakes by Stake ID
    mapping(bytes32 => Stake) stakes;

    // @notice Stores Stakes for user
    mapping(address => Stake[]) addressStakes;

    // @notice Stores Stakes for tickers
    mapping(string => Stake[]) tickerStakes;

    constructor(address _address) {
        owner = msg.sender; 
        emit EgcEvents.OwnerSet(address(0), owner);
        egcTokenAdd = _address;
    }


    // @notice Only contract address may interact with functions modified by this command
    modifier onlyOwner {
        require(msg.sender == owner, "Caller must be the owner");
        _;
    }

    /**
     * @dev Creates a stacking program.
     * @param duration - int value represent the duration of the program
     * @param rewardBasisPoint - int value representing the reward of project in bases 1% == 100 basis
     * @param tokenTicker - the ticker representation of a token
     */
    function createStakeProgram(uint duration, uint rewardBasisPoint, string memory tokenTicker) public onlyOwner returns (bytes32) {
        require(rewardBasisPoint > 0, "reward basis point must be positive");
        uint tickerIDX;

        bytes32 id = ProgramID.compute(ProgramID.ProgramKeyData(duration * 1 days, owner, rewardBasisPoint, tokenTicker));

        // if program id exists cancel operation
        require(!programExists[id], "Program Already Created");
        programExists[id] = true;
        programsID.push(id);
        if (tickerPrograms[tokenTicker].length == 0) {
            tickerIDX = 0;
        } else {
            tickerIDX = tickerPrograms[tokenTicker].length;
        }
        Program memory program = Program(id, 0, 0, duration, owner, rewardBasisPoint, tokenTicker, tickerIDX);
        programs[id] = program;
        tickerPrograms[tokenTicker].push(program);
        emit EgcEvents.ProgramCreated(owner, id, duration);
        return id;
    }

    /**
    * @dev List all program ids
    * @return a list of program IDs
    */
    function listProgramIDs() public view onlyOwner returns (bytes32[] memory)  {
        return programsID;
    }

    /**
    * @dev Gets a program by its id
    * @param id - the id of the program too get
    * @return a program
    */
    function getProgramByID(bytes32 id) public view onlyOwner returns (Program memory) {
        return programs[id];
    }

    /**
    * @dev Gets programs by token ticker
    * @param tokenTicker - the ticker of the token
    * @return a list of programs
    */
    function getProgramsByTicker(string memory tokenTicker) public view onlyOwner returns (Program[] memory) {
        return tickerPrograms[tokenTicker];
    }

    /**
    * @dev ALlows a user to make a token stake
    * @param programID - the id of the program to stake 
    * @param rewardWallet - the egc wallet to recieve rewards
    */
    function makeStake(bytes32 programID, address rewardWallet) public payable {
        require(msg.value > 0, "stacked value must be positive");
        uint tokenTickerIDX;
        uint addressTickerIDX;
        Program memory program = programs[programID];
        bytes32 id = StakeUtilities.compute(StakeUtilities.StakeKeyData(msg.sender, msg.value, block.timestamp));
        if (tickerStakes[program.tokenTicker].length == 0) {
            tokenTickerIDX = 0;
        } else {
            tokenTickerIDX = tickerStakes[program.tokenTicker].length;
        }

        if (addressStakes[msg.sender].length == 0) {
            addressTickerIDX = 0;
        } else {
            addressTickerIDX = addressStakes[msg.sender].length;
        }

        Stake memory staked = Stake(id, msg.sender, block.timestamp, block.timestamp + program.stakeDuration, msg.value, StakeUtilities.calculate(msg.value, program.rewardPercentage), programID, rewardWallet, program.tokenTicker, false, tokenTickerIDX, addressTickerIDX);
        programs[programID].unclaimedTokens += msg.value;
        programs[programID].stakesHistory += 1;
        tickerPrograms[program.tokenTicker][programs[programID].tickerProgramIDX] = programs[programID];
        stakes[id] = staked;
        addressStakes[msg.sender].push(staked);
        tickerStakes[program.tokenTicker].push(staked);
        emit EgcEvents.TokenStaked(msg.sender, programID, id, msg.value);
    }

    /**
    * @dev User stake history
    * @param _owner - the address to view stake history
    * @return a list of user stake data
    */
    function addressStakeHistory(address _owner) public view returns (Stake[] memory) {
        return addressStakes[_owner];
    }

    /**
    * @dev Tiker stake history
    * @param ticker - the ticker to check stake histories
    * @return a list of ticker stakes
    */
    function tickerStakeHistory(string memory ticker) public view returns (Stake[] memory) {
        return tickerStakes[ticker];
    }

    /**
    * @dev Send the right token to user address
    * @param ticker - the staked token
    * @param reciever - the wallet to send tokens to 
    * @return sent - bool value confirming token transfer
    */
    function sendBackStake(address reciever, string memory ticker, uint256 _value) internal returns (bool sent) {
        bool matchEGC = StakeUtilities.compareStrings(ticker, "EGC");

        if (matchEGC) {
            ERC20 token = ERC20(egcTokenAdd);
            return token.transfer(reciever, _value);
        }

        (bool maticSent, ) = reciever.call{value: _value}("");
        return maticSent;
    }

    /**
    * @dev Unstake a users tokens 
    * @param stakeID - the user stake to 
    */
    function unstakeToken(bytes32 stakeID) public {
        Stake memory stake = stakes[stakeID];
        require(stake.owner == msg.sender, "only the address that stacked the token can unstake it");
        require(stake.claimAt < block.timestamp, "staked tokens have not yet matured");
        require(!stake.claimed, "rewards have been claimed");
        require(stake.rewardsToClaim > 0, "rewards must be greater then zero");
        bool sentStaked = sendBackStake(stake.owner, stake.tokenTicker, stake.amountStacked);
        require(sentStaked, "the staked token failed to send");
        ERC20 token = ERC20(egcTokenAdd);
        bool sentReward = token.transfer(stake.rewardWallet, stake.rewardsToClaim);
        require(sentReward, "earned rewards failed to send");
        programs[stake.programID].unclaimedTokens -= stake.amountStacked;
        stakes[stakeID].rewardsToClaim = 0;
        stakes[stakeID].claimed = true;
        addressStakes[stake.owner][stakes[stakeID].addressStakeIDX] = stakes[stakeID];
        tickerStakes[stake.tokenTicker][stakes[stakeID].tickerStakeIDX] = stakes[stakeID];
        emit EgcEvents.TokenUnstaked(stake.rewardWallet, stake.programID, stakeID, stake.rewardsToClaim);     
    }

    receive() external payable {
    // This function is executed when a contract receives plain Ether (without data)
    }
}