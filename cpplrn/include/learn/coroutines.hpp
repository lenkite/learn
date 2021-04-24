//
// Created by Elankath, Tarun Ramakrishna on 24/04/21.
//

#ifndef CPPLRN_COROUTINES_HPP
#define CPPLRN_COROUTINES_HPP

#include <coroutine>
#include <iostream>
#include <memory>

template<typename T>
class nosuspend_task {
private:
    std::shared_ptr<T> value_;
public:

    explicit nosuspend_task(const std::shared_ptr<T> value) :value_{value} {
        std::cout << "(" << this << ") nosuspend_task(" << *value_ << ")" << std::endl;
    }

    T get() {
        std::cout << "(" << this << ") nosuspend_task::get() = " << *value_ << std::endl;
        return *value_;
    }

    ~nosuspend_task() {
        std::cout << "(" << this << ") ~nosuspend_tasktask()" << std::endl;
    }

    struct promise_type {
        std::shared_ptr<T> value_ = std::make_shared<T>();

        promise_type() {
            std::cout << "promise_type()" << std::endl;
        }

        ~promise_type() {
            std::cout << "~promise_type()" << std::endl;
        }

        nosuspend_task<T> get_return_object() {
            std::cout << "get_return_object()" << std::endl;
            return nosuspend_task{value_};
        }

        void return_value(T const& v) {
            std::cout << "return_value(" << v << ")" << std::endl;
            *value_ = v;
        }

        std::suspend_never initial_suspend() {          // (5)
            std::cout << "initial_suspend()" << std::endl;
            return {};
        }

        std::suspend_never final_suspend() noexcept {  // (6)
            std::cout << "final_suspend()" << std::endl;
            return {};
        }

        void unhandled_exception() {
            std::cerr << "Got unhandled exception :(" << std::endl;
        }
    };
};

nosuspend_task<int> create_eager_task() {
    std::cout << "create_eager_task() Executing: co_return 2021" << std::endl;
    co_return 2021;
}

#include <coroutine>
#include <iostream>
#include <optional>
#include <string_view>
#include <thread>
#include <vector>

std::jthread *thread;

template <typename T> struct future {
    struct promise_type {
        T value;
        future get_return_object() {
            return {std::coroutine_handle<promise_type>::from_promise(*this)};
        }
        std::suspend_always initial_suspend() noexcept {
            std::cout << "initial" << std::endl;
            return {};
        }
        std::suspend_always final_suspend() noexcept {
            std::cout << "final" << std::endl;
            return {};
        }
        void return_value(T x) {
            std::cout << "return value" << std::endl;
            value = std::move(x);
        }
        void unhandled_exception() noexcept {}

        ~promise_type() { std::cout << "future ~promise_type" << std::endl; }
    };

    struct awaiter {
        future &m_future;
        bool await_ready() const noexcept { return false; }

        void await_suspend(std::coroutine_handle<> handle) {
            std::cout << "await_suspend" << std::endl;
            *thread = std::jthread([this, handle] {
                std::cout << "Launch thread: " << std::this_thread::get_id()
                          << std::endl;
                m_future.coro.resume();
                handle.resume();
            });
        }

        T await_resume() {
            std::cout << "await_resume" << std::endl;
            return m_future.coro.promise().value;
        }

        ~awaiter() { std::cout << "awaiter::endl; }
    };

    std::coroutine_handle<promise_type> coro;

    future(std::coroutine_handle<promise_type> coro) : coro{coro} {}

    ~future() {
        std::cout << "~future" << std::endl;
        if (coro)
            coro.destroy();
    }

    awaiter operator co_await() {
        std::cout << "co_await" << std::endl;
        return {*this};
    }
};

template <typename F, typename... Args>
future<std::invoke_result_t<F, Args...>> async(F f, Args... args) {
    std::cout << "async" << std::endl;
    co_return f(args...);
}

struct task {

    struct promise_type {
        task get_return_object() { return {}; }
        std::suspend_never initial_suspend() noexcept { return {}; }
        std::suspend_never final_suspend() noexcept { return {}; }
        void return_void() {}
        void unhandled_exception() noexcept {}
        ~promise_type() { std::cout << "~task promise_type" << std::endl; }
    };

    ~task() { std::cout << "~task" << std::endl; }
};

int square(int x) {
    std::cout << "square in thread id " << std::this_thread::get_id()
              << std::endl;
    return x * x;
}

task f() {
    auto squared6 = co_await async(square, 6);

    std::cout << "Write " << squared6
              << " from thread: " << std::this_thread::get_id() << std::endl;
}

void demo_coroutines1() {
    std::jthread thread_a;
    ::thread = &thread_a;
    f();
}


void demo_coroutines() {
    std::cout << "(demo_coroutines) BEGIN" << std::endl;
    auto fut = create_eager_task();
    auto v = fut.get();
    std::cout << "(demo_coroutines) v: " << v << '\n';
    std::cout << "(demo_coroutines) END" << std::endl;
}

#endif //CPPLRN_COROUTINES_HPP
