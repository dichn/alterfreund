#include <iostream>
#include <cxxopts.hpp>
#include <libnotify/notify.h>

int main(int argc, char *argv[]) {
    notify_init("Sample");
    NotifyNotification *n = notify_notification_new(
        "Hello world", "some message text... bla bla", 0);

    notify_notification_set_timeout(n, 10000); // 10 seconds

    if (!notify_notification_show(n, 0)) {
        std::cerr << "show has failed" << std::endl;
        return -1;
    }
    return 0;
}

// g++ hello_world.cc -o hello_world `pkg-config --cflags --libs libnotify
