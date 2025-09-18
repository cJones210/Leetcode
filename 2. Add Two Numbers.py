class Solution(object):
    def addTwoNumbers(self, l1, l2):
        carry=0
        # first time using linked lists so this might not be optimal
        dummy = ListNode(0)
        curr_node = dummy
        while l1 is not None or l2 is not None or carry != 0:
            value1 = l1.val if l1 else 0
            value2 = l2.val if l2 else 0
            sum = value1 + value2 + carry
            carry = sum // 10
            digit = sum % 10
            # i also don't fully understand how linked list classes work yet so below is partly repeated from linked list documentation
            curr_node.next = ListNode(digit)
            curr_node = curr_node.next

            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next
        return dummy.next
    
    #beats 100% with a 0ms execution time but for this problem execution time seems very random because of python