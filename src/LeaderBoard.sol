// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
// import "../lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

/**
 * @title Game records
 * @dev Getter and setter methods
 */
contract GameHistory {
    struct GameSession {
        uint256 gid; //game ID
        string gtid; //game ID
        string uid; //user ID
        string data; //game data
        uint256 time; //game time stamp
    }

    struct Payment {
        address sender; //sender's address
        uint256 amount; //total amount sent
        uint256 time; //timestamp of the payment
    }

    address private immutable owner;
    uint256 public oneEGC;
    address public tokenAddressEGC;
    address public USDCTokenAddress;
    address public USDTTokenAddress;

    ERC20 private egcToken;
    ERC20 private USDTToken;
    ERC20 private USDCToken;

    // event for EVM logging
    event OwnerSet(address indexed oldOwner, address indexed newOwner);
    event ReceivedLessThanTarget(address indexed sender, uint256 amount);
    event Received(address indexed sender, uint256 amount);

    mapping(uint256 => GameSession[]) gameHistory; //historical data for each game
    mapping(string => GameSession[]) userHistory; //historical data for each user
    mapping(address => Payment[]) payments; //historical data for each payment
    mapping(address => uint256) totalPaid; //total amount paid by each sender

    constructor(address _egcAddress, address _usdcAddress, address _usdtAddress) {
        owner = msg.sender;
        emit OwnerSet(address(0), owner);
        tokenAddressEGC = _egcAddress;
        USDCTokenAddress = _usdcAddress;
        USDTTokenAddress = _usdtAddress;
        egcToken = ERC20(_egcAddress);
        USDCToken = ERC20(_usdcAddress);
        USDTToken = ERC20(_usdtAddress);
        oneEGC = 1 * (10 ** egcToken.decimals());
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Caller is not the owner");
        _;
    }

    /**
     * @dev Stores the game data for the given game ID, game type ID, user ID, data, and time.
     * @param _gid The unique game ID.
     * @param _gtid The game type ID.
     * @param _uid The user ID.
     * @param _data The game data to be stored.
     * @param _time The timestamp of when the game data is stored.
     */
    function storeGameData(uint256 _gid, string memory _gtid, string memory _uid, string memory _data, uint256 _time)
        public
    {
        GameSession memory currentGame_ = GameSession(_gid, _gtid, _uid, _data, _time);

        gameHistory[_gid].push(currentGame_);
        userHistory[_uid].push(currentGame_);
    }

    /**
     * @dev Retrieves the game history for a specific game ID.
     * @param _gid The unique game ID for which the game history is to be retrieved.
     * @return An array of GameSession structs representing the game history for the specified game ID.
     */
    function getGameHistory(uint256 _gid) public view returns (GameSession[] memory) {
        return gameHistory[_gid];
    }

    /**
     * @dev Retrieves the user history for a specific user ID.
     * @param _uid The unique user ID for which the user history is to be retrieved.
     * @return An array of GameSession structs representing the user history for the specified user ID.
     */
    function getUserHistory(string memory _uid) public view returns (GameSession[] memory) {
        return userHistory[_uid];
    }

    /**
     * @dev Sends a specified amount of EGC tokens to a given user address.
     * @param _user The address of the user to send the tokens to.
     * @param _amount The amount of EGC tokens to send.
     * @notice Requires that the contract has enough tokens to send.
     */
    function sendEgc(address _user, uint256 _amount) private {
        
        bool success = egcToken.transfer(_user, _amount);
        require(success, "swapped failed to send easy get coin tokens");
    }

    /**
     * @dev Retrieves the total amount paid by a specific user.
     * @param _user The address of the user.
     * @return The total amount paid by the user.
     */
    function userTotal(address _user) public view returns (uint256) {
        return totalPaid[_user];
    }

    /**
     * @dev Retrieves the stake history for a specific user.
     * @param _user The address of the user.
     * @return An array of Payment structs representing the stake history for the specified user.
     */
    function userStakeHistory(address _user) public view returns (Payment[] memory) {
        return payments[_user];
    }

    /**
     * @dev swapUSDC function to swap USDC for EasyGetCoin.
     * @param _amount The amount of USDC to be swapped.
     */
    function swapUSDT(uint256 _amount) public {
        require(_amount > 0, "swap value must be positive");
        bool usdtSwapped = USDTToken.transferFrom(msg.sender, address(this), _amount);
        require(usdtSwapped, "approve the amount to swap and ensure balance is sufficient");
        payments[msg.sender].push(Payment({sender: msg.sender, amount: _amount, time: block.timestamp}));
        totalPaid[msg.sender] += _amount;
        sendEgc(msg.sender, _amount);
    }

    /**
     * @dev swapUSDC function to swap USDC for EasyGetCoin.
     * @param _amount The amount of USDC to be swapped.
     */
    function swapUSDC(uint256 _amount) public {
        require(_amount > 0, "swap value must be positive");
        bool usdcSwapped = USDCToken.transferFrom(msg.sender, address(this), _amount);
        require(usdcSwapped, "approve the amount to swap and ensure balance is sufficient");
        payments[msg.sender].push(Payment({sender: msg.sender, amount: _amount, time: block.timestamp}));
        totalPaid[msg.sender] += _amount;
        sendEgc(msg.sender, _amount);
    }

    /**
     * @dev Fallback function to receive Ether.
     * This function is called when the contract receives Ether without a function being explicitly called.
     */
    receive() external payable {
        emit Received(msg.sender, msg.value);
        if (msg.value < oneEGC) {
            emit ReceivedLessThanTarget(msg.sender, msg.value);
            revert("Received amount is less than the target amount");
        }
        emit Received(msg.sender, msg.value);
        payments[msg.sender].push(Payment({sender: msg.sender, amount: msg.value, time: block.timestamp}));
        totalPaid[msg.sender] += msg.value;
        sendEgc(msg.sender, msg.value);
    }
}
