// first solution with minLefts record
class Solution {
    public int maxProfit(int[] prices) {
        if (prices.length == 0 || prices.length == 1) return 0;

        int[] minLefts = new int[prices.length];
        int maxProfit = 0;

        for (int i = 0; i < prices.length; i++) {
            if (i == 0) {
                minLefts[i] = -1;
                continue;   
            }
            
            if (minLefts[i - 1] == -1) {
                minLefts[i] = prices[i] > prices[i - 1] ? prices[i - 1] : -1;
            } else {
                minLefts[i] = prices[i] < minLefts[i - 1] ? -1 : minLefts[i - 1];
            }

            if (minLefts[i] != -1) {
                maxProfit = Math.max(maxProfit, prices[i] - minLefts[i]);
            }
        }

        return maxProfit;
    }
}

// sliding window solution
class Solution {
    public int maxProfit(int[] prices) {
        int l = 0;
        int maxProfit = 0;
        for(int r = 1; r < prices.length; r++) {
            if(prices[l] >= prices[r]) {
                l = r;
            } else {
                maxProfit = Math.max(maxProfit, prices[r] - prices[l]);
            }
        }

        return maxProfit;
    }
}