The 12 Projects of Codemas
==========================

# Description
This repo is based on the tutorials for small simple and nice 1 day projects with Raspberry Pi Pico and different components (LED, buzzer, sensors, screen, etc.).
Original projects from that Advent calendar can be found [here](https://thepihut.com/pages/maker-advent-2022-guides).
My repository consists of folders representing days in the original Advent Calendar which is written in TinyGo instead of MicroPython as it's done in the tutorial. However, some days are skipped, some duplicated with different variations of components interaction. Some projects are more advanced version of original ones, other are basically same projects, just rewritten in Golang syntax.

You can find video recording of projects inside `demo` folder.

# Projects
* day3 - actually contains summary of [Day 1](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-1-getting-started), [Day 2](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-2-let-s-get-blinky) and [Day 3](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-3-bashing-buttons). It has constant LED blinking (heartbeat), printing useful information, lights up LEDs on appropriate button click
* day4-1 - based on [Day 4](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-4-amazing-analogue) uses potentiometer to light up certain LEDs depending on potentiometer position (think of volume indicator)
* day4-2 - similarly, based on [Day 4](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-4-amazing-analogue) but applies intensity of LED light depending on potentiometer position
* day5 - contains project from [Day 5](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-5-hear-my-code). It uses buzzer to play "Jingle Bells", while you can set volume with potentiometer
* day6 - adds to previous project, the one from [Day 6](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-6-looking-for-light). It regulated LED lights intensity based on lightsensor (less light is brighter, like on smartphones)
* day7 - replica of [Day 7](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-7-monitoring-motion) in golang. It enables alarm (buzzer sound and LEDs blinking) on motion detection
* day8 - based on [Day 8](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-8-tracking-temps) it enables alarm (similar from previous project) when temperature is too low (less than 18 degrees)
* day9 - based on [Day 9](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-9-full-tilt) it enables buzzer and lights LEDs when tilt is too high
* day10 - replica of [Day 10](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-10-breaking-beams). This project lights LEDs when there's some object in between of two sensors. It also contains game counting taps, however it's not visible in demo
* day11-1 - project of [Day 11](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-11-omg-oled). This project shows on screen 2 words and depending on moment of time on button click stops at one of them
* day11-2 - project of [Day 11](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-11-omg-oled). It prints light sensor value on screen.
* day12-1 - project of [Day 12](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-12-rgb-led-strip). It implements light chaser, speed of which can be controlled by potentiometer
* day12-2 - project of [Day 12](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-12-rgb-led-strip). It shows fade in & fade out effect with LED strip

# How to run
Here's example which runs day3 project
```sh
export DIR=./day3
cd $DIR && tinygo flash -target=pico && tinygo monitor
```