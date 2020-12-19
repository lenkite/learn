//
// Created by Elankath, Tarun Ramakrishna on 19/12/20.
//

#ifndef CPPLRN_VECT_HPP
#define CPPLRN_VECT_HPP

#include <vector>
#include <iostream>
#include <algorithm>

void demo_erase_remove() {
    using namespace std;
    vector<int> v{1, 2, 3, 2, 5, 2, 6, 2, 4, 8};
    const auto new_end(remove(begin(v), end(v), 2));

}

#endif //CPPLRN_VECT_HPP
