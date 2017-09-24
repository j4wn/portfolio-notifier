# Portfolio Notifier

With the [deprecation of Google Finance Portfolios](https://support.google.com/finance/answer/7534448?hl=en) I had to move my portfolio to Google Sheets. This is a little app that sends me an email every day with my portfolio balance.

## Authentication

This app expects [Application Default Credentials](https://developers.google.com/identity/protocols/application-default-credentials) to be in place and for those credentials to have [Domain Wide Delegation](https://developers.google.com/identity/protocols/OAuth2ServiceAccount#delegatingauthority) in place with the correct scopes.

:rotating_light: **These credentials are sensitive, they can be used to send email on your domain.**

The required scopes are:
- https://www.googleapis.com/auth/gmail.send
- https://www.googleapis.com/auth/spreadsheets.readonly

# Usage

```raw
 Usage of portfolio-notifier:

  -from string
        From address (default "me@example.com")
  -to string
        Recipient of the email (default "you@example.com")
  -sheet string
        Google Sheet to search (default "1lTQNpSixfQwS3y6XrTc-GH9mJQKpf82toB3yjuUNCQQ")
  -cell string
        Sheet Name!Cell containing the value (default "Main Portfolio!P8")
```