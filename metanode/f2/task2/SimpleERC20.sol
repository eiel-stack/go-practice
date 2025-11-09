// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleERC20 {
    // 代币元数据
    string public name;
    string public symbol;
    uint8 public decimals = 18; // 标准18位小数
    uint256 public totalSupply;

    // 余额映射：地址 => 余额
    mapping(address => uint256) public balanceOf;
    // 授权映射：授权人 => (被授权人 => 授权额度)
    mapping(address => mapping(address => uint256)) public allowance;

    // 合约所有者（用于增发权限控制）
    address public owner;

    // 转账事件（标准ERC20事件）
    event Transfer(address indexed from, address indexed to, uint256 value);
    // 授权事件（标准ERC20事件）
    event Approval(address indexed owner, address indexed spender, uint256 value);

    // 构造函数：初始化代币名称、符号、初始供应量
    constructor(
        string memory _name,
        string memory _symbol,
        uint256 _initialSupply
    ) {
        name = _name;
        symbol = _symbol;
        owner = msg.sender;
        // 初始供应分配给部署者
        _mint(msg.sender, _initialSupply * (10 **uint256(decimals)));
    }

    // 修饰符：仅所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    // 转账功能
    function transfer(address to, uint256 amount) public returns (bool) {
        require(to != address(0), "Transfer to zero address");
        require(balanceOf[msg.sender] >= amount, "Insufficient balance");

        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    // 授权功能
    function approve(address spender, uint256 amount) public returns (bool) {
        require(spender != address(0), "Approve to zero address");

        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    // 授权转账功能
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public returns (bool) {
        require(from != address(0), "Transfer from zero address");
        require(to != address(0), "Transfer to zero address");
        require(balanceOf[from] >= amount, "Insufficient balance");
        require(allowance[from][msg.sender] >= amount, "Allowance exceeded");

        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        allowance[from][msg.sender] -= amount;
        emit Transfer(from, to, amount);
        return true;
    }

    // 增发功能（仅所有者）
    function mint(address to, uint256 amount) public onlyOwner {
        require(to != address(0), "Mint to zero address");
        _mint(to, amount);
    }

    // 内部增发逻辑
    function _mint(address to, uint256 amount) internal {
        totalSupply += amount;
        balanceOf[to] += amount;
        emit Transfer(address(0), to, amount); // 从0地址转账表示增发
    }
}