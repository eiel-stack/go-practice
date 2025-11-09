// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 导入OpenZeppelin的ERC721标准实现
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
// 导入计数器工具（用于自动生成tokenId）
import "@openzeppelin/contracts/utils/Counters.sol";

contract MyImageNFT is ERC721 {
    // 引入计数器库
    using Counters for Counters.Counter;
    // 声明tokenId计数器（从1开始）
    Counters.Counter private _tokenIdCounter;

    // 构造函数：初始化NFT名称和符号
    constructor(
        string memory _name, 
        string memory _symbol
    ) ERC721(_name, _symbol) {
        // 初始计数器从0开始，首次铸造会自动+1
    }

    // 铸造NFT函数：接收者地址 + 元数据URI
    function mintNFT(address recipient, string memory tokenURI) 
        public 
        returns (uint256) 
    {
        // 生成新的tokenId（每次+1）
        uint256 newTokenId = _tokenIdCounter.current() + 1;
        _tokenIdCounter.increment();

        // 铸造NFT给接收者
        _safeMint(recipient, newTokenId);
        // 设置该NFT的元数据URI
        _setTokenURI(newTokenId, tokenURI);

        return newTokenId; // 返回铸造的tokenId
    }

    // 重写tokenURI函数（OpenZeppelin v4+需要显式实现）
    function tokenURI(uint256 tokenId) 
        public 
        view 
        override 
        returns (string memory) 
    {
        return super.tokenURI(tokenId);
    }
}