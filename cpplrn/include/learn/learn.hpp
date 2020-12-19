#ifndef MYAPP_APP_H
#define MYAPP_APP_H

#include <iostream>
#include <string>
#include <tuple>
struct person
{
  std::string name;
  int age;
  friend std::ostream& operator<<(std::ostream& os, const person& person);
  [[nodiscard]] auto tie() const
  {
    return std::tie(name);  // comparisons are on name
  }
  auto operator<=>(person& o) const { return this->tie() <=> o.tie(); }
};
struct my_container
{
  std::string s;

  my_container()
  {
    s = "Init";
    std::cout << "Constructed" << std::endl;
  }

  ~my_container()
  {
    s = "UNSAFE";
    std::cout << "Destructed" << std::endl;
  }

  [[nodiscard]] const std::string& get_s() const { return s; }
};
#endif
