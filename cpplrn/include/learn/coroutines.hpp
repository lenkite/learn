//
// Created by Elankath, Tarun Ramakrishna on 24/04/21.
//

#ifndef CPPLRN_COROUTINES_HPP
#define CPPLRN_COROUTINES_HPP

#include <coroutine>
#include <iostream>
#include <memory>

template<typename T>
struct eager_future {
    std::shared_ptr<T> value;                           // (3)
    explicit eager_future(std::shared_ptr<T> p): value(p) {}
    ~eager_future() = default;
    T get() {                                          // (10)
        return *value;
    }

    struct promise_type {
        std::shared_ptr<T> ptr = std::make_shared<T>(); // (4)
        ~promise_type() = default;
        eager_future<T> get_return_object() {              // (7)
            return eager_future{ptr};
        }
        void return_value(T v) {
            *ptr = v;
        }
        std::suspend_never initial_suspend() {          // (5)
            return {};
        }
        std::suspend_never final_suspend() noexcept {  // (6)
            return {};
        }
        void unhandled_exception() {
            std::exit(1);
        }
    };
};

eager_future<int> createFuture() {                         // (1)
    co_return 2021;                                    // (9)
}

#endif //CPPLRN_COROUTINES_HPP
