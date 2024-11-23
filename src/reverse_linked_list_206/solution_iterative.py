# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        current = head
        prev = None

        while current is not None:
            _next = current.next
            current.next = prev
            prev = current
            current = _next

        return prev

    def print_list(self, head: Optional[ListNode]) -> None:
        while head is not None:
            print(head.val, end="->")
            head = head.next


def test():
    obj = Solution()
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)
    head.next.next.next = ListNode(4)
    head.next.next.next.next = ListNode(5)
    new_head = obj.reverseList(head)
    print(obj.print_list(new_head))
