/**
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].

Solution Analysis:
Runtime: 1 ms, faster than 99.84% of Java online submissions for Min Cost Climbing Stairs.
Memory Usage: 36.9 MB, less than 100.00% of Java online submissions for Min Cost Climbing Stairs.
**/

class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int twoStepBefore = cost[0];
        int oneStepBefore = cost[1];
        int current =0;
        for(int i =2;i<cost.length;i++){
            current = Math.min(oneStepBefore,twoStepBefore)+cost[i];
            twoStepBefore = oneStepBefore;
            oneStepBefore = current;
        }
        return Math.min(oneStepBefore, twoStepBefore);
    }
}