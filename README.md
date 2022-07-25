# ssh-notify

ssh-notify is a Go program, originally written in Bash, that sends a message to a Discord webhook whenever an SSH login occurs, and when the user logs out again.

The format of the message is as follows:
```
{hostname}: {user} logged in (remote host: {remote_ip}).
```

## Configuration

To configure ssh-notify, copy the example config (`ssh-notify.json.example`) to the same directory. If your webhook URL is `https://discord.com/api/webhooks/374355038877978415/jzR_6rd2USWNq2XdXvnfC5QCXaH7xxAKZDQCXaWh-w5H7xxFAuuIDlzPCSKuH7xxLFD3` then the `WebhookId` is `374355038877978415` and the `WebhookToken` is `jzR_6rd2USWNq2XdXvnfC5QCXaH7xxAKZDQCXaWh-w5H7xxFAuuIDlzPCSKuH7xxLFD3`. You can also configure ssh-notify to not notify for logins to certain users, useful if you have automated scripts logging into your servers.

You will also need to tell SSH to use ssh-notify. Do this by adding the following at the bottom of the `/etc/pam.d/sshd` file. Ensure that ssh-notify is executable.

```
session    optional    pam_exec.so    /path/to/ssh-notify --config /path/to/ssh-notify.json
```