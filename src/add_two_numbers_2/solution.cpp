#include "solution.h"

#include <iostream>
#include <memory>
#include <utility>

std::unique_ptr<ListNode>
Solution::addTwoNumbers(const std::unique_ptr<ListNode>& l1, const std::unique_ptr<ListNode>& l2) {
    auto dummy = std::make_unique<ListNode>(0);


    auto* tail = dummy.get();
    auto carry = 0;

    auto* p1 = l1.get();
    auto* p2 = l2.get();

    while (p1 || p2 || carry) {
        auto new_val = (p1 ? p1->val : 0) + (p2 ? p2->val : 0) + carry;

        carry = new_val / 10;
        new_val = new_val % 10;

        tail->next = std::make_unique<ListNode>(new_val);
        tail = tail->next.get();

        if (p1) {
            p1 = p1->next.get();
        }
        if (p2) {
            p2 = p2->next.get();
        }
    }

    return std::move(dummy->next);
}

void test() {
    auto l1 = std::make_unique<ListNode>(2);
    l1->next = std::make_unique<ListNode>(4);
    l1->next->next = std::make_unique<ListNode>(3);

    auto l2 = std::make_unique<ListNode>(5);
    l2->next = std::make_unique<ListNode>(6);
    l2->next->next = std::make_unique<ListNode>(4);

    const auto result = Solution::addTwoNumbers(l1, l2);

    for (ListNode* it = result.get(); it != nullptr; it = it->next.get()) {
        std::cout << it->val << " ";
    }
    std::cout << "\n";
}
