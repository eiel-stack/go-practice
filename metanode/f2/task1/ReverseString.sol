// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ReverseString {

    function reverseString(string memory _str) public pure returns (string memory) {
        bytes memory strBytes = bytes(_str);
        bytes memory reversedBytes = new bytes(strBytes.length);
        for (uint i = 0; i < strBytes.length; i++) {
            reversedBytes[i] = strBytes[strBytes.length - 1 - i];
        }
        return string(reversedBytes);
    }
    
}