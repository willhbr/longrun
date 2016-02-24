# LongRun

_Run commands that take ages without looking_

LongRun lets you run a command that may take a few minutes or hours on your server and receive a push notification from [PushBullet](https://www.pushbullet.com) on any of your devices when the command completes with the output of the command.

To install it, just download and run the `install.sh` script, or run this:

    curl https://raw.githubusercontent.com/JavaNut13/longrun/master/install.sh | bash

This will download the correct binary for your system and move it to `/usr/local/bin`. If there isn't a binary for your OS or it doesn't detect it correctly, just download the binary from the `builds` folder.

Your Pushbullet API key ([from your settings page](https://www.pushbullet.com/#settings/account)) must be stored in `~/.longrun-token` - this is setup by the install script, but if you installed LongRun by building it yourself or you changed your API key then you can just modify this file.

Once LongRun is installed, you can run a command like so:

    lrun apt-get install ruby
    
Or, if you don't care about the output:

    lrun "apt-get update > /dev/null"

> Note: The command will not read from `STDIN`, so only run commands that require no input.
> 
> To be able to run a command and disconnect from SSH without killing the process, use something like [screen](https://www.gnu.org/software/screen/) or [tmux](http://tmux.github.io). I personally prefer tmux.
