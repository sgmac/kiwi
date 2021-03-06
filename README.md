# KiwiVM

Yet another not very useful CLI tool. Why not? It turns out I got a very cheap server long time ago, just 10USD per year, OpenVZ not very useful to be honest. 
I took a look to the API and decided to implement first a package to interact with the VPS. Now it's time for the tool.

This API doesn't have much sense, you need a VeID and PrivateKey for each VPS you have, that means you will need to use different clients when using the [API](https://github.com/sgmac/bandwagon).

![kiwivm](./kiwi.png)

# Setup

First time you run, creates a default directory and an empty configuration.

**~/.kiwi/config**

```bash
[VPS]
  VeID   = "<ADD_VPS_ID>"
  APIKey = "<PRIVATE_KEY>"
```

At this point you can request actions on that VPS.


# Usage

Some of the enabled actions are listed below:

```bash
NAME:
   kiwi - Manage your Bandwagon VPS

USAGE:
   kiwi [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     images, i    List OS images.
     status, f    Status VPS.
     start, s     Start VPS.
     stop, t      Stop VPS.
     kill, k      Kill VPS.
     reboot, r    Reboot VPS.
     hostname, h  Set hostname.
     install, l   Install OS.
     help, h      Shows a list of commands or help for one command
```
