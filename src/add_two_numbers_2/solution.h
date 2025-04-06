#ifndef SOLUTION_H
#define SOLUTION_H
#include <memory>

struct ListNode {
    int val;
    std::unique_ptr<ListNode> next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, std::unique_ptr<ListNode> next) : val(x), next(std::move(next)) {}
};

class Solution {
public:
    static std::unique_ptr<ListNode>
    addTwoNumbers(const std::unique_ptr<ListNode>& l1, const std::unique_ptr<ListNode>& l2);
};

void test();

#endif  // SOLUTION_H
