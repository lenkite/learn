//
// Created by Elankath, Tarun Ramakrishna on 29/11/20.
//
#ifndef CPPLRN_UTILITIES_HPP
#define CPPLRN_UTILITIES_HPP

#include <atomic>
#include <concepts>
template <typename C, typename... Args>
requires std::invocable<C, Args...> decltype(auto) call(C&& callable,
                                                        Args&&... args) {
  return std::invoke(std::forward<C>(callable), std::forward<Args>(args)...);
}

template <typename T> void print(const T& coll) {
  std::cout << "elems: ";
  for (const auto& elem : coll) {
    std::cout << elem << ' ';
  }
  std::cout << '\n';
}
void demo_std_invoke() {
  std::vector<int> vals{0, 8, 15, 42, 13, -1, 0};
  call([&vals] { std::cout << "size: " << vals.size() << '\n'; });
  call(print<std::vector<int>>, vals);
  call(&decltype(vals)::pop_back, vals);
  call(print<std::vector<int>>, vals);
  auto ai = std::atomic<int>(42);  // error
  std::cout << ai.operator int() << std::endl;
}

#endif  // CPPLRN_UTILITIES_HPP
