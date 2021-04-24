//
// Created by Elankath, Tarun Ramakrishna on 12/03/21.
//

#ifndef CPPLRN_SMARTPOINTERS_HPP
#define CPPLRN_SMARTPOINTERS_HPP

#include <memory>
#include "structs.hpp"

std::unique_ptr<person> create_person() {
    return std::make_unique<person>(person{"bingo", 20});
}

void demo_smart_pointers() {
    using namespace std;
    cout << "(demo_smart_pointers) Calling create_person .." << endl;
    auto person = create_person();
    cout << "(demo_smart_pointers) Exiting scope " << endl;
}

#endif //CPPLRN_SMARTPOINTERS_HPP
