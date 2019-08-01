/***
There are N children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?

Example 1:

Input: [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
Example 2:

Input: [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
             The third child gets 1 candy because it satisfies the above two conditions.

**/

class Solution {
    public int candy(int[] ratings) {
		int numCandy = 0;
		int can[]=new int[ratings.length];
		
        if(ratings.length == 0)
			return 0;
		else{
			
			for(int j = 0;j<can.length;j++){
				can[j] = 1;
			}
			
			for(int i = 0;i<ratings.length-1;i++){
				if(ratings[i+1] > ratings[i])
					can[i+1] = can[i]+1;
			}
			for(int k = ratings.length-1; k>=1; k--){
				if(ratings[k-1]> ratings[k]){
					if(can[k-1] < (can[k]+1))
						can[k-1] = can[k]+1;
				}
		    }
			
			for(int j = 0;j<can.length;j++){
				numCandy = numCandy+ can[j];
			}
		}
    return numCandy;
	}
	
}