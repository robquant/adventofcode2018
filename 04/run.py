import re
from collections import defaultdict

def read_minute(line):
    minute_start = line.find(":") + 1
    return int(line[minute_start:minute_start+2])

guard_pattern = re.compile("Guard #(\d+)")
lines = sorted(open("input.txt").readlines())

sleep_times = defaultdict(lambda: [0] * 60)

for line in lines:
    if guard_pattern.search(line):
        active_guard = int(guard_pattern.search(line).groups()[0])
    if "falls asleep" in line:
        minute_start = read_minute(line)
    if "wakes up" in line:
        minute_stop = read_minute(line)
        guard = sleep_times[active_guard]
        for minute in range(minute_start, minute_stop):
            guard[minute] += 1

sleepy_guard = max(sleep_times.items(), key=lambda guard: sum(guard[1]))
sleepy_guard_id = sleepy_guard[0]
sleepy_guard_schedule = sleep_times[sleepy_guard[0]]
print(sleepy_guard_id * sleepy_guard_schedule.index(max(sleepy_guard_schedule)))

max_often_sleep = 0
for guard_id, schedule in sleep_times.items():
    if max(schedule) > max_often_sleep:
        max_often_sleep = max(schedule)
        max_often_guard_id = guard_id
        max_often_minute = schedule.index(max_often_sleep)
print(max_often_guard_id * max_often_minute)