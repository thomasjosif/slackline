# SLACKLINE FORK
### This was used to link our slack channels when we were using slack for communications.

It's a fork of the original slackline, but I added avatar support without having to pay for it (yuck)

 1. Create a channel you want to share with another team.
 2. Create an Incoming WebHook integration and select the channel you created.
 3. Copy the Incoming WebHook token (you can find it in the left column
    from the integration page).
 4. Create a URL with the following format: ```https://hellsgamers-slack-bridge.herokuapp.com/bridge/?token=[TOKEN]&domain=[YOUR_SLACK_DOMAIN]``` send it to the person setting up the other team.
 5. The person setting up the other team will send you a similar
    URL with their domain and token, create an Outgoing WebHook with
    that URL and the channel you created in step 1.
