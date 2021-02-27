// This file include a external header cxxopts.hpp
// Imitate the tool "notify-send foo bar"
// make sure you have the lib "libnotify-devel"

#include <chrono>
#include <cxxopts.hpp>
#include <iostream>
#include <libnotify/notify.h>
#include <string>
#include <thread>

int main(int argc, char *argv[]) {
    notify_init("Sample");
    cxxopts::Options options("test", "A brief description");

    NotifyNotification *n = notify_notification_new(
        "Hello world", "some message text... bla bla", 0);

    options.add_options()
    (
        "t,time", "set timer", cxxopts::value<int>()->default_value("5"))(
        "h,help", "Print usage"
    );
    auto result = options.parse(argc, argv);
    if (result.count("help")) {
        std::cout << options.help() << std::endl;
        exit(0);
    }

    int time_i = result["time"].as<int>();
    std::cout << "notify in " << std::to_string(time_i) << "seconds;"
              << std::endl;
    std::this_thread::sleep_for(std::chrono::seconds(time_i));

    notify_notification_set_timeout(n, 10000); // 10 seconds

    if (!notify_notification_show(n, 0)) {
        std::cerr << "show has failed" << std::endl;
        return -1;
    }
    return 0;
}

// g++ desk_timer.cc -o desk_timer `pkg-config --cflags --libs libnotify`
