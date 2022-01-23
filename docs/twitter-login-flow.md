# Twitter Login Flow

_view on large screen or with wrap off_

_Mobile flow_

**mobile**                          **tribist.com**                             **twitter.com**
    |     /users/login?source=twitter     |                                           |
    --------------->----------------------+                                           |
    |                                     |                                           |
    |                                     |     `POST /oauth/request_token`           |                <---- passes the `oauth_callback`
    |                                     ------------------->------------------------+
    |                                     |                                           |
    |                                     |                                           |
    |                                     +------------------<-------------------------                <---- returns `oauth_token`, `oauth_token_secret` `oauth_callback_confirmed`
    |                                     |                                           |
    | 302 to `/oauth/authenticate?oauth_token`                                        |
    +--------------<-----------------------                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                `GET  /oauth/authenticate?oauth_token`                           |
    -------------------------------------->-------------------------------------------+
    |                                     |                                           |
    |                `302 /users/authorize/twitter?oauth_token,oauth-verifier`        |
    +-------------------------------------<--------------------------------------------
    |                                     |                                           |
    |   `GET /users/authorize/twitter?oauth_token,oauth-verifier`                     |
    --------------->----------------------+                                           |
    |                                     |                                           |
    |                                     |    `POST /oauth/access_token`             |
    |                                     -------------------->-----------------------+                <---- passes `oauth_verifier`
    |                                     |                                           |
    |                                     |                                           |
    |                                     +-------------------<------------------------                <---- returns `oauth_token` `oauth_token_secret` `user_id` `screen_name`
    |                                     |                                           |
    |    `302 found`                      |                                           |
    +--------------<-----------------------                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
