//
// Created by Elankath, Tarun Ramakrishna on 29/11/20.
//

#ifndef CPPLRN_STRUCTS_HPP
#define CPPLRN_STRUCTS_HPP
#include <iostream>
#include <string>
struct foo4 {  // size = 24, alignment = 8
               // foo4: +--------+--------+--------+--------+ //
  int a;       // members: |aaaab...|cccc....|dddddddd|e.......|
  char b;      // . represents a byte of padding
  float c;
  double d;
  bool e;
};

struct foo5 {  // size = 24, alignment = 8
               // foo5: +--------+--------+--------+--------+ //
               // members: |dddddddd|aaaacccc|be......|
  double d;
  int a;
  float c;
  char b;  // . represents a byte of padding
  bool e;
};

struct foo6 {
  int id;
  std::string name;
  int age;
};

struct foo7 {
  int id;
  int age;
  std::string name;
};
void demo_structs1() {
  using namespace std;
  cout << "size(foo4)=" << sizeof(foo4) << endl;
  cout << "size(foo5)=" << sizeof(foo5) << endl;
  cout << "size(foo6)=" << sizeof(foo6) << endl;
  cout << "size(foo7)=" << sizeof(foo7) << endl;
  cout << "size(std::string)=" << sizeof(std::string) << endl;
}

namespace bingo {
    struct address {
        std::string city;
    };

    struct person {
        address addr;
    };

}
void structs() {
    std::cout << "strucs called" << std::endl;
}
#endif  // CPPLRN_STRUCTS_HPP
