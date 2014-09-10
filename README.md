# Share channels between Slack accounts!

## Why did we develop Slackline?

We are big fans of [Slack][slack] and have been using it for a while
now. Our friends at [Vizzuality][vizzuality] started using it recently
and we missed not being able to speak with them. We started with a light
integration with Slack using 4 webhooks and this project.

This project evolved into a simpler and better version you can pay for: [slackline.io](http://slackline.io)

We are maintaining this small project free and open source for those who don't need slackline.io.

## What is the difference between the free and paid Slackline?

The free version is a light integration with the Slack API just using incoming and outgoing webhooks.

[slackline.io](http://slackline.io) allows you to:
 - See the avatars from the users in the other team.
 - Easier to setup: create your shared channel, share the URL to the shared channel with somebody in the other team and they can connect their team by themselves. No need to exchange tokens or manually craft URLs.
 - Support for more than two teams connected to the same shared channel.
 - Any team can change the channel they are using with no need to change anything in the other teams.

[slackline.io](http://slackline.io) is still under development, but you can try it for free connecting to our [#slackline shared channel](http://slackline.io/shared_channels/slackline) and we'll notify you as soon as it's available.

## How do I use slackline for free?

We would obviously recommend you use [slackline.io](http://slackline.io) since it's easier to setup and seeing avatars from those on the other side really rocks. But, if you want to use this version, that's cool with us too! :)

You just need to follow the following steps to setup a channel.

 1. Create a channel you want to share with another team.
 2. Create an Incoming WebHook integration and select the channel you created.
 3. Copy the Incoming WebHook token (you can find it in the left column
    from the integration page).
 4. Create a URL with the following format: ```http://slackline.herokuapp.com/bridge/?token=[TOKEN]&domain=[YOUR_SLACK_DOMAIN]``` send it to the person setting up the other team.
 5. The person setting up the other team will send you a similar
    URL with their domain and token, create an Outgoing WebHook with
    that URL and the channel you created in step 1.

Once you have done this in both teams, you will have a chat-room
shared by both teams.

Here you have an example of a Outgoing WebHook URL:

```
http://slackline.herokuapp.com/bridge/?token=bcaa5867b1d42142b74eDVA4&domain=avengers.slack.com
```

## How does it work?

We are just bridging hooks, we don't store any messages going through
the bridge.

## DISCLAIMER

This project is not officially supported by [Slack][slack] and they are
not responsible for the use you make of this and won't give you any
support related to this integration.

Now that we are talking disclaimers... I'm not responsible either for
any use of the software. Use at your own risk and to be safe it might be
good if you deploy this yourself rather than using my Heroku deployment ;)

## MIT LICENSE

Copyright (C) <2014> Ernesto Jimenez <erjica@gmail.com>


Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


[slack]: http://slack.com
[vizzuality]: http://vizzuality.com
