class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        mixed=sorted(nums1+nums2)
        mixedlen=len(mixed)
        if mixedlen % 2 == 0:
            even = float(mixed[mixedlen/2]+mixed[mixedlen/2-1])
            if even == 0:
                return(0)
            else:
                return(even/2)
        else:
            return(mixed[mixedlen/2])    
    
#I understand this isnt the algorithmically fastest solution but it is still relatively fast for only needing to merge sort
#this one took two days because i tried to figure out a better way then this

#0ms Beats 100.00%