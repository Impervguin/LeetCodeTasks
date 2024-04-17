class Solution:
    def customSortString(self, order: str, s: str) -> str:
        d = dict()
        for i in range(len(order)):
            d[order[i]] = i # dictionary for order in string. The lower the index, the higher the order.
        def key(x):
            return d[x] if x in d.keys() else -1 # Key func using dictionary.
        return "".join(sorted(s, key=key)) # sorting with this key function.