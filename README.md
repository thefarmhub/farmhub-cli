# FarmHub CLI

[![Version](https://badge.fury.io/gh/thefarmhub%2Ffarmhub-cli.svg)](https://badge.fury.io/gh/thefarmhub%2Ffarmhub-cli)
![Test](https://github.com/thefarmhub/farmhub-cli/workflows/test/badge.svg)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)

Handling your IoT devices and data with ease.

## Getting Started

### Login to your FarmHub account

```
farmhub login
```

### Generating the code for your device

Once you have a project on your FarmHub dashboard, just run this command to walk through a sensor setup process with code generation

```
farmhub generate -o sketch.ino
```

### Putting code on your device

You've probably generated the code on your [dashboard](https://my.farmhub.ag) for your Aquaponics or Hydroponics kit from Atlas Scientific and need to put it on the device.  Simply download that file from your dashboard and run the following commands:

```
farmhub flash <pathtosketch>

# Example if you used the above generate command
farmhub flash sketch.ino
```

### Calibrating a device

You can interact with your device by running:

```
farmhub monitor
```

It will connect to your device via a serial monitor where you can run commands.

## Installing

### Homebrew

Installing on a Mac with [homebrew](https://brew.sh/):

```
brew tap farmhub/famrhub-cli https://github.com/thefarmhub/farmhub-cli
brew install farmhub
```

### Scoop

Installing on a windows machine with [scoop](https://scoop.sh/):

```
scoop bucket add farmhub-cli https://github.com/thefarmhub/farmhub-cli
scoop install farmhub
```

### Manual Download

If you're not using a package manager you can just download the binary from our [releases](https://github.com/thefarmhub/farmhub-cli/releases) page or use the script below to automatically download it per your operating system requirements.

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/thefarmhub/farmhub-cli/main/download.sh)"
```
