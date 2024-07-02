/* Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order. */


class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        sort(nums1.begin(), nums1.end());
        sort(nums2.begin(), nums2.end());
        
        std::vector<int> res;
        
        auto it1 = nums1.begin(), it2 = nums2.begin();
        while (it1 < nums1.end() && it2 < nums2.end()) {
            if (*it1 == *it2) {
                res.push_back(*it1);
                it1++;
                it2++;
            } else if (*it1 > *it2) {
                it2++;
            } else {
                it1++;
            }
        }
        
        return res;
    }
};
