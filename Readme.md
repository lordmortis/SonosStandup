# Overview

This is an CLI tool that can handle timed standup playback on a sonos device.
It allows you to add links to audio files that will be randomly selected when you run standup.
It maintains state to ensure every track in the rotation is played before another track is played again.
It also will set the state of the sonos playback back to what it was before standup.
Each of the above is handled by running various commands.

## Configuration

[config.yaml.sample](config.yaml.sample) contains the configuration template. 

 * **SonosIP** is the ip address/hostname of the sonos device. It does not have to be on the local network as we aren't using zeroconf/bonjour to find it.
 * **StatePath** is where to store the current state of standup playback and the state of playback to restore
 * **Volume** is what to set the volume to (out of 100) when playing standup - often you want it louder than normal music playback

## Commands

You can run with the argument `--help` to get info about the commands.

### `info`

This command outputs the current state - what tracks are listed, which ones are played or not, and what the state will be restored upon running `postStandup`

### `addFile` 

This command will add the specified link (the supplied argument) to the playback rotation and also to the unplayed list

### `removeFile`

This command will remove the specified link (the supplied argument) from playback rotation, played and unplayed lists

### `runStandup`

This will pause the currently playing music, record the position and volume, and pick a random track from the available ones to play for standup at the specified volume in the config file

### `postStandup`

This will return the playhead to the position and state before playing standup, and restore the correct volume.

# Developing

## Prerequisites

- [Golang development environment](https://golang.org/dl/)

## Config file setup

1. copy config.yaml.sample to config.yaml, and set the options for your environment

