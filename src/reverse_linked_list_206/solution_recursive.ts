// Definition for singly-linked list.

class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function reverseList(head: ListNode | null): ListNode | null {
  if (head == null || head.next == null) {
    return head;
  }

  const newHead = reverseList(head.next);
  head.next.next = head;
  head.next = null;
  return newHead;
}

function printList(head: ListNode | null) {
  let current: ListNode | null = head;

  while (current !== null) {
    process.stdout.write(`${current.val} -> `);
    current = current.next;
  }
}

export function test() {
  const head = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5)))));

  const newHead = reverseList(head);
  printList(newHead);
}
