#include <experimental/source_location>
#include <iostream>
#include <iterator>
#include <set>
#include <sstream>
#include <vector>

#include "learn/prettyprint.hpp"
#include "learn/structs.hpp"
#include "learn/utilities.hpp"
#include "learn/smartpointers.hpp"
#include "learn/coroutines.hpp"

void log(std::string_view message,
         const std::experimental::source_location& location =
             std::experimental::source_location::current()) {
  std::cout << "info:" << location.file_name() << ":" << location.line() << " "
            << message << '\n';
}

std::ostream& operator<<(std::ostream& os, const person& person) {
  os << "name: " << person.name << " age: " << person.age;
  return os;
}

[[maybe_unused]] void designated_initializers() {
  std::cout << "-- designated initializer for Person" << std::endl;
  auto p = person{.name = "bingo", .age = 23};
  std::cout << p << std::endl;
}

[[maybe_unused]] void source_location() { log("Logging Hello world!"); }

[[maybe_unused]] void set_remove_dups_demo() {
  using namespace std;
  string const input{"a a a b c foo bar foobar foo bar bar"};
  set<string> s1;
  std::istringstream instream(input);
  istream_iterator<string> it1{cin};
  istream_iterator<string> end;
  copy(it1, end, inserter(s1, s1.end()));
}

[[maybe_unused]] void vector_remove_erase_idiom() {
  using namespace std;
  vector<int> v{1, 2, 3, 2, 5, 2, 6, 2, 4, 8};
  cout << "Original      : " << v << endl;
  const auto new_end = remove(begin(v), end(v), 2);
  cout << "After Remove  : " << v << endl;
  v.erase(new_end, end(v));
  cout << "After Erase   : " << v << endl;
  const auto is_odd = [](int i) { return i % 2 != 0; };
  v.erase(remove_if(begin(v), end(v), is_odd), end(v));
  cout << "After Odd Del : " << v << endl;
}

[[maybe_unused]] void const_lifetime_extension2() {
  const std::string& s = my_container().get_s();
  std::cout << s << std::endl;
}

class address {
 public:
  address& city(std::string&& city) {
    city_ = std::move(city);
    return *this;
  }

  address& state(std::string&& state) {
    state_ = state;
    return *this;
  }
  // See cppmovepdf: 5.1.3 Using Move Semantics to Solve the Dilemma
  [[nodiscard]] const std::string& city() const& {
    std::cout << "const std::string& city() const &" << std::endl;
    return city_;
  }
  [[nodiscard]] std::string city() && {
    std::cout << "std::string city() &&" << std::endl;
    return std::move(city_);
  }
  [[maybe_unused]] [[nodiscard]] const std::string& state() const {
    return state_;
  }

  std::string city_;
  std::string state_;

  friend std::ostream& operator<<(std::ostream& os, const address& address) {
    os << "city_: " << address.city_ << " state_: " << address.state_;
    return os;
  }
};

address& create_address_p() {
  return address{}.city("bangalore").state("karnataka");
}

address create_address_q() {
  address a{"mumbai", "maharashtra"};
  return a;
}

[[maybe_unused]] void demo_rvalues()  // BAD CODE
{
  std::cout << "entering demo_rvalues" << std::endl;
  auto& addr = create_address_p();
  std::cout << "0. temp stuff to override stack"
            << std::endl;  // if you comment this it works
  std::cout << "0. addr.city=" << addr.city()
            << std::endl;  // prints garbage if prev line is not commented

  auto&& addr1 = create_address_q();
  std::cout << "1. temp stuff to override stack" << std::endl;
  std::cout << "1. addr1: " << addr1 << std::endl;

  auto&& city2 = create_address_q().city_;
  std::cout << "2. temp stuff to override stack" << std::endl;
  std::cout << "2. city2: " << city2 << std::endl;

  auto&& city3 = create_address_q().city();
  std::cout << "3. temp stuff to override stack" << std::endl;
  std::cout << "3. city3: " << city3 << std::endl;

  auto city4 =
      create_address_q().city();  // this is technically BAD but it works ??
  std::cout << "4. temp stuff to override stack" << std::endl;
  std::cout << "4. temp stuff to override stack" << std::endl;
  std::cout << "4. city4: " << city4 << std::endl;
}

[[maybe_unused]] void demo_sort_with_tie() {
  using namespace std;
  vector people = {
      person{"tre", 40}, {"Madhav", 42}, {"Vinay", 35}, {"Marco", 9}};
  cout << "Unsorted:" << people << endl;
  sort(begin(people), end(people));
  cout << "Sorted  :" << people << endl;
}

[[maybe_unused]] void demo_auto_decls() {
  using namespace std;
  vector v1{2, 3, 4};
  auto v2 = vector<int>{2, 3, 4};
  [[maybe_unused]] auto v3 = {2, 3, 4};
  [[maybe_unused]] auto v4{1};
  auto v5{vector<int>{1}};
  cout << "typeid(v1).name: " << typeid(v1).name() << endl;
  cout << "typeid(v2).name: " << typeid(v2).name() << endl;
  cout << "typeid(v3).name: " << typeid(v3).name() << endl;
  cout << "typeid(v4).name: " << typeid(v4).name() << endl;
  cout << "v4: " << v4 << endl;
  cout << "v5: " << v5 << endl;
  cout << "typeid(v5).name: " << typeid(v5).name() << endl;
}
int main() {
  //    designated_initializers();
  //    set_remove_dups_demo();
  //    vector_remove_erase_idiom();
  //  const_lifetime_extension2();
  //  demo_rvalues();
  //  demo_auto_decls();
  //  demo_sort_with_tie();
  //  demo_std_invoke();
  // demo_structs1();
    demo_smart_pointers();
  // return 0;
}

class flat_land {
 public:
  auto operator<=>(const flat_land&) const = default;
  int x;
  int y;

 private:
  int z;
};
