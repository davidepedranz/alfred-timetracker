# Alfred TimeTracker

[![Build Status](https://travis-ci.com/davidepedranz/alfred-timetracker.svg?branch=master)](https://travis-ci.com/davidepedranz/alfred-timetracker)
[![GolangCI](https://golangci.com/badges/github.com/davidepedranz/alfred-timetracker.svg)](https://golangci.com/r/github.com/davidepedranz/alfred-timetracker)

Track your time on Google Calendar 📅 using [Alfred](https://www.alfredapp.com/).

Have you ever asked yourself the question, `What did I do today?`.
If you did and could not answer, this Alfred workflow is for you.
Sometimes we have so many different tasks and interrupt that it is easy to lose track.
You can try to track them manually, but this approach is error-prone, time-consuming, and tedious.
This workflow allows you to leverage Alfred's power to track your work easily.

*Disclaimer*. The workflow is already working, but it is still under construction. Use it at your own risk ⚡.

## Installation

1. Download the latest workflow from the [releases](https://github.com/davidepedranz/timetracker/releases) page
2. Add it to Alfred (double-click is usually enough)
3. Run `tt authorize` & `tt setup`
4. Track your time like a pro 😎

## Usage

| Command        | Explanation                                                         |
| -------------- | ------------------------------------------------------------------- |
| tt authorize   | Authorize Alfred to access your Google Calendar.                    |
| tt deauthorize | Revoke the access from your Google Calendar.                        |
| tt setup       | Create a `Tracking` calendar and store its ID in the configuration. |
| tt start       | Start tracking a new task.                                          |
| tt stop        | Stop tracking an ongoing task and add it to Google Calendar.        |
| tt list        | List the ongoing tasks.                                             |
| tt cancel      | Cancel an ongoing task.                                             |
| tt track       | Track a new task with already known duration.                       |
| tt update      | Check if there are updates available.                               |

*Pro trick*: you can omit the `tt` prefix.

## Contribute

If you find any bug or want to propose a new feature, please open an issue to discuss it.

## License

This repository contains free software released under the MIT Licence.
Please check out the [LICENSE](./LICENSE) file for details.
