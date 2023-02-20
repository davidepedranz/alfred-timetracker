# Alfred TimeTracker

[![Build Status](https://travis-ci.com/davidepedranz/alfred-timetracker.svg?branch=master)](https://travis-ci.com/davidepedranz/alfred-timetracker)
[![GolangCI](https://golangci.com/badges/github.com/davidepedranz/alfred-timetracker.svg)](https://golangci.com/r/github.com/davidepedranz/alfred-timetracker)

Track your time on Google Calendar 📅 using [Alfred](https://www.alfredapp.com/).

Have you ever asked yourself the question, `What did I do today?`.
If you did and could not answer, this Alfred workflow is for you.
Sometimes we have so many different tasks and interrupt that it is easy to lose track.
You can track them manually, but this approach is error-prone, time-consuming, and tedious.
This workflow lets you to leverage Alfred's power to track your work easily.

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

_Pro trick_: you can omit the `tt` prefix.

## Contribute

If you find any bugs or want to propose a new feature, please open an issue to discuss it.

## Security

On February. 2022, Google deprecated the use of Loopback IP addresses on the Chrome OAuth client type.
The change forced us to migrate to the Desktop OAuth client type.
Desktop clients require a `client_secret` as part of the OAuth authorization flow, which we need to distribute together as part of the Alfred package.
Without appropriate protection, an attacker able to sniff HTTP requests on the machine might try to intercept the OAuth `code` and exchange it for a valid `access_token`.
This project prevents this and similar attacks by implementing the OAuth [Proof Key for Code Exchange (PKCE) extension](https://www.rfc-editor.org/rfc/rfc7636).

## License

This repository contains free software released under the MIT Licence.
Please check out the [LICENSE](./LICENSE) file for details.
