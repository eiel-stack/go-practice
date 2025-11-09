// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract MyImageNFT is ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIdCounter;

    constructor(string memory _name, string memory _symbol)
        ERC721(_name, _symbol)
    {}

    function mintNFT(address recipient, string memory _tokenURI)
        public
        returns (uint256)
    {
        uint256 newTokenId = _tokenIdCounter.current() + 1;
        _tokenIdCounter.increment();

        _safeMint(recipient, newTokenId);
        _setTokenURI(newTokenId, _tokenURI); // Updated parameter name

        return newTokenId;
    }
    function tokenURI(uint256 tokenId)
        public
        view
        override
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}