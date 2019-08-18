/**
Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.

You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.

Example 1:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".
Example 2:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).

**/

/**
Runtime: 8 ms, faster than 83.96% of Java online submissions for Minimum Index Sum of Two Lists.
Memory Usage: 38.6 MB, less than 100.00% of Java online submissions for Minimum Index Sum of Two Lists.

**/

class Solution {
    public String[] findRestaurant(String[] list1, String[] list2) {
        HashMap<String,Integer> countMap = new HashMap<>();
        List<String> list = new ArrayList<>();
        int minValue = Integer.MAX_VALUE;
        for(int i =0;i<list1.length;i++)
            countMap.put(list1[i],i);
        for(int j =0;j<list2.length;j++){
            if(countMap.containsKey(list2[j]) && countMap.get(list2[j])+j <=minValue)
            {
                if(countMap.get(list2[j])+j <minValue){
                    list.clear();
                    minValue = countMap.get(list2[j])+j;
                }
                list.add(list2[j]);
            }
        }
        return list.toArray(new String[list.size()]);
    }
}