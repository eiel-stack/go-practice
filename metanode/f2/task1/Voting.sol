// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {

    //  存储候选人的票数：候选人地址 => 映射得票数
    mapping(address => uint256) private candidateVotes;

    // 投票函数：向指定候选人投票（每次投一票）
    function vote(address candidate) external {
        require(candidate != address(0), "Invalid candidate address");
        candidateVotes[candidate] += 1;
    }

    // 获取候选人的票数
    function getVotes(address candidate) external view returns (uint256) {
        return candidateVotes[candidate];
    }

    // 重置所有候选人的票数
    function resetVotes() external {
        for (uint256 i = 0; i < candidateVotes.length; i++) {
            candidateVotes[i] = 0;
        }
    }
}