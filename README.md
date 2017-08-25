# Webhook Forwarding
## From Bitbucket to Glip

Glip is not easy to integrate with. This simple server listens for bitbucket to send a post request indicating that a pull request has been opened. Next a post request is made to glip in a format that will display the author, pull request title, and link to the pull request. There are certainly better ways to acomplish the goal of bitbucket to glip integration, but this was a good toy project to learn new things.
