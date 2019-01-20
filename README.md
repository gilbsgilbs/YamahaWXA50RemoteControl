# YamahaWXA50RemoteControl

Very basic CLI written in Go that can communicate with a Yamaha WXA-50
amplifier.

It just fulfills my very basic needs. I may or may not improve it over time
depending on how my needs change.

## Usage

0. Configure your network so that your amplifier gets a static IP address lease.
0. Create a config file in `$HOME/.config/wxa50/config.yml`:
   ```
   endpoint: http://<IP address of the amplifier>
   ```
   *Note*: You can omit this config file, but you'll have to set the `--endpoint`
   flag each time you use the CLI.
0. Get usage information by running the main file.

## What's done?

- Power on/off
- Increase/Decrease/Get volume
- Mute/Unmute/Toggle mute
- Get/Change current audio source (aka Input Selection)

## How's it done?

This project uses [Cobra](https://github.com/spf13/cobra) for CLI interactions.

Apart from [this document](https://goo.gl/kL9igU), I didn't find any API
documentation. I didn't test the API routes provided in the document because
reverse engineering isn't hard at all in this case and seems like it would
allow more improvements in the future. Just go on the HTTP endpoint of the
amplifier and sniff the packets through wireshark, or even simpler, using the
"Network" tab in the Developer Console of any modern browser.

## TODO

- Automate releases
- Add more features
- Check if it works with more Yamaha devices
- First go project, style could probably be improved
- Maybe some basic regression tests (mocking the server isn't really interesting)
